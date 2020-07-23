package layers

import (
	"math"
	"math/rand"
	"time"

	"github.com/xxnmxx/goda/nn"
)

type Layer interface {
	Forward() []float64
	Backward() []float64
}

// Relu
type Relu []float64

func NewRelu() *Relu {
	return &Relu{}
}

func (r *Relu) Forward(x []float64) []float64 {
	out := make([]float64, len(x))
	for i, v := range x {
		if v <= 0 {
			out[i] = 0
		} else {
			out[i] = v
		}
	}
	return out
}

func (r *Relu) Backward(y []float64) []float64 {
	dx := make([]float64, len(y))
	for i, v := range y {
		if v == 0 {
			dx[i] = 0
		} else {
			dx[i] = 1
		}
	}
	*r = dx
	return dx
}

// Sigmoid
type Sigmoid []float64

func NewSigmoid() *Sigmoid {
	return &Sigmoid{}
}

func (s *Sigmoid) Forward(x []float64) []float64 {
	out := make([]float64, len(x))
	for i, v := range x {
		out[i] = 1 / (1 + math.Exp(-v))
	}
	*s = out
	return out
}

func (s *Sigmoid) Backward(y []float64) []float64 {
	dx := make([]float64, len(y))
	for i, v := range y {
		dx[i] = v * (1.0 - (*s)[i]*(*s)[i])
	}
	return dx
}

// Affine
type Affine struct {
	w  [][]float64
	b  []float64
	x  []float64
	dw [][]float64
	db []float64
}

// FixMe
func NewAffine(inputSize, layerSize int) *Affine {
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	// make random weights
	w := nn.NewMatrix(inputSize, layerSize)
	for i, v := range w {
		for j, w := range v {
			w[i][j] = r.NormFloat64()
		}
	}
	// make random baiases
	b := make([]float64, inputSize)
	for i, _ := range b {
		b[i] = r.NormFloat64()
	}
	return &Affine{
		w: w,
		b: b,
	}
}

func (a *Affine) Forward(x []float64) []float64 {
	a.x = x
	iped := nn.InnerProduct(x, a.w)
	added := nn.Add(iped, a.b)
	return added
}

func (a *Affine) Backward(y []float64) [][]float64 {
	dx := nn.InnerProduct(y, nn.Transpose(a.w))
	a.dw = nn.InnerProduct(nn.Transpose(a.x), y)
	a.db = nn.Sum(y)
	return dx
}
