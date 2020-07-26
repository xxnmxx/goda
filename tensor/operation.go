package tensor

import "fmt"

// Matrix:
// dot(a, b)[i,m] = sum(a[i,:] * b[:,m])
// Nd array:
// dot(a, b)[i,j,k,m] = sum(a[i,j,:] * b[k,:,m])
// Dot Product
func Dot(a, b *Tensor) *Tensor {
	if !checkDim(a, b) {
		fmt.Println("dimension does not match")
		return nil
	}
	if !checkMat(a, b) {
		fmt.Println("dimension over flow(d > 2)")
		return nil
	}
	if !checkShape(a, b) {
		fmt.Println("shepe does not match")
		return nil
	}
	shape := []int{a.Shape[0], b.Shape[len(b.Shape)-1]}
	dot := NewZeros(shape...)
	for i := 0; i < a.Shape[1]; i++ {
		posA := 0
		for j := 0; j < len(dot.Data); j++ {
			posB := j % dot.Shape[1]
			dot.Data[j] += a.Data[posA + i*dot.Stride[1]] * b.Data[posB+i*dot.Stride[0]]
			if j % dot.Shape[1] != 0 {
				posA += a.Stride[0]
			}
		}
	}
	return dot
}

// WIP should implement overflow error.
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

// Check dimension for dot production.
func checkDim(a, b *Tensor) bool {
	return len(a.Shape) == len(b.Shape)
}

// Check tensors are 2d or less d.
func checkMat(a, b *Tensor) bool {
	return len(a.Shape) <= 2 && len(b.Shape) <= 2
}

// Check shepe for dot production.
func checkShape(a, b *Tensor) bool {
	return a.Shape[len(a.Shape)-1] == b.Shape[0]
}
