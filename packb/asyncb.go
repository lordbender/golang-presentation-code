package packb

import (
	"fmt"
	"golang-presentation-code/core"
	"sync"
	"time"
)

type rowHelper struct {
	row      []int
	position int
}

// SquareMatrixParallel takes a matrix and returns
// a matrix of the same size, but with every
// index squared.
//
// bit overflow possible when squaring 32bit integers: be careful.
func SquareMatrixParallel(size int) [][]int {

	a := core.GetMatrix(size)
	c := core.GetMatrix(size)

	matrixSquareStart := time.Now()
	respond := make(chan rowHelper, size)
	var wg sync.WaitGroup
	wg.Add(size)

	for i := 0; i < size; i++ {
		go squareIt(respond, &wg, rowHelper{row: a[i], position: i})
	}

	wg.Wait()
	close(respond)

	for queryResp := range respond {
		c[queryResp.position] = queryResp.row
	}
	matrixSquareElapsed := time.Since(matrixSquareStart)

	fmt.Printf("Time to Multiply Matrix %s on length %d\n\n", matrixSquareElapsed, len(c[0]))
	return c
}

func squareIt(respond chan<- rowHelper, wg *sync.WaitGroup, a rowHelper) {
	defer wg.Done()
	size := len(a.row)
	c := make([]int, size)
	for j := 0; j < size; j++ {
		c[j] = a.row[j] ^ 2
	}
	respond <- rowHelper{row: c, position: a.position}
}
