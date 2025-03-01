package main

import (
	"fmt"
	"math"
)

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string { return s.msg }

func raizQuadrada(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{"não existe raiz quadrada de numero negativo"}
	}
	resultado := math.Sqrt(x)
	return resultado, nil
}

func main() {
	x := -10
	res, err := raizQuadrada(float64(x))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
