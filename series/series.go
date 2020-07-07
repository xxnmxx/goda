package series

type Series interface {
	Type() string
	Len() int
}

type floatSeries []float64

func (fs *floatSeries) Type() string { return "floatSeries" }
func (fs *floatSeries) Len() float64 { return float64(len(*fs)) }
func (fs *floatSeries) Sum() float64 {
	sum := 0.0
	for _, v := range *fs {
		sum += v
	}
	return sum
}

type stringSeries []string

func (ss *stringSeries) Type() string { return "stringSeries" }
func (ss *stringSeries) Len() float64 { return float64(len(*ss)) }

