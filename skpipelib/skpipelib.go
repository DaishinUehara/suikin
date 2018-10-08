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

// SkPipeIf Pipeへのアクセスインターフェース
type SkPipeIf interface {
	//	GetStdin() (io.Writer, error)
	AddExec(SkExecIf) error
	//	GetStdout() io.Reader
	//	GetStderr() io.Reader
	PipeExec(io.Reader, io.Writer, io.Writer) error
}

// SkPipe is Struct of SkPipeIf
type SkPipe struct {
	skExecAr           []SkExecIf
	skPipeErrReaderArr []*io.PipeReader
	infieldArr         [][]string
	outfieldArr        [][]string
}

// AddExec pipeで実行する処理を追加する
func (sp SkPipe) AddExec(skexec SkExecIf, infield []string, outfield []string) {
	if sp.skExecAr == nil {
		sp.skExecAr = make([]SkExecIf, 0, 5)
		sp.infieldArr = make([][]string, 0, 5)
		sp.outfieldArr = make([][]string, 0, 5)
	}
	sp.skExecAr = append(sp.skExecAr, skexec)
	sp.infieldArr = append(sp.infieldArr, infield)
	sp.outfieldArr = append(sp.outfieldArr, outfield)

}

// PipeExec PipeExec make go channel and execute skexec
func (sp SkPipe) PipeExec(iosr io.Reader, ioso io.Writer, iose io.Writer) (err error) {

	var (
		pipeReader    *io.PipeReader
		pipeWriter    *io.PipeWriter
		pipeErrReader *io.PipeReader
		pipeErrWriter *io.PipeWriter
	)

	errWriteBuffer := bufio.NewWriter(iose)
	stdWriteBuffer := bufio.NewWriter(ioso)

	pipeReaderArr := make([]*io.PipeReader, 0, 5)
	pipeWriterArr := make([]*io.PipeWriter, 0, 5)

	fmt.Printf("%v", pipeReaderArr)

	pipeErrReaderArr := make([]*io.PipeReader, 0, 5)
	pipeErrWriterArr := make([]*io.PipeWriter, 0, 5)

	for i, skexec := range sp.skExecAr {

		// Pipe for Stdion Stdout
		pipeReader, pipeWriter = io.Pipe()
		pipeReaderArr = append(pipeReaderArr, pipeReader)
		sp.skPipeErrReaderArr = pipeErrReaderArr // append時実体が変わる可能性があるため再代入
		pipeWriterArr = append(pipeWriterArr, pipeWriter)

		// Pipe for Error
		pipeErrReader, pipeErrWriter = io.Pipe()
		pipeErrReaderArr = append(pipeErrReaderArr, pipeErrReader)
		pipeErrWriterArr = append(pipeErrWriterArr, pipeErrWriter)

		goexe := skexec
		goi := i

		go func() {
			if goi == 0 {
				goexe.Exec(iosr, pipeWriter, pipeErrWriter, sp.infieldArr[goi], sp.outfieldArr[goi])
			} else {
				goexe.Exec(pipeReaderArr[goi-1], pipeWriter, pipeErrWriter, sp.infieldArr[goi], sp.outfieldArr[goi])
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
	return err
}
