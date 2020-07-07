package series

import (
	"testing"
	"bytes"
)

func TestReadCsv(t *testing.T) {
	in := bytes.NewBufferString("1,2,3,4,5")
	e := []string{"1","2","3","4","5"}
	recs, err := readCsv(in)
	if err != nil {
		t.Errorf("read error:%v",err)
	}
	for i, v := range recs {
		if string(v[0][i]) != e[i] {
			t.Errorf("i:%v, e:%v, a:%v",i,e[i],string(v[0][i]))
		}
	}
}
