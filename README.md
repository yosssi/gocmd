# gocmd - Command utility functions in Go

[![Build Status](https://drone.io/github.com/yosssi/gocmd/status.png)](https://drone.io/github.com/yosssi/gocmd/latest)

## Usage

### func Pipe (commands ...*exec.Cmd) ([]byte, error)

Pipe executes the commands passed as the arguments connecting with the pipe and returns the output and error.

#### Example
	import (
		"github.com/yosssi/gocmd"
		"os/exec"
	)

	output, err := gocmd.Pipe(exec.Command("find", "."), exec.Command("grep", "\\.go"))
