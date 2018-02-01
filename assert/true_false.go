package assert

import "testing"

// True asserts that the specified value is true.
//
//    assert.True(t, myBool)
//
func True(t *testing.T, value bool) {
	if value != true {
		fail(t, "Expected: true, Actual: %#v", value)
	}
}

// False asserts that the specified value is false.
//
//    assert.False(t, myBool)
//
func False(t *testing.T, value bool) {
	if value != false {
		fail(t, "Expected: false, Actual: %#v", value)
	}
}
