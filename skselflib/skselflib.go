package skselflib

import (
	"bufio"
	"fmt"
	"io"

	"github.com/DaishinUehara/suikin/skcmnlib"
	"github.com/DaishinUehara/suikin/skerrlib"
)

// SkSelf is
type SkSelf struct {
	InColumnName []string
	OutColumName []string
}

// Exec はselfを実行する
func (ss *SkSelf) Exec(stdin io.Reader, stdout io.Writer, stderr io.Writer) (err error) {

	// ここから入力チェック
	if stdin == nil {
		return skerrlib.ErrNotInitialized{NoInitializedItem: "stdin io.Reader", StackTrace: skerrlib.PrintCallStack()}
	}

	if stdout == nil {
		return skerrlib.ErrNotInitialized{NoInitializedItem: "stdout io.Writer", StackTrace: skerrlib.PrintCallStack()}
	}

	if stderr == nil {
		return skerrlib.ErrNotInitialized{NoInitializedItem: "stderr io.Writer", StackTrace: skerrlib.PrintCallStack()}
	}

	incolumnlen := len(ss.InColumnName)
	outcolumnlen := len(ss.OutColumName)

	if incolumnlen != outcolumnlen {
		// 入力と出力が1対1で対応していない場合、エラーとする
		return skerrlib.ErrInFieldCntNotEqualOutFieldCnt{InFieldCount: incolumnlen, OutFieldCount: outcolumnlen, StackTrace: skerrlib.PrintCallStack()}
	}

	if incolumnlen == 0 && outcolumnlen == 0 {
		// 入力フィールドも出力フィールドもない場合、なにもしない
		return nil
	}
	// ここまで入力チェック

	scanner := bufio.NewScanner(stdin)
	stdoutBuffer := bufio.NewWriter(stdout)
	stderrBuffer := bufio.NewWriter(stderr)

	var line0 string // 1 行目
	if scanner.Scan() {
		line0 = scanner.Text()
	} else {
		// 入力フィールドと出力フィールドが指定されているにも関わらず
		// 1行目(ヘッダ)が読めない場合エラーとする
		return skerrlib.ErrNoHeaderRecord{FieldName: ss.InColumnName[0], StackTrace: skerrlib.PrintCallStack()}
	}

	// 1行目をセパレートする
	var headerFields []string
	headerFields, err = skcmnlib.SeparateField(line0) // 本来ならここでerrは返ってこない。
	if err != nil {
		return skerrlib.ErrUnexpected{Err: err, StackTrace: skerrlib.PrintCallStack()}
	}

	// fieldをどの順番で出力するかのインデックスを作成する
	var fieldIndex []int
	fieldIndex, err = skcmnlib.GetFieldIndexArray(headerFields, ss.InColumnName)
	if err != nil {
		switch err.(type) {
		case skerrlib.ErrNoInputFieldName:
			return err
		default:
			return skerrlib.ErrUnexpected{Err: err, StackTrace: skerrlib.PrintCallStack()}
		}
	}

	// 1行目(ヘッダ)の出力
	headerstr := skcmnlib.ConnectFields(ss.OutColumName, " ")
	if len(headerstr) > 0 {
		fmt.Fprintln(stdoutBuffer, headerstr)
	}

	var selfields []string
	for scanner.Scan() {
		// レコードを読み取りフィールドに分割
		fields, err1 := skcmnlib.SeparateField(scanner.Text()) // 本来ならここでerrは返ってこない。
		if err1 != nil {
			return skerrlib.ErrUnexpected{Err: err1, StackTrace: skerrlib.PrintCallStack()}
		}

		// 分割されたフィールドから出力する文字列の順番に配列に格納
		selfields, err = skcmnlib.SortByIndex(fields, fieldIndex)
		if err != nil {
			switch err.(type) {
			case skerrlib.ErrOutOfIndex:
				return err
			default:
				return skerrlib.ErrUnexpected{Err: err, StackTrace: skerrlib.PrintCallStack()}
			}
		}
		fmt.Fprintln(stdoutBuffer, skcmnlib.ConnectFields(selfields, " "))
	}

	if err = scanner.Err(); err != nil {
		return skerrlib.ErrScan{Err: err, StackTrace: skerrlib.PrintCallStack()}
	}

	// 結果をflushする。
	err = stdoutBuffer.Flush()
	if err != nil {
		fmt.Fprintf(stderrBuffer, "Stdout Flush error: %v\n", err)
		return skerrlib.ErrFlushBuffer{ErrorItem: "stdoutBuffer", StackTrace: skerrlib.PrintCallStack()}
	}
	err = stderrBuffer.Flush()
	if err != nil {
		return skerrlib.ErrFlushBuffer{ErrorItem: "stderrBuffer", StackTrace: skerrlib.PrintCallStack()}
	}

	return nil
}
