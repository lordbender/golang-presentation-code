package packb

import (
	"errors"
	"fmt"
	"golang-presentation-code/core"
	"time"
)

// MatrixSquare sequentially squares a matrix
func MatrixSquare(size int) [][]int {
	a := core.GetMatrix(size)
	c := make([][]int, size)
	for i := 0; i < size; i++ {
		c[i] = make([]int, size)
	}

	matrixSquareStart := time.Now()
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			c[i][j] = a[i][j] ^ 2
		}
	}
	matrixSquareElapsed := time.Since(matrixSquareStart)

	fmt.Printf("\tSequential Time to Square Matrix, A: O(n^2) -> %s on %d\n", matrixSquareElapsed, len(a[0]))

	return c
}

// MatrixMultiplication illustrates one way to perform
// Ascnc Behaviors in GOLANG
func MatrixMultiplication(size int) {
	a := core.GetMatrix(size)
	b := core.GetMatrix(size)

	matrixMultiplicationStart := time.Now()
	_, err := dot(a, transpose(b))
	matrixMultiplicationElapsed := time.Since(matrixMultiplicationStart)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\tSequential Time to Multiply Matrix, A x B = C: O(n^3) -> %s on %d x %d matrix\n", matrixMultiplicationElapsed, size, size)
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
