package packb

import (
	"errors"
	"fmt"
	"golang-presentation-code/core"
	"time"
)

// MatrixMultiplication illustrates one way to perform
// Ascnc Behaviors in GOLANG
func MatrixMultiplication(size int) {
	a := core.GetMatrix(size)
	b := core.GetMatrix(size)

	matrixMultiplicationStart := time.Now()
	c, err := dot(a, transpose(b))
	matrixMultiplicationElapsed := time.Since(matrixMultiplicationStart)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Time to Multiply Matrix %s on length %d\n\n", matrixMultiplicationElapsed, len(c[0]))
}

func transpose(x [][]int) [][]int {
	out := make([][]int, len(x[0]))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func dot(x, y [][]int) ([][]int, error) {
	if len(x[0]) != len(y) {
		return nil, errors.New("ERROR: dx/dy does not support matrix multiplication.")
	}

	out := make([][]int, len(x))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			if len(out[i]) < 1 {
				out[i] = make([]int, len(y))
			}
			out[i][j] += x[i][j] * y[j][i]
		}
	}

	return out, nil
}
