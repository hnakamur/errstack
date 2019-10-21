package errstack_test

import (
	"errors"
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"

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

func TestErrorWithLV(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).String("reqID", "req1")
		if got, want := errstack.LV(err), []string{"reqID", "req1"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Stringer", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Stringer("addr", net.IPv4(192, 0, 2, 1))
		if got, want := errstack.LV(err), []string{"addr", "192.0.2.1"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("HexByte", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).HexByte("byte", '\xba')
		if got, want := errstack.LV(err), []string{"byte", "ba"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("HexBytes", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).HexBytes("bytes", []byte{'\xba', '\xbe'})
		if got, want := errstack.LV(err), []string{"bytes", "babe"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Fmt", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Fmt("strings", "%v", []string{"a", "b"})
		if got, want := errstack.LV(err), []string{"strings", "[a b]"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Bool", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Bool("bool", true)
		if got, want := errstack.LV(err), []string{"bool", "true"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Int", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Int("int", -12)
		if got, want := errstack.LV(err), []string{"int", "-12"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Int8", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Int8("int8", -12)
		if got, want := errstack.LV(err), []string{"int8", "-12"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Int16", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Int16("int16", -12)
		if got, want := errstack.LV(err), []string{"int16", "-12"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Int32", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Int32("int32", -12)
		if got, want := errstack.LV(err), []string{"int32", "-12"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Int64", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Int64("int64", -12)
		if got, want := errstack.LV(err), []string{"int64", "-12"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Uint", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Uint("uint", 12)
		if got, want := errstack.LV(err), []string{"uint", "12"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Uint8", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Uint8("uint8", 12)
		if got, want := errstack.LV(err), []string{"uint8", "12"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Uint16", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Uint16("uint16", 12)
		if got, want := errstack.LV(err), []string{"uint16", "12"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Uint32", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Uint32("uint32", 12)
		if got, want := errstack.LV(err), []string{"uint32", "12"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Uint64", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Uint64("uint64", 12)
		if got, want := errstack.LV(err), []string{"uint64", "12"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Float32", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Float32("float32", 1.2)
		if got, want := errstack.LV(err), []string{"float32", "1.2"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Float64", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).Float64("float64", 1.2)
		if got, want := errstack.LV(err), []string{"float64", "1.2"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("Time", func(t *testing.T) {
		loc, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			t.Fatal(err)
		}
		err = errstack.WithLV(errors.New("my error")).Time("time", time.Date(2019, 10, 22, 5, 31, 53, 123456789, loc), time.RFC3339Nano)
		if got, want := errstack.LV(err), []string{"time", "2019-10-22T05:31:53.123456789+09:00"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("UTCTime", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).UTCTime("utctime", time.Date(2019, 10, 22, 5, 31, 53, 123456789, time.UTC))
		if got, want := errstack.LV(err), []string{"utctime", "2019-10-22T05:31:53.123456Z"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
	t.Run("StringInt64", func(t *testing.T) {
		err := errstack.WithLV(errors.New("my error")).String("reqID", "req1").Int64("userID", 1)
		if got, want := errstack.LV(err), []string{"reqID", "req1", "userID", "1"}; !reflect.DeepEqual(got, want) {
			t.Errorf("lv unmatch, got:%v, want:%v", got, want)
		}
	})
}
