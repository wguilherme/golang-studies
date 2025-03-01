package main

import (
	"embedding/embedder"
)

type User struct {
	embedder.Embedder
	Name string
	ID   uint64
}

func main() {
	user := User{Name: "any_name", ID: 1}
	// Nesse caso, embedamos, o que é diferente de herança, que é o que acontece em outras linguagens
	// Aqui, o método GetTitle é promovido para o tipo User
	// Isso significa que podemos chamar o método GetTitle diretamente no tipo User
	// Isso é chamado de composição
	// Composição é uma forma de reuso de código em Go
	// Composição é mais flexível que herança
	user.GetTitle("Title")
}
