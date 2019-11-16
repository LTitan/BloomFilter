package datastruct

import "math"

import "unsafe"

import "fmt"

var seeds = []uint{13, 31, 131, 1313, 13131}

// SimpleHash hash
type SimpleHash struct {
	Cap, Seed uint
}

// Hash .
func (s *SimpleHash) Hash(val string) uint {
	var result uint
	for i := 0; i < len(val); i++ {
		result = result*s.Seed + uint(val[i])
	}
	return (s.Cap - 1) & result
}

// BloomFilter .
type BloomFilter struct {
	Set   *Bitmap
	Funcs []SimpleHash
	size  uint
	used  uint
}

// New .
func New(size uint) *BloomFilter {
	bf := new(BloomFilter)
	for i := 0; i < len(seeds); i++ {
		sh := SimpleHash{Cap: 0x1f3fffff, Seed: seeds[i]}
		bf.Funcs = append(bf.Funcs, sh)
	}
	bf.Set = NewBitmap(size)
	bf.size = size
	bf.used = 0
	return bf
}

// Add .
func (bf *BloomFilter) Add(val string) {
	for i := 0; i < len(bf.Funcs); i++ {
		bf.Set.Add(bf.Funcs[i].Hash(val))
	}
	bf.used += uint(len(bf.Funcs))
}

// Has .
func (bf *BloomFilter) Has(val string) bool {
	if len(val) == 0 {
		return false
	}
	ret := true
	for i := 0; i < len(bf.Funcs); i++ {
		ret = ret && bf.Set.Has(bf.Funcs[i].Hash(val))
	}
	return ret
}

// Delete delete a key
func (bf *BloomFilter) Delete(val string) {
	for i := 0; i < len(bf.Funcs); i++ {
		bf.Set.Clear(bf.Funcs[i].Hash(val))
	}
	bf.used -= uint(len(bf.Funcs))
}

// Avaliable .
func (bf *BloomFilter) Avaliable() string {
	var a int64
	ptr := unsafe.Sizeof(a)
	return fmt.Sprintf("%vM", uint(math.Ceil(float64((bf.size-bf.used)*uint(ptr))/1048576.0)))
}

// Release .
func (bf *BloomFilter) Release() {
	bf.Set.Release()
	bf.Set = nil
	return
}
