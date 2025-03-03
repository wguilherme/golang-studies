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
	writer := MyWriter{}
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
		writer.Write(buffer[:n])
	}
}

type MyWriter struct{}

func (MyWriter) Write(b []byte) (int, error) {
	// neste caso, estamos escrevendo os bytes lidos
	// usando o fmt.Print ao invés do fmt.Println
	// para que a saída seja na mesma linha
	fmt.Print(string(b))
	return len(b), nil
}
