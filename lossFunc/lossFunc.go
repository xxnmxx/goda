package lossFunc

import (
	"log"
	"math"
)

func MeanSquaredError(y, t []float64) float64 {
	if len(y) != len(t) {
		log.Fatalf("length unmatch: y=%v,t=%v", len(y), len(t))
	}
	sm := 0.0
	for i := range y {
		sm += math.Pow((y[i] - t[i]), 2)
	}
	return 0.5 * sm
}

func CrossEntropyError(y,t []float64) float64 {
	delta := 1e-7
	if len(y) != len(t) {
		log.Fatalf("length unmatch: y=%v,t=%v", len(y), len(t))
	}
	ce := 0.0
	for i := range y {
		ce += t[i] * math.Log(y[i] + delta)
	}
	return -ce
}
