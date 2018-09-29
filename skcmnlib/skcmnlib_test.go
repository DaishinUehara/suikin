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

	if ans1 != "" {
		t.Errorf("[NG]:skcmnlib.ConnectFields(%v,%v):ans1=%s\n", instrArr1, spc1, ans1)
	} else {
		t.Logf("[OK]:skcmnlib.ConnectFields(%v,%v):ans1=%s\n", instrArr1, spc1, ans1)
	}

	var instrArr2 []string
	var spc2 string
	ans2 := skcmnlib.ConnectFields(instrArr2, spc2)

	if ans1 != "" {
		t.Errorf("[NG]:skcmnlib.ConnectFields(%v,%v):ans2=%s\n", instrArr2, spc2, ans2)
	} else {
		t.Logf("[OK]:skcmnlib.ConnectFields(%v,%v):ans2=%s\n", instrArr2, spc2, ans2)
	}

	instrArr3 := make([]string, 0, 5)
	spc3 := " "
	instrArr3 = append(instrArr3, "てすてすと")

	ans3 := skcmnlib.ConnectFields(instrArr3, spc3)
	if ans3 != "てすてすと" {
		t.Errorf("[NG]:skcmnlib.ConnectFields(%v,%v):ans3=%s\n", instrArr3, spc3, ans3)
	}
	t.Logf("[OK]:skcmnlib.ConnectFields(%v,%v):ans3=%s\n", instrArr3, spc3, ans3)

	instrArr4 := make([]string, 0, 5)
	spc4 := " "
	instrArr4 = append(instrArr4, "てすてすと1")
	instrArr4 = append(instrArr4, "てすてすと2")

	ans4 := skcmnlib.ConnectFields(instrArr4, spc4)
	if ans4 != "てすてすと1 てすてすと2" {
		t.Errorf("[NG]:skcmnlib.ConnectFields(%v,%v):ans4=%s\n", instrArr4, spc4, ans4)
	}
	t.Logf("[OK]:skcmnlib.ConnectFields(%v,%v):ans4=%s\n", instrArr4, spc4, ans4)

	instrArr5 := make([]string, 0, 5)
	spc5 := ","
	instrArr5 = append(instrArr5, "てすてすと1")
	instrArr5 = append(instrArr5, "てすてすと2")

	ans5 := skcmnlib.ConnectFields(instrArr5, spc5)
	if ans5 != "てすてすと1,てすてすと2" {
		t.Errorf("[NG]:skcmnlib.ConnectFields(%v,%v):ans5=%s\n", instrArr5, spc5, ans5)
	}
	t.Logf("[OK]:skcmnlib.ConnectFields(%v,%v):ans5=%s\n", instrArr5, spc5, ans5)

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
		t.Logf("[OK]:skcmnlib.GetFieldIndexArray(%v,%v):idx1=%v:err1=%v\n", header1, select1, idx1, err1)
	} else {
		t.Errorf("[NG]:skcmnlib.GetFieldIndexArray(%v,%v):idx1=%v:err1=%v\n", header1, select1, idx1, err1)
	}

	header2 := make([]string, 0, 5)
	header2 = append(header2, "項目1")
	header2 = append(header2, "項目2")
	header2 = append(header2, "項目3")

	select2 := make([]string, 0, 5)

	idx2, err2 := skcmnlib.GetFieldIndexArray(header2, select2)
	if len(idx2) == 0 {
		t.Logf("[OK]:skcmnlib.GetFieldIndexArray(%v,%v):idx2=%v:err2=%v\n", header2, select2, idx2, err2)
	} else {
		t.Errorf("[NG]:skcmnlib.GetFieldIndexArray(%v,%v):idx2=%v:err2=%v\n", header2, select2, idx2, err2)
	}

	header3 := make([]string, 0, 5)
	header3 = append(header3, "項目1")
	header3 = append(header3, "項目2")
	header3 = append(header3, "項目3")

	select3 := make([]string, 0, 5)
	select3 = append(select3, "項目4")

	idx3, err3 := skcmnlib.GetFieldIndexArray(header3, select3)
	if err3 == nil || idx3[0] != -1 {
		t.Errorf("[NG]:skcmnlib.GetFieldIndexArray(%v,%v):idx3=%v:err3=%v\n", header3, select3, idx3, err3)
	} else {
		t.Logf("[OK]:skcmnlib.GetFieldIndexArray(%v,%v):idx3=%v:err3=%v\n", header3, select3, idx3, err3)
	}

	header4 := make([]string, 0, 5)
	header4 = append(header4, "項目1")
	header4 = append(header4, "項目2")
	header4 = append(header4, "項目3")
	header4 = append(header4, "項目4")

	select4 := make([]string, 0, 5)
	select4 = append(select4, "項目4")
	select4 = append(select4, "項目1")

	idx4, err4 := skcmnlib.GetFieldIndexArray(header4, select4)
	if err4 == nil && idx4[0] == 3 && idx4[1] == 0 && len(idx4) == 2 {
		t.Logf("[OK]:skcmnlib.GetFieldIndexArray(%v,%v):idx4=%v:err4=%v\n", header4, select4, idx4, err4)
	} else {
		t.Errorf("[NG]:skcmnlib.GetFieldIndexArray(%v,%v):idx4=%v:err4=%v\n", header4, select4, idx4, err4)
	}

}

func TestSeparateField(t *testing.T) {

	line1 := ""
	st1, err1 := skcmnlib.SeparateField(line1)
	if len(st1) == 0 && err1 == nil {
		t.Logf("[OK]:skcmnlib.SeparateField(%v):st1=%v:err1=%v\n", line1, st1, err1)
	} else {
		t.Errorf("[NG]:skcmnlib.SeparateField(%v):st1=%v:err1=%v\n", line1, st1, err1)
	}

	line2 := "項目1"
	st2, err2 := skcmnlib.SeparateField(line2)
	if len(st2) == 1 && err2 == nil && st2[0] == "項目1" {
		t.Logf("[OK]:skcmnlib.SeparateField(%v):st2=%v:err2=%v\n", line2, st2, err2)
	} else {
		t.Errorf("[NG]:skcmnlib.SeparateField(%v):st2=%v:err2=%v\n", line2, st2, err2)
	}

	line3 := "項目1 項目2"
	st3, err3 := skcmnlib.SeparateField(line3)
	if len(st3) == 2 && err3 == nil && st3[0] == "項目1" && st3[1] == "項目2" {
		t.Logf("[OK]:skcmnlib.SeparateField(%v):st3=%v:err3=%v\n", line3, st3, err3)
	} else {
		t.Errorf("[NG]:skcmnlib.SeparateField(%v):st3=%v:err3=%v\n", line3, st3, err3)
	}

	line4 := "項目1  項目2"
	st4, err4 := skcmnlib.SeparateField(line4)
	if len(st4) == 2 && err4 == nil && st4[0] == "項目1" && st4[1] == "項目2" {
		t.Logf("[OK]:skcmnlib.SeparateField(%v):st4=%v:err4=%v\n", line4, st4, err4)
	} else {
		t.Errorf("[NG]:skcmnlib.SeparateField(%v):st4=%v:err4=%v\n", line4, st4, err4)
	}

	line5 := "項目1 項目2 "
	st5, err5 := skcmnlib.SeparateField(line5)
	if len(st5) == 2 && err5 == nil && st5[0] == "項目1" && st5[1] == "項目2" {
		t.Logf("[OK]:skcmnlib.SeparateField(%v):st5=%v:err5=%v\n", line5, st5, err5)
	} else {
		t.Errorf("[NG]:skcmnlib.SeparateField(%v):st5=%v:err5=%v\n", line5, st5, err5)
	}

	line6 := "項目1 項目2\t 項目3\t"
	st6, err6 := skcmnlib.SeparateField(line6)
	if len(st6) == 3 && err6 == nil && st6[0] == "項目1" && st6[1] == "項目2" && st6[2] == "項目3" {
		t.Logf("[OK]:skcmnlib.SeparateField(%v):st6=%v:err6=%v\n", line6, st6, err6)
	} else {
		t.Errorf("[NG]:skcmnlib.SeparateField(%v):st6=%v:err6=%v\n", line6, st6, err6)
	}

	line7 := "項目1\\t項目2\t 項目3\t"
	st7, err7 := skcmnlib.SeparateField(line7)
	if len(st7) == 2 && err7 == nil && st7[0] == "項目1\t項目2" && st7[1] == "項目3" {
		t.Logf("[OK]:skcmnlib.SeparateField(%v):st7=%v:err7=%v\n", line7, st7, err7)
	} else {
		t.Errorf("[NG]:skcmnlib.SeparateField(%v):st7=%v:err7=%v\n", line7, st7, err7)
	}
}
