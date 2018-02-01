package conv

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/subchen/gstack/assert"
)

func TestConvertPointerToObject(t *testing.T) {
	assert := assert.New(t)

	a, b, c := "123", "true", "9.9"

	assert.Equal(ConvertAs(&a, TYPE_INT), 123)
	assert.Equal(ConvertAs(&b, TYPE_BOOL), true)
	assert.Equal(ConvertAs(&c, TYPE_FLOAT64), 9.9)
}

func TestConvertObjectToPointer(t *testing.T) {
	assert := assert.New(t)

	a, b, c := 123, true, 9.9

	assert.Equal(ConvertAs("123", reflect.PtrTo(TYPE_INT)), &a)
	assert.Equal(ConvertAs("true", reflect.PtrTo(TYPE_BOOL)), &b)
	assert.Equal(ConvertAs("9.9", reflect.PtrTo(TYPE_FLOAT64)), &c)
}

func TestAsBool(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AsBool(0), false)
	assert.Equal(AsBool(nil), false)

	assert.Equal(AsBool("false"), false)
	assert.Equal(AsBool("FALSE"), false)
	assert.Equal(AsBool("False"), false)
	assert.Equal(AsBool("f"), false)
	assert.Equal(AsBool("F"), false)
	assert.Equal(AsBool(false), false)
	assert.Equal(AsBool("foo"), false)

	assert.Equal(AsBool("true"), true)
	assert.Equal(AsBool("TRUE"), true)
	assert.Equal(AsBool("True"), true)
	assert.Equal(AsBool("t"), true)
	assert.Equal(AsBool("T"), true)
	assert.Equal(AsBool(1), true)
	assert.Equal(AsBool(true), true)
	assert.Equal(AsBool(-1), true)
}

func TestAsInt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AsInt(nil), 0)
	assert.Equal(AsInt(1), 1)
	assert.Equal(AsInt(int64(1)), 1)
	assert.Equal(AsInt(uint(1)), 1)
	assert.Equal(AsInt("1"), 1)
	assert.Equal(AsInt(true), 1)
}

func TestAsInt64(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AsInt64(nil), int64(0))
	assert.Equal(AsInt64(1), int64(1))
	assert.Equal(AsInt64(int64(1)), int64(1))
	assert.Equal(AsInt64(uint(1)), int64(1))
	assert.Equal(AsInt64("1"), int64(1))
	assert.Equal(AsInt64(true), int64(1))
}

func TestAsUint(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AsUint(nil), uint(0))
	assert.Equal(AsUint(1), uint(1))
	assert.Equal(AsUint(uint(1)), uint(1))
	assert.Equal(AsUint("1"), uint(1))
	assert.Equal(AsUint(true), uint(1))
}

func TestAsUint64(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AsUint64(nil), uint64(0))
	assert.Equal(AsUint64(1), uint64(1))
	assert.Equal(AsUint64(uint(1)), uint64(1))
	assert.Equal(AsUint64(uint64(1)), uint64(1))
	assert.Equal(AsUint64("1"), uint64(1))
	assert.Equal(AsUint64(true), uint64(1))
}

func TestAsFloat64(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AsFloat64(nil), float64(0))
	assert.Equal(AsFloat64(1), float64(1))
	assert.Equal(AsFloat64(float32(1.5)), float64(1.5))
	assert.Equal(AsFloat64(uint(1)), float64(1))
	assert.Equal(AsFloat64("1.9"), float64(1.9))
	assert.Equal(AsFloat64(true), float64(1))
}

func TestAsString(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AsString(nil), "")
	assert.Equal(AsString(1), "1")
	assert.Equal(AsString(int64(2)), "2")
	assert.Equal(AsString(float64(1.9)), "1.9")
	assert.Equal(AsString(true), "true")
	assert.Equal(AsString(errors.New("error")), "error")
	assert.Equal(AsString(bytes.NewBufferString("buffer")), "buffer")
	assert.Equal(AsString([]byte("buffer")), "buffer")
}

func TestAsDuration(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AsDuration("2h45m36s55ms66ns"), time.Duration(9936055000066))
	assert.Equal(AsDuration(9936055000066), time.Duration(9936055000066))
	assert.Equal(AsDuration(uint64(9936055000066)), time.Duration(9936055000066))
}

func TestAsTime(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AsTime("2006-01-02T15:04:05.999+00:00").Unix(), time.Date(2006, 1, 2, 15, 4, 5, 999, time.UTC).Unix())
	assert.Equal(AsTime("2006-01-02T15:04:05+00:00").Unix(), time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC).Unix())
	assert.Equal(AsTime("2006-01-02T15:04:05"), time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC))
	assert.Equal(AsTime("2006-01-02"), time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC))
	assert.Equal(AsTime("2006-01-02 15:04:05"), time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC))
	assert.Equal(AsTime("2006-01-02 15:04:05+00:00").Unix(), time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC).Unix())
	assert.Equal(AsTime("2006-01-02 15:04:05 +00:00").Unix(), time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC).Unix())
	assert.Equal(AsTime("2006-01-02 15:04:05 +0000").Unix(), time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC).Unix())
	assert.Equal(AsTime("02 Jan 06 15:04 UTC"), time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC))
	assert.Equal(AsTime("02 Jan 2006"), time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC))
	assert.Equal(AsTime("Mon, 02 Jan 2006 15:04:05 UTC"), time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC))
}
