package main

import (
	"fmt"
	"github.com/xxnmxx/goda/tensor"
)

func main() {
	z := tensor.NewZeros(4092,4092)
	z.Randomize()
	mul := tensor.Mul(z,z)
	fmt.Println(mul)
}
