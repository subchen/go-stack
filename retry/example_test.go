package retry_test

import (
	"fmt"
	"time"

	"github.com/subchen/go-stack/retry"
)

func ExampleOperation_Attempt() {
	operation := &retry.Operation{
		Retries: 3,
		Sleep:   1 * time.Second,
		Step:    200 * time.Millisecond,
	}

	err := operation.Attempt(func() error {
		fmt.Printf("%v do something\n", time.Now())

		return fmt.Errorf("some error")
	})

	if err != nil {
		panic(err)
	}
}
