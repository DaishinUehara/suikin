package skpipelib

import (
	"bufio"
	"io"
	"sync"

	"github.com/DaishinUehara/suikin/skerrlib"
)

// TODO 諸々

// SkExecIf Pipeの実行対象となるインターフェース
type SkExecIf interface {
	Exec(io.Reader, io.Writer, io.Writer) error
}

// SkMultiIf Pipeへのアクセスインターフェース
type SkMultiIf interface {
	AddExec(SkExecIf)
	AddMultiExec(SkMultiIf)
	MultiExec(io.Reader, io.Writer, io.Writer) ([]error, error)
}

// SkExecInfo SkMultiの実行に必要な構造体を作成
//type SkExecInfo struct {
//	skexec   SkExecIf
//	infield  []string
//	outfield []string
//}

// SkMulti is Struct of SkMultiIf
type SkMulti struct {
	pSkExecInfoArr *([]SkExecIf)
}

// AddExec pipeで実行する処理を追加する
func (sp *SkMulti) AddExec(skexec SkExecIf) {
	var (
		skexecinfoarr []SkExecIf
	)
	if sp.pSkExecInfoArr == nil {
		skexecinfoarr = make([]SkExecIf, 0, 5)
		sp.pSkExecInfoArr = &skexecinfoarr
	} else {
		skexecinfoarr = *sp.pSkExecInfoArr
	}
	skexecinfoarr = append(skexecinfoarr, skexec)
	sp.pSkExecInfoArr = &skexecinfoarr

}

// AddMultiExec pipeで実行する処理を複数追加する
func (sp *SkMulti) AddMultiExec(skemultixec SkMultiIf) {
	var (
		skexecinfoarr []SkExecIf
	)
	if sp.pSkExecInfoArr == nil {
		skexecinfoarr = make([]SkExecIf, 0, 5)
		sp.pSkExecInfoArr = &skexecinfoarr
	} else {
		skexecinfoarr = *sp.pSkExecInfoArr
	}
	for _, skexec := range skexecinfoarr {
		skexecinfoarr = append(skexecinfoarr, skexec)
	}
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
		execlen       int
		skexecinfoar  []SkExecIf
		wg            sync.WaitGroup
	)

	errWriteBuffer := bufio.NewWriter(iose)
	stdWriteBuffer := bufio.NewWriter(ioso)

	if sp.pSkExecInfoArr != nil {
		skexecinfoar = *sp.pSkExecInfoArr
		execlen = len(skexecinfoar)
	} else {
		return errAr, skerrlib.ErrNotInitialized{NoInitializedItem: "sp.pSkExecInfoArr", StackTrace: skerrlib.PrintCallStack()}
	}

	pipeReaderArr := make([]*io.PipeReader, 0, execlen)
	pipeWriterArr := make([]*io.PipeWriter, 0, execlen)

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

		// Folow lines plug given loop values for the valiable to use go-routine
		if i == 0 {
			wg.Add(1)
			go func(ir io.Reader, iw *io.PipeWriter, ier *io.PipeWriter, skexec SkExecIf) {
				err := skexec.Exec(ir, iw, ier)
				errAr = append(errAr, err)
				iw.Close()
				ier.Close()
				wg.Done()
			}(iosr, pipeWriter, pipeErrWriter, execinfo)
		} else {
			wg.Add(1)
			go func(ir io.Reader, iw *io.PipeWriter, ier *io.PipeWriter, skexec SkExecIf) {
				err := skexec.Exec(ir, iw, ier)
				errAr = append(errAr, err)
				iw.Close()
				ier.Close()
				wg.Done()
			}(pipeReaderArr[i-1], pipeWriter, pipeErrWriter, execinfo)

		}
		wg.Add(1)
		go func(pErrR *io.PipeReader, errW *bufio.Writer) {
			// This Go Routine Read Error From Last Execute And Write Buffer.
			sc := bufio.NewScanner(pErrR)
			for sc.Scan() {
				errW.Write(sc.Bytes())
				errW.WriteString("\n")
			}
			pErrR.Close()
			errW.Flush()
			wg.Done()
		}(pipeErrReader, errWriteBuffer)

	}

	// This Go Routine Read Standard Output from Last Execute And Write Buffer.
	sc := bufio.NewScanner(pipeReaderArr[execlen-1])
	for sc.Scan() {
		stdWriteBuffer.Write(sc.Bytes())
		stdWriteBuffer.WriteString("\n")
	}
	pipeReader.Close()
	stdWriteBuffer.Flush()

	errAr = errAr[:cap(errAr)]
	wg.Wait()
	return errAr, nil
}
