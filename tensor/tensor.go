package tensor

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Tensor is multi dimensional array.
type Tensor struct {
	Data   []float64
	Shape  []int
	Stride []int
}

// New functions
// NewTensor makes tensor from given data.
func NewTensor(data []float64, shape ...int) *Tensor {
	if !okLength(data, shape) {
		log.Fatalf("valError:invalid shape len:%v,shape:%v\n", len(data), shape)
	}
	return &Tensor{
		Data:   data,
		Shape:  shape,
		Stride: stride(shape),
	}
}

// NewZeros returns zeros tensor.
func NewZeros(shape ...int) *Tensor {
	return &Tensor{
		Data:   make([]float64, length(shape)),
		Shape:  shape,
		Stride: stride(shape),
	}
}

// NewOnes returns ones tensor.
func NewOnes(shape ...int) *Tensor {
	ones := nums(1, shape)
	if !okLength(ones, shape) {
		log.Fatalf("valError:invalid shape len:%v,shape:%v\n", len(ones), shape)
	}
	return &Tensor{
		Data:   ones,
		Shape:  shape,
		Stride: stride(shape),
	}
}

func okLength(data []float64, shape []int) bool {
	return len(data) == length(shape)
}

func length(shape []int) int {
	length := 0
	for i, v := range shape {
		if i == 0 {
			length += v
		} else {
			length *= v
		}
	}
	return length
}

func stride(shape []int) []int {
	stride := make([]int, len(shape))
	for i := len(shape) - 1; i >= 0; i-- {
		if i == len(shape)-1 {
			stride[i] = 1
		} else {
			stride[i] = shape[i+1] * stride[i+1]
		}
	}
	return stride
}

func nums(num float64, shape []int) []float64 {
	out := make([]float64, length(shape))
	for i := 0; i < length(shape); i++ {
		out[i] = num
	}
	return out
}

// ReadCsv reads data from file.
func ReadCsv(path string) *Tensor {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("file open failed:%v", err)
	}
	r := csv.NewReader(f)
	recs := make([][]string, 0)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("file read failed:%v", err)
		}
		recs = append(recs, rec)
	}
	flt := flatten(recs)
	data := parseFloat(flt)
	shape := []int{len(recs), len(recs[0])}
	return &Tensor{
		Data:   data,
		Shape:  shape,
		Stride: stride(shape),
	}
}

func parseFloat(s []string) []float64 {
	floats := make([]float64, len(s))
	for i, v := range s {
		float, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatalf("parse float failed:%v", err)
		}
		floats[i] = float
	}
	return floats
}

func flatten(s [][]string) []string {
	flt := []string{}
	for _, v := range s {
		for _, w := range v {
			flt = append(flt, w)
		}
	}
	return flt
}

// Methods
// ToDo:Reshape
// Inspect returns formatted array.
func (ts *Tensor) Inspect() string {
	out := strConv(ts.Data)
	for i := len(ts.Shape) - 1; i >= 0; i-- {
		out = agg(out, ts.Shape[i])
	}
	return format(out)
}

func strConv(in []float64) []string {
	out := make([]string, len(in))
	for i, v := range in {
		s := fmt.Sprint(v)
		out[i] = s
	}
	return out
}

func agg(in []string, interval int) []string {
	var out []string
	var b strings.Builder
	p := 0
	for p < len(in) {
		b.WriteString("[")
		for _, v := range in[p : p+interval] {
			b.WriteString(v + " ")
		}
		b.WriteString("]")
		out = append(out, b.String())
		b.Reset()
		p = p + interval
	}
	return out
}

func format(in []string) string {
	var b strings.Builder
	for _, v := range in {
		rep := strings.Replace(v, " ]", "]", -1)
		//rep = strings.Replace(rep,"] ","]\n",-1)
		b.WriteString(rep)
	}
	return b.String()
}

// Len returns length(number of elements) of the tensor.
func (ts *Tensor) Len() int {
	return len(ts.Data)
}

// Shape returns shape of the tensor.
//func (ts *Tensor) Shape() []int {
//	return ts.Shape
//}

// Stride returns stride of the tensor.
//func (ts *Tensor) Stride() []int {
//	return ts.Stride
//}

// Randomize
func (ts *Tensor) Randomize() {
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	for i, _ := range ts.Data {
		ts.Data[i] = r.NormFloat64()
	}
}
