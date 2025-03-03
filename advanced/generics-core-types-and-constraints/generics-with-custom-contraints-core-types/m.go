package main

import (
	"fmt"
)

func main() {
	foo(123)
	foo("")
	foo([]int{})

	// neste caso, o tipo MeuTipo não implementa a interface MyConstraint
	// pois MeuTipo é um tipo básico (string)
	// logo, não é possível usar MeuTipo como argumento
	// nesse caso, para que a interface aceite um CORE TYPE, é necessário
	// adicionar o caractere ~ antes do tipo
	// por ex.: ~string
	var mt MeuTipo = ""
	foo(mt)
}

type MeuTipo string

type MyConstraint interface {
	// ~string é um CORE TYPE
	// ou seja, aceita tipos básicos
	// e tipos que implementam a interface Foo
	int | ~string | []int

	// também poderíamos usar um pacote padrão da linguagem
	// para auxiliar em casos como este, se trata do pacote constraints.
	/* ex.:

	constraints.Integer
	constraints.Complex
	constraints.Float

	ex.: de implementação do constraints.Signed:

	type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
	}

	*/

}

func foo[T MyConstraint](arg T) {
	fmt.Println(arg)
}
