package matrix

import (
	"bytes"
	"fmt"
)

type Matrix interface {
	Shape() []int
	Data() []float64
	Stride() int
}

type DenseMatrix struct {
	data   []float64
	r, c   int
	stride int
}

func NewDenseMatrix(r, c int, data []float64) *DenseMatrix {
	return &DenseMatrix{
		data:   data,
		r:      r,
		c:      c,
		stride: c,
	}
}

func (dm *DenseMatrix) Shape() []int { return []int{dm.r, dm.c} }
func (dm *DenseMatrix) Data() []float64 { return dm.data }
func (dm *DenseMatrix) Stride() int { return dm.stride }

func PrintMat(m Matrix) {
	var out bytes.Buffer
	shape := m.Shape()
	data := m.Data()
	stride := m.Stride()
	for i := 0; i <= shape[0]-1; i++ {
		out.WriteString("[")
		for j := 0; j <= shape[1]-1; j++ {
			out.WriteString(fmt.Sprint(data[j+stride*i]))
			out.WriteString("\t")
		}
		out.WriteString("]\n")
	}
	fmt.Println(out.String())
}
