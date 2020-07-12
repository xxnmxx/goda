package goda

import (
	"encoding/csv"
)

type CsvReader struct {
	r *csv.Reader
}

func (cr *CsvReader) NewSeries(col int) *Series {}

func (cr *CsvReader) typeSelector() {
}

func (cr *CsvReader) NewDataFrame() *DataFrame {}
