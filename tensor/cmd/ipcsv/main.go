package main

import (
	"fmt"
	"github.com/xxnmxx/goda/array"
)

func main() {
	ar := array.ReadCsv(`.\data`)
	out := ar.Inspect()
	fmt.Println(out)
}
