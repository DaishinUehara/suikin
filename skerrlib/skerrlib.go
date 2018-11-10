package skerrlib

import (
	"fmt"
)

// E0001MSG Out of Index Error Message.
const E0001MSG = "E0001:%s[%d] is Out of Index. %v\n"

// E0002MSG Not Initialized Error Message.
const E0002MSG = "E0002:%s is not initialized! %v\n"

// E0003MSG Buffer flush Error Message.
const E0003MSG = "E0003:Buffer %s flush Error! %v\n"

// E0004MSG Input Field Number is not Output Field Number Error Message.
const E0004MSG = "E0004:Input filds count %d is not output fild count %d! %v\n"

// E0005MSG Argument Error Message.
const E0005MSG = "E0005:Argument Error %v! %v\n"

// E0006MSG Input Data Header Record Error Message.
const E0006MSG = "E0006:No Header Record Field %s Error! %v\n"

// E0007MSG Standard Input File Open Error Message
const E0007MSG = "E0007:Standard Input File %s Open Error!:%v %v\n"

// E0008MSG Standard Output File Open Error Message
const E0008MSG = "E0008:Standard Output File %s Open Error!:%v %v\n"

// E0009MSG Standard Error Output File Open Error Message
const E0009MSG = "E0009:Standard Error Output File %s Open Error!:%v %v\n"

// E0010MSG Input Output Column Name Format Error Message
const E0010MSG = "E0010:Input/Output Column Name Format Error:%s. %v\n"

// E0011MSG No Input Filld Name Error Message
const E0011MSG = "E0011:No Input Field Name:%s. %v\n"

// E0012MSG Scan Error Message
const E0012MSG = "E0012:Scan Error:%v. %v\n"

// E0013MSG LoggerGenerate Error Message
const E0013MSG = "E0013:Logger Generate Error:%v. %v\n"

// E0014MSG Read File Error Message
const E0014MSG = "E0014:Read File Error:%v. %v\n"

// E0015MSG Read File Error Message
const E0015MSG = "E0015:Get Absolute Path Error:%v. %v\n"

// E0016MSG YAML Mapping Error Message
const E0016MSG = "E0016:YAML Mapping Error:%v. %v\n"

// E0100MSG DateTime Format Error Message
const E0100MSG = "E0100:DateTime %s is not Format yyyy[mm[dd[hhmiss]]]:%v. %v\n"

// E0110MSG Time Format Error Message
const E0110MSG = "E0110:Time %s is not Format hhmiss:%v. %v\n"

// E9999MSG Unexpected Error Message
const E9999MSG = "E9999:Unexpected Error:%v. %v\n"

// ErrOutOfIndex スライス(配列)の範囲外へのアクセスを実施した場合
type ErrOutOfIndex struct {
	ArrayName  string
	Index      int
	StackTrace []string
}

func (e *ErrOutOfIndex) Error() string {
	return fmt.Sprintf(E0001MSG, e.ArrayName, e.Index, e.StackTrace)
}

// ErrNotInitialized 初期化されていない場合のエラー
type ErrNotInitialized struct {
	// 	PkgMethodName     string
	NoInitializedItem string
	StackTrace        []string
}

func (e *ErrNotInitialized) Error() string {
	return fmt.Sprintf(E0002MSG, e.NoInitializedItem, e.StackTrace)
}

// ErrFlushBuffer バッファのflushがエラーとなった場合
type ErrFlushBuffer struct {
	ErrorItem  string
	StackTrace []string
}

func (e *ErrFlushBuffer) Error() string {
	return fmt.Sprintf(E0003MSG, e.ErrorItem, e.StackTrace)
}

// ErrInFieldCntNotEqualOutFieldCnt 入力フィールド数と出力フィールド数が一致すべきときに一致しない場合
type ErrInFieldCntNotEqualOutFieldCnt struct {
	InFieldCount  int
	OutFieldCount int
	StackTrace    []string
}

func (e *ErrInFieldCntNotEqualOutFieldCnt) Error() string {
	return fmt.Sprintf(E0004MSG, e.InFieldCount, e.OutFieldCount, e.StackTrace)
}

// ErrArgument 入力フィールド数と出力フィールド数が一致すべきときに一致しない場合
type ErrArgument struct {
	Argv       []string
	StackTrace []string
}

func (e *ErrArgument) Error() string {
	return fmt.Sprintf(E0005MSG, e.Argv, e.StackTrace)
}

// ErrNoHeaderRecord ヘッダレコードがない場合
type ErrNoHeaderRecord struct {
	FieldName  string
	StackTrace []string
}

func (e *ErrNoHeaderRecord) Error() string {
	return fmt.Sprintf(E0006MSG, e.FieldName, e.StackTrace)
}

// ErrStdInputFileOpen 標準入力ファイルオープンエラー
type ErrStdInputFileOpen struct {
	FileName   string
	Err        error
	StackTrace []string
}

func (e *ErrStdInputFileOpen) Error() string {
	return fmt.Sprintf(E0007MSG, e.FileName, e.Err, e.StackTrace)
}

// ErrStdOutputFileOpen 標準出力ファイルオープンエラー
type ErrStdOutputFileOpen struct {
	FileName   string
	Err        error
	StackTrace []string
}

func (e *ErrStdOutputFileOpen) Error() string {
	return fmt.Sprintf(E0008MSG, e.FileName, e.Err, e.StackTrace)
}

// ErrStdErrOutputFileOpen 標準エラー出力ファイルオープンエラー
type ErrStdErrOutputFileOpen struct {
	FileName   string
	Err        error
	StackTrace []string
}

func (e *ErrStdErrOutputFileOpen) Error() string {
	return fmt.Sprintf(E0009MSG, e.FileName, e.Err, e.StackTrace)
}

// ErrInputOutputColumNameFormat 標準エラー出力ファイルオープンエラー
type ErrInputOutputColumNameFormat struct {
	ColumnName string
	StackTrace []string
}

func (e *ErrInputOutputColumNameFormat) Error() string {
	return fmt.Sprintf(E0010MSG, e.ColumnName, e.StackTrace)
}

// ErrNoInputFieldName 存在しないフィールド名が指定されたときのエラー
type ErrNoInputFieldName struct {
	FieldName  string
	StackTrace []string
}

func (e *ErrNoInputFieldName) Error() string {
	return fmt.Sprintf(E0011MSG, e.FieldName, e.StackTrace)
}

// ErrScan スキャンエラー
type ErrScan struct {
	Err        error
	StackTrace []string
}

func (e *ErrScan) Error() string {
	return fmt.Sprintf(E0012MSG, e.Err, e.StackTrace)
}

// ErrLoggerGenerate ロガー生成エラー
type ErrLoggerGenerate struct {
	Err        error
	StackTrace []string
}

func (e *ErrLoggerGenerate) Error() string {
	return fmt.Sprintf(E0013MSG, e.Err, e.StackTrace)
}

// ErrReadFile ファイルリードエラー
type ErrReadFile struct {
	Err        error
	StackTrace []string
}

func (e *ErrReadFile) Error() string {
	return fmt.Sprintf(E0014MSG, e.Err, e.StackTrace)
}

// ErrGetAbsolutePath 絶対パス取得エラー
type ErrGetAbsolutePath struct {
	Err        error
	StackTrace []string
}

func (e *ErrGetAbsolutePath) Error() string {
	return fmt.Sprintf(E0015MSG, e.Err, e.StackTrace)
}

// ErrYamlMapping YAMLマッピングエラー
type ErrYamlMapping struct {
	Err        error
	StackTrace []string
}

func (e *ErrYamlMapping) Error() string {
	return fmt.Sprintf(E0016MSG, e.Err, e.StackTrace)
}

// ErrDateTimeFormat 日時フォーマットエラー
type ErrDateTimeFormat struct {
	DateTimeStr string
	Err         error
	StackTrace  []string
}

func (e *ErrDateTimeFormat) Error() string {
	return fmt.Sprintf(E0100MSG, e.DateTimeStr, e.Err, e.StackTrace)
}

// ErrTimeFormat 時刻フォーマットエラー
type ErrTimeFormat struct {
	TimeStr    string
	Err        error
	StackTrace []string
}

func (e *ErrTimeFormat) Error() string {
	return fmt.Sprintf(E0110MSG, e.TimeStr, e.Err, e.StackTrace)
}

// ErrUnexpected 予期しないエラー
type ErrUnexpected struct {
	Err        error
	StackTrace []string
}

func (e *ErrUnexpected) Error() string {
	return fmt.Sprintf(E9999MSG, e.Err, e.StackTrace)
}
