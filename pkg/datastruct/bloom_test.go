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
	for i := 0; i < 10000000; i++ {
		s := fmt.Sprintf("%d.%d.%d.%d", i, i/10, i/100, i/1000)
		if bf.Has(s) {
			cnt++
		}
	}
	t.Logf("accuracy: %.4f", 10000000.0/float64(cnt))
}
func Benchnew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(0)
	}
}
