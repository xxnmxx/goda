package tensor

import (
	"testing"
)


//var z1 *Tensor = NewZeros(4096,4096)
//var z2 *Tensor = NewZeros(4096,4096)
var m1 *Tensor = NewZeros(20000,20000)
var m2 *Tensor = NewZeros(20000,20000)


func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	NewZeros(20000,20000)
}

//func BenchmarkRand(b *testing.B) {
//	b.ResetTimer()
//	m1.Randomize()
//}

func BenchmarkMul(b *testing.B) {
	b.ResetTimer()
	Mul(m1,m2)
}
