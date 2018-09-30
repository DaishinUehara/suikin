package skselflib_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/DaishinUehara/suikin/skselflib"
)

func TestExec(t *testing.T) {
	var stdin io.Reader
	var stdout io.Writer
	var stderr io.Writer
	var stdin1 *bytes.Buffer
	var stdout1 *bytes.Buffer
	var stderr1 *bytes.Buffer
	var err error
	incolumnname := make([]string, 0)
	outcolumnname := make([]string, 0)

	err = skselflib.Exec(stdin, stdout, stderr, incolumnname, outcolumnname)
	if err != nil {
		t.Logf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin, stdout, stderr, incolumnname, outcolumnname, err)
	} else {
		t.Errorf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin, stdout, stderr, incolumnname, outcolumnname, err)
	}

	stdin1 = bytes.NewBufferString("")
	err = skselflib.Exec(stdin1, stdout, stderr, incolumnname, outcolumnname)
	if err != nil {
		t.Logf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin1, stdout, stderr, incolumnname, outcolumnname, err)
	} else {
		t.Errorf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin1, stdout, stderr, incolumnname, outcolumnname, err)
	}

	stdin1 = bytes.NewBufferString("")
	stdout1 = new(bytes.Buffer)
	err = skselflib.Exec(stdin1, stdout1, stderr, incolumnname, outcolumnname)
	if err != nil {
		t.Logf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin1, stdout1, stderr, incolumnname, outcolumnname, err)
	} else {
		t.Errorf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin1, stdout1, stderr, incolumnname, outcolumnname, err)
	}

	stdin1 = bytes.NewBufferString("")
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	err = skselflib.Exec(stdin1, stdout1, stderr1, incolumnname, outcolumnname)
	if err == nil && stdout1.String() == "" && stderr1.String() == "" {
		t.Logf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, err)
	} else {
		t.Errorf("[NG]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, err)
	}

}
