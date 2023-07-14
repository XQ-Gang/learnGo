package main

import (
	"errors"
	"fmt"
	"io"
)

type MyError struct {
	msg string
	err error
}

func (e *MyError) Error() string {
	return e.msg + ": " + e.err.Error()
}

func (e *MyError) Unwrap() error {
	return e.err
}

func NewErrorWithCode(msg string, err error) error {
	return &MyError{
		msg: msg,
		err: err,
	}
}

// AddContext don’t just check errors, handle them gracefully – add context (wrap).
func AddContext(err error) error {
	// Bad example
	// if err != nil {
	//	return err
	// }

	// The Best example
	if err != nil {
		return fmt.Errorf("context: %w", err)
	}
	return nil
}

// CheckMatch check if error match.
func CheckMatch(err error) {
	// Bad example (Go < 1.13)
	if err == io.EOF {
		fmt.Println("err == io.EOF")
	}

	// Good example
	// error could be wrapped
	if errors.Is(err, io.EOF) {
		fmt.Println("errors.Is(err, io.EOF)")
	}
}

// Casting cast error to MyError.
func Casting(err error) {
	// Bad example (Go < 1.13)
	if e, ok := err.(*MyError); ok {
		fmt.Println("Casting v1", e)
	}

	// Good example
	var e *MyError
	if errors.As(err, &e) {
		fmt.Println("Casting v2", e)
	}
}

func main() {
	err := NewErrorWithCode("biz error", io.EOF)
	fmt.Println(err)
	err = AddContext(err)
	fmt.Println(err)
	CheckMatch(err)
	Casting(err)

	// Output:
	// biz error: EOF
	// context: biz error: EOF
	// errors.Is(err, io.EOF)
	// Casting v2 biz error: EOF
}
