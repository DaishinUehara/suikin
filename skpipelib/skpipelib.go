package skpipelib

import "io"

// SkExec Pipeの実行対象となるインターフェース
type SkExec interface {
	Exec(io.Reader, io.Writer, io.Writer, []string, []string) error
}
