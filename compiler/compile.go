package compiler

import (
	"bytes"
	"ksp.sk/proboj/web/config"
	"os/exec"
)

func Compile(root string) (string, error) {
	cmd := exec.Command(config.Configuration.MakeCommand, "player")
	cmd.Dir = root
	var b bytes.Buffer
	cmd.Stderr = &b

	err := cmd.Run()
	_, isItExitError := err.(*exec.ExitError)
	if isItExitError {
		return b.String(), nil
	}
	return "", err
}
