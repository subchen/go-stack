package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecOutput(name string, arg ...string) (output string, err error) {
	cmd := exec.Command(name, arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func ShellOutput(shell string) (output string, err error) {
	return ExecOutput("sh", "-c", shell)
}

func Exec(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Shell(shell string) error {
	return Exec("sh", "-c", shell)
}
