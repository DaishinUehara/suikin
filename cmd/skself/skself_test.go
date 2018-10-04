package main

import (
	"testing"
)

func TestSelExec(t *testing.T) {
	var err error
	var argv []string

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
}
