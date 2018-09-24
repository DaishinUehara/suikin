package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func printUsage() {
	_usage := "plus <n1> <n2> ...\n"
	_version := "Thu Dec  8 10:46:51 JST 2018\n"
	_code := "Suikin 2.4 Go (windows/go-lang 1.9/UTF-8)\n"
	_description := []string{
		"<n1> <n2> ... must be int.\n",
		"if <n1> <n2> ... are not int then exit status will be 2.\n",
		"\n",
	}

	os.Stderr.Write([]byte("Usage    :" + _usage))
	os.Stderr.Write([]byte("Version  :" + _version))
	os.Stderr.Write([]byte("           " + _code))
	os.Stderr.Write([]byte("Description\n"))
	for _, desc := range _description {
		os.Stderr.Write([]byte(desc))
	}
	os.Exit(1)
}

func die(message string) {
	os.Stderr.Write([]byte("Error[plus] :" + message))
	os.Exit(1)
}

func main() {
	var tmp1, tmp2 *big.Int
	var l int
	var err1 error

	var _errmsg = []byte("Parse Error!")

	if len(os.Args) < 2 {
		printUsage()
	}

	tmp1 = big.NewInt(0)
	for i, agrStr := range os.Args {
		if i == 0 {
			continue
		}
		l, err1 = strconv.Atoi(agrStr)
		if err1 != nil {
			os.Stderr.Write(_errmsg)
			os.Exit(2)
		}
		tmp2 = big.NewInt(int64(l))
		tmp1.Add(tmp1, tmp2)
	}
	fmt.Println(tmp1)
}
