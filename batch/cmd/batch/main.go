package main

import (
	"github.com/xxnmxx/goda/batch"
	"fmt"
)

func main() {
	input := []float64{1,2,3,4,5,6,7,8,9,10,11,12,3}
	batches := batch.Batch(input,2)
	fmt.Println(batches)
	batch.ShuffledBatch(input,2)
	fmt.Println(batches)
	batch.RandomBatch(input,2)
	fmt.Println(batches)
}
