package errstack

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// ErrorWithLV is an error with methods for adding a pair of label and value.
type ErrorWithLV interface {
	error

	String(label, value string) ErrorWithLV
	Stringer(label string, value fmt.Stringer) ErrorWithLV
	HexByte(label string, value byte) ErrorWithLV
	HexBytes(label string, value []byte) ErrorWithLV
	Fmt(label string, format string, a ...interface{}) ErrorWithLV
	Bool(label string, value bool) ErrorWithLV
	Int(label string, value int) ErrorWithLV
	Int8(label string, value int8) ErrorWithLV
	Int16(label string, value int16) ErrorWithLV
	Int32(label string, value int32) ErrorWithLV
	Int64(label string, value int64) ErrorWithLV
	Uint(label string, value uint) ErrorWithLV
	Uint8(label string, value uint8) ErrorWithLV
	Uint16(label string, value uint16) ErrorWithLV
	Uint32(label string, value uint32) ErrorWithLV
	Uint64(label string, value uint64) ErrorWithLV
	Float32(label string, value float32) ErrorWithLV
	Float64(label string, value float64) ErrorWithLV
	Time(label string, value time.Time, format string) ErrorWithLV
	UTCTime(label string, value time.Time) ErrorWithLV
}

type errorWithLV struct {
	err error
	lv  []string
}

// WithLV wraps the error and attach pairs of labels and values
// to the wrapped error.
//
// If any of the err's chain has the LV() []string method,
// the pairs of labeld and values of the wrapped error will be
// the concatination of the pairs of labels and values of the
// first error which has LV() []string method and the lv argument.
//
// The pairs of labels and values can be get calling the LV
// function to the wrapped error.
func WithLV(err error, lv ...string) ErrorWithLV {
	if err == nil {
		panic("err must not be nil")
	}
	if len(lv)%2 == 1 {
		panic("lv must be label and value pairs")
	}

	e2 := &errorWithLV{err: err}
	if lv2 := LV(err); len(lv2) > 0 {
		e2.lv = make([]string, len(lv2)+len(lv))
		copy(e2.lv, lv2)
		copy(e2.lv[len(lv2):], lv)
	} else {
		e2.lv = lv
	}
	return e2
}

// LV get the pairs of labels and values of the first error
// which has LV() []string method in the err's chain.
func LV(err error) []string {
	for {
		if e2, ok := err.(interface{ LV() []string }); ok {
			return e2.LV()
		}

		if err = errors.Unwrap(err); err == nil {
			return nil
		}
	}
}

func (e *errorWithLV) Error() string {
	return e.err.Error()
}

func (e *errorWithLV) Unwrap() error {
	return e.err
}

func (e *errorWithLV) LV() []string {
	return e.lv
}

func (e *errorWithLV) String(label, value string) ErrorWithLV {
	e.lv = append(e.lv, label, value)
	return e
}

func (e *errorWithLV) Stringer(label string, value fmt.Stringer) ErrorWithLV {
	e.lv = append(e.lv, label, value.String())
	return e
}

func (e *errorWithLV) HexByte(label string, value byte) ErrorWithLV {
	e.lv = append(e.lv, label, hex.EncodeToString([]byte{value}))
	return e
}

func (e *errorWithLV) HexBytes(label string, value []byte) ErrorWithLV {
	e.lv = append(e.lv, label, hex.EncodeToString(value))
	return e
}

func (e *errorWithLV) Fmt(label string, format string, a ...interface{}) ErrorWithLV {
	e.lv = append(e.lv, label, fmt.Sprintf(format, a...))
	return e
}

func (e *errorWithLV) Bool(label string, value bool) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.FormatBool(value))
	return e
}

func (e *errorWithLV) Int(label string, value int) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.Itoa(value))
	return e
}

func (e *errorWithLV) Int8(label string, value int8) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.Itoa(int(value)))
	return e
}

func (e *errorWithLV) Int16(label string, value int16) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.Itoa(int(value)))
	return e
}

func (e *errorWithLV) Int32(label string, value int32) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.FormatInt(int64(value), 10))
	return e
}

func (e *errorWithLV) Int64(label string, value int64) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.FormatInt(int64(value), 10))
	return e
}

func (e *errorWithLV) Uint(label string, value uint) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.FormatUint(uint64(value), 10))
	return e
}

func (e *errorWithLV) Uint8(label string, value uint8) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.FormatUint(uint64(value), 10))
	return e
}

func (e *errorWithLV) Uint16(label string, value uint16) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.FormatUint(uint64(value), 10))
	return e
}

func (e *errorWithLV) Uint32(label string, value uint32) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.FormatUint(uint64(value), 10))
	return e
}

func (e *errorWithLV) Uint64(label string, value uint64) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.FormatUint(uint64(value), 10))
	return e
}

func (e *errorWithLV) Float32(label string, value float32) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.FormatFloat(float64(value), 'g', -1, 32))
	return e
}

func (e *errorWithLV) Float64(label string, value float64) ErrorWithLV {
	e.lv = append(e.lv, label, strconv.FormatFloat(value, 'g', -1, 64))
	return e
}

func (e *errorWithLV) Time(label string, value time.Time, format string) ErrorWithLV {
	if format == "" {
		format = time.RFC3339
	}
	e.lv = append(e.lv, label, value.Format(format))
	return e
}

func (e *errorWithLV) UTCTime(label string, value time.Time) ErrorWithLV {
	e.lv = append(e.lv, label, value.Format("2006-01-02T15:04:05.999999Z"))
	return e
}
