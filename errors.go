package gstack

import (
	"fmt"
)

// PanicIfErr throws a panic if any
// of the passed args is a non nil error
func PanicIfErr(args ...interface{}) {
	for _, v := range args {
		if err, _ := v.(error); err != nil {
			panic(err)
		}
	}
}

// AsError returns r as error, converting it when necessary
func AsError(r interface{}) error {
	if r == nil {
		return nil
	}
	if err, ok := r.(error); ok {
		return err
	}
	return fmt.Errorf("%v", r)
}

// AsErrorString returns error message, converting it when necessary
func AsErrorString(r interface{}) string {
	if r == nil {
		return ""
	}
	if err, ok := r.(error); ok {
		return err.Error()
	}
	if err, ok := r.(string); ok {
		return err
	}
	return fmt.Sprintf("%v", r)
}
