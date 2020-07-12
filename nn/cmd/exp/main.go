package main

import (
	"fmt"

	"github.com/xxnmxx/goda/nn"
)

func main() {
	//input := []float64{1, 0.5}

	w1 := [][]float64{
		[]float64{0.1, 0.3, 0.5},
		[]float64{0.2, 0.4, 0, 6},
	}
	b1 := []float64{0.1, 0.2, 0.3}
	a := nn.Sigmoid
	n1 := nn.NewNetwork()
	n1.SetWeight(w1)
	n1.SetBias(b1)
	n1.SetActivator(nn.Sigmoid)
	n2 := nn.NewNetwork()
	n2.SetAll(w1,b1,a)
	//wed1 := nn.InnerProduct(input, w1)
	//bed1 := nn.Add(wed1, b1)
	//acted := nn.Activate(bed1, nn.Sigmoid)
	fmt.Println(n1)
	fmt.Println(n2)
}
