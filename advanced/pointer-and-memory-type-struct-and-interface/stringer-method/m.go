package main

import "fmt"

type EmptyStruct struct{}

func (EmptyStruct) String() string {
	return "empty struct"
}

func main() {
	var e EmptyStruct
	fmt.Println(e)
}
