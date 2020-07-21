package main

import (
	"github.com/xxnmxx/goda/tensor"
	"fmt"
)

func main() {
	shape := []int{2,2,3}
	t := tensor.NewRandomTensor(shape)
	fmt.Println(t)
}
