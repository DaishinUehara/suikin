package skpipelib

import (
	"bufio"
	"fmt"
	"io"
)

// TODO 諸々

// SkExecIf Pipeの実行対象となるインターフェース
type SkExecIf interface {
	Exec(io.Reader, io.Writer, io.Writer, []string, []string) error
}

// SkMultiIf Pipeへのアクセスインターフェース
type SkMultiIf interface {
	AddExec(SkExecIf) error
	MultiExec(io.Reader, io.Writer, io.Writer) []error
}

// SkMulti is Struct of SkMultiIf
type SkMulti struct {
	// SkExecを格納した配列
	skExecAr []SkExecIf

	// 入力フィールド用の配列
	infieldArr [][]string

	// 出力フィールド用の配列
	outfieldArr [][]string
}

// AddExec pipeで実行する処理を追加する
func (sp SkMulti) AddExec(skexec SkExecIf, infield []string, outfield []string) {
	if sp.skExecAr == nil {
		sp.skExecAr = make([]SkExecIf, 0, 5)
		sp.infieldArr = make([][]string, 0, 5)
		sp.outfieldArr = make([][]string, 0, 5)
	}
	sp.skExecAr = append(sp.skExecAr, skexec)
	sp.skExecAr = sp.skExecAr[:cap(sp.skExecAr)]
	sp.infieldArr = append(sp.infieldArr, infield)
	sp.infieldArr = sp.infieldArr[:cap(sp.infieldArr)]
	sp.outfieldArr = append(sp.outfieldArr, outfield)
	sp.outfieldArr = sp.outfieldArr[:cap(sp.outfieldArr)]

}

// MultiExec PipeExec make go channel and execute skexec
func (sp SkMulti) MultiExec(iosr io.Reader, ioso io.Writer, iose io.Writer) []error {

	var (
		pipeReader    *io.PipeReader
		pipeWriter    *io.PipeWriter
		pipeErrReader *io.PipeReader
		pipeErrWriter *io.PipeWriter
		errAr         []error
		goi           int
		goexe         SkExecIf
	)

	errWriteBuffer := bufio.NewWriter(iose)
	stdWriteBuffer := bufio.NewWriter(ioso)

	ln := len(sp.skExecAr)

	pipeReaderArr := make([]*io.PipeReader, 0, ln)
	pipeWriterArr := make([]*io.PipeWriter, 0, ln)

	fmt.Printf("%v", pipeReaderArr)

	// 実体が変わらないようあらかじめサイズ指定して確保
	pipeErrReaderArr := make([]*io.PipeReader, 0, ln)
	pipeErrWriterArr := make([]*io.PipeWriter, 0, ln)

	errAr = make([]error, 0, ln)

	for i, skexec := range sp.skExecAr {

		// Pipe for Stdion Stdout
		pipeReader, pipeWriter = io.Pipe()
		pipeReaderArr = append(pipeReaderArr, pipeReader)
		pipeWriterArr = append(pipeWriterArr, pipeWriter)

		// Pipe for Error
		pipeErrReader, pipeErrWriter = io.Pipe()
		pipeErrReaderArr = append(pipeErrReaderArr, pipeErrReader)
		pipeErrWriterArr = append(pipeErrWriterArr, pipeErrWriter)

		goexe = skexec
		goi = i

		go func() {
			if goi == 0 {
				err := goexe.Exec(iosr, pipeWriter, pipeErrWriter, sp.infieldArr[goi], sp.outfieldArr[goi])
				errAr = append(errAr, err)
			} else {
				err := goexe.Exec(pipeReaderArr[goi-1], pipeWriter, pipeErrWriter, sp.infieldArr[goi], sp.outfieldArr[goi])
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

		if goi == len(sp.skExecAr)-1 {
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
	return errAr
}
