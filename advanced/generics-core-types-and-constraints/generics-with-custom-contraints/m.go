package main

import "fmt"

type MeuTipo string

func (MeuTipo) Foo() {}

func main() {
	var mt MeuTipo = "meu tipo"
	foo(mt)
}

type MyConstraint interface {
	Foo()
}

func foo[T MyConstraint](arg T) {
	fmt.Println(arg)
}
