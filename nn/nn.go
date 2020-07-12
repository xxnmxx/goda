package nn

import (
	"log"
	"math"
)

// Network
type Network struct {
	W [][]float64
	B []float64
	A activator
}

func NewNetwork() *Network {
	return &Network{}
}

// Method
// Setter
func (n *Network) SetAll(ws [][]float64, bs []float64, a activator) {
	n.W, n.B, n.A = ws, bs, a
}

func (n *Network) SetWeight(ws [][]float64) {
	n.W = ws
}

func (n *Network) SetBias(bs []float64) {
	n.B = bs
}

func (n *Network) SetActivator(a activator) {
	n.A = a
}

// Checker
func (n *Network) checkShape(ins []float64) bool {
	if len(ins) != n.W {
		fmt.Printf("shape unmatched: ins: %v, ws: [%v,%v]",len(ins),len(n.W),len(n.W[0]))
	}
}

// Movement
func (n *Network) Forword(ins []float64) *Netword {

}

// Activators
type activator func(float64) float64
func Sigmoid(in float64) float64 {
	return 1 / (1 + math.Pow(math.E, -in))
}

func ReLU(in float64) float64 {
	if in > 0 {
		return in
	}
	return 0
}

// Array Operation
func InnerProduct(ins []float64, ws [][]float64) []float64 {
	// Check the shape
	if len(ins) != len(ws) {
		log.Fatalf("shape unmatched:ins %v, ws %v,%v\n",len(ins),len(ws),len(ws[0]))
	}
	outs := make([]float64, len(ws[0]))
	for i, in := range ins {
		for j, w := range ws[i] {
			outs[j] += in*w
		}
	}
	return outs
}

//func Add(ins []float64, add []float64) []float64 {
//	for i, in := range ins {
//		ins[i] += add
//	}
//	return outs
//}

func Activate(ins []float64, f activator) []float64 {
	out := make([]float64,len(ins))
	for i, in := range ins {
		out[i] = f(in)
	}
	return out
}

