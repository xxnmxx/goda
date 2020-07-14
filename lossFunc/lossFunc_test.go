package lossFunc

import (
	"fmt"
	"testing"
)

func TestMSE(t *testing.T) {
	y := []float64{0.1, 0.05, 0.6, 0.0, 0.05, 0.1, 0.0, 0.1, 0.0, 0.0}
	tr := []float64{0, 0, 1, 0, 0, 0, 0, 0, 0, 0}
	fmt.Printf("MeanSquaredError:%v\n",MeanSquaredError(y, tr))
	if MeanSquaredError(y, tr) != 0.09750000000000003 {
		t.Error(MeanSquaredError(y, tr))
	}
}

func TestCEE(t *testing.T) {
	y := []float64{0.1, 0.05, 0.6, 0.0, 0.05, 0.1, 0.0, 0.1, 0.0, 0.0}
	tr := []float64{0, 0, 1, 0, 0, 0, 0, 0, 0, 0}
	fmt.Printf("CrossEntropyError:%v\n", CrossEntropyError(y, tr) )
	if CrossEntropyError(y, tr) != 0.510825457099338 {
		t.Error(CrossEntropyError(y, tr))
	}
}
