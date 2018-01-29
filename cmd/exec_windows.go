// +build windows
package exec

import (
	"os"
	"os/exec
)

func getShell() string {
	shell := "cmd"
	if other := os.Getenv("SHELL"); other != "" {
		shell = other
	}
	return shell
}

func ShellOutput(scripts string) (output string, err error) {
	script = "\"" + script + "\""
	return ExecOutput(getShell(), "/C", scripts)
}

func Shell(scripts string) error {
	script = "\"" + script + "\""
	return Exec(getShell(), "/C", scripts)
}
