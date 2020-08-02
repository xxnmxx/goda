package tensor

import (
	"fmt"
	"log"
)

// Dot Product
// Matrix:
// dot(a, b)[i,m] = sum(a[i,:] * b[:,m])
// Nd array:
// dot(a, b)[i,j,k,m] = sum(a[i,j,:] * b[k,:,m])
func Dot(a, b *Tensor) *Tensor {
	//if !checkDim(a, b) {
	//	log.Fatal("Error on Dot:dimension does not match")
	//}
	if !checkMat(a, b) {
		log.Fatal("Error on Dot:dimension over flow(d > 2)")
	}
	if !checkShape(a, b) {
		log.Fatal("Error on Dot:shape does not match")
	}
	// For 1D x 2D
	var shape []int
	if len(a.Shape) < 2 {
		shape = []int{b.Shape[len(b.Shape)-1]}
	} else {
		shape = []int{a.Shape[0], b.Shape[len(b.Shape)-1]}
	}
	dot := NewZeros(shape...)
	// recursive like method.
	// indexing by modulo(%).
	//   i-->
	// j a0b0 a1b2 a2b4
	// | a0b1 a1b3 a2b5
	// v a3b0 a4b2 a5b4
	//   a3b1 a4b2 a5b5
	//    .    .    .
	//    .    .    .
	for i := 0; i < a.Shape[len(a.Shape)-1]; i++ {
		posA := 0
		for j := 0; j < len(dot.Data); j++ {
			posB := j % dot.Shape[len(dot.Shape)-1]
			dot.Data[j] += a.Data[posA+i*a.Stride[len(a.Stride)-1]] * b.Data[posB+i*b.Stride[0]]
			if (j+1)%dot.Shape[0] == 0 {
				posA += a.Stride[0]
			}
		}
	}
	return dot
}

// Check dimension for dot production.
func checkDim(a, b *Tensor) bool {
	return len(a.Shape) == len(b.Shape)
}

// Check tensors are 2d or less d.
func checkMat(a, b *Tensor) bool {
	return len(a.Shape) <= 2 && len(b.Shape) <= 2
}

// Check shape for dot production.
func checkShape(a, b *Tensor) bool {
	return a.Shape[len(a.Shape)-1] == b.Shape[0]
}

// wip should implement overflow error.
// Ix select sublist.
func (ts *Tensor) Ix(ix ...int) *Tensor {
	dif := len(ts.Shape) - len(ix)
	shape := ts.subShape(ix, dif)
	return &Tensor{
		Data:   ts.subData(ix, dif),
		Shape:  shape,
		Stride: stride(shape),
	}
}

func (ts *Tensor) subData(ix []int, dif int) []float64 {
	start, end := ts.getStartEnd(ix, dif)
	if dif == 0 {
		return []float64{ts.Data[start]}
	}
	return ts.Data[start:end]
}

func (ts *Tensor) getStartEnd(ix []int, dif int) (start, end int) {
	for i := range ix {
		start += ix[i] * ts.Stride[i]
	}
	for i := len(ts.Shape) - dif; i < len(ts.Shape); i++ {
		if i == len(ts.Shape)-dif {
			end += ts.Shape[i]
		} else {
			end *= ts.Shape[i]
		}
	}
	end += start
	return start, end
}

func (ts *Tensor) subShape(ix []int, dif int) []int {
	return ts.Shape[len(ts.Shape)-dif:]
}

// Slicing returns vector.
// axis must be one and invoked -1.
func (ts *Tensor) Slicing(ix ...int) []float64 {
	axs, count := lookUpAxis(ix)
	out := make([]float64, ts.Shape[axs])
	start := 0
	for i, _ := range ts.Shape {
		if i != axs {
			start += ix[i] * ts.Stride[i]
		}
	}
	if count == 1 {
		for i := 0; i < ts.Shape[axs]; i++ {
			out[i] = ts.Data[start]
			start += ts.Stride[axs]
		}
	} else {
		fmt.Println("too many axises")
		return nil
	}
	return out
}

func lookUpAxis(ix []int) (axs, count int) {
	axs = 0
	count = 0
	for i, v := range ix {
		if v == -1 {
			axs += i
			count++
		}
	}
	return axs, count
}

// Reshape reshapes the tensor.
func (ts *Tensor) Reshape(shape ...int) bool {
	if length(shape) != ts.Len() {
		fmt.Println("shape unmatch")
		return false
	}
	ts.Shape = shape
	return true
}

// BroadCasting wip
// Add,Multi ...
func Add(a, b *Tensor) *Tensor {
	shape, ok := okBroadCasting(a, b)
	if !ok {
		log.Fatalf("could not be broadcast.\na:\n%v\nb:\n%v\n", a.Shape, b.Shape)
	}
	return NewTensor(bcAdd(shape,a,b),shape...)
}

func Mul(a, b *Tensor) *Tensor {
	shape, ok := okBroadCasting(a, b)
	if !ok {
		log.Fatalf("could not be broadcast.\na:\n%v\nb:\n%v\n", a.Shape, b.Shape)
	}
	return NewTensor(bcMul(shape,a,b),shape...)
}

func Div(a, b *Tensor) *Tensor {
	shape, ok := okBroadCasting(a, b)
	if !ok {
		log.Fatalf("could not be broadcast.\na:\n%v\nb:\n%v\n", a.Shape, b.Shape)
	}
	return NewTensor(bcDiv(shape,a,b),shape...)
}

// wip
func bcAdd(shape []int, a,b *Tensor) []float64 {
	bc := make([]float64,length(shape))
	if b.Shape[len(b.Shape)-1] == 1 && len(b.Data) != 1{
		for i := range a.Data {
			stretch := a.Shape[len(a.Shape)-1]
			rep := len(b.Data)
			bc[i] = a.Data[i] + b.Data[i/stretch%rep]
		}
		return bc
	}
	for i := range a.Data {
		bc[i] = a.Data[i] + b.Data[i%len(b.Data)]
	}
	return bc
}

func bcMul(shape []int, a,b *Tensor) []float64 {
	bc := make([]float64,length(shape))
	if b.Shape[len(b.Shape)-1] == 1 && len(b.Data) != 1{
		for i := range a.Data {
			stretch := a.Shape[len(a.Shape)-1]
			rep := len(b.Data)
			bc[i] = a.Data[i] + b.Data[i/stretch%rep]
		}
		return bc
	}
	for i := range a.Data {
		bc[i] = a.Data[i] * b.Data[i%len(b.Data)]
	}
	return bc
}

func bcDiv(shape []int, a,b *Tensor) []float64 {
	bc := make([]float64,length(shape))
	if b.Shape[len(b.Shape)-1] == 1 && len(b.Data) != 1{
		for i := range a.Data {
			stretch := a.Shape[len(a.Shape)-1]
			rep := len(b.Data)
			bc[i] = a.Data[i] + b.Data[i/stretch%rep]
		}
		return bc
	}
	for i := range a.Data {
		bc[i] = a.Data[i] / b.Data[i%len(b.Data)]
	}
	return bc
}

func okBroadCasting(a, b *Tensor) ([]int, bool) {
	if len(a.Shape) == len(b.Shape) {
		return compSameDim(a, b)
	}
	return compDifDim(a, b)
}

func compSameDim(a, b *Tensor) ([]int, bool) {
	shape := make([]int, len(a.Shape))
	for i := range a.Shape {
		switch {
		case a.Shape[i] == b.Shape[i]:
			shape[i] = a.Shape[i]
		case a.Shape[i] == 1:
			shape[i] = b.Shape[i]
		case b.Shape[i] == 1:
			shape[i] = a.Shape[i]
		default:
			return nil, false
		}
	}
	return shape, true
}

func compDifDim(a, b *Tensor) ([]int, bool) {
	shape := make([]int, len(a.Shape))
	dif := len(a.Shape) - len(b.Shape)
	if dif < 0 {
		fmt.Println("a must have more or same Dim than that b have")
		return nil, false
	}
	for i := range a.Shape {
		switch {
		case i < dif:
			shape[i] = a.Shape[i]
		case a.Shape[i] == b.Shape[i-dif]:
			shape[i] = a.Shape[i]
		case a.Shape[i] == 1:
			shape[i] = b.Shape[i-dif]
		case b.Shape[i-dif] == 1:
			shape[i] = a.Shape[i]
		default:
			return nil, false
		}
	}
	return shape, true
}

// Add
//func (ts *Tensor)
