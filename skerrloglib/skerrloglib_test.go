package skerrloglib_test

import (
	"testing"

	"github.com/DaishinUehara/suikin/skerrloglib"
)

func TestErrOutOfIndex(t *testing.T) {
	var skerr = skerrloglib.ErrOutOfIndex{ArrayName: "test", Index: 1, StackTrace: []string{"abc", "def"}}
	skerr.LogOutput()
}
