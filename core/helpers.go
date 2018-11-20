package core

import (
	"github.com/alok87/goutils/pkg/random"
)

// GetMatrix creates a 2d matrix.
func GetMatrix(size int) [][]int {
	a := make([][]int, size)

	for i := 0; i < size; i++ {
		a[i] = random.RangeInt(2, 100, size)
	}

	return a
}
