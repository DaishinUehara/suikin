package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/DaishinUehara/suikin/skcmnlib"
	"github.com/DaishinUehara/suikin/skerrlib"
	"github.com/DaishinUehara/suikin/skselflib"
)

func printUsage() {
	_usage := "self <infile|-> <outfile|-> <errfile|-> [incolumname1[,newcolumnname1]] [incolumname2[,newcolumnname2]] ...\n"
	_version := "Thu Dec  8 10:46:51 JST 2018\n"
	_code := "Suikin 2.4 Go (windows/go-lang 1.9/UTF-8)\n"
	_description := []string{
		"\n",
	}

	os.Stderr.Write([]byte("Usage    :" + _usage))
	os.Stderr.Write([]byte("Version  :" + _version))
	os.Stderr.Write([]byte("           " + _code))
	os.Stderr.Write([]byte("Description\n"))
	for _, desc := range _description {
		os.Stderr.Write([]byte(desc))
	}
}

/*
select column

読み込むファイルの1行はbufio.Scannerの制限により64Kbyte以内でなければならない。
selc 入力ファイル|- 出力ファイル|- エラー出力ファイル|- [入力ファイル出力項目1[,出力ファイル出力項目名1]] ...

例1)
標準入力から読み取り、結果を標準出力に返す。

*/
func main() {
	err := selfExec(os.Args)
	if err != nil {
		switch err.(type) {
		case skerrlib.ErrArgument:
			printUsage()
		default:
			fmt.Printf("Got unexpected error!\n")
		}
		os.Exit(1)
	}
	os.Exit(0)
}

func selfExec(argv []string) error {
	arglen := len(argv)
	if arglen < 4 {
		return skerrlib.ErrArgument{Argv: argv}
	}

	var selectColumnName []string
	selectColumnName = make([]string, 0, 20)

	for i, str := range argv {
		if i >= 4 {
			// 最初と最後以外
			selectColumnName = append(selectColumnName, str)
		}
	}

	var err error
	var stdin io.Reader
	var stdout io.Writer
	var stderr io.Writer

	var infile *os.File
	var outfile *os.File
	var errfile *os.File

	if argv[1] != "-" {
		// ファイルを開く場合
		infile, err = os.Open(argv[1])
		if err != nil {
			return skerrlib.ErrStdInputFileOpen{FileName: argv[1], Err: err}
			// fmt.Fprintf(os.Stderr, "Input File %s open error: %v\n", argv[1], err)
		}
		defer infile.Close() // 関数return時に閉じる
		stdin = bufio.NewReader(infile)
	} else {
		// 標準入力の場合
		stdin = bufio.NewReader(os.Stdin)
	}

	if argv[2] != "-" {
		// ファイルを開く場合
		outfile, err = os.Open(argv[2])
		if err != nil {
			return skerrlib.ErrStdOutputFileOpen{FileName: argv[2], Err: err}
			// fmt.Fprintf(os.Stderr, "Output File %s open error: %v\n", argv[2], err)
		}
		defer outfile.Close() // 関数return時に閉じる
		stdout = bufio.NewWriter(outfile)
	} else {
		// 標準入力の場合
		stdout = bufio.NewWriter(os.Stdout)
	}

	if argv[3] != "-" {
		// ファイルを開く場合
		errfile, err = os.Open(os.Args[3])
		if err != nil {
			return skerrlib.ErrStdErrOutputFileOpen{FileName: argv[3], Err: err}
			// fmt.Fprintf(os.Stderr, "Standard Error File %s open error: %v\n", argv[3], err)
		}
		defer errfile.Close() // 関数return時に閉じる
		stderr = bufio.NewWriter(errfile)
	} else {
		// 標準入力の場合
		stderr = bufio.NewWriter(os.Stderr)
	}

	var incolumnname []string
	var outcolumnname []string

	incolumnname, outcolumnname, err = skcmnlib.CammaDivide(selectColumnName)
	if err != nil {
		switch err.(type) {
		case skerrlib.ErrInputOutputColumNameFormat:
			// fmt.Fprintf(stderr, "Select Column Arguments Error selectColumnName=%v:err=%v\n", selectColumnName, err)
			// return err
			return err
		default:
			return skerrlib.ErrUnexpected{Err: err}
		}
	}

	err = skselflib.Exec(stdin, stdout, stderr, incolumnname, outcolumnname)
	if err != nil {
		fmt.Fprintf(stderr, "Select Field Processing Error file=%v:err=%v\n", argv[3], err)
	}
	return err
}
