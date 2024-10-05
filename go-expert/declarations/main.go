package main

const a = "Hello, World!"

// declarar múltiplas variáveis de uma vez
var (
	b bool    = true
	c int     = 10
	d string  = "any_string"
	e float64 = 1.2
)

func main() {
	a := "X" // string
	// o go infere o tipo na variável.
	println(a)
}

func x() {
}
