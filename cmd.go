package gocmd

import (
	"os/exec"
)

func Pipe(commands ...*exec.Cmd) ([]byte, error) {
	for i, command := range commands[:len(commands)-1] {
		out, err := command.StdoutPipe()
		if err != nil {
			return nil, err
		}
		command.Start()
		commands[i+1].Stdin = out
	}
	final, err := commands[len(commands)-1].CombinedOutput()
	if err != nil {
		return nil, err
	}
	return final, nil
}
