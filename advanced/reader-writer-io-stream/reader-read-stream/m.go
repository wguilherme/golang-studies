package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello, world"
	reader := strings.NewReader(str)

	// se limitarmos os bytes do slice abaixo, limitaria a quantidade de bytes lidos
	// por exemplo, se colocarmos 2 bytes, o output seria "he" e não "hello, world"
	buffer := make([]byte, 10000)
	n, err := reader.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(string(buffer[:n]))
}
