package main

import (
	"fmt"
	"github.com/xxnmxx/goda/tensor"
)

func main() {
	in := []float64{1,2,3,4,5,6,7,8,9,10,11,12}
	ts := tensor.NewTensor(in,4,3)
	fmt.Println(ts.Inspect())
	sub := ts.Ix(1,1)
	fmt.Println(sub)
	fmt.Printf("shape:%v\ndata:%v\n",sub.Shape,sub.Inspect())
	sl := ts.Slicing(-1,1)
	fmt.Println(sl)
	inA := in
	inB := []float64{1,2,3,4}
	a := tensor.NewTensor(inA,6,2)
	b := tensor.NewTensor(inB,2,2)
	fmt.Println(tensor.Dot(a,b))
}
