package main

import (
	"fmt"
	"github.com/xxnmxx/goda/tensor"
)

func main() {
	inA := []float64{1,2,3}
	inB := []float64{2}
	a := tensor.NewTensor(inA,3)
	b := tensor.NewTensor(inB,1)
	ok := tensor.Div(a,b)
	fmt.Println(a.Inspect())
	fmt.Println(b.Inspect())
	fmt.Println(ok.Inspect())

}
