package skselflib

import (
	"bufio"
	"fmt"
	"io"

	"github.com/DaishinUehara/suikin/skcmnlib"
	"github.com/DaishinUehara/suikin/skerrlib"
)

// Exec はselfを実行する
func Exec(stdin io.Reader, stdout io.Writer, stderr io.Writer, incolumnname []string, outcolumnname []string) (err error) {

	// ここから入力チェック
	if stdin == nil {
		//		err = fmt.Errorf("stdin io.Reader not initialized!")
		return skerrlib.ErrNotInitialized{NoInitializedItem: "stdin io.Reader"}
	}

	if stdout == nil {
		//		err = fmt.Errorf("stdout io.Writer not initialized!")
		return skerrlib.ErrNotInitialized{NoInitializedItem: "stdout io.Writer"}
	}

	if stderr == nil {
		//		err = fmt.Errorf("stderr io.Writer not initialized!")
		return skerrlib.ErrNotInitialized{NoInitializedItem: "stderr io.Writer"}
	}

	incolumnlen := len(incolumnname)
	outcolumnlen := len(outcolumnname)

	if incolumnlen != outcolumnlen {
		// 入力と出力が1対1で対応していない場合、エラーとする
		return skerrlib.ErrInFieldCnt_NE_OutFieldCnt{InFieldCount: incolumnlen, OutFieldCount: outcolumnlen}
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
		return err
	}

	// 1行目をセパレートする
	var headerFields []string
	headerFields, err = skcmnlib.SeparateField(line0)
	if err != nil {
		return err
	}

	// fieldをどの順番で出力するかのインデックスを作成する
	var fieldIndex []int
	fieldIndex, err = skcmnlib.GetFieldIndexArray(headerFields, incolumnname)
	if err != nil {
		return err
	}

	// 1行目(ヘッダ)の出力
	headerstr := skcmnlib.ConnectFields(outcolumnname, " ")
	if len(headerstr) > 0 {
		fmt.Fprintln(stdoutBuffer, headerstr)
	}

	var selfields []string
	for scanner.Scan() {
		// レコードを読み取りフィールドに分割
		fields, err1 := skcmnlib.SeparateField(scanner.Text())
		if err1 != nil {
			return err1
		}

		// 分割されたフィールドから出力する文字列の順番に配列に格納
		selfields, err = skcmnlib.SortByIndex(fields, fieldIndex)
		if err != nil {
			return err
		}
		fmt.Fprintln(stdoutBuffer, skcmnlib.ConnectFields(selfields, " "))
	}
	if err = scanner.Err(); err != nil {
		return err
	}

	// 結果をflushする。
	err = stdoutBuffer.Flush()
	if err != nil {
		fmt.Fprintf(stderrBuffer, "Stdout Flush error: %v\n", err)
		return skerrlib.ErrFlushBuffer{ErrorItem: "stdoutBuffer"}
	}
	err = stderrBuffer.Flush()
	if err != nil {
		return skerrlib.ErrFlushBuffer{ErrorItem: "stderrBuffer"}
	}

	return nil
}
