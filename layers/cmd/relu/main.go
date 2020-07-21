package main

import (
	"github.com/xxnmxx/goda/layers"
	//"github.com/xxnmxx/goda/nn"
	"fmt"
)

func main() {
	// reru
	//rl := nn.NewRelu()
	//x := []float64{1,-1,0,11}
	//y := rl.Forward(x)
	//dx := rl.Backward(y)
	//fmt.Println("Relu:")
	//fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(dx)
	// Sig
	sl := layers.NewSigmoid()
	x := []float64{1,-1,0,11,50}
	y := sl.Forward(x)
	dx := sl.Backward(y)
	fmt.Println("Sigmoid:")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(dx)

}
