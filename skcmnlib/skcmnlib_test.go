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

func TestSortByIndex(t *testing.T) {
	inarr := make([]string, 0, 10)
	idxarr := make([]int, 0, 10)

	srtarr, err := skcmnlib.SortByIndex(inarr, idxarr)
	if len(srtarr) == 0 && err == nil {
		t.Logf("[OK]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	} else {
		t.Errorf("[NG]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	}

	inarr = make([]string, 0, 10)
	idxarr = make([]int, 0, 10)
	idxarr = append(idxarr, 0)

	srtarr, err = skcmnlib.SortByIndex(inarr, idxarr)
	if len(srtarr) == 0 && err != nil {
		t.Logf("[OK]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	} else {
		t.Errorf("[NG]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	}

	inarr = make([]string, 0, 10)
	idxarr = make([]int, 0, 10)
	inarr = append(inarr, "")

	srtarr, err = skcmnlib.SortByIndex(inarr, idxarr)
	if len(srtarr) == 0 && err == nil {
		t.Logf("[OK]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	} else {
		t.Errorf("[NG]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	}

	inarr = make([]string, 0, 10)
	idxarr = make([]int, 0, 10)
	inarr = append(inarr, "")
	idxarr = append(idxarr, 0)

	srtarr, err = skcmnlib.SortByIndex(inarr, idxarr)
	if len(srtarr) == 1 && srtarr[0] == "" && err == nil {
		t.Logf("[OK]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	} else {
		t.Errorf("[NG]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	}

	inarr = make([]string, 0, 10)
	idxarr = make([]int, 0, 10)
	inarr = append(inarr, "a")
	inarr = append(inarr, "b")
	idxarr = append(idxarr, 0)

	srtarr, err = skcmnlib.SortByIndex(inarr, idxarr)
	if len(srtarr) == 1 && srtarr[0] == "a" && err == nil {
		t.Logf("[OK]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	} else {
		t.Errorf("[NG]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	}

	inarr = make([]string, 0, 10)
	idxarr = make([]int, 0, 10)
	inarr = append(inarr, "a")
	inarr = append(inarr, "b")
	idxarr = append(idxarr, 1)

	srtarr, err = skcmnlib.SortByIndex(inarr, idxarr)
	if len(srtarr) == 1 && srtarr[0] == "b" && err == nil {
		t.Logf("[OK]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	} else {
		t.Errorf("[NG]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	}

	inarr = make([]string, 0, 10)
	idxarr = make([]int, 0, 10)
	inarr = append(inarr, "a")
	inarr = append(inarr, "b")
	idxarr = append(idxarr, 1)
	idxarr = append(idxarr, 0)
	idxarr = append(idxarr, 1)

	srtarr, err = skcmnlib.SortByIndex(inarr, idxarr)
	if len(srtarr) == 3 && srtarr[0] == "b" && srtarr[1] == "a" && srtarr[2] == "b" && err == nil {
		t.Logf("[OK]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	} else {
		t.Errorf("[NG]:skcmnlib.SortByIndex(%v,%v):srtarr=%v:err=%v\n", inarr, idxarr, srtarr, err)
	}

}

func TestDateToUnixSec(t *testing.T) {
	str1 := "19700101000000"
	sec1, err1 := skcmnlib.DateToUnixSec(str1) // YYYYMMDDhhmmss
	if sec1 == 0 && err1 == nil {
		t.Logf("[OK]:skcmnlib.DateToUnixSec(%v):sec1=%v:err1=%v\n", str1, sec1, err1)
	} else {
		t.Errorf("[NG]:skcmnlib.DateToUnixSec(%v):sec1=%v:err1=%v\n", str1, sec1, err1)
	}

	str2 := "19700101000001"
	sec2, err2 := skcmnlib.DateToUnixSec(str2) // YYYYMMDDhhmmss
	if sec2 == 1 && err2 == nil {
		t.Logf("[OK]:skcmnlib.DateToUnixSec(%v):sec2=%v:err2=%v\n", str2, sec2, err2)
	} else {
		t.Errorf("[NG]:skcmnlib.DateToUnixSec(%v):sec2=%v:err2=%v\n", str2, sec2, err2)
	}

	str3 := "19700101000101"
	sec3, err3 := skcmnlib.DateToUnixSec(str3) // YYYYMMDDhhmmss
	if sec3 == 61 && err3 == nil {
		t.Logf("[OK]:skcmnlib.DateToUnixSec(%v):sec3=%v:err3=%v\n", str3, sec3, err3)
	} else {
		t.Errorf("[NG]:skcmnlib.DateToUnixSec(%v):sec3=%v:err3=%v\n", str3, sec3, err3)
	}

	str4 := "19700101010101"
	sec4, err4 := skcmnlib.DateToUnixSec(str4) // YYYYMMDDhhmmss
	if sec4 == 3661 && err4 == nil {
		t.Logf("[OK]:skcmnlib.DateToUnixSec(%v):sec4=%v:err4=%v\n", str4, sec4, err4)
	} else {
		t.Errorf("[NG]:skcmnlib.DateToUnixSec(%v):sec4=%v:err4=%v\n", str4, sec4, err4)
	}

	str5 := "197001010000"
	sec5, err5 := skcmnlib.DateToUnixSec(str5) // YYYYMMDDhhmm
	if sec5 == 0 && err5 == nil {
		t.Logf("[OK]:skcmnlib.DateToUnixSec(%v):sec5=%v:err5=%v\n", str5, sec5, err5)
	} else {
		t.Errorf("[NG]:skcmnlib.DateToUnixSec(%v):sec5=%v:err5=%v\n", str5, sec5, err5)
	}

	str6 := "197001010001"
	sec6, err6 := skcmnlib.DateToUnixSec(str6) // YYYYMMDDhhmm
	if sec6 == 60 && err6 == nil {
		t.Logf("[OK]:skcmnlib.DateToUnixSec(%v):sec6=%v:err6=%v\n", str6, sec6, err6)
	} else {
		t.Errorf("[NG]:skcmnlib.DateToUnixSec(%v):sec6=%v:err6=%v\n", str6, sec6, err6)
	}

	str7 := "1970010101"
	sec7, err7 := skcmnlib.DateToUnixSec(str7) // YYYYMMDDhh
	if sec7 == 3600 && err7 == nil {
		t.Logf("[OK]:skcmnlib.DateToUnixSec(%v):sec7=%v:err7=%v\n", str7, sec7, err7)
	} else {
		t.Errorf("[NG]:skcmnlib.DateToUnixSec(%v):sec7=%v:err7=%v\n", str7, sec7, err7)
	}

	str8 := "19700102"
	sec8, err8 := skcmnlib.DateToUnixSec(str8) // YYYYMMDD
	if sec8 == 86400 && err8 == nil {
		t.Logf("[OK]:skcmnlib.DateToUnixSec(%v):sec8=%v:err8=%v\n", str8, sec8, err8)
	} else {
		t.Errorf("[NG]:skcmnlib.DateToUnixSec(%v):sec8=%v:err8=%v\n", str8, sec8, err8)
	}

	str9 := "197002"
	sec9, err9 := skcmnlib.DateToUnixSec(str9) // YYYYMM
	if sec9 == 2678400 && err9 == nil {
		t.Logf("[OK]:skcmnlib.DateToUnixSec(%v):sec9=%v:err9=%v\n", str9, sec9, err9)
	} else {
		t.Errorf("[NG]:skcmnlib.DateToUnixSec(%v):sec9=%v:err9=%v\n", str9, sec9, err9)
	}

}

func TestTimeToUnixSec(t *testing.T) {
	str1 := "000000"
	sec1, err1 := skcmnlib.TimeToUnixSec(str1) // YYYYMMDDhhmmss
	if sec1 == 0 && err1 == nil {
		t.Logf("[OK]:skcmnlib.TimeToUnixSec(%v):sec1=%v:err1=%v\n", str1, sec1, err1)
	} else {
		t.Errorf("[NG]:skcmnlib.TimeToUnixSec(%v):sec1=%v:err1=%v\n", str1, sec1, err1)
	}
}
