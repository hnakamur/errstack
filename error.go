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
// If any of arguments was created with New or Errorf in
// this package, then stack call frames are taken from the
// last argument of those having stack call frames and
// set to the wrapped error.
//
// If none of arguments are created with New or Errorf in
// this package, then stack call frames are generated and
// set to the wrapped error.
//
// The original error returned by errors.New can be
// obtained by calling Unwrap method of the wrapped error.
//
// Call stack frames can be obtained by calling the Stack
// function later at the upper call frame.
func Errorf(format string, a ...interface{}) error {
	var s []Frame
	for i := len(a) - 1; i >= 0; i-- {
		if e2, ok := a[i].(*errorWithStack); ok {
			s = e2.stack
			break
		}
	}
	if s == nil {
		s = stacks(3)
	}
	return &errorWithStack{
		err:   fmt.Errorf(format, a...),
		stack: s,
	}
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
		if e2, ok := err.(*errorWithStack); ok {
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

func stacks(skip int) []Frame {
	var ss []Frame
	pcs := make([]uintptr, atomic.LoadUint32(&MaxFrames))
	n := runtime.Callers(0, pcs[:])
	for i := skip; i < n; i++ {
		pc := pcs[i]
		fn := runtime.FuncForPC(pc)
		path, line := fn.FileLine(pc)
		name := fn.Name()
		ss = append(ss, Frame{Name: name, Line: line, Path: path})
	}
	return ss
}