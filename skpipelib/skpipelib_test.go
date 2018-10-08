package skpipelib_test

// TODO 諸々

import (
	"io"
	"testing"

	"github.com/DaishinUehara/suikin/skpipelib"
)

type Exe1 struct {
}

func (ex1 Exe1) Exec(io.Reader, io.Writer, io.Writer, []string, []string) error {
	return nil
}

func TestAddExec(t *testing.T) {
	exe1 := new(Exe1)
	pe := new(skpipelib.SkPipe)
	inst1 := make([]string, 0, 5)
	outst1 := make([]string, 0, 5)
	pe.AddExec(exe1, inst1, outst1)

}

func TestPipeExec(t *testing.T) {
}
