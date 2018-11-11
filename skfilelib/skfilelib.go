package skfilelib

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// RmFile remove file function.
func RmFile(rmfilepath string) error {
	var err error
	var abslog string
	abslog, err = filepath.Abs(rmfilepath)
	if _, err = os.Stat(abslog); err != nil {
		return err
	}
	err = os.Remove(abslog)
	return err
}

// Last Return not blank last line.
// If the last line is blank , return the one before line.
func Last(readfilepath string) (string, error) {
	var last string
	var tmp string
	var err error
	var abspath string
	abspath, err = filepath.Abs(readfilepath)
	if _, err = os.Stat(abspath); err != nil {
		return "", err
	}

	fp, err := os.Open(abspath)
	if err != nil {
		return "", fmt.Errorf(readfilepath + " can't be opened")
	}

	last = ""
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		tmp = last
		last = scanner.Text()
	}
	if last == "" {
		last = tmp
	}
	return last, err
}
