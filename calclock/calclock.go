package main

import (
	"fmt"
	"os"
)

func main() {
	//	flag.Parse()
	yyyymmdd := "19700101"
	t, err := uspgo.DateToUnixSec(yyyymmdd)
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Errorf("calckock : %d", err))
		os.Exit(1)
	}
	fmt.Printf("%d", t)
}
