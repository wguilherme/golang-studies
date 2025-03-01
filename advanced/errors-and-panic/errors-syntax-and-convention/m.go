package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 10
	b := 0
	res, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
