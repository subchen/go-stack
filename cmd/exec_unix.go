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

func ShellOutput(scripts string) (output string, err error) {
	return ExecOutput(getShell(), "-c", scripts)
}

func Shell(scripts string) error {
	return Exec(getShell(), "-c", scripts)
}
