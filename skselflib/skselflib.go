package skselflib

import (
	"bufio"
	"fmt"
	"io"
	"github.com/DaishinUehara/suikin/skcmnlib"
)

// Exec はselfを実行する
func Exec(stdin io.Reader, stdout io.Writer, stderr io.Writer, incolumnname []string, outcolumnname []string) (err error) {

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
	fmt.Fprintln(stdoutBuffer, skcmnlib.ConnectFields(outcolumnname, ' '))

	var selfields []string
	for scanner.Scan() {
		// レコードを読み取りフィールドに分割
		fields, err1 := skcmnlib.SeparateField(scanner.Text())
		if err1 != nil {
			return err1
		}

		// 分割されたフィールドから出力する文字列の順番に配列に格納
		selfields, err = selectFields(fields, fieldIndex)
		if err != nil {
			return err
		}
		fmt.Fprintln(stdoutBuffer, skcmnlib.ConnectFields(selfields, ' '))
	}
	if err = scanner.Err(); err != nil {
		return err
	}

	// 結果をflushする。
	stdoutBuffer.Flush()
	if err != nil {
		fmt.Fprintf(stderrBuffer, "Stdout Flush error: %v\n", err)
		return err
	}
	stderrBuffer.Flush()
	if err != nil {
		fmt.Fprintf(stderrBuffer, "Stderr Flush error: %v\n", err)
		return err
	}

	return nil
}

// 区切られたフィールドをfieldIndexに並び替えselfield配列に格納して返す。
func selectFields(fields []string, fieldIndex []int) (selfields []string, err error) {
	selfields = make([]string, 0, 50)
	for _, fi := range fieldIndex {
		selfields = append(selfields, fields[fi])
	}
	return selfields, err
}
