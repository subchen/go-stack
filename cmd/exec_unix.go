// +build !windows

package cmd

import (
	"os"
)

func getShell() string {
	shell := "/bin/sh"
	if other := os.Getenv("SHELL"); other != "" {
		shell = other
	}
	return shell
}

func ShellOutput(script string) (output string, err error) {
	return ExecOutput(getShell(), "-c", script)
}

func Shell(script string) error {
	return Exec(getShell(), "-c", script)
}
