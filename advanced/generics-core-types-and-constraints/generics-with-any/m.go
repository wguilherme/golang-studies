package main

import "fmt"

func main() {
	foo("")
	foo(123)
	foo([]int{1})

}

func foo[T any](arg T) {
	fmt.Println(arg)
}

/*
os genéricos em Go são executados em tempo de compilação,
o que significa que o código genérico é gerado para cada tipo de argumento que você passa para a função genérica.
Isso é diferente de linguagens como Java e C#, onde o código genérico é gerado em tempo de execução.

por exemplo, o código acima, quando compilado, gera algo semelhante a isso:

func foo1(arg string) {}
func foo2(arg int) {}
func foo3(arg []int) {}

*/
