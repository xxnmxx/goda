package series

import (
	"encoding/csv"
	"io"
)

func readCsv(in io.Reader) ([][]string, error) {
	recs := make([][]string, 0)
	r := csv.NewReader(in)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		recs = append(recs,rec)
	}
	return recs,nil
}
