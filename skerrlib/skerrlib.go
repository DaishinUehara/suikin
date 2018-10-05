package skerrlib

import "fmt"

// E0001MSG Out of Index Error Message.
const E0001MSG = "E0001:%s[%d] is Out of Index.\n"

// E0002MSG Not Initialized Error Message.
const E0002MSG = "E0002:%s is not initialized!\n"

// E0003MSG Buffer flush Error Message.
const E0003MSG = "E0003:Buffer %s flush Error!\n"

// E0004MSG Input Field Number is not Output Field Number Error Message.
const E0004MSG = "E0004:Input filds count %d is not output fild count %d!\n"

// E0005MSG Argument Error Message.
const E0005MSG = "E0005:Argument Error %v!\n"

// E0006MSG Input Data Header Record Error Message.
const E0006MSG = "E0006:No Header Record Error!\n"

// ErrOutOfIndex スライス(配列)の範囲外へのアクセスを実施した場合
type ErrOutOfIndex struct {
	ArrayName string
	Index     int
}

func (e ErrOutOfIndex) Error() string {
	return fmt.Sprintf(E0001MSG, e.ArrayName, e.Index)
}

// ErrNotInitialized 初期化されていない場合のエラー
type ErrNotInitialized struct {
	NoInitializedItem string
}

func (e ErrNotInitialized) Error() string {
	return fmt.Sprintf(E0002MSG, e.NoInitializedItem)
}

// ErrFlushBuffer バッファのflushがエラーとなった場合
type ErrFlushBuffer struct {
	ErrorItem string
}

func (e ErrFlushBuffer) Error() string {
	return fmt.Sprintf(E0003MSG, e.ErrorItem)
}

// ErrInFieldCntNotEqualOutFieldCnt 入力フィールド数と出力フィールド数が一致すべきときに一致しない場合
type ErrInFieldCntNotEqualOutFieldCnt struct {
	InFieldCount  int
	OutFieldCount int
}

func (e ErrInFieldCntNotEqualOutFieldCnt) Error() string {
	return fmt.Sprintf(E0004MSG, e.InFieldCount, e.OutFieldCount)
}

// ErrArgument 入力フィールド数と出力フィールド数が一致すべきときに一致しない場合
type ErrArgument struct {
	Argv []string
}

func (e ErrArgument) Error() string {
	return fmt.Sprintf(E0005MSG, e.Argv)
}

// ErrNoHeaderRecord ヘッダレコードがない場合
type ErrNoHeaderRecord struct {
}

func (e ErrNoHeaderRecord) Error() string {
	return fmt.Sprintf(E0006MSG)
}
