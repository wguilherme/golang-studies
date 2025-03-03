package main

import "fmt"

func main() {
	foo("")
	foo(123)

	// neste caso, []int não é comparável (X === Y por exemplo)
	// logo, não é possível usar []int como argumento
	// para a função genérica foo
	// foo([]int{1})

}

func foo[T comparable](arg T) {
	fmt.Println(arg)
}
