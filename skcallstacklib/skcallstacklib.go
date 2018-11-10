package skcallstacklib

import (
	"bytes"
	"fmt"
	"regexp"
	"runtime"
)

// SkCallInfo コール情報を格納する構造体
type SkCallInfo struct {
	PkgName  string
	FuncName string
	FileName string
	FileLine int
}

// パッケージ名と関数名を分離する正規表現
var regPkgFunc = regexp.MustCompile(`^(\S.+)\.(\S.+)$`)

// SkCallStack コール情報をダンプしコールスタックとして格納する。
func SkCallStack(i int) (callstack []*SkCallInfo) {
	for ; ; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			break
		}

		fn := runtime.FuncForPC(pc)
		fileName, fileLine := fn.FileLine(pc)
		_fn := regPkgFunc.FindStringSubmatch(fn.Name())
		callstack = append(callstack, &SkCallInfo{
			PkgName:  _fn[1],
			FuncName: _fn[2],
			FileName: fileName,
			FileLine: fileLine,
		})
	}
	return
}

// PrintCallStack is reutrn CallStackString
func PrintCallStack() []string {
	callstack := SkCallStack(2)
	var ret = bytes.NewBuffer(make([]byte, 0, 100))
	var stack []string
	stack = make([]string, 1)
	for _, callinfo := range callstack {
		ret.WriteString("filename=")
		ret.WriteString(callinfo.FileName)
		ret.WriteString(",line=")
		ret.WriteString(fmt.Sprintf("%d", callinfo.FileLine))
		ret.WriteString(",method=")
		ret.WriteString(callinfo.PkgName)
		ret.WriteString(".")
		ret.WriteString(callinfo.FuncName)
		stack = append(stack, ret.String())
		ret.Reset()
	}
	return stack
}
