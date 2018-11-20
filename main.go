package main

import (
	"fmt"
	"golang-presentation-code/packa"
	"golang-presentation-code/packb"
)

func main() {
	packa.CoolPublicFunction("Look ma, no hands")

	packb.MatrixMultiplication(10000)

	fmt.Printf("\n-------------------------------------------------------------------------------\n")

	// P Versus S Example
	packb.MatrixSquare(10000)
	packb.SquareMatrixParallel(10000)

	fmt.Printf("\n-------------------------------------------------------------------------------\n")
	packb.MatrixSquare(20000)
	packb.SquareMatrixParallel(20000)

	fmt.Printf("\n-------------------------------------------------------------------------------\n")
	// packb.MatrixSquare(30000)
	// packb.SquareMatrixParallel(30000)
}
