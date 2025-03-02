package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()
	fmt.Println(err)
	// neste caso, os erros estão agrupados
	// através da função errors.Join
	// e poddem ser comparados com errors.Is
	// para verificar se o erro original está presente
	// no caso abaixo, ambos casos retornam true
	fmt.Println(errors.Is(err, ErrQualquer))
	fmt.Println(errors.Is(err, ErrQualquer2))

}

var ErrQualquer = errors.New("error")
var ErrQualquer2 = errors.New("error2")

func a() error { return ErrQualquer }
func b() error { return ErrQualquer2 }

func foo() error {

	var errorResult error

	if err := a(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}

	if err := b(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}

	return errorResult
}
