package skpipelib_test

// TODO 諸々

import (
	"bytes"
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
	pe := new(skpipelib.SkMulti)
	infield1 := make([]string, 0, 5)
	outfiled1 := make([]string, 0, 5)
	pe.AddExec(exe1, infield1, outfiled1)
	stdin1 := bytes.NewBufferString("テスト")
	stdout1 := new(bytes.Buffer)
	stderr1 := new(bytes.Buffer)

	errAr1 := pe.MultiExec(stdin1, stdout1, stderr1)
	if 1 == len(errAr1) &&
		errAr1[0] == nil &&
		stdout1.String() == "" &&
		stderr1.String() == "" {
		t.Logf("[OK]:skselflib.MultiExec(%v,%v,%v,%v,%v):errAr1=%v,len(errAr1)=%d\n", stdin1, stdout1, stderr1, infield1, infield1, errAr1, len(errAr1))
	} else {
		t.Errorf("[NG]:skselflib.MultiExec(%v,%v,%v,%v,%v):errAr1=%v,len(errAr1)=%d\n", stdin1, stdout1, stderr1, infield1, infield1, errAr1, len(errAr1))
	}

}

func TestPipeExec(t *testing.T) {
}
