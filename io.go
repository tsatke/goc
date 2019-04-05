package goc

import (
	"fmt"
	"io"
	"os"
	"strings"

	"gitlab.com/TimSatke/abc"
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
	_, err := io.WriteString(Out, fmt.Sprint(args...))
	if err != nil {
		abc.Errorf("Unable to write to desired output: %v", err)
	}
}

func Println(args ...interface{}) {
	_, err := io.WriteString(Out, fmt.Sprintln(args...))
	if err != nil {
		abc.Errorf("Unable to write line to desired output: %v", err)
	}
}

func Printf(format string, args ...interface{}) {
	_, err := io.WriteString(Out, fmt.Sprintf(format, args...))
	if err != nil {
		abc.Errorf("Unable to write format to desired output: %v", err)
	}
}

func Prompt(text string) bool {
	Print(text)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		abc.Errorf("Unable to from desired input: %v", err)
	}
	return strings.ToLower(input) == "y"
}
