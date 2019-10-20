errstack
========

A Go library to attach call stack frames to errors.

## Warning

* This project is open source, but closed development, no support, no pull request welcome. If you are unsatisfied, feel free to fork it, pick another library, or roll your own.
* I understand this may be a "Don't do that" style in the Go error handling best practices.
* I don't promise the future compatibility. Also this library may not work in the future versions of Go, and a migration path may not be provided, leading to the dead end, users of this library will have to rewrite all of your code. You are on your own with no help.

## How to install

```
go get -u github.com/hnakamur/errstack
```

## Usage

```
func ExampleStack() {
	level1 := func() error { return os.ErrNotExist }
	level2 := func() error {
		// NOTE: You can use "%s" here just to get call stack frames.
		// However you cannot use errors.Is() or errors.As().
		return errstack.Errorf("level2: %s", level1())
	}
	level3 := func() error { return fmt.Errorf("level3: %w", level2()) }
	err := level3()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		fmt.Printf("is NotExist: %v\n", errors.Is(err, os.ErrNotExist))
		fmt.Printf("stack: %v\n", errstack.Stack(err))
	}
}

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
```

## License

MIT License
