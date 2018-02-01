package assert

import "testing"

func Panic(t *testing.T, action func()) {
	defer func() {
		recovered := recover()
		if recovered == nil {
			fail(t, "Should to panic, but didn't")
		}
	}()

	action()
}

func NoPanic(t *testing.T, action func()) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			fail(t, "Unexpected panic found: %#v", recovered)
		}
	}()

	action()
}
