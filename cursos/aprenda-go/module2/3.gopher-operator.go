package main

import "fmt"

func main() {

	x := 10
	y := "good morning"

	fmt.Printf("x: %v, %T", x, x)
	fmt.Printf("\n")
	fmt.Printf("y: %v, %T", y, y)
	
}

// short operator := is used to declare and assign a value to a variable
// need to declare at least one variable with :=
// can't use := to redeclare a variable
// can't use := to declare a variable without assigning a value
// short operator is block scoped