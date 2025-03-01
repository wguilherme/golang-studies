package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}

func (*Dog) Sound() string {
	fmt.Println("woof")
	return "woof"
}

type Cat struct{}

func (*Cat) Sound() string {
	fmt.Println("meow")
	return "meow"
}

func takeAnimal(a Animal) {
	// descobrir o tipo concreto entre vários tipos
	// comparar com ponteiro, pois os métodos são definidos com pointer receivers
	switch t := a.(type) {
	case *Dog:
		t.Sound()
	case *Cat:
		t.Sound()
	default:
		fmt.Println("unknown animal")
	}
}

func main() {

	dog := Dog{}
	cat := Cat{}

	takeAnimal(&dog)
	takeAnimal(&cat)

}
