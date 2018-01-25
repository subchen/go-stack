package retry

import (
	"time"
)

type Operation struct {
	Retries int
	Sleep   time.Duration
	Step    time.Duration
}

type RetryFunc func() error

func (o *Operation) Attempt(fn RetryFunc) error {
	retries := o.Retries
	sleep := o.Sleep
	step := o.Step

	if retries <= 0 {
		panic("retries must be large than zero")
	}
	if sleep <= 0 {
		panic("sleep must be large than zero")
	}
	if step < 0 {
		panic("step must be large or equals than zero")
	}

	for true {
		err := fn()
		if err == nil {
			return nil
		}

		if retries <= 0 {
			return err
		}

		retries--
		sleep += step
		time.Sleep(sleep)
	}

	panic("unreachable")
}
