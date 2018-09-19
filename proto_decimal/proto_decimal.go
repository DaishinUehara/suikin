package main

import (
	"fmt"
	"math/big"
	"os"
)

func getNeedBitSize(st string) (bitsize uint, err error) {

	return 8, nil
}

func main() {
	var _errmsg = []byte("Parse Error!")

	var bitsize uint
	cf := "0.25622222222222"
	bitsize, err1 := getNeedBitSize(cf)
	if err1 != nil {
		os.Stderr.Write(_errmsg)
		os.Exit(1)
	}

	ftmp, _, err2 := big.ParseFloat(cf, 10, bitsize, big.ToNearestEven)
	if err2 != nil {
		os.Stderr.Write(_errmsg)
		os.Exit(1)
	}

	fmt.Println(ftmp)
	os.Exit(0)
}
