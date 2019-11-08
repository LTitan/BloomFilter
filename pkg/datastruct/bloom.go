package datastruct

const DEFUALT_SIZE = 0x01 << 26

var seeds = []uint{13, 31, 131, 1313, 13131}

type SimpleHash struct {
	Cap, Seed uint
}

func (s *SimpleHash) Hash(val string) uint {
	var result uint
	for i := 0; i < len(val); i++ {
		result = result*s.Seed + uint(val[i])
	}
	return (s.Cap - 1) & result
}

type BloomFilter struct {
	Set   *bitmap.Bitmap
	Funcs []SimpleHash
}

func New() *BloomFilter {
	bf := new(BloomFilter)
	for i := 0; i < len(seeds); i++ {
		sh := SimpleHash{Cap: DEFUALT_SIZE, Seed: seeds[i]}
		bf.Funcs = append(bf.Funcs, sh)
	}
	bf.Set = bitmap.New(DEFUALT_SIZE)
	return bf
}
func (bf *BloomFilter) Add(val string) {
	for i := 0; i < len(bf.Funcs); i++ {
		bf.Set.Add(bf.Funcs[i].Hash(val))
	}
}
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