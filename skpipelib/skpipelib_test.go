package skpipelib_test

// TODO 諸々

import (
	"bufio"
	"bytes"
	"io"
	"testing"

	"github.com/DaishinUehara/suikin/skerrlib"
	"github.com/DaishinUehara/suikin/skselflib"

	"github.com/DaishinUehara/suikin/skpipelib"
)

// Exe1 テスト用のスタブ
type Exe1 struct {
}

func (ex1 *Exe1) Exec(io.Reader, io.Writer, io.Writer) error {
	return nil
}

// Exe2 テスト用のスタブ
type Exe2 struct {
}

func (ex2 *Exe2) Exec(io.Reader, io.Writer, io.Writer) error {
	return nil
}

// Exe3 テスト用のスタブ
type Exe3 struct {
}

func (ex3 *Exe3) Exec(io.Reader, io.Writer, io.Writer) error {
	return nil
}

func TestAddExec(t *testing.T) {
	exe1 := new(Exe1)
	pe1 := new(skpipelib.SkMulti)
	infield1 := make([]string, 0, 5)
	outfield1 := make([]string, 0, 5)
	pe1.AddExec(exe1)
	stdin1 := bytes.NewBufferString("テスト")
	stdout1 := new(bytes.Buffer)
	stderr1 := new(bytes.Buffer)

	errAr1, err1 := pe1.MultiExec(stdin1, stdout1, stderr1)
	if 1 == len(errAr1) &&
		errAr1[0] == nil &&
		err1 == nil &&
		stdout1.String() == "" &&
		stderr1.String() == "" {
		t.Logf("[OK]:skselflib.MultiExec(%v,%v,%v):infield1=%v,outfield1=%v,errAr1=%v,len(errAr1)=%d,err1=%v\n", stdin1, stdout1, stderr1, infield1, outfield1, errAr1, len(errAr1), err1)
	} else {
		t.Errorf("[NG]:skselflib.MultiExec(%v,%v,%v):infield1=%v,outfield1=%v,errAr1=%v,len(errAr1)=%d,err1=%v\n", stdin1, stdout1, stderr1, infield1, outfield1, errAr1, len(errAr1), err1)
	}

	// var exe2 Exe2
	var (
		pe2       skpipelib.SkMulti
		stdin2    io.Reader
		stdout2   *bytes.Buffer
		stderr2   *bytes.Buffer
		infield2  []string
		outfield2 []string
	)
	errAr2, err2 := pe2.MultiExec(stdin2, stdout2, stderr2)
	if err2 != nil {
		switch err2.(type) {
		case skerrlib.ErrNotInitialized:
			t.Logf("[OK]:skselflib.MultiExec(%v,%v,%v):infield2=%v,outfield2=%v,errAr2=%v,len(errAr2)=%d,err2=%v\n", stdin2, stdout2, stderr2, infield2, outfield2, errAr2, len(errAr2), err2)
		default:
			t.Errorf("[NG]:skselflib.MultiExec(%v,%v,%v):infield2=%v,outfield2=%v,errAr2=%v,len(errAr2)=%d,err2=%v\n", stdin2, stdout2, stderr2, infield2, outfield2, errAr2, len(errAr2), err2)
		}
	} else {
		t.Errorf("[NG]:skselflib.MultiExec(%v,%v,%v):infield2=%v,outfield2=%v,errAr2=%v,len(errAr2)=%d,err2=%v\n", stdin2, stdout2, stderr2, infield2, outfield2, errAr2, len(errAr2), err2)
	}

	var (
		exe3        *skselflib.SkSelf  // メソッド実行する構造体
		pe3         *skpipelib.SkMulti // 実行パイプ
		stdin3      io.Reader
		stdout3     *bytes.Buffer
		stderr3     *bytes.Buffer
		infield3    []string
		outfield3   []string
		line3       string
		scannerout3 *bufio.Scanner
		scannererr3 *bufio.Scanner
	)

	pe3 = new(skpipelib.SkMulti) // 実行パイプ
	infield3 = make([]string, 0, 5)
	infield3 = append(infield3, "項目1")
	infield3 = append(infield3, "項目3")
	outfield3 = make([]string, 0, 5)
	outfield3 = append(outfield3, "a")
	outfield3 = append(outfield3, "c")

	exe3 = new(skselflib.SkSelf) // メソッド実行する構造体
	exe3.InColumnName = infield3
	exe3.OutColumName = outfield3

	pe3.AddExec(exe3) // パイプに実行構造体を追加
	stdin3 = bytes.NewBufferString("項目1 項目2 項目3\n1 2 3\n4 5 6\n7 8 9")
	stdout3 = new(bytes.Buffer)
	stderr3 = new(bytes.Buffer)

	errAr3, err3 := pe3.MultiExec(stdin3, stdout3, stderr3)
	strStdOut3 := stdout3.String()
	strStdErr3 := stderr3.String()
	if err3 == nil && strStdOut3 == "a c\n1 3\n4 6\n7 9\n" && strStdErr3 == "" {
		t.Logf("[OK]:skselflib.MultiExec(%v,%v,%v):infield3=%v,outfield3=%v,errAr3=%v,len(errAr3)=%d,err3=%v\n", stdin3, stdout3, stderr3, infield3, outfield3, errAr3, len(errAr3), err3)

		scannerout3 = bufio.NewScanner(stdout3)
		t.Logf("stdout3=")
		for scannerout3.Scan() {
			line3 = scannerout3.Text()
			t.Logf("%s\n", line3)
		}

		scannererr3 = bufio.NewScanner(stderr3)
		t.Logf("stderr3=")
		for scannererr3.Scan() {
			line3 = scannererr3.Text()
			t.Logf("%s\n", line3)
		}

	} else {
		t.Errorf("[NG]:skselflib.MultiExec(%v,%v,%v):infield3=%v,outfield3=%v,errAr3=%v,len(errAr3)=%d,err3=%v\n", stdin3, stdout3, stderr3, infield3, outfield3, errAr3, len(errAr3), err3)
	}

	var (
		exe40       *skselflib.SkSelf
		exe41       *skselflib.SkSelf
		pe4         *skpipelib.SkMulti
		stdin4      io.Reader
		stdout4     *bytes.Buffer
		stderr4     *bytes.Buffer
		infield40   []string
		outfield40  []string
		infield41   []string
		outfield41  []string
		line4       string
		scannerout4 *bufio.Scanner
		scannererr4 *bufio.Scanner
	)

	pe4 = new(skpipelib.SkMulti) // 実行パイプ
	infield40 = make([]string, 0, 5)
	infield40 = append(infield40, "項目1")
	infield40 = append(infield40, "項目3")
	outfield40 = make([]string, 0, 5)
	outfield40 = append(outfield40, "a")
	outfield40 = append(outfield40, "c")

	infield41 = make([]string, 0, 5)
	infield41 = append(infield41, "c")
	infield41 = append(infield41, "a")
	infield41 = append(infield41, "c")
	outfield41 = make([]string, 0, 5)
	outfield41 = append(outfield41, "C")
	outfield41 = append(outfield41, "A")
	outfield41 = append(outfield41, "C")

	exe40 = new(skselflib.SkSelf) // メソッド実行する構造体1
	exe40.InColumnName = infield40
	exe40.OutColumName = outfield40
	pe4.AddExec(exe40) // パイプに実行構造体を追加

	exe41 = new(skselflib.SkSelf) // メソッド実行する構造体2
	exe41.InColumnName = infield41
	exe41.OutColumName = outfield41
	pe4.AddExec(exe41) // パイプに実行構造体を追加

	stdin4 = bytes.NewBufferString("項目1 項目2 項目3\n1 2 3\n4 5 6\n7 8 9")
	stdout4 = new(bytes.Buffer)
	stderr4 = new(bytes.Buffer)

	t.Logf("[CALL]:skselflib.MultiExec(%v,%v,%v)\n", stdin4, stdout4, stderr4)
	t.Logf("[CALL]:exe40.InColumnName=%v,exe40.OutColumName=%v\n", exe40.InColumnName, exe40.OutColumName)
	t.Logf("[CALL]:exe41.InColumnName=%v,ouexe41.OutColumNametfield41=%v\n", exe41.InColumnName, exe41.OutColumName)
	errAr4, err4 := pe4.MultiExec(stdin4, stdout4, stderr4)
	strStdOut4 := stdout4.String()
	strStdErr4 := stderr4.String()
	t.Logf("[RETURN]:errAr4=%v,len(errAr4)=%d,err4=%v\n", errAr4, len(errAr4), err4)
	t.Logf("[RETURN]:strStdOut4=%v,strStdErr4=%v\n", strStdOut4, strStdErr4)

	if err4 == nil && strStdOut4 == "C A C\n3 1 3\n6 4 6\n9 7 9\n" && strStdErr4 == "" {
		t.Logf("[OK]:skselflib.MultiExec(%v,%v,%v):errAr4=%v,len(errAr4)=%d,err4=%v\n", stdin4, stdout4, stderr4, errAr4, len(errAr4), err4)
		t.Logf("[OK]:infield40=%v,outfield40=%v\n", infield40, outfield40)
		t.Logf("[OK]:infield41=%v,outfield41=%v\n", infield41, outfield41)

		scannerout4 = bufio.NewScanner(stdout4)
		t.Logf("stdout4=")
		for scannerout4.Scan() {
			line3 = scannerout4.Text()
			t.Logf("%s\n", line3)
		}

		scannererr4 = bufio.NewScanner(stderr4)
		t.Logf("stderr4=")
		for scannererr4.Scan() {
			line4 = scannererr4.Text()
			t.Logf("%s\n", line4)
		}

	} else {
		t.Errorf("[NG]:skselflib.MultiExec(%v,%v,%v):errAr4=%v,len(errAr4)=%d,err4=%v\n", stdin4, stdout4, stderr4, errAr4, len(errAr4), err4)
		t.Errorf("[NG]:infield40=%v,outfield40=%v\n", infield40, outfield40)
		t.Errorf("[NG]:infield41=%v,outfield41=%v\n", infield41, outfield41)

		scannerout4 = bufio.NewScanner(stdout4)
		t.Logf("stdout4=")
		for scannerout4.Scan() {
			line3 = scannerout4.Text()
			t.Logf("%s\n", line3)
		}

		scannererr4 = bufio.NewScanner(stderr4)
		t.Logf("stderr4=")
		for scannererr4.Scan() {
			line4 = scannererr4.Text()
			t.Logf("%s\n", line4)
		}

	}
}

func TestPipeExec(t *testing.T) {
}
