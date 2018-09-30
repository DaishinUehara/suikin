package skselflib_test

import (
	"io"
	"testing"

	"github.com/DaishinUehara/suikin/skselflib"
)

func TestExec(t *testing.T) {
	var stdin io.Reader
	var stdout io.Writer
	var stderr io.Writer
	incolumnname := make([]string, 0)
	outcolumnname := make([]string, 0)

	err := skselflib.Exec(stdin, stdout, stderr, incolumnname, outcolumnname)
	if err != nil {
		t.Logf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v)\n", stdin, stdout, stderr, incolumnname, outcolumnname, err)
	} else {
		t.Errorf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v)\n", stdin, stdout, stderr, incolumnname, outcolumnname, err)
	}

}
