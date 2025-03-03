package main

import "fmt"

func main() {
	foo(123)
	foo("")
	foo([]int{})
}

type MyConstraint interface {
	int | string | []int
}

func foo[T MyConstraint](arg T) {
	fmt.Println(arg)
}
