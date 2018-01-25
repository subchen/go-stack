package retry_test

import (
	"fmt"
	"time"

	"github.com/subchen/go-stack/retry"
)

func ExampleOperation_Attempt() {
	operation := &retry.Operation{
		Retries:    3,
		MinInteval: 1 * time.Second,
		Factor:     1.2,
		Randomize:  true,
	}

	err := operation.Attempt(func(attempt int) error {
		fmt.Printf("%v: %d: do something\n", time.Now(), attempt)

		return fmt.Errorf("some error")
	})

	if err != nil {
		panic(err)
	}
}
