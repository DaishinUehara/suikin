package skstublib

import (
	"io/ioutil"
	"os"
)

// SkStdStub Stdin stdout stderr stub for func().
func SkStdStub(inbuf string, argv []string, fn func([]string) error) (strStdOut string, strStdErr string, err error) {
	inrp, inwp, _ := os.Pipe()   // Input Pipe
	outrp, outwp, _ := os.Pipe() // Std Output Pipe
	errrp, errwp, _ := os.Pipe() // Std Error Pipe

	// backup stdin, stdout and atderr
	stdinBak := os.Stdin
	stdoutBak := os.Stdout
	stderrBak := os.Stderr

	// write inbuf and close
	inwp.Write([]byte(inbuf))
	inwp.Close()

	// connect pipe to stdin, stdout and stderror.
	os.Stdin = inrp
	os.Stdout = outwp
	os.Stderr = errwp

	fnerr := fn(argv) // Execute Function

	// restore stdin, stdout and stderr
	os.Stdin = stdinBak
	os.Stdout = stdoutBak
	os.Stderr = stderrBak

	// close stdout and stderr write pipe connected function.
	outwp.Close()
	errwp.Close()

	// read stdout and stderr pipe connected function.
	stdOutBuf, _ := ioutil.ReadAll(outrp)
	stdErrBuf, _ := ioutil.ReadAll(errrp)

	// convert Buffer to String
	strStdOut = string(stdOutBuf)
	strStdErr = string(stdErrBuf)

	return strStdOut, strStdErr, fnerr

}
