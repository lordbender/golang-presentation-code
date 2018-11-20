package main

import "golang-presentation-code/packa"
import "golang-presentation-code/packb"

func main() {
	packa.CoolPublicFunction("Look ma, no hands")

	packb.MatrixMultiplication(10000)
	packb.SquareMatrixParallel(10000)
}
