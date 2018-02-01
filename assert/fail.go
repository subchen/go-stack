package assert

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

// Fail reports a failure through
func fail(t *testing.T, msgAndArgs ...interface{}) {
	message := messageFromMsgAndArgs(msgAndArgs...)
	t.Errorf("\r\t%s\n%s\n", getCallerFileAndLine(), message)
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		return msgAndArgs[0].(string)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}

func getCallerFileAndLine() string {
	_, file, line, ok := runtime.Caller(3)

	if strings.HasSuffix(file, "assert.go") {
		_, file, line, ok = runtime.Caller(4)
	}

	if !ok {
		file = "???"
		line = 1
	} else {
		// Truncate file name at last file name separator.
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
	}

	return fmt.Sprintf("%s:%d", file, line)
}
