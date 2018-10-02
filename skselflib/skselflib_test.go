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
	var stdout1_str string
	var stderr1 *bytes.Buffer
	var stderr1_str string
	var err error
	var incolumnname []string
	var outcolumnname []string
	incolumnname = make([]string, 0)
	outcolumnname = make([]string, 0)

	// ↓initial error

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

	// ↑initial error

	stdin1 = bytes.NewBufferString("")
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	err = skselflib.Exec(stdin1, stdout1, stderr1, incolumnname, outcolumnname)
	if err == nil && stdout1.String() == "" && stderr1.String() == "" {
		t.Logf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, err)
	} else {
		t.Errorf("[NG]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, err)
	}

	stdin1 = bytes.NewBufferString("")
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	incolumnname = make([]string, 0)
	incolumnname = append(incolumnname, "項目1")

	err = skselflib.Exec(stdin1, stdout1, stderr1, incolumnname, outcolumnname)
	if err != nil && stdout1.String() == "" && stderr1.String() == "" {
		t.Logf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, err)
	} else {
		t.Errorf("[NG]:skselflib.Exec(%v,%v,%v,%v,%v):err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, err)
	}

	stdin1 = bytes.NewBufferString("項目1 項目2\n")
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	incolumnname = make([]string, 0)
	incolumnname = append(incolumnname, "項目1")

	err = skselflib.Exec(stdin1, stdout1, stderr1, incolumnname, outcolumnname)
	stdout1_str = stdout1.String()
	stderr1_str = stderr1.String()
	if err != nil && stdout1_str == "" && stderr1_str == "" {
		t.Logf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, stdout1_str, stderr1_str, err)
	} else {
		t.Errorf("[NG]:skselflib.Exec(%v,%v,%v,%v,%v):stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, stdout1_str, stderr1_str, err)
	}

	stdin1 = bytes.NewBufferString("項目1 項目2\n")
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	incolumnname = make([]string, 0)
	incolumnname = append(incolumnname, "項目1")
	outcolumnname = make([]string, 0)
	outcolumnname = append(outcolumnname, "項目a")

	err = skselflib.Exec(stdin1, stdout1, stderr1, incolumnname, outcolumnname)
	stdout1_str = stdout1.String()
	stderr1_str = stderr1.String()
	if err == nil && stdout1_str == "項目a\n" && stderr1_str == "" {
		t.Logf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, stdout1_str, stderr1_str, err)
	} else {
		t.Errorf("[NG]:skselflib.Exec(%v,%v,%v,%v,%v):stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, stdout1_str, stderr1_str, err)
	}

	stdin1 = bytes.NewBufferString("項目1 項目2") // 改行がない場合
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	incolumnname = make([]string, 0)
	incolumnname = append(incolumnname, "項目1")
	outcolumnname = make([]string, 0)
	outcolumnname = append(outcolumnname, "項目a")

	err = skselflib.Exec(stdin1, stdout1, stderr1, incolumnname, outcolumnname)
	stdout1_str = stdout1.String()
	stderr1_str = stderr1.String()
	if err == nil && stdout1_str == "項目a\n" && stderr1_str == "" {
		t.Logf("[OK]:skselflib.Exec(%v,%v,%v,%v,%v):stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, stdout1_str, stderr1_str, err)
	} else {
		t.Errorf("[NG]:skselflib.Exec(%v,%v,%v,%v,%v):stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, incolumnname, outcolumnname, stdout1_str, stderr1_str, err)
	}

}
