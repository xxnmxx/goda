package layers

type Layers interface {
	Forward() *Tensor
	Backward() *Tensor
}

type Relu *Tensor

func (rl *Relu) Forward() *Tensor {
	// wip
	return nil
}
