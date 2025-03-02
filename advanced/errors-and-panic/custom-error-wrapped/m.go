package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()
	// importante sempre usar o errors.Is para comparar erros
	// pois se o erro for encapsulado, o erro original será
	// comparado com o erro encapsulado
	// enquanto que se usarmos ==, o erro encapsulado não será comparado
	// com o erro original, fazendo com que a comparação falhe
	// ex.: err == ErrQualquer não funcionaria neste caso
	if err != nil && errors.Is(err, ErrQualquer) {
		fmt.Println("error:", err)
		return
	}
}

func foo() error {
	err := bar()
	if err != nil {
		// neste caso, o erro é encapsulado
		// e será possível verificar com errors.Is na func main.
		return fmt.Errorf("foo error: %w", err)
	}
	return nil
}

var ErrQualquer = errors.New("qualquer erro")

func bar() error {
	return ErrQualquer
}
