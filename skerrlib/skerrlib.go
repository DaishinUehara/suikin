package skerrlib

import "fmt"

const E0001_MSG = "E0001:%s[%d] is Out of Index.\n"
const E0002_MSG = "E0002:%s is not initialized!\n"
const E0003_MSG = "E0003:Buffer %s flush Error!\n"
const E0004_MSG = "E0004:Input filds count %d is not output fild count %d!\n"

// スライス(配列)の範囲外へのアクセスを実施した場合
type ErrOutOfIndex struct {
	ArrayName string
	Index     int
}

func (e ErrOutOfIndex) Error() string {
	return fmt.Sprintf(E0001_MSG, e.ArrayName, e.Index)
}

// 初期化されていない場合のエラー
type ErrNotInitialized struct {
	NoInitializedItem string
}

func (e ErrNotInitialized) Error() string {
	return fmt.Sprintf(E0002_MSG, e.NoInitializedItem)
}

// バッファのflushがエラーとなった場合
type ErrFlushBuffer struct {
	ErrorItem string
}

func (e ErrFlushBuffer) Error() string {
	return fmt.Sprintf(E0003_MSG, e.ErrorItem)
}

// 入力フィールド数と出力フィールド数が一致すべきときに一致しない場合
type ErrInFieldCnt_NE_OutFieldCnt struct {
	InFieldCount  int
	OutFieldCount int
}

func (e ErrInFieldCnt_NE_OutFieldCnt) Error() string {
	return fmt.Sprintf(E0004_MSG, e.InFieldCount, e.OutFieldCount)
}
