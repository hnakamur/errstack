package errstack

import "errors"

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
func WithLV(err error, lv ...string) error {
	if err == nil {
		panic("err must not be nil")
	}
	if len(lv) == 0 || len(lv)%2 == 1 {
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
