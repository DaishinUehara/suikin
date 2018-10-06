package main

import (
	"fmt"
	"os"
)

func main() {
	err := sample01Exec(os.Args)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func sample01Exec(argv []string) error {
	return fmt.Errorf("Err")
}
