package errstack_test

import (
	"errors"
	"fmt"
	"os"

	"github.com/hnakamur/errstack"
)

func Example_is() {
	level1 := func() error { return os.ErrNotExist }
	level2 := func() error { return errstack.Errorf("level2: %w", level1()) }
	level3 := func() error { return fmt.Errorf("level3: %w", level2()) }
	err := level3()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		fmt.Printf("is NotExist: %v\n", errors.Is(err, os.ErrNotExist))
		fmt.Printf("stack: %v\n", errstack.Stack(err))
	}
}
