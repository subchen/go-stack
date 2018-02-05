package ss

import (
	"strings"
)

func Indent(v string, spaces int) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func IndentAll(v string, spaces int) string {
	return "\n" + Indent(v, spaces)
}
