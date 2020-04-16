package vrb

import (
	"fmt"
	"io"
	"io/ioutil"
)

var vrbWriter = ioutil.Discard

// Public method to initialize the verbose io.Writer
func Init(wrtr io.Writer) {
	vrbWriter = wrtr
}

// Print usage is equivalent to https://golang.org/pkg/fmt/#Print
func Print(a ...interface{}) (int, error) {
	return fmt.Fprint(vrbWriter, a...)
}

// Printf usage is equivalent to https://golang.org/pkg/fmt/#Printf
func Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(vrbWriter, format, a...)
}

// Println usage is equivalent to https://golang.org/pkg/fmt/#Println
func Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(vrbWriter, a...)
}
