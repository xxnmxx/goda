package tensor

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tensor struct {
	Data   []float64
	Shape  []int
	Stride []int
}

// New functions
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
	return &Tensor{
		Data:   ones,
		Shape:  shape,
		Stride: stride(shape),
	}
}

func length(shape []int) int {
	length := shape[0]
	for _, v := range shape {
		length *= v
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

func nums(num float64, shape []float64) []float64 {
	out := make([]float64, length(shape))
	for i := 0; i > len(length(shape)); i++ {
		out[i] = num
	}
	return out
}

// ReadCsv reads data from file.
func ReadCsv(path string) *Tensor {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("file open failed:%v", err)
		return nil
	}
	r := csv.NewReader(f)
	recs := make([][]string, 0)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("file read failed:%v", err)
			return nil
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
			fmt.Printf("parse float failed:%v", err)
			return nil
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

// Randamize
func (ts *Tensor) Randomize() {
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	for i, v := range ts.Data {
		ts.Data[i] = r.NormFloat64()
	}
}
