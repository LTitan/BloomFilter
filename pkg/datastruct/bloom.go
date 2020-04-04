package datastruct

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"unsafe"

	fileop "github.com/LTitan/BloomFilter/pkg/files"
)

var seeds = []uint{13, 31, 131, 1313, 13131}

const basePath = "./dump_record"

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
		sh := SimpleHash{Cap: 0x1111ffff, Seed: seeds[i]}
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

// Dump .
func (bf *BloomFilter) Dump(key string) {
	str := bf.Set.String()
	if err := fileop.CreateFileDir(basePath); err != nil {
		log.Fatalln(err)
		return
	}
	if err := fileop.WtiteString(basePath+"/"+key, str); err != nil {
		log.Fatalln(err)
		return
	}
}

// Load return map[string]*filter
func Load() (ret map[string]*BloomFilter, err error) {
	files, err := fileop.GetAllFile(basePath)
	if err != nil {
		return nil, err
	}
	if len(files) <= 0 {
		return nil, nil
	}
	ret = make(map[string]*BloomFilter, len(files))
	for _, f := range files {
		data, err := fileop.ReadFile(basePath, f)
		if err != nil {
			return nil, err
		}
		ret[f] = stringToBloomfilter(data)
	}
	return
}

func stringToBloomfilter(str string) *BloomFilter {
	ret := New(0x01 << 20)
	if len(str) == 0 {
		return ret
	}
	nums := strings.Split(str, ",")
	for _, num := range nums {
		n, err := strconv.ParseUint(num, 10, 64)
		if err != nil {
			continue
		}
		ret.Set.Add(uint(n))
	}
	return ret
}
