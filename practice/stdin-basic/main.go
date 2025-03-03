package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	para testar no MAC:
	go run main.go
	hello    (Enter)
	world    (Enter)
	[Ctrl+D]
*/

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// Guarda as linhas em uma slice para referência
	var lines []string
	for input.Scan() {
		text := input.Text()
		lines = append(lines, text)
		counts[text]++
	}

	fmt.Println("Lines read:", lines)
	fmt.Println("Counts:", counts)
}
