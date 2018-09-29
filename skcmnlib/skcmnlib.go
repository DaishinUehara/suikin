package skcmnlib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// CammaDivide カンマで区切って2つの配列に分けて格納する
func CammaDivide(selectColumnName []string) ([]string, []string, error) {
	var incolumnname []string
	var outcolumnname []string
	var err error
	var columnarray []string

	incolumnname = make([]string, 0, 20)
	outcolumnname = make([]string, 0, 20)

	for _, column := range selectColumnName {
		columnarray = strings.Split(column, ",")
		l := len(columnarray)
		if l == 1 {
			incolumnname = append(incolumnname, columnarray[0])
			outcolumnname = append(outcolumnname, columnarray[0])
		} else if l == 2 {
			incolumnname = append(incolumnname, columnarray[0])
			outcolumnname = append(outcolumnname, columnarray[1])
		} else {
			err = fmt.Errorf("Input/Output Column Name Format Error: %s", column)
		}
	}
	return incolumnname, outcolumnname, err
}

// ConnectFields fieldsをconnectStrで結合し返す。
func ConnectFields(fields []string, connectStr string) string {
	// メモリ確保
	buf := make([]byte, 0, 100)
	for i, field := range fields {
		if i > 0 {
			buf = append(buf, connectStr...)
		}
		buf = append(buf, field...)
	}
	return string(buf)
}

// GetFieldIndex 配列headerFieldsからfieldNameと完全一致する文字列をリニアサーチしインデックスを返す。
// エラーの場合fieldIndexには-1が戻り、errが設定される。
// fieldNameがheaderFieldsの複数に一致する場合には、一番若い番号のインデックスが戻る
func GetFieldIndex(headerFields []string, fieldName string) (fieldIndex int, err error) {
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

// GetFieldIndexArray headerFieldsからselectFieldNamesに格納した文字列配列に一致するヘッダ文字列を検索し位置配列を返す。
func GetFieldIndexArray(headerFields []string, selectFieldNames []string) ([]int, error) {
	fieldIndex := make([]int, 0, len(selectFieldNames))
	for _, selectFieldName := range selectFieldNames {
		l, err := GetFieldIndex(headerFields, selectFieldName)
		fieldIndex = append(fieldIndex, l)
		if err != nil {
			return fieldIndex, err
		}
	}
	return fieldIndex, nil
}

// SeparateField stringをスペースもしくはタブで区切ったstring配列に格納します。
// '\ 'は区切り文字ではなくデータのスペースとして扱われます。
// '\0'は長さ0のブランク文字列として扱います。
func SeparateField(line string) (st []string, err error) {
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
	if len(rword) > 0 {
		st = append(st, string(rword))
	}
	return st, err
}

///////////////////////////////

// DateToUnixSec yyyymmddをunix時間に直す
func DateToUnixSec(timestr string) (int64, error) {
	l := len(timestr)
	layoutBase := "20060102030405"
	layout := layoutBase[0:l]
	t1, e1 := time.Parse(layout, timestr)
	if e1 != nil {
		return 0, fmt.Errorf("usp.DateToUnixSec: %v", e1)
	}
	return t1.Unix(), nil
}

// RowToTabCol 行をtabによりカラム区切りに変更する
func RowToTabCol(row string) []string {

	f := func(c rune) bool {
		if c == '\t' {
			return true
		}
		return false
	}

	s2 := strings.FieldsFunc(row, f)
	if 0 == len(s2) {
		ret := []string{row}
		return ret
	}
	return s2
}

// StdinRowToCol 標準入力から1行読み込み配列に格納
func StdinRowToCol() []string {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		r1 := strings.Fields(s.Text())
		return r1
	}
	return nil
}

// min iとjのうち最小値の値を返す
func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

// Column カラムはユニコード文字の配列
type Column []rune

// Record 1行は複数のカラム
type Record []Column

// Records 複数行
type Records []Record

// Len ソートインターフェース用のメソッド追加
func (r Records) Len() int {
	return len(r)
}

// Swap ソートインターフェース用のメソッド追加
func (r Records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Less ソートインターフェース用の大小比較追加
func (r Records) Less(i, j int) bool {
	// 比較対象となるレコードをセット
	reci := r[i]
	recj := r[j]

	// ▼項目数を求める
	wordcnt := min(len(reci), len(recj))
	// ▲項目数を求める

	for l := 0; l < wordcnt; l++ {
		// 比較対象となる項目をセット
		wordi := reci[l]
		wordj := recj[l]
		wl := WordComp(wordi, wordj)
		if wl < 0 {
			return true
		}
		if wl > 0 {
			return false
		}
		l++
	}
	if len(reci) == wordcnt {
		return true
	}
	return false
}

// WordComp w1,w2をルーンで比較しw1のほうが大きい場合には1をw2のほうが大きい場合には-1を等しい場合には0を返す
func WordComp(w1 []rune, w2 []rune) int {
	rlen := min(len(w1), len(w2))
	for i := 0; i < rlen; i++ {
		if w1[i] < w2[i] {
			return -1
		}
		if w1[i] > w2[i] {
			return 1
		}
	}
	if len(w1) == len(w2) {
		return 0
	}
	if len(w1) < len(w2) {
		return -1
	}
	return 1
}
