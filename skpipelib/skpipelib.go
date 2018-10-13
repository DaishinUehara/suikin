package skpipelib

import (
	"bufio"
	"fmt"
	"io"

	"github.com/DaishinUehara/suikin/skerrlib"
)

// TODO 諸々

// SkExecIf Pipeの実行対象となるインターフェース
type SkExecIf interface {
	Exec(io.Reader, io.Writer, io.Writer, []string, []string) error
}

// SkMultiIf Pipeへのアクセスインターフェース
type SkMultiIf interface {
	AddExec(SkExecIf, []string, []string)
	MultiExec(io.Reader, io.Writer, io.Writer) ([]error, error)
}

// SkExecInfo SkMultiの実行に必要な構造体を作成
type SkExecInfo struct {
	skexec   SkExecIf
	infield  []string
	outfield []string
}

// SkMulti is Struct of SkMultiIf
type SkMulti struct {
	pSkExecInfoArr *([]SkExecInfo)
}

// AddExec pipeで実行する処理を追加する
func (sp *SkMulti) AddExec(skexec SkExecIf, infield []string, outfield []string) {
	var (
		skexecinfoarr []SkExecInfo
		pskeinfo      *SkExecInfo
	)
	if sp.pSkExecInfoArr == nil {
		skexecinfoarr = make([]SkExecInfo, 0, 5)
		sp.pSkExecInfoArr = &skexecinfoarr
	} else {
		skexecinfoarr = *sp.pSkExecInfoArr
	}
	pskeinfo = new(SkExecInfo)
	pskeinfo.skexec = skexec
	pskeinfo.infield = infield
	pskeinfo.outfield = outfield
	skexecinfoarr = append(skexecinfoarr, *pskeinfo)
	sp.pSkExecInfoArr = &skexecinfoarr

}

// MultiExec PipeExec make go channel and execute skexec
func (sp *SkMulti) MultiExec(iosr io.Reader, ioso io.Writer, iose io.Writer) ([]error, error) {

	var (
		pipeReader    *io.PipeReader
		pipeWriter    *io.PipeWriter
		pipeErrReader *io.PipeReader
		pipeErrWriter *io.PipeWriter
		errAr         []error
		goi           int
		execlen       int
		skexecinfoar  []SkExecInfo
		goinfo        SkExecInfo
	)

	errWriteBuffer := bufio.NewWriter(iose)
	stdWriteBuffer := bufio.NewWriter(ioso)

	if sp.pSkExecInfoArr != nil {
		skexecinfoar = *sp.pSkExecInfoArr
		execlen = len(skexecinfoar)
	} else {
		return errAr, skerrlib.ErrNotInitialized{PkgMethodName: "skpipelib.MultiExec", NoInitializedItem: "sp.pSkExecInfoArr"}
	}

	pipeReaderArr := make([]*io.PipeReader, 0, execlen)
	pipeWriterArr := make([]*io.PipeWriter, 0, execlen)

	fmt.Printf("%v", pipeReaderArr)

	// 実体が変わらないようあらかじめサイズ指定して確保
	pipeErrReaderArr := make([]*io.PipeReader, 0, execlen)
	pipeErrWriterArr := make([]*io.PipeWriter, 0, execlen)

	errAr = make([]error, 0, execlen)

	for i, execinfo := range skexecinfoar {

		// Pipe for Stdion Stdout
		pipeReader, pipeWriter = io.Pipe()
		pipeReaderArr = append(pipeReaderArr, pipeReader)
		pipeWriterArr = append(pipeWriterArr, pipeWriter)

		// Pipe for Error
		pipeErrReader, pipeErrWriter = io.Pipe()
		pipeErrReaderArr = append(pipeErrReaderArr, pipeErrReader)
		pipeErrWriterArr = append(pipeErrWriterArr, pipeErrWriter)

		goi = i           // goルーチンに格納するため、ループ変数から代入
		goinfo = execinfo // goルーチンに格納するため、ループ変数から代入

		go func() {
			if goi == 0 {
				err := goinfo.skexec.Exec(iosr, pipeWriter, pipeErrWriter, goinfo.infield, goinfo.outfield)
				errAr = append(errAr, err)
			} else {
				err := goinfo.skexec.Exec(pipeReaderArr[goi-1], pipeWriter, pipeErrWriter, goinfo.infield, goinfo.outfield)
				errAr = append(errAr, err)
			}
			pipeWriter.Close()
			pipeErrWriter.Close()
		}()

		go func() {
			sc := bufio.NewScanner(pipeErrReader)
			for sc.Scan() {
				errWriteBuffer.Write(sc.Bytes())
			}
			pipeErrReader.Close()
			errWriteBuffer.Flush()
		}()

		if goi == execlen-1 {
			// 最後の場合、パイプからバッファに出力
			go func() {
				sc := bufio.NewScanner(pipeReader)
				for sc.Scan() {
					stdWriteBuffer.Write(sc.Bytes())
				}
				pipeReader.Close()
				stdWriteBuffer.Flush()
			}()
		}

	}
	errAr = errAr[:cap(errAr)]
	return errAr, nil
}
