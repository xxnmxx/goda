package nn

import (
	"fmt"
	"log"
	"math"
)

// Network
type Network struct {
	Weights   [][]float64
	Baiases   []float64
	Activator activator
}

func NewNetwork() *Network {
	return &Network{}
}

// Method
// Setter
func (n *Network) SetAll(ws [][]float64, bs []float64, a activator) {
	n.Weights, n.Baiases, n.Activator = ws, bs, a
}

func (n *Network) SetWeight(ws [][]float64) {
	n.Weights = ws
}

func (n *Network) SetBias(bs []float64) {
	n.Baiases = bs
}

func (n *Network) SetActivator(a activator) {
	n.Activator = a
}

// Checker
// FixMe need to implement shape chacker etc...

// Movement
func (n *Network) Forword(ins []float64) []float64 {
	if n.Weights == nil {
		fmt.Println("Weights are nil.")
		return nil
	}
	if n.Baiases == nil {
		fmt.Println("Baiases are nil.")
		return nil
	}
	if n.Activator == nil {
		fmt.Println("Activator is nil.")
		return nil
	}
	iped := InnerProduct(ins,n.Weights)
	added := Add(iped,n.Baiases)
	outs := Activate(added,n.Activator)
	return outs
}

// Activators not use now.
type activator func(float64) float64

func Sigm0id(in float64) float64 {
	return 1 / (1 + math.Pow(math.E, -in))
}

func ReLU(in float64) float64 {
	if in > 0 {
		return in
	}
	return 0
}

func Identity(in float64) float64 {
	return in
}

// Outputter
type outputter func([]float64) []float64

func SoftMax(ins []float64) []float64 {
	max := 0.0
	for _, in := range ins {
		max = math.Max(max,in)
	}
	sum := 0.0
	for _, in := range ins {
		sum += math.Exp(in-max)
	}
	outs := make([]float64,len(ins))
	for i, in := range ins {
		outs[i] = math.Exp(in-max) / sum
	}
	return outs
}

// Matrix Operation
func InnerProduct(ins []float64, ws [][]float64) []float64 {
	// Check the shape
	if len(ins) != len(ws) {
		log.Fatalf("shape unmatched:ins %v, ws [%v:%v]\n", len(ins), len(ws), len(ws[0]))
	}
	outs := make([]float64, len(ws[0]))
	for i, in := range ins {
		for j, w := range ws[i] {
			outs[j] += in * w
		}
	}
	return outs
}

func Add(ins []float64, add []float64) []float64 {
	for i := range ins {
		ins[i] += add[i]
	}
	return ins
}

func Transpose(ins [][]float64) [][]float64 {
	outs := NewMatrix(len(ins[0]),len(ins))
	for i, in := range ins {
		for j, v := range in {
			outs[j][i] = v
		}
	}
	return outs
}

func NewMatrix(r,c int) [][]float64 {
	outs := make([][]float64,r)
	for i := 0; i < r; i++ {
		outs[i] = make([]float64,c)
	}
	return outs
}

func SumRows(ins [][]float64) []float64 {
	outs := make([]float64,len(ins))
	for i, in := range ins {
		for _, v := range in {
			outs[i] += v
		}
	}
	return outs
}

func SumCols(ins [][]float64) []float64 {
	outs := make([]float64,len(ins[0]))
	insT := Transpose(ins)
	for i, in := range insT {
		for _, v := range in {
			outs[i] += v
		}
	}
	return outs
}

func Sum(ins []float64) float64 {
	out := 0.0
	for _, v := range ins {
		out += v
	}
	return out
}

func Activate(ins []float64, f activator) []float64 {
	out := make([]float64, len(ins))
	for i, in := range ins {
		out[i] = f(in)
	}
	return out
}
