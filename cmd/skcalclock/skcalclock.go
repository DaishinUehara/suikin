package main

import (
	"fmt"
	"os"
	"github.com/DaishinUehara/suikin/skcmnlib"
)

func main() {
	//	flag.Parse()
	yyyymmdd := "19700101"
	t, err := skcmnlib.DateToUnixSec(yyyymmdd)
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Errorf("calckock : %d", err))
		os.Exit(1)
	}
	fmt.Printf("%d", t)
}
