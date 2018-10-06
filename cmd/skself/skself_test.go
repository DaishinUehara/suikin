package main

import (
	"testing"

	"github.com/DaishinUehara/suikin/skstublib"
)

func TestSelExec(t *testing.T) {
	var err error
	var argv []string
	var strStdOut string
	var strStdErr string
	var strStdIn string

	argv = make([]string, 0, 4)
	err = selfExec(argv)
	if err != nil {
		t.Logf("[OK]:main.selfExec(%v):err=%v\n", argv, err)
	} else {
		t.Errorf("[NG]:main.selfExec(%v):err=%v\n", argv, err)
	}

	argv = make([]string, 0, 4)
	argv = append(argv, "skself")
	err = selfExec(argv)
	if err != nil {
		t.Logf("[OK]:main.selfExec(%v):err=%v\n", argv, err)
	} else {
		t.Errorf("[NG]:main.selfExec(%v):err=%v\n", argv, err)
	}

	argv = make([]string, 0, 4)
	argv = append(argv, "skself")
	argv = append(argv, "-")
	err = selfExec(argv)
	if err != nil {
		t.Logf("[OK]:main.selfExec(%v):err=%v\n", argv, err)
	} else {
		t.Errorf("[NG]:main.selfExec(%v):err=%v\n", argv, err)
	}

	argv = make([]string, 0, 4)
	argv = append(argv, "skself")
	argv = append(argv, "-")
	argv = append(argv, "-")
	err = selfExec(argv)
	if err != nil {
		t.Logf("[OK]:main.selfExec(%v):err=%v\n", argv, err)
	} else {
		t.Errorf("[NG]:main.selfExec(%v):err=%v\n", argv, err)
	}

	argv = make([]string, 0, 4)
	argv = append(argv, "skself")
	argv = append(argv, "-")
	argv = append(argv, "-")
	argv = append(argv, "-")
	strStdIn = ""
	strStdOut, strStdErr, err = skstublib.SkStdStub(strStdIn, argv, func(argv1 []string) error { err1 := selfExec(argv1); return err1 })
	if err == nil && strStdOut == "" && strStdErr == "" {
		t.Logf("[OK]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	} else {
		t.Errorf("[NG]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	}

	argv = make([]string, 0, 4)
	argv = append(argv, "skself")
	argv = append(argv, "-")
	argv = append(argv, "-")
	argv = append(argv, "-")
	argv = append(argv, "項目1")
	strStdIn = ""
	strStdOut, strStdErr, err = skstublib.SkStdStub(strStdIn, argv, func(argv1 []string) error { err1 := selfExec(argv1); return err1 })
	if err != nil && strStdOut == "" && strStdErr == "" {
		t.Logf("[OK]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	} else {
		t.Errorf("[NG]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	}

	argv = make([]string, 0, 4)
	argv = append(argv, "skself")
	argv = append(argv, "-")
	argv = append(argv, "-")
	argv = append(argv, "-")
	argv = append(argv, "項目1")
	strStdIn = "項目2"
	strStdOut, strStdErr, err = skstublib.SkStdStub(strStdIn, argv, func(argv1 []string) error { err1 := selfExec(argv1); return err1 })
	if err != nil && strStdOut == "" && strStdErr == "" {
		t.Logf("[OK]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	} else {
		t.Errorf("[NG]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	}

	argv = make([]string, 0, 4)
	argv = append(argv, "skself")
	argv = append(argv, "-")
	argv = append(argv, "-")
	argv = append(argv, "-")
	strStdIn = "項目2"
	strStdOut, strStdErr, err = skstublib.SkStdStub(strStdIn, argv, func(argv1 []string) error { err1 := selfExec(argv1); return err1 })
	if err == nil && strStdOut == "" && strStdErr == "" {
		t.Logf("[OK]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	} else {
		t.Errorf("[NG]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	}

	argv = make([]string, 0, 4)
	argv = append(argv, "skself")
	argv = append(argv, "-")
	argv = append(argv, "-")
	argv = append(argv, "-")
	argv = append(argv, "項目1")
	strStdIn = "項目1"
	strStdOut, strStdErr, err = skstublib.SkStdStub(strStdIn, argv, func(argv1 []string) error { err1 := selfExec(argv1); return err1 })
	if err == nil && strStdOut == "項目1\n" && strStdErr == "" {
		t.Logf("[OK]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	} else {
		t.Errorf("[NG]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	}

	argv = make([]string, 0, 4)
	argv = append(argv, "skself")
	argv = append(argv, "-")
	argv = append(argv, "-")
	argv = append(argv, "-")
	argv = append(argv, "項目3,c")
	argv = append(argv, "項目1")
	strStdIn = "項目1 項目2 項目3\n1 2 3\n4 5 6"
	strStdOut, strStdErr, err = skstublib.SkStdStub(strStdIn, argv, func(argv1 []string) error { err1 := selfExec(argv1); return err1 })
	if err == nil && strStdOut == "c 項目1\n3 1\n6 4\n" && strStdErr == "" {
		t.Logf("[OK]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	} else {
		t.Errorf("[NG]:main.selfExec(%v):err=%v,stdin=%v,stdout=%v,stderr=%v\n", argv, err, strStdIn, strStdOut, strStdErr)
	}
}
