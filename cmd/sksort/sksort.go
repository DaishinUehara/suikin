package main

import (
	"os"

	"github.com/DaishinUehara/suikin/skerrlib"
)

func printUsage() {
	_usage := "sksort <infile|-> <outfile|-> <errfile|-> [sortcolumname1[,datatype1[,sorttype1]]] [sortcolumname2[,datatype2[,sorttype2]]] ...\n"
	_version := "Thu Dec  8 10:46:51 JST 2018\n"
	_code := "Suikin 0.0 Go (windows/go-lang 1.11.2/UTF-8)\n"
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
sksort column

読み込むファイルの1行はbufio.Scannerの制限により64Kbyte以内でなければならない。
sksort 入力ファイル|- 出力ファイル|- エラー出力ファイル|- [ソート項目1[,データ種別1[,ソート順1(昇順、降順)]]]  [ソート項目2[,データ種別2[,ソート順2(昇順、降順)]]] ...

*/

func main() {
	err := sortExec(os.Args)
	if err != nil {
		switch err.(type) {
		case *skerrlib.ErrArgument:
			// TODO エラー処理追加
			printUsage()
		case *skerrlib.ErrUnexpected:
		}
		os.Exit(1)
	}
	os.Exit(0)
}

func sortExec(argv []string) error {
	return nil
}
