package main

import (
	"github.com/xxnmxx/goda/tensor"
)

type neuralNet struct {
	config  neuralNetConfig
	wHidden *tensor.Tensor
	bHidden *tensor.Tensor
	wOut    *tensor.Tensor
	bOut    *tensor.Tensor
}

type neuralNetConfig struct {
	inputNeurons  int
	outputNeurons int
	hiddenNeurons int
	epochs        int
	learningRate  float64
}

func main() {
	// do something
}

func newNetwork(config *neuralNetConfig) *neuralNet {
	return &neuralNet{config: config}
}

func (nn *neuralNet) train(x, y *tensor.Tensor) error {
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)
	wHidden := tensor.NewZeros(nn.config.inputNeurons, nn.config.hiddenNeurons)
	bHidden := tensor.NewZeros(1, nn.config.hiddenNeurons)
	wOut := tensor.NewZeros(nn.config.hiddenNeurons, nn.config.outputNeurons)
	bOut := tensor.NewZeros(1, nn.config.outputNeurons)
	wHiddenRaw := wHidden.Data
	bHiddenRaw := bHidden.Data
	wOutRaw := wOut.Data
	bOutRaw := bOut.Data

	for _, param := range [][]float64{
		wHiddenRaw,
		bHiddenRaw,
		wOutRaw,
		bOutRaw,
	} {
		for i := range param {
			param[i] = randGen.Float64()
		}
	}
	output := new(tensor.Tensor) // must to fix

	if err := nn.backpropagate(x, y, wHidden, bHidden, wOut, bOut, output); err != nil {
		return err
	}
	nn.wHidden = wHidden
	nn.bHidden = bHidden
	nn.wOut = wOut
	nn.bOut = bOut

	return nil
}
