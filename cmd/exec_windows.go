// +build windows

package cmd

import (
	"os"
)

func getShell() string {
	shell := "cmd"
	if other := os.Getenv("SHELL"); other != "" {
		shell = other
	}
	return shell
}

func ShellOutput(script string) (output string, err error) {
	script = "\"" + script + "\""
	return ExecOutput(getShell(), "/C", script)
}

func Shell(script string) error {
	script = "\"" + script + "\""
	return Exec(getShell(), "/C", script)
}
