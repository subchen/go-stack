package assert

import (
	"strings"
	"testing"
)

func HasPrefix(t *testing.T, str string, prefix string) {
	if !strings.HasPrefix(str, prefix) {
		fail(t, "Excepted '%s' starts with '%s', but didn't!", str, prefix)
	}
}

func HasNotPrefix(t *testing.T, str string, prefix string) {
	if strings.HasPrefix(str, prefix) {
		fail(t, "Excepted '%s' NOT starts with '%s', but did!", str, prefix)
	}
}
func HasSuffix(t *testing.T, str string, prefix string) {
	if !strings.HasSuffix(str, prefix) {
		fail(t, "Excepted '%s' ends with '%s', but didn't!", str, prefix)
	}
}

func HasNotSuffix(t *testing.T, str string, prefix string) {
	if strings.HasSuffix(str, prefix) {
		fail(t, "Excepted '%s' NOT ends with '%s', but did!", str, prefix)
	}
}
