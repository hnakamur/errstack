package errstack_test

import (
	"errors"
	"fmt"
	"os"

	"github.com/hnakamur/errstack"
)

func Example_as() {
	level1 := func() error {
		return &os.PathError{Op: "create", Path: "/tmp/foo", Err: os.ErrExist}
	}
	level2 := func() error { return errstack.Errorf("level2: %w", level1()) }
	level3 := func() error { return fmt.Errorf("level3: %w", level2()) }
	err := level3()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		var e2 *os.PathError
		if errors.As(err, &e2) {
			fmt.Printf("pathError: %v\n", e2)
		}
		fmt.Printf("stack: %v\n", errstack.Stack(err))
	}
}
