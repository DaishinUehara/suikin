package skcmnlib_test

import (
	"testing"

	"github.com/DaishinUehara/suikin/skcmnlib"
)

func TestCammaDivide(t *testing.T) {
	selectColumnName1 := make([]string, 0, 5)
	selectColumnName1 = append(selectColumnName1, "工場")
	selectColumnName1 = append(selectColumnName1, "売上")
	selectColumnName1 = append(selectColumnName1, "原価")
	selectColumnName1 = append(selectColumnName1, "売上総利益")

	incolumnname1, outcolumnname1, err1 := skcmnlib.CammaDivide(selectColumnName1)

	if err1 != nil {
		t.Errorf("skcmnlib.CammaDivide:想定できないエラー: %v", err1)
	}

	for i, ans1 := range selectColumnName1 {
		if incolumnname1[i] != ans1 {
			t.Errorf("skcmnlib.CammaDivide:incolumnname1[%d]=[戻値:%s]:[想定:%s]\n", i, incolumnname1[i], ans1)
			t.Errorf("skcmnlib.CammaDivide:outcolumnname1[%d]=[戻値:%s]:[想定:%s]\n", i, outcolumnname1[i], ans1)
		}
	}

	selectColumnName2 := make([]string, 0, 5)
	selectColumnName2 = append(selectColumnName2, "工場,工場1")
	selectColumnName2 = append(selectColumnName2, "売上,売上(合計)")
	selectColumnName2 = append(selectColumnName2, "原価,原価(合計)")
	selectColumnName2 = append(selectColumnName2, "粗利,粗利(合計)")

	incolumnname2, outcolumnname2, err2 := skcmnlib.CammaDivide(selectColumnName2)

	if err2 != nil {
		t.Errorf("skcmnlib.CammaDivide:想定できないエラー: %v", err2)
	}

	inans2 := make([]string, 0, 5)
	inans2 = append(inans2, "工場")
	inans2 = append(inans2, "売上")
	inans2 = append(inans2, "原価")
	inans2 = append(inans2, "粗利")

	outans2 := make([]string, 0, 5)
	outans2 = append(outans2, "工場1")
	outans2 = append(outans2, "売上(合計)")
	outans2 = append(outans2, "原価(合計)")
	outans2 = append(outans2, "粗利(合計)")

	for i, input := range selectColumnName2 {
		if incolumnname2[i] != inans2[i] {
			t.Errorf("skcmnlib.CammaDivide:incolumnname1[%d]=[入力値:%s]:[戻値:%s]:[想定:%s]\n", i, input, incolumnname2[i], inans2)
		}
		if outcolumnname2[i] != outans2[i] {
			t.Errorf("skcmnlib.CammaDivide:outcolumnname1[%d]=[入力値:%s]:[戻値:%s]:[想定:%s]\n", i, input, outcolumnname2[i], outans2)
		}
	}
}

func TestConnectFields(t *testing.T) {

	instrArr1 := make([]string, 0, 5)
	spc1 := " "
	ans1 := skcmnlib.ConnectFields(instrArr1, spc1)
	if len(instrArr1) == 0 {
		t.Log("instrArr1 size is 0")
	} else {
		for i, instr := range instrArr1 {
			t.Logf("[OK]:skcmnlib.ConnectFields:入力1 instrArr1[%d]=%s\n", i, instr)
		}
	}
	t.Logf("[OK]:skcmnlib.ConnectFields:入力1 spc1=%v\n", spc1)

	if ans1 != "" {
		t.Errorf("[NG]:skcmnlib.ConnectFields:戻値1: %s\n", ans1)
	} else {
		t.Logf("[OK]:skcmnlib.ConnectFields:戻値1: %s\n", ans1)
	}

	var instrArr2 []string
	var spc2 string
	ans2 := skcmnlib.ConnectFields(instrArr2, spc2)
	for i, instr := range instrArr2 {
		t.Logf("[OK]:skcmnlib.ConnectFields:入力2 instArr[%d]=%s\n", i, instr)
	}
	t.Logf("[OK]:skcmnlib.ConnectFields:入力2 spc1=%v\n", spc2)

	if ans1 != "" {
		t.Errorf("[NG]:skcmnlib.ConnectFields:戻値2: %s\n", ans2)
	} else {
		t.Logf("[OK]:skcmnlib.ConnectFields:戻値2: %s\n", ans2)
	}

	instrArr3 := make([]string, 0, 5)
	spc3 := " "
	instrArr3 = append(instrArr3, "てすてすと")

	ans3 := skcmnlib.ConnectFields(instrArr3, spc3)
	if ans3 != "てすてすと" {
		t.Errorf("[NG]:skcmnlib.ConnectFields:戻値3: %s\n", ans3)
	}
	t.Logf("[OK]:skcmnlib.ConnectFields:戻値3: %s\n", ans3)

	instrArr4 := make([]string, 0, 5)
	spc4 := " "
	instrArr4 = append(instrArr4, "てすてすと1")
	instrArr4 = append(instrArr4, "てすてすと2")

	ans4 := skcmnlib.ConnectFields(instrArr4, spc4)
	if ans4 != "てすてすと1 てすてすと2" {
		t.Errorf("[NG]:skcmnlib.ConnectFields:戻値4: %s\n", ans4)
	}
	t.Logf("[OK]:skcmnlib.ConnectFields:戻値4: %s\n", ans4)

	instrArr5 := make([]string, 0, 5)
	spc5 := ","
	instrArr5 = append(instrArr5, "てすてすと1")
	instrArr5 = append(instrArr5, "てすてすと2")

	ans5 := skcmnlib.ConnectFields(instrArr5, spc5)
	if ans5 != "てすてすと1,てすてすと2" {
		t.Errorf("[NG]:skcmnlib.ConnectFields:戻値5: %s\n", ans5)
	}
	t.Logf("[OK]:skcmnlib.ConnectFields:戻値5: %s\n", ans5)

}

func TestGetFieldIndex(t *testing.T) {

	input1 := make([]string, 0, 5)
	ans1, err1 := skcmnlib.GetFieldIndex(input1, "")
	if err1 != nil || ans1 != -1 {
		t.Logf("[OK]:skcmnlib.GetFieldIndex(%v,\"\"):ans1=%d:err1=%v\n", input1, ans1, err1)
	} else {
		t.Errorf("[OK]:skcmnlib.GetFieldIndex(%v,\"\"):ans1=%d:err1=%v\n", input1, ans1, err1)
	}

	input2 := make([]string, 0, 5)
	input2 = append(input2, "あいう")
	ans2, err2 := skcmnlib.GetFieldIndex(input2, "")
	if err2 != nil || ans2 != -1 {
		t.Logf("[OK]:skcmnlib.GetFieldIndex(%v,\"あいう\"):ans2=%d:err2=%v\n", input2, ans2, err2)
	} else {
		t.Errorf("[OK]:skcmnlib.GetFieldIndex(%v,\"あいう\"):ans2=%d:err2=%v\n", input2, ans2, err2)
	}

	input3 := make([]string, 0, 5)
	input3 = append(input3, "あいう")
	input3 = append(input3, "えおか")
	ans3, err3 := skcmnlib.GetFieldIndex(input3, "えおか")
	if err3 == nil || ans3 == 1 {
		t.Logf("[OK]:skcmnlib.GetFieldIndex(%v,\"えおか\"):ans3=%d:err3=%v\n", input3, ans3, err3)
	} else {
		t.Errorf("[NG]:skcmnlib.GetFieldIndex(%v,\"えおか\"):ans3=%d:err3=%v\n", input3, ans3, err3)
	}

	input4 := make([]string, 0, 5)
	input4 = append(input4, "あいう")
	input4 = append(input4, "えおか")
	input4 = append(input4, "あいう")
	ans4, err4 := skcmnlib.GetFieldIndex(input4, "あいう")
	if err4 == nil || ans4 == 0 {
		t.Logf("[OK]:skcmnlib.GetFieldIndex(%v,\"あいう\"):ans4=%d:err4=%v\n", input4, ans4, err4)
	} else {
		t.Errorf("[NG]:skcmnlib.GetFieldIndex(%v,\"あいう\"):ans4=%d:err4=%v\n", input4, ans4, err4)
	}
}

func TestGetFieldIndexArray(t *testing.T) {
	header1 := make([]string, 0, 5)
	select1 := make([]string, 0, 5)

	idx1, err1 := skcmnlib.GetFieldIndexArray(header1, select1)
	if len(idx1) == 0 {
		t.Logf("[OK]:skcmnlib.GetFieldIndex:header1=%v:select1=%v:idx1=%v:err1=%v\n", header1, select1, idx1, err1)
	} else {
		t.Errorf("[NG]:skcmnlib.GetFieldIndex:header1=%v:select1=%v:idx1=%v:err1=%v\n", header1, select1, idx1, err1)
	}
}
