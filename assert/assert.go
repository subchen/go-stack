package assert

import (
	"testing"
)

type Assert struct {
	t *testing.T
}

func New(t *testing.T) *Assert {
	return &Assert{t}
}

func (a *Assert) True(value bool) {
	True(a.t, value)
}

func (a *Assert) False(value bool) {
	False(a.t, value)
}

func (a *Assert) Nil(object interface{}) {
	Nil(a.t, object)
}

func (a *Assert) NotNil(object interface{}) {
	NotNil(a.t, object)
}

func (a *Assert) Empty(object interface{}) {
	Empty(a.t, object)
}

func (a *Assert) NotEmpty(object interface{}) {
	NotEmpty(a.t, object)
}

func (a *Assert) Zero(object interface{}) {
	Zero(a.t, object)
}

func (a *Assert) NotZero(object interface{}) {
	NotZero(a.t, object)
}

func (a *Assert) Error(error interface{}) {
	Error(a.t, error)
}

func (a *Assert) NoError(error interface{}) {
	NoError(a.t, error)
}

func (a *Assert) Equal(actual, expected interface{}) {
	Equal(a.t, actual, expected)
}

func (a *Assert) NotEqual(actual, expected interface{}) {
	NotEqual(a.t, actual, expected)
}

func (a *Assert) EqualVal(actual, expected interface{}) {
	EqualVal(a.t, actual, expected)
}

func (a *Assert) NotEqualVal(actual, expected interface{}) {
	NotEqualVal(a.t, actual, expected)
}

func (a *Assert) Contains(list interface{}, element interface{}) {
	Contains(a.t, list, element)
}

func (a *Assert) NotContains(list interface{}, element interface{}) {
	NotContains(a.t, list, element)
}

func (a *Assert) HasPrefix(str, prefix string) {
	HasPrefix(a.t, str, prefix)
}

func (a *Assert) HasNotPrefix(str, prefix string) {
	HasNotPrefix(a.t, str, prefix)
}

func (a *Assert) HasSuffix(str, suffix string) {
	HasSuffix(a.t, str, suffix)
}

func (a *Assert) HasNotSuffix(str, suffix string) {
	HasNotSuffix(a.t, str, suffix)
}

func (a *Assert) Len(object interface{}, length int) {
	Len(a.t, object, length)
}

func (a *Assert) SameType(object interface{}, anotherObject interface{}) {
	SameType(a.t, object, anotherObject)
}

func (a *Assert) Implements(object interface{}, interfaceObject interface{}) {
	Implements(a.t, object, interfaceObject)
}

func (a *Assert) Panic(action func()) {
	Panic(a.t, action)
}

func (a *Assert) NoPanic(action func()) {
	NoPanic(a.t, action)
}
