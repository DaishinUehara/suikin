package skpipelib

import (
	"io"
)

// SkExecIf Pipeの実行対象となるインターフェース
type SkExecIf interface {
	Exec(io.Reader, io.Writer, io.Writer, []string, []string) error
}

// SkPipeIf Pipeへのアクセスインターフェース
type SkPipeIf interface {
	GetStdin() (io.Writer, error)
	AddExec(SkExecIf) error
	GetStdout() io.Reader
	GetStderr() io.Reader
	Exec(io.Reader, io.Writer, io.Writer) error
}

// SkPipe is Struct of SkPipeIf
type SkPipe struct {
	skExecAr []SkExecIf
}

// AddExec pipeで実行する処理を追加する
func (sp SkPipe) AddExec(skexec SkExecIf) {
	if sp.skExecAr == nil {
		sp.skExecAr = make([]SkExecIf, 0, 5)
	}
	sp.skExecAr = append(sp.skExecAr, skexec)
}
