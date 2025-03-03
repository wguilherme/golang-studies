package main

func main() {
	wrongWayFoo(Foo{})
	wrongWayFoo(Bar{})
	rightWayFoo(Foo{})
	rightWayFoo(Bar{})
}

type Foo struct{}

func (Foo) Do() {}

type Bar struct{}

func (Bar) Do() {}

func wrongWayFoo[T Foo | Bar](arg T) {
	// neste caso, não é possível chamar o método Do
	// se trata de uma limitação dos genéricos em Go
	// o mesmo é válido para cammpos das structs
	// arg.Do()
}

// neste caso, o correto seria implementar uma interface que
// contenha o método Do e embedda os tipos Foo e Bar
// para que seja possível chamar o método Do.

// mas também teria um jeito de fazer isso, declarando a interface Do() diretamente
// seria algo semelhante a:
// func wrongWayFoo[T interface{ Do() }](arg T) {}
// mas isso não é uma boa prática, pois não é possível garantir que o método Do
// seja implementado corretamente em todos os tipos que implementam a interface

// sendo assim, o correto seria fazer algo semelhante a:

type MinhaInterface interface {
	// neste caso, somente os tipos Foo e Bar implementam a interface
	// se por exemplo, Bar não estivesse declarado aqui daria erro na func main
	// ao tentar chamar rightWayFoo(Bar{})
	// pois Bar não implementa a interface MinhaInterface
	Foo | Bar
	Do()
}

func rightWayFoo[T MinhaInterface](arg T) {
	arg.Do()
}
