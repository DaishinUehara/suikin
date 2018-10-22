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
	//var incolumnname []string
	//var outcolumnname []string
	var skself = new(skselflib.SkSelf)
	//incolumnname = make([]string, 0)
	//outcolumnname = make([]string, 0)
	skself.InColumnName = make([]string, 0)
	skself.OutColumName = make([]string, 0)

	// ↓initial error

	// err = skselflib.Exec(stdin, stdout, stderr, incolumnname, outcolumnname)
	err = skself.Exec(stdin, stdout, stderr)
	if err != nil {
		t.Logf("[OK]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,err=%v\n", stdin, stdout, stderr, skself.InColumnName, skself.OutColumName, err)
	} else {
		t.Errorf("[NG]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,err=%v\n", stdin, stdout, stderr, skself.InColumnName, skself.OutColumName, err)
	}

	stdin1 = bytes.NewBufferString("")
	// err = skselflib.Exec(stdin1, stdout, stderr, incolumnname, outcolumnname)
	err = skself.Exec(stdin1, stdout, stderr)
	if err != nil {
		t.Logf("[OK]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,err=%v\n", stdin1, stdout, stderr, skself.InColumnName, skself.OutColumName, err)
	} else {
		t.Errorf("[NG]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,err=%v\n", stdin1, stdout, stderr, skself.InColumnName, skself.OutColumName, err)
	}

	stdin1 = bytes.NewBufferString("")
	stdout1 = new(bytes.Buffer)
	err = skself.Exec(stdin1, stdout1, stderr)
	if err != nil {
		t.Logf("[OK]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,err=%v\n", stdin1, stdout1, stderr, skself.InColumnName, skself.OutColumName, err)
	} else {
		t.Errorf("[NG]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,err=%v\n", stdin1, stdout1, stderr, skself.InColumnName, skself.OutColumName, err)
	}

	// ↑initial error

	stdin1 = bytes.NewBufferString("")
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	err = skself.Exec(stdin1, stdout1, stderr1)
	if err == nil && stdout1.String() == "" && stderr1.String() == "" {
		t.Logf("[OK]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, err)
	} else {
		t.Errorf("[NG]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, err)
	}

	stdin1 = bytes.NewBufferString("")
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	skself.InColumnName = make([]string, 0)
	skself.InColumnName = append(skself.InColumnName, "項目1")

	err = skself.Exec(stdin1, stdout1, stderr1)
	if err != nil && stdout1.String() == "" && stderr1.String() == "" {
		t.Logf("[OK]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, err)
	} else {
		t.Errorf("[NG]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, err)
	}

	stdin1 = bytes.NewBufferString("項目1 項目2\n")
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	skself.InColumnName = make([]string, 0)
	skself.InColumnName = append(skself.InColumnName, "項目1")

	err = skself.Exec(stdin1, stdout1, stderr1)
	stdout1_str = stdout1.String()
	stderr1_str = stderr1.String()
	if err != nil && stdout1_str == "" && stderr1_str == "" {
		t.Logf("[OK]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, stdout1_str, stderr1_str, err)
	} else {
		t.Errorf("[NG]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, stdout1_str, stderr1_str, err)
	}

	stdin1 = bytes.NewBufferString("項目1 項目2\n")
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	skself.InColumnName = make([]string, 0)
	skself.InColumnName = append(skself.InColumnName, "項目1")
	skself.OutColumName = make([]string, 0)
	skself.OutColumName = append(skself.OutColumName, "項目a")

	err = skself.Exec(stdin1, stdout1, stderr1)
	stdout1_str = stdout1.String()
	stderr1_str = stderr1.String()
	if err == nil && stdout1_str == "項目a\n" && stderr1_str == "" {
		t.Logf("[OK]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, stdout1_str, stderr1_str, err)
	} else {
		t.Errorf("[NG]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, stdout1_str, stderr1_str, err)
	}

	stdin1 = bytes.NewBufferString("項目1 項目2") // 改行がない場合
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	skself.InColumnName = make([]string, 0)
	skself.InColumnName = append(skself.InColumnName, "項目1")
	skself.OutColumName = make([]string, 0)
	skself.OutColumName = append(skself.OutColumName, "項目a")

	err = skself.Exec(stdin1, stdout1, stderr1)
	stdout1_str = stdout1.String()
	stderr1_str = stderr1.String()
	if err == nil && stdout1_str == "項目a\n" && stderr1_str == "" {
		t.Logf("[OK]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, stdout1_str, stderr1_str, err)
	} else {
		t.Errorf("[NG]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, stdout1_str, stderr1_str, err)
	}

	stdin1 = bytes.NewBufferString("項目1 項目2 項目3\n1 2 3")
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	skself.InColumnName = make([]string, 0)
	skself.InColumnName = append(skself.InColumnName, "項目1")
	skself.InColumnName = append(skself.InColumnName, "項目3")
	skself.OutColumName = make([]string, 0)
	skself.OutColumName = append(skself.OutColumName, "項目a")
	skself.OutColumName = append(skself.OutColumName, "項目c")

	err = skself.Exec(stdin1, stdout1, stderr1)
	stdout1_str = stdout1.String()
	stderr1_str = stderr1.String()
	if err == nil && stdout1_str == "項目a 項目c\n1 3\n" && stderr1_str == "" {
		t.Logf("[OK]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, stdout1_str, stderr1_str, err)
	} else {
		t.Errorf("[NG]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, stdout1_str, stderr1_str, err)
	}

	stdin1 = bytes.NewBufferString("項目1 項目2 項目3\n1 2 3\n4 5 6")
	stdout1 = new(bytes.Buffer)
	stderr1 = new(bytes.Buffer)
	skself.InColumnName = make([]string, 0)
	skself.InColumnName = append(skself.InColumnName, "項目3")
	skself.InColumnName = append(skself.InColumnName, "項目1")
	skself.InColumnName = append(skself.InColumnName, "項目3")
	skself.InColumnName = append(skself.InColumnName, "項目2")
	skself.OutColumName = make([]string, 0)
	skself.OutColumName = append(skself.OutColumName, "項目c")
	skself.OutColumName = append(skself.OutColumName, "項目a")
	skself.OutColumName = append(skself.OutColumName, "項目c")
	skself.OutColumName = append(skself.OutColumName, "項目b")

	err = skself.Exec(stdin1, stdout1, stderr1)
	stdout1_str = stdout1.String()
	stderr1_str = stderr1.String()
	if err == nil && stdout1_str == "項目c 項目a 項目c 項目b\n3 1 3 2\n6 4 6 5\n" && stderr1_str == "" {
		t.Logf("[OK]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, stdout1_str, stderr1_str, err)
	} else {
		t.Errorf("[NG]:skself.Exec(%v,%v,%v):skself.InColumnName=%v,skself.OutColumName=%v,stdout1_str=%v,stderr1_str=%v,err=%v\n", stdin1, stdout1, stderr1, skself.InColumnName, skself.OutColumName, stdout1_str, stderr1_str, err)
	}

}
