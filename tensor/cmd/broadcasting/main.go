package main

import (
	"fmt"
	"github.com/xxnmxx/goda/tensor"
)

func main() {
	inA := []float64{1,2,3,4,5,6}
	inB := []float64{1,2,3,4,5,6}
	a := tensor.NewTensor(inA,3,2)
	b := tensor.NewTensor(inB,2,3)
	dot := tensor.Dot(a,b)
	//ok := tensor.BroadCasting(a,b)
	fmt.Println(a.Inspect())
	fmt.Println(b.Inspect())
	fmt.Println(dot.Inspect())
	//fmt.Println(ok)
}
