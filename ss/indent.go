package ss

import (
	"strings"
)

func IndentLines(v string, spaces int) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func IndentLinesFull(v string, spaces int) string {
	return "\n" + IndentLines(v, spaces)
}
