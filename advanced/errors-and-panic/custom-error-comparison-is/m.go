package main

import (
	"errors"
)

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string { return s.msg }

var ErrNotFound = errors.New("not found")

func main() {
	err := foo()
	if err != nil && errors.Is(err, ErrNotFound) {
		println("not found error")
		return
	}
	println("another error")
}

func foo() error {
	return ErrNotFound
}
