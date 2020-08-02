package main

import (
	"github.com/xxnmxx/goda/tensor"
)

type neuralNet struct {
	config  neuranNetConfig
	wHidden *tensor.Tensor
	bHidden *tensor.Tensor
	wOut    *tensor.Tensor
	bOut    *tensor.Tensor
}
