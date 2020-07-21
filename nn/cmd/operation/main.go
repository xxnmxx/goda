package main

import (
	"fmt"
	"github.com/xxnmxx/goda/nn"
)

func main() {
	ins := [][]float64{
		[]float64{1,2,3},
		[]float64{4,5,6},
		[]float64{7,8,9},
		[]float64{10,11,12},
	}
	//t := nn.Transpose(ins)
	//fmt.Println(ins)
	//fmt.Println(t)
	//m := nn.NewMatrix(2,3)
	//fmt.Println(m)
	fmt.Println(nn.SumRows(ins))
	fmt.Println(nn.SumCols(ins))
}
