package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	str := "hello, world\n"
	reader := strings.NewReader(str)
	buffer := make([]byte, 2)
	// para detectarmos o fim do stream, observaremos o erro EOF
	// que é retornado quando não há mais bytes para serem lidos
	// EOF: end of file
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(n)
		fmt.Println(buffer[:n])
	}
}
