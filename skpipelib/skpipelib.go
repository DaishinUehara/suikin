package skpipelib

import (
	"io"
)

// SkExecIf Pipeの実行対象となるインターフェース
type SkExecIf interface {
	Exec(io.Reader, io.Writer, io.Writer, []string, []string) error
}

// SkPipeIF Pipeへのアクセスインターフェース
type SkPipeIf interface {
	GetStdin() (io.Writer, error)
	AddExec(SkExecIf) error
	GetStdout() io.Reader
	GetStderr() io.Reader
}
