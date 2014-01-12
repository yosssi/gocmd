package gocmd

import (
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

const (
	lenExpected = 2
)

func TestPipe(t *testing.T) {
	ls := exec.Command("ls", "-ltr")
	grep := exec.Command("grep", "cmd_test.go")
	output, err := Pipe(ls, grep)
	if err != nil {
		panic(err)
	}
	if len(strings.Split(string(output), "\n")) != lenExpected {
		t.Error("The length of the results should be " + strconv.Itoa(lenExpected))
	}
}
