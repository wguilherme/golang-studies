package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

func main() {
	str := "hello, world\n"
	reader := strings.NewReader(str)
	writer := MyWriter{os.Stdout}
	buffer := make([]byte, 2)
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

type MyWriter struct {
	w io.Writer
}

func (mw MyWriter) Write(b []byte) (int, error) {
	for i, bb := range b {
		b[i] = bb + 10
	}
	return mw.w.Write(b)
}
