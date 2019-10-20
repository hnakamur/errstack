package errstack_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/hnakamur/errstack"
)

func TestWithLV(t *testing.T) {
	t.Run("wrapWithErrstackErrorfW", func(t *testing.T) {
		inner := func() error {
			return errstack.WithLV(errors.New("my error"), "reqID", "req1")
		}
		outer := func() error {
			return errstack.Errorf("outer: %w", inner())
		}
		err := outer()
		if got, want := errstack.LV(err), []string{"reqID", "req1"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("wrapWithErrstackErrorfs", func(t *testing.T) {
		inner := func() error {
			return errstack.WithLV(errors.New("my error"), "reqID", "req1")
		}
		outer := func() error {
			return errstack.Errorf("outer: %s", inner())
		}
		err := outer()
		if got, want := errstack.LV(err), []string{"reqID", "req1"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("wrapWithFmtErrorfW", func(t *testing.T) {
		inner := func() error {
			return errstack.WithLV(errors.New("my error"), "reqID", "req1")
		}
		outer := func() error {
			return fmt.Errorf("outer: %w", inner())
		}
		err := outer()
		if got, want := errstack.LV(err), []string{"reqID", "req1"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("wrapWithFmtErrorfs", func(t *testing.T) {
		inner := func() error {
			return errstack.WithLV(errors.New("my error"), "reqID", "req1")
		}
		outer := func() error {
			return fmt.Errorf("outer: %s", inner())
		}
		err := outer()
		var want []string
		if got := errstack.LV(err); !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("concaatWithErrstackErrorfW", func(t *testing.T) {
		inner := func() error {
			return errstack.WithLV(errors.New("my error"), "reqID", "req1")
		}
		outer := func() error {
			return errstack.WithLV(errstack.Errorf("outer: %w", inner()), "userID", "user1")
		}
		err := outer()
		if got, want := errstack.LV(err), []string{"reqID", "req1", "userID", "user1"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("concaatWithErrstackErrorfs", func(t *testing.T) {
		inner := func() error {
			return errstack.WithLV(errors.New("my error"), "reqID", "req1")
		}
		outer := func() error {
			return errstack.WithLV(errstack.Errorf("outer: %s", inner()), "userID", "user1")
		}
		err := outer()
		if got, want := errstack.LV(err), []string{"reqID", "req1", "userID", "user1"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("concaatWithFmtErrorfW", func(t *testing.T) {
		inner := func() error {
			return errstack.WithLV(errors.New("my error"), "reqID", "req1")
		}
		outer := func() error {
			return errstack.WithLV(fmt.Errorf("outer: %w", inner()), "userID", "user1")
		}
		err := outer()
		if got, want := errstack.LV(err), []string{"reqID", "req1", "userID", "user1"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("concaatWithFmtErrorfs", func(t *testing.T) {
		inner := func() error {
			return errstack.WithLV(errors.New("my error"), "reqID", "req1")
		}
		outer := func() error {
			return errstack.WithLV(fmt.Errorf("outer: %s", inner()), "userID", "user1")
		}
		err := outer()
		if got, want := errstack.LV(err), []string{"userID", "user1"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
}
