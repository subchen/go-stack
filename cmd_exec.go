package gstack

import (
	"os"
	"os/exec"
)

func ExecCommandOutput(name string, arg ...string) (output string, err error) {
	cmd := exec.Command(name, arg...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func ExecCommandStdpipe(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
