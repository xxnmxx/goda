package main

import (
	//"fmt"
	"github.com/xxnmxx/goda/matrix"
)

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	r, c := 3, 4
	d := matrix.NewDenseMatrix(r, c, data)
	matrix.PrintMat(d)
}
