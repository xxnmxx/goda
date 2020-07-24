package main

import (
	"fmt"
	"github.com/xxnmxx/goda/tensor"
)

func main() {
	zeros := tensor.NewZeros(3,3)
	ones := tensor.NewOnes(3,3,2)
	fmt.Println(zeros.Inspect())
	fmt.Println(ones.Inspect())
	zeros.Randomize()
	fmt.Println(zeros.Inspect())
}
