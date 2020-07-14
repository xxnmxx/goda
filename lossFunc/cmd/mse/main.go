package main

import (
	"fmt"

	"github.com/xxnmxx/goda/lossFunc"
)

func main() {
	y := []float64{0.1, 0.05, 0.6, 0.0, 0.05, 0.1, 0.0, 0.1, 0.0, 0.0}
	t := []float64{0, 0, 1, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(lossFunc.MeanSquaredError(y, t))
}
