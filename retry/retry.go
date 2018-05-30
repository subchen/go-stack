// The retry package encapsulates the mechanism around retrying operation.
//
// It is a golang implementation for nodejs: https://www.npmjs.com/package/retry.
package retry

import (
	"math"
	"math/rand"
	"time"
)

// Operation defines the options for retry
// sleepInteval = min(random * minInteval * pow(factor, attempt), maxInteval)
type Operation struct {
	Retries    int           // The maximum amount of times to retry the operation
	MinInteval time.Duration // default is 1s
	MaxInteval time.Duration // defaults is equals to MinInteval
	Factor     float64       // The exponential factor to use, default is 1
	Randomize  bool          // Randomizes the timeouts by multiplying with a factor between 0.5 to 1.5
}

// RetryFunc is func, attampt is 0-based
type RetryFunc func(attempt int) error

// Attampt accepts the function fn that is to be retried and executes it.
func (o *Operation) Attempt(fn RetryFunc) error {
	retries := o.Retries
	minInteval := o.MinInteval
	maxInteval := o.MaxInteval
	factor := o.Factor

	if retries < 0 {
		retries = 0
	}
	if minInteval <= 0 {
		minInteval = 1 * time.Second
	}
	if maxInteval < minInteval {
		maxInteval = minInteval
	}
	if factor < 1 {
		factor = 1
	}

	attempt := 0
	for true {
		err := fn(attempt)
		if err == nil {
			return nil
		}

		attempt++
		if attempt > retries {
			return err
		}

		sleep := float64(minInteval) * math.Pow(factor, float64(attempt))
		sleep = math.Min(sleep, float64(maxInteval))
		if o.Randomize {
			sleep = (rand.Float64() + 0.5) * sleep
		}
		time.Sleep(time.Duration(sleep))
	}

	panic("unreachable")
}
