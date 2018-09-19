package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
 select field
*/
func main() {
	var fileName string
	var selectFieldName []string
	selectFieldName = make([]string, 0, 20)
	for i, str := range os.Args {
		if i == 0 {

		} else if i == len(os.Args)-1 {
			// 最後の引数
			fileName = str
		} else {
			// 最初と最後以外
			selectFieldName = append(selectFieldName, str)
		}
	}
	var f *os.File
	var err error
	if fileName != "-" {
		// ファイルを開く
		f, err = os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", fileName, err)
		}
		defer f.Close()
	}
	// 関数return時に閉じる
	var scanner *bufio.Scanner
	if fileName != "-" {
		scanner = bufio.NewScanner(f)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	self(fileName, selectFieldName)
}

func self(stdin io.Reader, stdout io.Writer, selectField []string) (err error) {

	var line0 string // 1 行目

	if scanner.Scan() {
		line0 = scanner.Text()
	} else {
		return err
	}

	// 1行目をセパレートする
	var headerFields []string
	headerFields, err = separateField(line0)
	if err != nil {
		return err
	}

	// fieldをどの順番で出力するかのインデックスを作成する
	var fieldIndex []int
	fieldIndex, err = getFieldIndexArray(headerFields, selectField)
	if err != nil {
		return err
	}

	var wtr = bufio.NewWriter(os.Stdout)
	var selfields []string
	for scanner.Scan() {
		// レコードを読み取りフィールドに分割
		fields, err1 := separateField(scanner.Text())
		if err1 != nil {
			return err1
		}

		// 分割されたフィールドから結合する文字列の順番に配列に格納
		selfields, err = selectFields(fields, fieldIndex)
		if err != nil {
			return err
		}
		fmt.Fprintln(wtr, connectFields(selfields, ' '))
	}

	if err = scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", inputFilePath, err)
		return err
	}

	// 結果を標準出力にflushする。
	err = wtr.Flush()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Stdout Flush error: %v\n", err)
		return err
	}

	return nil

}

// fieldsをconnectCharで結合し返す
func connectFields(fields []string, connectChar byte) string {
	// メモリ確保
	buf := make([]byte, 0, 100)
	for i, field := range fields {
		if i > 0 {
			buf = append(buf, ' ')
		}
		buf = append(buf, field...)
	}
	return string(buf)
}

// フィールド名のインデックスを返す。
func getFieldIndexArray(headerFields []string, selectFieldNames []string) ([]int, error) {
	fieldIndex := make([]int, 0, len(selectFieldNames))
	for i, selectFieldName := range selectFieldNames {
		l, err := getFieldIndex(headerFields, selectFieldName)
		fieldIndex[i] = l
		if err != nil {
			return fieldIndex, err
		}
	}
	return fieldIndex, nil
}

// ヘッダから文字列をリニアサーチしインデックスを返す。
func getFieldIndex(headerFields []string, fieldName string) (fieldIndex int, err error) {
	fieldIndex = -1
	for i, headerFieldName := range headerFields {
		if headerFieldName == fieldName {
			fieldIndex = i
			break
		}
	}
	if fieldIndex == -1 {
		//err = errors.New(fmt.Sprintf("No FieldName: %s", fieldName))
		err = fmt.Errorf("No FieldName: %s", fieldName)
		return fieldIndex, err
	}
	return fieldIndex, nil
}

// 区切られたフィールドをfieldIndexに並び替えselfield配列に格納して返す。
func selectFields(fields []string, fieldIndex []int) (selfields []string, err error) {
	selfields = make([]string, 0, 50)
	for _, fi := range fieldIndex {
		selfields = append(selfields, fields[fi])
	}
	return selfields, err
}

// stringをスペースもしくはタブで区切ったstring配列に格納します。
// ただし'\ 'は区切り文字ではなくデータのスペースとして扱われます。
// '\0'は長さ0のブランク文字列として扱います。
func separateField(line string) (st []string, err error) {
	st = make([]string, 0, 50) // makeで初期capacityを指定して予めメモリを確保

	var isBackSlash = false
	var isSpace = false
	var rword []rune
	rword = make([]rune, 0, 255)
	for _, unc := range line {
		if isBackSlash {
			// 前の文字がバックスラッシュだった場合
			if unc == rune('\\') {
				if isSpace {
					st = append(st, string(rword))
					rword = make([]rune, 0, 255)
				}
				rword = append(rword, unc)
			} else if unc == rune(' ') || unc == rune('t') {
				if isSpace {
					st = append(st, string(rword))
					rword = make([]rune, 0, 255)
				}
				rword = append(rword, unc)
			} else if unc == rune('0') {
				if isSpace {
					st = append(st, string(rword))
					rword = make([]rune, 0, 255)
				}
				rword = append(rword, rune(0)) // \0 はrune(0)とする
			}
			isBackSlash = false
			isSpace = false
		} else {
			// 前の文字がバックスラッシュ以外の場合
			if unc == rune('\\') {
				// バックスラッシュの場合
				isBackSlash = true
			} else if unc != rune(' ') && unc != rune('\t') {
				// 通常文字の場合
				if isSpace {
					st = append(st, string(rword))
					rword = make([]rune, 0, 255)
				}
				rword = append(rword, unc)
				isBackSlash = false
				isSpace = false
			} else {
				// スペース文字の場合
				isSpace = true
			}
		}
	}
	return st, err
}
