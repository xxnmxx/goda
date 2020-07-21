package main

import (
	"fmt"
	"github.com/xxnmxx/goda/nn"
	"github.com/xxnmxx/goda/layers"
)

func main() {
	al := layers.NewAffine(3,5)
	fmt.Println(al)
}
