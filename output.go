package goc

import (
	"fmt"
	"io"
	"os"
)

var (
	Out io.Writer
	In  io.Reader
)

func init() {
	Out = os.Stdout
	In = os.Stdin
}

func Print(args ...interface{}) {
	io.WriteString(Out, fmt.Sprint(args...))
}

func Println(args ...interface{}) {
	io.WriteString(Out, fmt.Sprintln(args...))
}

func Printf(format string, args ...interface{}) {
	io.WriteString(Out, fmt.Sprintf(format, args...))
}
