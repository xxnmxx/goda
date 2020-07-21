package tensor

import (
	"math/rand"
	"time"
)

type Tensor struct {
	Data  []float64
	Shape []int
}

//func NewTensor(data []float64, shape []float64) {}
func NewRandomTensor(shape []int) *Tensor {
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	rng := 0
	for _, v := range shape {
		if rng == 0 {
			rng += v
		} else {
			rng *= v
		}
	}
	out := make([]float64, rng)
	for i := 0; i < rng; i++ {
		out[i] = r.NormFloat64()
	}
	return &Tensor{
		Data:  out,
		Shape: shape,
	}
}

func (t *Tensor) Print() {
	var b strings.Builder
	dim := len(t.Shape)
	b.WriteString(strings.Repeat("[",len(t.Shape)))
	for i := len(t.Shape); i > 0; i-- {
		for i := // do something
		b.WriteString()
	}
}
