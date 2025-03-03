package main

import (
	"fmt"
)

func main() {
	foo(123)
	foo("")
	foo([]int{})
	var mt MeuTipo = ""
	foo(mt)

	var ms MyStruct[string] = MyStruct[string]{}

	fmt.Println(ms)
}

type MeuTipo string

type MyConstraint interface {
	int | ~string | []int
}

func foo[T MyConstraint](arg T) {
	fmt.Println(arg)
}

type MyStruct[T any] struct {
	Foo T
}

// limitação, genéricos em Go não suportam métodos

// o código abaixo seria inválido. Error: method must have no types parameters
// func (MyStruct[T]) foo[A any](a A) {}

// porém, é possível fazer isso:
// um método usar o genérico da struct
// func (MyStruct[T]) foo(a T) {}
