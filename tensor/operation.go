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
	itvA := a.Stride[len(a.Stride)-1]
	itvB := b.Stride[0]
	posA := 0
	posB := 0
	for i := 0; i < len(dot.Data); i++ {
		if posA > len(a.Data) || posB > len(b.Data) {
		}
		dot.Data[i] = a.Data[posA] * b.Data[posB]
		posA += itvA
		posB += itvB
	}
	return dot
}

func dotInterval(ts *Tensor) int {
	return 0
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
