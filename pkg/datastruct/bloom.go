package datastruct

// DEFUALTSIZE .
const DEFUALTSIZE = 0x01 << 26

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
	Set   Bitmap
	Funcs []SimpleHash
}

// New .
func New() *BloomFilter {
	bf := new(BloomFilter)
	for i := 0; i < len(seeds); i++ {
		sh := SimpleHash{Cap: DEFUALTSIZE, Seed: seeds[i]}
		bf.Funcs = append(bf.Funcs, sh)
	}
	bf.Set = NewBitmap(DEFUALTSIZE)
	return bf
}

// Add .
func (bf *BloomFilter) Add(val string) {
	for i := 0; i < len(bf.Funcs); i++ {
		bf.Set.Add(bf.Funcs[i].Hash(val))
	}
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
