package main

import (
	"fmt"

	"github.com/xxnmxx/goda/nn"
)

func main() {
	input := []float64{1, 0.5}

	w1 := [][]float64{
		[]float64{0.1, 0.3, 0.5},
		[]float64{0.2, 0.4, 0.6},
	}
	b1 := []float64{0.1, 0.2, 0.3}
	n1 := nn.NewNetwork()
	n1.SetAll(w1, b1, nn.Sigmoid)
	out1 := n1.Forword(input)

	w2 := [][]float64{
		[]float64{0.1, 0.4},
		[]float64{0.2, 0.5},
		[]float64{0.3, 0.6},
	}
	b2 := []float64{0.1, 0.2}

	n2 := nn.NewNetwork()
	n2.SetAll(w2, b2, nn.Sigmoid)
	out2 := n2.Forword(out1)

	w3 := [][]float64{
		[]float64{0.1,0.3},
		[]float64{0.2,0.4},
	}
	b3 := []float64{0.1,0.2}
	n3 := nn.NewNetwork()
	n3.SetAll(w3,b3,nn.Identity)
	out3 := n3.Forword(out2)
	final := nn.SoftMax(out3)
	fmt.Println(out1)
	fmt.Println(out2)
	fmt.Println(out3)
	fmt.Println(final)

	t := []float64{0.3,2.9,4.0}
	testMax := nn.SoftMax(t)
	fmt.Println(testMax)

}
