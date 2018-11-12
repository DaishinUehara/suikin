package skerrlib

import (
	"fmt"
)

// MsgMap is Error Message Map.
var MsgMap = map[string]string{
	// E0001 Out of Index Error Mesage.
	"E0001": "E0001:%s[%d] is Out of Index. %v\n",
	// E0002 Not Initialized Error Message.
	"E0002": "E0002:%s is not initialized! %v\n",
	// E0003 Buffer flush Error Message.
	"E0003": "E0003:Buffer %s flush Error! %v\n",
	// E0004 Input Field Number is not Output Field Number Error Message.
	"E0004": "E0004:Input fields count %d is not equal output fields count %d! %v\n",
	// E0005 Argument Error Message.
	"E0005": "E0005:Argument Error %v! %v\n",
	// E0006 Input Data Header Record Error Message.
	"E0006": "E0006:No Header Record Field %s Error! %v\n",
	// E0007 Standard Input File Open Error Message.
	"E0007": "E0007:Standard Input File %s Open Error!:%v %v\n",
	// E0008 Standard Output File Open Error Message.
	"E0008": "E0008:Standard Output File %s Open Error!:%v %v\n",
	// E0009 Standard Error Output File Open Error Message.
	"E0009": "E0009:Standard Error Output File %s Open Error!:%v %v\n",
	// E0010 Input Output Column Name Format Error Message.
	"E0010": "E0010:Input/Output Column Name Format Error:%s. %v\n",
	// E0011 No Input Filld Name Error Message.
	"E0011": "E0011:No Input Field Name:%s. %v\n",
	// E0012 Scan Error Message.
	"E0012": "E0012:Scan Error:%v. %v\n",
	// E0013 LoggerGenerate Error Message.
	"E0013": "E0013:Logger Generate Error:%v. %v\n",
	// E0014 Read File Error Message.
	"E0014": "E0014:Read File Error:%v. %v\n",
	// E0015 Read File Error Message.
	"E0015": "E0015:Get Absolute Path Error:%v. %v\n",
	// E0016 YAML Mapping Error Message.
	"E0016": "E0016:YAML Mapping Error:%v. %v\n",
	// E0100 DateTime Format Error Message.
	"E0100": "E0100:DateTime %s is not Format yyyy[mm[dd[hhmiss]]]:%v. %v\n",
	// E0110 Time Format Error Message.
	"E0110": "E0110:Time %s is not Format hhmiss:%v. %v\n",
	// E9999 Unexpected Error Message.
	"E9999": "E9999:Unexpected Error:%v. %v\n",
}

// ErrOutOfIndex is Out of Index Access.
type ErrOutOfIndex struct {
	ArrayName  string
	Index      int
	StackTrace []string
}

func (e *ErrOutOfIndex) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.ArrayName, e.Index, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrOutOfIndex) MsgCd() string {
	return "E0001"
}

// ErrNotInitialized 初期化されていない場合のエラー
type ErrNotInitialized struct {
	NoInitializedItem string
	StackTrace        []string
}

func (e *ErrNotInitialized) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.NoInitializedItem, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrNotInitialized) MsgCd() string {
	return "E0002"
}

// ErrFlushBuffer バッファのflushがエラーとなった場合
type ErrFlushBuffer struct {
	ErrorItem  string
	StackTrace []string
}

func (e *ErrFlushBuffer) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.ErrorItem, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrFlushBuffer) MsgCd() string {
	return "E0003"
}

// ErrInFieldCntNotEqualOutFieldCnt 入力フィールド数と出力フィールド数が一致すべきときに一致しない場合
type ErrInFieldCntNotEqualOutFieldCnt struct {
	InFieldCount  int
	OutFieldCount int
	StackTrace    []string
}

func (e *ErrInFieldCntNotEqualOutFieldCnt) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.InFieldCount, e.OutFieldCount, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrInFieldCntNotEqualOutFieldCnt) MsgCd() string {
	return "E0004"
}

// ErrArgument 入力フィールド数と出力フィールド数が一致すべきときに一致しない場合
type ErrArgument struct {
	Argv       []string
	StackTrace []string
}

func (e *ErrArgument) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.Argv, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrArgument) MsgCd() string {
	return "E0005"
}

// ErrNoHeaderRecord ヘッダレコードがない場合
type ErrNoHeaderRecord struct {
	FieldName  string
	StackTrace []string
}

func (e *ErrNoHeaderRecord) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.FieldName, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrNoHeaderRecord) MsgCd() string {
	return "E0006"
}

// ErrStdInputFileOpen 標準入力ファイルオープンエラー
type ErrStdInputFileOpen struct {
	FileName   string
	Err        error
	StackTrace []string
}

func (e *ErrStdInputFileOpen) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.FileName, e.Err, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrStdInputFileOpen) MsgCd() string {
	return "E0007"
}

// ErrStdOutputFileOpen 標準出力ファイルオープンエラー
type ErrStdOutputFileOpen struct {
	FileName   string
	Err        error
	StackTrace []string
}

func (e *ErrStdOutputFileOpen) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.FileName, e.Err, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrStdOutputFileOpen) MsgCd() string {
	return "E0008"
}

// ErrStdErrOutputFileOpen 標準エラー出力ファイルオープンエラー
type ErrStdErrOutputFileOpen struct {
	FileName   string
	Err        error
	StackTrace []string
}

func (e *ErrStdErrOutputFileOpen) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.FileName, e.Err, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrStdErrOutputFileOpen) MsgCd() string {
	return "E0009"
}

// ErrInputOutputColumNameFormat 標準エラー出力ファイルオープンエラー
type ErrInputOutputColumNameFormat struct {
	ColumnName string
	StackTrace []string
}

func (e *ErrInputOutputColumNameFormat) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.ColumnName, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrInputOutputColumNameFormat) MsgCd() string {
	return "E0010"
}

// ErrNoInputFieldName 存在しないフィールド名が指定されたときのエラー
type ErrNoInputFieldName struct {
	FieldName  string
	StackTrace []string
}

func (e *ErrNoInputFieldName) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.FieldName, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrNoInputFieldName) MsgCd() string {
	return "E0011"
}

// ErrScan スキャンエラー
type ErrScan struct {
	Err        error
	StackTrace []string
}

func (e *ErrScan) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.Err, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrScan) MsgCd() string {
	return "E0012"
}

// ErrLoggerGenerate ロガー生成エラー
type ErrLoggerGenerate struct {
	Err        error
	StackTrace []string
}

func (e *ErrLoggerGenerate) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.Err, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrLoggerGenerate) MsgCd() string {
	return "E0013"
}

// ErrReadFile ファイルリードエラー
type ErrReadFile struct {
	Err        error
	StackTrace []string
}

func (e *ErrReadFile) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.Err, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrReadFile) MsgCd() string {
	return "E0014"
}

// ErrGetAbsolutePath 絶対パス取得エラー
type ErrGetAbsolutePath struct {
	Err        error
	StackTrace []string
}

func (e *ErrGetAbsolutePath) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.Err, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrGetAbsolutePath) MsgCd() string {
	return "E0015"
}

// ErrYamlMapping YAMLマッピングエラー
type ErrYamlMapping struct {
	Err        error
	StackTrace []string
}

func (e *ErrYamlMapping) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.Err, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrYamlMapping) MsgCd() string {
	return "E0016"
}

// ErrDateTimeFormat 日時フォーマットエラー
type ErrDateTimeFormat struct {
	DateTimeStr string
	Err         error
	StackTrace  []string
}

func (e *ErrDateTimeFormat) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.DateTimeStr, e.Err, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrDateTimeFormat) MsgCd() string {
	return "E0100"
}

// ErrTimeFormat 時刻フォーマットエラー
type ErrTimeFormat struct {
	TimeStr    string
	Err        error
	StackTrace []string
}

func (e *ErrTimeFormat) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.TimeStr, e.Err, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrTimeFormat) MsgCd() string {
	return "E0110"
}

// ErrUnexpected 予期しないエラー
type ErrUnexpected struct {
	Err        error
	StackTrace []string
}

func (e *ErrUnexpected) Error() string {
	return fmt.Sprintf(MsgMap[e.MsgCd()], e.Err, e.StackTrace)
}

// MsgCd is Get MsgCd
func (e *ErrUnexpected) MsgCd() string {
	return "E9999"
}
