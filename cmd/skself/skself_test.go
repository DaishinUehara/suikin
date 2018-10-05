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
	strStdOut, strStdErr, err = skstublib.SkStdStub("", argv, func(argv1 []string) { err1 := selfExec(argv1); return err1 })
	if err == nil {
		t.Logf("[OK]:main.selfExec(%v):err=%v,stdout=%v,stderr=%v\n", argv, err, strStdOut, strStdErr)
	} else {
		t.Errorf("[NG]:main.selfExec(%v):err=%v,stdout=%v,stderr=%v\n", argv, err, strStdOut, strStdErr)
	}

}
