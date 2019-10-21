package errstack_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/hnakamur/errstack"
)

func TestStack(t *testing.T) {
	t.Run("wrapAtTop", func(t *testing.T) {
		err := testStackWrapAtTopLevel2()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, []string{
			"github.com/hnakamur/errstack_test.testStackWrapAtTopLevel2",
			"github.com/hnakamur/errstack_test.TestStack.func1",
		})
	})
	t.Run("wrapAtMiddle", func(t *testing.T) {
		err := testStackWrapAtMiddleLevel3()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, []string{
			"github.com/hnakamur/errstack_test.testStackWrapAtMiddleLevel2",
			"github.com/hnakamur/errstack_test.testStackWrapAtMiddleLevel3",
			"github.com/hnakamur/errstack_test.TestStack.func2",
		})
	})
	t.Run("wrapAtBottom", func(t *testing.T) {
		err := testStackWrapAtBottomLevel2()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, []string{
			"github.com/hnakamur/errstack_test.testStackWrapAtBottomLevel1",
			"github.com/hnakamur/errstack_test.testStackWrapAtBottomLevel2",
			"github.com/hnakamur/errstack_test.TestStack.func3",
		})
	})
	t.Run("wrapOnlyAtMiddle", func(t *testing.T) {
		err := testStackWrapOnlyAtMiddleLevel3()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, []string{
			"github.com/hnakamur/errstack_test.testStackWrapOnlyAtMiddleLevel2",
			"github.com/hnakamur/errstack_test.testStackWrapOnlyAtMiddleLevel3",
			"github.com/hnakamur/errstack_test.TestStack.func4",
		})
	})
	t.Run("wrapOnlyAtBottom", func(t *testing.T) {
		err := testStackWrapOnlyAtBottomLevel2()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, []string{
			"github.com/hnakamur/errstack_test.testStackWrapOnlyAtBottomLevel1",
			"github.com/hnakamur/errstack_test.testStackWrapOnlyAtBottomLevel2",
			"github.com/hnakamur/errstack_test.TestStack.func5",
		})
	})
	t.Run("wrapOnlyAtMiddleNoGood", func(t *testing.T) {
		err := testStackWrapOnlyAtMiddleNoGoodLevel2()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, nil)
	})
	t.Run("wrapOnlyAtBottomNoGood", func(t *testing.T) {
		err := testStackWrapOnlyAtBottomNoGoodLevel2()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, nil)
	})
}

func testStackWrapAtTopLevel2() error {
	return errstack.Errorf("outer: %s", testStackWrapAtTopLevel1())
}
func testStackWrapAtTopLevel1() error { return os.ErrExist }

func testStackWrapAtMiddleLevel3() error {
	return errstack.Errorf("top: %w", testStackWrapAtMiddleLevel2())
}
func testStackWrapAtMiddleLevel2() error {
	return errstack.Errorf("middle: %s", testStackWrapAtMiddleLevel1())
}
func testStackWrapAtMiddleLevel1() error { return os.ErrExist }

func testStackWrapAtBottomLevel2() error {
	return errstack.Errorf("outer: %s", testStackWrapAtBottomLevel1())
}
func testStackWrapAtBottomLevel1() error { return errstack.New("my error") }

func testStackWrapOnlyAtMiddleLevel3() error {
	return fmt.Errorf("top: %w", testStackWrapOnlyAtMiddleLevel2())
}
func testStackWrapOnlyAtMiddleLevel2() error {
	return errstack.Errorf("middle: %s", testStackWrapOnlyAtMiddleLevel1())
}
func testStackWrapOnlyAtMiddleLevel1() error { return os.ErrExist }

func testStackWrapOnlyAtBottomLevel2() error {
	return fmt.Errorf("outer: %w", testStackWrapOnlyAtBottomLevel1())
}
func testStackWrapOnlyAtBottomLevel1() error { return errstack.New("my error") }

func testStackWrapOnlyAtBottomNoGoodLevel2() error {
	return fmt.Errorf("outer: %s", testStackWrapOnlyAtBottomNoGoodLevel1())
}
func testStackWrapOnlyAtBottomNoGoodLevel1() error { return errstack.New("my error") }

func testStackWrapOnlyAtMiddleNoGoodLevel3() error {
	return fmt.Errorf("top: %s", testStackWrapOnlyAtMiddleNoGoodLevel2())
}
func testStackWrapOnlyAtMiddleNoGoodLevel2() error {
	return errstack.Errorf("middle: %s", testStackWrapOnlyAtMiddleNoGoodLevel1())
}
func testStackWrapOnlyAtMiddleNoGoodLevel1() error { return os.ErrExist }

func TestStackWithLV(t *testing.T) {
	t.Run("wrapAtTop", func(t *testing.T) {
		err := testStackWithLVWrapAtTopLevel2()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, []string{
			"github.com/hnakamur/errstack_test.testStackWithLVWrapAtTopLevel2",
			"github.com/hnakamur/errstack_test.TestStackWithLV.func1",
		})
	})
	t.Run("wrapAtMiddle", func(t *testing.T) {
		err := testStackWithLVWrapAtMiddleLevel3()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, []string{
			"github.com/hnakamur/errstack_test.testStackWithLVWrapAtMiddleLevel2",
			"github.com/hnakamur/errstack_test.testStackWithLVWrapAtMiddleLevel3",
			"github.com/hnakamur/errstack_test.TestStackWithLV.func2",
		})
	})
	t.Run("wrapAtBottom", func(t *testing.T) {
		err := testStackWithLVWrapAtBottomLevel2()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, []string{
			"github.com/hnakamur/errstack_test.testStackWithLVWrapAtBottomLevel1",
			"github.com/hnakamur/errstack_test.testStackWithLVWrapAtBottomLevel2",
			"github.com/hnakamur/errstack_test.TestStackWithLV.func3",
		})
	})
	t.Run("wrapOnlyAtMiddle", func(t *testing.T) {
		err := testStackWithLVWrapOnlyAtMiddleLevel3()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, []string{
			"github.com/hnakamur/errstack_test.testStackWithLVWrapOnlyAtMiddleLevel2",
			"github.com/hnakamur/errstack_test.testStackWithLVWrapOnlyAtMiddleLevel3",
			"github.com/hnakamur/errstack_test.TestStackWithLV.func4",
		})
	})
	t.Run("wrapOnlyAtBottom", func(t *testing.T) {
		err := testStackWithLVWrapOnlyAtBottomLevel2()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, []string{
			"github.com/hnakamur/errstack_test.testStackWithLVWrapOnlyAtBottomLevel1",
			"github.com/hnakamur/errstack_test.testStackWithLVWrapOnlyAtBottomLevel2",
			"github.com/hnakamur/errstack_test.TestStackWithLV.func5",
		})
	})
	t.Run("wrapOnlyAtMiddleNoGood", func(t *testing.T) {
		err := testStackWithLVWrapOnlyAtMiddleNoGoodLevel2()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, nil)
	})
	t.Run("wrapOnlyAtBottomNoGood", func(t *testing.T) {
		err := testStackWithLVWrapOnlyAtBottomNoGoodLevel2()
		s := errstack.Stack(err)
		testStackFrameNames(t, s, nil)
	})
}

func testStackWithLVWrapAtTopLevel2() error {
	return errstack.WithLV(
		errstack.Errorf("outer: %s", testStackWithLVWrapAtTopLevel1()),
	).Int64("userID", 1)
}
func testStackWithLVWrapAtTopLevel1() error { return os.ErrExist }

func testStackWithLVWrapAtMiddleLevel3() error {
	return errstack.WithLV(
		errstack.Errorf("top: %w", testStackWithLVWrapAtMiddleLevel2()),
	).String("reqID", "req1")
}
func testStackWithLVWrapAtMiddleLevel2() error {
	return errstack.WithLV(
		errstack.Errorf("middle: %s", testStackWithLVWrapAtMiddleLevel1()),
	).Int64("userID", 1)
}
func testStackWithLVWrapAtMiddleLevel1() error { return os.ErrExist }

func testStackWithLVWrapAtBottomLevel2() error {
	return errstack.WithLV(
		errstack.Errorf("outer: %s", testStackWithLVWrapAtBottomLevel1()),
	).String("reqID", "req1")
}
func testStackWithLVWrapAtBottomLevel1() error {
	return errstack.WithLV(errstack.New("my error")).Int64("userID", 1)
}

func testStackWithLVWrapOnlyAtMiddleLevel3() error {
	return fmt.Errorf("top: %w", testStackWithLVWrapOnlyAtMiddleLevel2())
}
func testStackWithLVWrapOnlyAtMiddleLevel2() error {
	return errstack.WithLV(
		errstack.Errorf("middle: %s", testStackWithLVWrapOnlyAtMiddleLevel1()),
	).Int64("userID", 1)
}
func testStackWithLVWrapOnlyAtMiddleLevel1() error { return os.ErrExist }

func testStackWithLVWrapOnlyAtBottomLevel2() error {
	return fmt.Errorf("outer: %w", testStackWithLVWrapOnlyAtBottomLevel1())
}
func testStackWithLVWrapOnlyAtBottomLevel1() error {
	return errstack.WithLV(errstack.New("my error")).Int64("userID", 1)
}

func testStackWithLVWrapOnlyAtBottomNoGoodLevel2() error {
	return fmt.Errorf("outer: %s", testStackWithLVWrapOnlyAtBottomNoGoodLevel1())
}
func testStackWithLVWrapOnlyAtBottomNoGoodLevel1() error {
	return errstack.WithLV(errstack.New("my error")).Int64("userID", 1)
}

func testStackWithLVWrapOnlyAtMiddleNoGoodLevel3() error {
	return fmt.Errorf("top: %s", testStackWithLVWrapOnlyAtMiddleNoGoodLevel2())
}
func testStackWithLVWrapOnlyAtMiddleNoGoodLevel2() error {
	return errstack.WithLV(
		errstack.Errorf("middle: %s", testStackWithLVWrapOnlyAtMiddleNoGoodLevel1()),
	).Int64("userID", 1)
}
func testStackWithLVWrapOnlyAtMiddleNoGoodLevel1() error { return os.ErrExist }

func TestIs(t *testing.T) {
	t.Run("wrapAtTop", func(t *testing.T) {
		err := testIsWrapAtTopLevel2()
		if got, want := errors.Is(err, os.ErrExist), true; got != want {
			t.Errorf("unmatch Is result, got:%v, want:%v", got, want)
		}
	})
	t.Run("wrapAtMiddle", func(t *testing.T) {
		err := testIsWrapAtMiddleLevel3()
		if got, want := errors.Is(err, os.ErrExist), true; got != want {
			t.Errorf("unmatch Is result, got:%v, want:%v", got, want)
		}
	})
	t.Run("wrapAtBottom", func(t *testing.T) {
		err := testIsWrapAtBottomLevel2()
		if got, want := errors.Is(err, os.ErrExist), false; got != want {
			t.Errorf("unmatch Is result, got:%v, want:%v", got, want)
		}
	})
	t.Run("wrapOnlyAtMiddle", func(t *testing.T) {
		err := testIsWrapOnlyAtMiddleLevel3()
		if got, want := errors.Is(err, os.ErrExist), true; got != want {
			t.Errorf("unmatch Is result, got:%v, want:%v", got, want)
		}
	})
	t.Run("wrapOnlyAtBottom", func(t *testing.T) {
		err := testIsWrapOnlyAtBottomLevel2()
		if got, want := errors.Is(err, os.ErrExist), false; got != want {
			t.Errorf("unmatch Is result, got:%v, want:%v", got, want)
		}
	})
	t.Run("wrapOnlyAtMiddleNoGood", func(t *testing.T) {
		err := testIsWrapOnlyAtMiddleNoGoodLevel3()
		if got, want := errors.Is(err, os.ErrExist), false; got != want {
			t.Errorf("unmatch Is result, got:%v, want:%v", got, want)
		}
	})
	t.Run("wrapOnlyAtBottomNoGood", func(t *testing.T) {
		err := testIsWrapOnlyAtBottomNoGoodLevel2()
		if got, want := errors.Is(err, os.ErrExist), false; got != want {
			t.Errorf("unmatch Is result, got:%v, want:%v", got, want)
		}
	})
}

func testIsWrapAtTopLevel2() error {
	return errstack.Errorf("outer: %w", testIsWrapAtTopLevel1())
}
func testIsWrapAtTopLevel1() error { return os.ErrExist }

func testIsWrapAtMiddleLevel3() error {
	return errstack.Errorf("top: %w", testIsWrapAtMiddleLevel2())
}
func testIsWrapAtMiddleLevel2() error {
	return errstack.Errorf("middle: %w", testIsWrapAtMiddleLevel1())
}
func testIsWrapAtMiddleLevel1() error { return os.ErrExist }

func testIsWrapAtBottomLevel2() error {
	return errstack.Errorf("outer: %w", testIsWrapAtBottomLevel1())
}
func testIsWrapAtBottomLevel1() error { return errstack.New("my error") }

func testIsWrapOnlyAtMiddleLevel3() error {
	return fmt.Errorf("top: %w", testIsWrapOnlyAtMiddleLevel2())
}
func testIsWrapOnlyAtMiddleLevel2() error {
	return errstack.Errorf("middle: %w", testIsWrapOnlyAtMiddleLevel1())
}
func testIsWrapOnlyAtMiddleLevel1() error { return os.ErrExist }

func testIsWrapOnlyAtBottomLevel2() error {
	return fmt.Errorf("outer: %w", testIsWrapOnlyAtBottomLevel1())
}
func testIsWrapOnlyAtBottomLevel1() error { return errstack.New("my error") }

func testIsWrapOnlyAtBottomNoGoodLevel2() error {
	return fmt.Errorf("outer: %s", testIsWrapOnlyAtBottomNoGoodLevel1())
}
func testIsWrapOnlyAtBottomNoGoodLevel1() error { return errstack.New("my error") }

func testIsWrapOnlyAtMiddleNoGoodLevel3() error {
	return fmt.Errorf("top: %s", testIsWrapOnlyAtMiddleNoGoodLevel2())
}
func testIsWrapOnlyAtMiddleNoGoodLevel2() error {
	return errstack.Errorf("middle: %w", testIsWrapOnlyAtMiddleNoGoodLevel1())
}
func testIsWrapOnlyAtMiddleNoGoodLevel1() error { return os.ErrExist }

func TestAs(t *testing.T) {
	t.Run("wrapAtTop", func(t *testing.T) {
		err := testAsWrapAtTopLevel2()
		var e2 *os.PathError
		if errors.As(err, &e2) {
			if e2 != errCreatePath {
				t.Errorf("unmatch err, got:%v, want:%v", e2, errCreatePath)
			}
		} else {
			t.Errorf("unmatch As result, got:%v, want:%v", false, true)
		}
	})
	t.Run("wrapAtMiddle", func(t *testing.T) {
		err := testAsWrapAtMiddleLevel2()
		var e2 *os.PathError
		if errors.As(err, &e2) {
			if e2 != errCreatePath {
				t.Errorf("unmatch err, got:%v, want:%v", e2, errCreatePath)
			}
		} else {
			t.Errorf("unmatch As result, got:%v, want:%v", false, true)
		}
	})
	t.Run("wrapAtBottom", func(t *testing.T) {
		err := testAsWrapAtBottomLevel2()
		var e2 *os.PathError
		if errors.As(err, &e2) {
			if e2 != errCreatePath {
				t.Errorf("unmatch err, got:%v, want:%v", e2, errCreatePath)
			}
		} else {
			t.Errorf("unmatch As result, got:%v, want:%v", false, true)
		}
	})
	t.Run("wrapOnlyAtMiddle", func(t *testing.T) {
		err := testAsWrapOnlyAtMiddleLevel2()
		var e2 *os.PathError
		if errors.As(err, &e2) {
			if e2 != errCreatePath {
				t.Errorf("unmatch err, got:%v, want:%v", e2, errCreatePath)
			}
		} else {
			t.Errorf("unmatch As result, got:%v, want:%v", false, true)
		}
	})
	t.Run("wrapOnlyAtBottom", func(t *testing.T) {
		err := testAsWrapOnlyAtBottomLevel2()
		var e2 *os.PathError
		if errors.As(err, &e2) {
			if e2 != errCreatePath {
				t.Errorf("unmatch err, got:%v, want:%v", e2, errCreatePath)
			}
		} else {
			t.Errorf("unmatch As result, got:%v, want:%v", false, true)
		}
	})
	t.Run("wrapOnlyAtMiddleNoGood", func(t *testing.T) {
		err := testAsWrapOnlyAtMiddleNoGoodLevel3()
		var e2 *os.PathError
		if errors.As(err, &e2) {
			t.Errorf("unmatch As result, got:%v, want:%v", true, false)
		}
	})
	t.Run("wrapOnlyAtBottomNoGood", func(t *testing.T) {
		err := testAsWrapOnlyAtBottomNoGoodLevel2()
		var e2 *os.PathError
		if errors.As(err, &e2) {
			t.Errorf("unmatch As result, got:%v, want:%v", true, false)
		}
	})
}

var errCreatePath = &os.PathError{Op: "create", Path: "foo", Err: os.ErrExist}

func testAsWrapAtTopLevel2() error {
	return errstack.Errorf("outer: %w", testAsWrapAtTopLevel1())
}
func testAsWrapAtTopLevel1() error { return errCreatePath }

func testAsWrapAtMiddleLevel3() error {
	return errstack.Errorf("top: %w", testAsWrapAtMiddleLevel2())
}
func testAsWrapAtMiddleLevel2() error {
	return errstack.Errorf("middle: %w", testAsWrapAtMiddleLevel1())
}
func testAsWrapAtMiddleLevel1() error { return errCreatePath }

func testAsWrapAtBottomLevel2() error {
	return errstack.Errorf("outer: %w", testAsWrapAtBottomLevel1())
}
func testAsWrapAtBottomLevel1() error {
	return errstack.Errorf("bottom: %w", errCreatePath)
}

func testAsWrapOnlyAtMiddleLevel3() error {
	return fmt.Errorf("top: %w", testAsWrapOnlyAtMiddleLevel2())
}
func testAsWrapOnlyAtMiddleLevel2() error {
	return errstack.Errorf("middle: %w", testAsWrapOnlyAtMiddleLevel1())
}
func testAsWrapOnlyAtMiddleLevel1() error { return errCreatePath }

func testAsWrapOnlyAtBottomLevel2() error {
	return fmt.Errorf("outer: %w", testAsWrapOnlyAtBottomLevel1())
}
func testAsWrapOnlyAtBottomLevel1() error {
	return errstack.Errorf("bottom: %w", errCreatePath)
}

func testAsWrapOnlyAtBottomNoGoodLevel2() error {
	return fmt.Errorf("outer: %s", testAsWrapOnlyAtBottomNoGoodLevel1())
}
func testAsWrapOnlyAtBottomNoGoodLevel1() error {
	return errstack.Errorf("bottom: %w", errCreatePath)
}

func testAsWrapOnlyAtMiddleNoGoodLevel3() error {
	return fmt.Errorf("top: %s", testAsWrapOnlyAtMiddleNoGoodLevel2())
}
func testAsWrapOnlyAtMiddleNoGoodLevel2() error {
	return errstack.Errorf("middle: %w", testAsWrapOnlyAtMiddleNoGoodLevel1())
}
func testAsWrapOnlyAtMiddleNoGoodLevel1() error {
	return errCreatePath
}

func testStackFrameNames(t *testing.T, frames []errstack.Frame, names []string) {
	if len(frames) < len(names) {
		t.Fatalf("stack depth too shallow, got=%d, want=%d", len(frames), len(names))
	}
	for i := 0; i < len(names); i++ {
		if got, want := frames[i].Name, names[i]; got != want {
			t.Errorf("unmatch frames[%d].Name, got:%s, want:%s", i, got, want)
		}
	}
}
