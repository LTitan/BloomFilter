package datastruct

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	bf := New(1024)
	//bf := make(map[string]bool)
	for i := 0; i < 10000000; i++ {
		s := fmt.Sprintf("%d.%d.%d.%d", i, i/10, i/100, i/1000)
		bf.Add(s)
	}
	cnt := 0
	// 负样本
	for i := 0; i < 10000000; i++ {
		s := fmt.Sprintf("%09d.%09d.%d.%09d", i%1000, i%10, i, i%100)
		if bf.Has(s) {
			cnt++
		}
	}
	t.Logf("error accuracy: %.9f", float64(cnt)/10000000.0)
}
func Benchnew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(0)
	}
}
