package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}

func (*Dog) Sound() string {
	// (!)  Nesse caso não dá erro pois o método não tenta acessar nenhum campo do receptor
	// (!)  Se tentássemos acessar um campo do receptor, daria erro:
	// (!)  invalid memory address or nil pointer dereference

	// Também é comum que muitos métodos em GO tenham um tratamento no receiver para evitar esse tipo de erro
	// Por exemplo, se o método Sound tentasse acessar um campo do receptor, poderíamos fazer algo assim:
	/*
			if d == nil {
			return "<nil>"
		}
	*/

	fmt.Println("woof")
	return "woof"
}

func main() {
	var a Animal
	var dog *Dog
	fmt.Println(a == nil)
	fmt.Println(dog == nil)
	a = dog
	fmt.Println(a == nil)
	a.Sound()

	/*
	 se chamasse a.Sound() diretamente de uma interface nula, ai sim daria erro
	 daria nil pointer dereference
	*/

}
