package main

import (
	"fmt"
	"github.com/xxnmxx/goda/tensor"
)

func main() {
	dataA := []float64{0,1,2}
	dataB := []float64{0,1,2,3,4,5}
	a := tensor.NewTensor(dataA,3)
	b := tensor.NewTensor(dataB,3,2)
	dot := tensor.Dot(a,b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(dot)
	fmt.Println(dot.Inspect())

}
