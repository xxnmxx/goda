package batch

import (
	"math/rand"
	"time"
)

func ShuffledBatch(src []float64, batchSize int) [][]float64 {
	rand.Shuffle(len(src), func(i, j int) {
		src[i], src[j] = src[j], src[i]
	})
	return Batch(src, batchSize)
}

func Batch(act []float64, batchSize int) [][]float64 {
	batches := make([][]float64, 0, (len(act)+batchSize-1)/batchSize)
	for batchSize < len(act) {
		act, batches = act[batchSize:], append(batches, act[0:batchSize:batchSize])
	}
	batches = append(batches, act)
	return batches
}

func RandomBatch(src []float64, batchSize int) [][]float64 {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := len(src) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		src[i], src[j] = src[j], src[i]
	}
	return Batch(src, batchSize)
}
