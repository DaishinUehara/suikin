package skfilelib_test

import (
	"testing"

	"github.com/DaishinUehara/suikin/skfilelib"
)

func TestTouch(t *testing.T) {
	err := skfilelib.Touch("../tmp/test01.txt")
	if err != nil {
		t.Errorf("[NG]:skfilelib.Touch(\"../tmp/test01.txt\"):err=%v\n", err)
	}
	if skfilelib.Exists("../tmp/test01.txt") {
		t.Log("[OK]skfilelib.Exists(\"../tmp/test01.txt\")")
	} else {
		t.Error("[NG]skfilelib.Exists(\"../tmp/test01.txt\")")
	}
	skfilelib.RmFile("../tmp/test01.txt")
}
