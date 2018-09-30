package skerrlib

import "fmt"

const E0001_MSG = "E0001:%s[%d] is Out of Index.\n"
const E0002_MSG = "E0002:%s is not initialized!\n"

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
