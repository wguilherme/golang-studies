package main

import (
	"errors"
)

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string { return s.msg }

// se o erro tem informações adicionais, é melhor criar um tipo de erro customizado
// para que possamos comparar o erro com o tipo de erro customizado
// se o erro é genérico para todos erros, podemos usar o errors.Is para analisar
// errors criados à partir do errors.New

func main() {
	err := foo()
	// se a func de erro fosse um pointer receiver, aqui deveríamos passar o ponteiro para o sqrtError
	// ex.: var sqrtError *SqrtError
	var sqrtError SqrtError
	// Neste caso precisamos passar o ponteiro pois o AS realiza mutação no valor
	// atualizando a variável sqrtError com o valor do erro
	// e sempre vamos chamar o pinteiro independente de já ser um ponteiro ou não
	// caso ela seja, teremos um cenário de ponteiro de ponteiro (double pointer)
	// ex.: &sqrtError da var *SqrtError que já é uma dereference
	// este cenário seria se a func de erro tivesse um pointer receiver por ex.
	if err != nil && errors.As(err, &sqrtError) {
		println(sqrtError.msg)
		return
	}
	println("another error")
}

func foo() error {
	// se a func de erro fosse um pointer receiver, aqui deveríamos retornar um ponteiro para o sqrtError
	// ex.: return &SqrtError{msg: "sqrt error message"}
	return SqrtError{msg: "sqrt error message"}
}
