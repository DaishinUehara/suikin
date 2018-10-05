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
const E0006MSG = "E0006:No Header Record Field %s Error!\n"

// E0007MSG Standard Input File Open Error Message
const E0007MSG = "E0007:Standard Input File %s Open Error!:%v\n"

// E0008MSG Standard Output File Open Error Message
const E0008MSG = "E0008:Standard Output File %s Open Error!:%v\n"

// E0009MSG Standard Error Output File Open Error Message
const E0009MSG = "E0009:Standard Error Output File %s Open Error!:%v\n"

// E0010MSG Input Output Column Name Format Error Message
const E0010MSG = "E0010:Input/Output Column Name Format Error:%s.\n"

// E0011MSG No Input Filld Name Error Message
const E0011MSG = "E0011:No Input Field Name:%s.\n"

// E9999MSG Unexpected Error Message
const E9999MSG = "E9999:Unexpected Error:%v.\n"

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
	FieldName string
}

func (e ErrNoHeaderRecord) Error() string {
	return fmt.Sprintf(E0006MSG, e.FieldName)
}

// ErrStdInputFileOpen 標準入力ファイルオープンエラー
type ErrStdInputFileOpen struct {
	FileName string
	Err      error
}

func (e ErrStdInputFileOpen) Error() string {
	return fmt.Sprintf(E0007MSG, e.FileName, e.Err)
}

// ErrStdOutputFileOpen 標準出力ファイルオープンエラー
type ErrStdOutputFileOpen struct {
	FileName string
	Err      error
}

func (e ErrStdOutputFileOpen) Error() string {
	return fmt.Sprintf(E0008MSG, e.FileName, e.Err)
}

// ErrStdErrOutputFileOpen 標準エラー出力ファイルオープンエラー
type ErrStdErrOutputFileOpen struct {
	FileName string
	Err      error
}

func (e ErrStdErrOutputFileOpen) Error() string {
	return fmt.Sprintf(E0009MSG, e.FileName, e.Err)
}

// ErrInputOutputColumNameFormat 標準エラー出力ファイルオープンエラー
type ErrInputOutputColumNameFormat struct {
	ColumnName string
}

func (e ErrInputOutputColumNameFormat) Error() string {
	return fmt.Sprintf(E0010MSG, e.ColumnName)
}

// ErrNoInputFieldName 存在しないフィールド名が指定されたときのエラー
type ErrNoInputFieldName struct {
	FieldName string
}

func (e ErrNoInputFieldName) Error() string {
	return fmt.Sprintf(E0011MSG, e.FieldName)
}

// ErrUnexpected 予期しないエラー
type ErrUnexpected struct {
	Err error
}

func (e ErrUnexpected) Error() string {
	return fmt.Sprintf(E9999MSG, e.Err)
}
