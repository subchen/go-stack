assert
======================

[![Go Report Card](https://goreportcard.com/badge/github.com/subchen/gstack/assert)](https://goreportcard.com/report/github.com/subchen/gstack/assert)
[![GoDoc](https://godoc.org/github.com/subchen/gstack/assert?status.svg)](https://godoc.org/github.com/subchen/gstack/assert)

The assert package provides some helpful methods that allow you to write better test code in Go.

* Prints friendly for read
* Readable code

Installation
---------------

```bash
$ go get github.com/subchen/gstack/assert
```

Usage
---------------

```go
import (
    "testing"
    "github.com/subchen/gstack/assert"
)

func TestToString(t *testing.T) {
    assert.Equal(t, ToString(nil), "")
    assert.Equal(t, ToString(true), "true")
    assert.Equal(t, ToString(0), "0")
}
```

if you assert many times, use the below:

```go
import (
    "testing"
    "github.com/subchen/gstack/assert"
)

func TestToString(t *testing.T) {
    assert := assert.New(t)

    assert.Equal(ToString(nil), "")
    assert.Equal(ToString(true), "true")
    assert.Equal(ToString(0), "0")
}
```

result on failure

```bash
$ go test
--- FAIL: TestToString (0.00s)
        to_string_test.go:12
                Expected: "true"
                Actual  : "false"
FAIL
exit status 1
```

### API on godoc.org

https://godoc.org/github.com/subchen/gstack/assert

