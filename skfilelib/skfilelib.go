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

// Last is return not blank last line of file.
// If the last line is blank, return the one before line.
// If file is blank, return blank.
func Last(readfilepath string) (string, error) {
	var last string
	var tmp string

	abspath, err := filepath.Abs(readfilepath)
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

// Touch is create blnak file.
func Touch(touchfilepath string) error {

	touchfilepath, err1 := filepath.Abs(touchfilepath)
	if err1 != nil {
		return err1
	}
	fp, err2 := os.OpenFile(touchfilepath, os.O_CREATE, 0666)
	if err2 != nil {
		return err2
	}
	//	_, err3 := fp.WriteString("")
	defer fp.Close()
	return err2
}

/*
Exists is file exists check.
If file exists, return true.
If file not exists, return false.
*/
func Exists(filename string) bool {
	fpath, err1 := filepath.Abs(filename)
	if err1 != nil {
		return false
	}
	_, err2 := os.Stat(fpath)
	return err2 == nil
}
