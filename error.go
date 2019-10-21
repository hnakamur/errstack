// Package errstack provides functions which creates or wraps
// an error from which you can get call stack frames later
// at an upper call frame.
// Functions New and Errorf are a drop-in replacement of
// errors.New and fmt.Errorf.
package errstack

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"sync/atomic"
)

// MaxFrames is the maximum stack frame count used in New and Errorf.
// Use functions in the sync/atomic package to modify this value.
var MaxFrames = uint32(128)

type errorWithStack struct {
	err   error
	stack []Frame
}

// Frame is a call stack frame.
type Frame struct {
	Name string
	Line int
	Path string
}

// New creates an error with errors.New and
// returns a wrapped error.
//
// Call stack frames are generated and set to the wrapped
// error.
//
// The original error returned by errors.New can be
// obtained by calling Unwrap method of the wrapped error.
//
// Call stack frames can be obtained by calling the Stack
// function later at the upper call frame.
func New(text string) error {
	s := stacks(3)
	return &errorWithStack{
		err:   errors.New(text),
		stack: s,
	}
}

// Errorf creates an error with fmt.Errorf and
// returns a wrapped error.
//
// If any of arguments has LV() []string method,
// then the error is wrapped with the result from the last
// argument which has LV() []string method. That result
// can be obtained later with the LV function.
//
// If any of arguments has the Stack() []Frame method,
// then stack call frames are taken from the last argument
// of those having stack call frames and set to the wrapped
// error.
//
// If none of arguments has the Stack() []Frame method,
// then stack call frames are generated and set to the
// wrapped error.
//
// The original error returned by fmt.Errorf can be
// obtained by calling Unwrap method of the wrapped error.
//
// Call stack frames can be obtained by calling the Stack
// function later at the upper call frame.
func Errorf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)

	var lv []string
	for i := len(a) - 1; i >= 0; i-- {
		if e2, ok := a[i].(error); ok {
			if lv2 := LV(e2); lv2 != nil {
				lv = lv2
				break
			}
		}
	}

	var s []Frame
	for i := len(a) - 1; i >= 0; i-- {
		if e2, ok := a[i].(error); ok {
			if s2 := Stack(e2); s2 != nil {
				s = s2
				break
			}
		}
	}
	if s == nil {
		s = stacks(3)
	}

	err = &errorWithStack{
		err:   err,
		stack: s,
	}
	if lv == nil {
		return err
	}
	return &errorWithLV{err: err, lv: lv}
}

// Stack finds the first error in err's chain that has stack call frames,
// and returns those if found.
//
// Note you need to build an error chain only with fmt.Errorf with "%w",
// errstack.Errorf, and errstack.New in order to get stack frames.
// Otherwise Stack returns nil.
func Stack(err error) []Frame {
	if err == nil {
		return nil
	}
	for {
		if e2, ok := err.(interface{ Stack() []Frame }); ok {
			return e2.Stack()
		}

		if e2, ok := err.(interface{ Unwrap() error }); ok {
			err = e2.Unwrap()
		} else {
			break
		}
	}
	return nil
}

func (e *errorWithStack) Error() string {
	return e.err.Error()
}

func (e *errorWithStack) Unwrap() error {
	return e.err
}

func (e *errorWithStack) Stack() []Frame {
	return e.stack
}

// String retruns a string representation for the stack.
func (f *Frame) String() string {
	var b []byte
	b = append(b, f.Name...)
	b = append(b, '@')
	b = append(b, f.Path...)
	b = append(b, ':')
	b = strconv.AppendInt(b, int64(f.Line), 10)
	return string(b)
}

func stacks(skip int) []Frame {
	var ss []Frame
	pcs := make([]uintptr, atomic.LoadUint32(&MaxFrames))
	runtime.Callers(skip, pcs)
	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		ss = append(ss, Frame{
			Name: frame.Function,
			Line: frame.Line,
			Path: frame.File,
		})
	}
	return ss
}
