package datastruct

import (
	"bytes"
	"strconv"
	"sync"
)

const (
	// Size .
	Size    uint = 64
	logSize uint = 6
)

// Bitmap bit map
type Bitmap struct {
	element []uint64
	length  uint
	rw      *sync.RWMutex
}

// NewBitmap .
func NewBitmap(length uint) (bmp *Bitmap) {
	defer func() {
		if r := recover(); r != nil {
			bmp = &Bitmap{
				make([]uint64, 0),
				0,
				&sync.RWMutex{},
			}
		}
	}()
	bmp = &Bitmap{
		element: make([]uint64, length),
		length:  length,
		rw:      &sync.RWMutex{},
	}
	return
}

// NeedSize .
func NeedSize(n uint) int {
	cap := ^uint(0)
	if n > (cap - Size + 1) {
		return int(cap >> logSize)
	}
	return int((n + (Size - 1)) >> logSize)
}
func (b *Bitmap) extend(n uint) {
	if n >= b.length {
		nsize := NeedSize(n + 1)
		if b.element == nil {
			b.element = make([]uint64, nsize)
		} else if cap(b.element) >= nsize {
			b.element = b.element[:nsize]
		} else if len(b.element) < nsize {
			newset := make([]uint64, nsize, 2*nsize)
			copy(newset, b.element)
			b.element = newset
		}
		b.length = n + 1
	}
}

// Add .
func (b *Bitmap) Add(n uint) {
	b.rw.Lock()
	b.extend(n)
	b.element[n>>logSize] |= 1 << (n & (Size - 1))
	b.rw.Unlock()
}

// Has .
func (b *Bitmap) Has(n uint) bool {
	if n >= b.length {
		return false
	}
	b.rw.RLock()
	res := b.element[n>>logSize]&(1<<(n&(Size-1))) != 0
	b.rw.RUnlock()
	return res
}

// Len .
func (b *Bitmap) Len() uint {
	return b.length
}

// Clear clear
func (b *Bitmap) Clear(i uint) {
	if i >= b.length {
		return
	}
	b.rw.Lock()
	b.element[i>>logSize] &^= 1 << (i & (Size - 1))
	b.rw.Unlock()
	return
}

// Release .
func (b *Bitmap) Release() {
	b.element = nil
	return
}

// func Int64ToBytes(i uint64) []byte {
// 	var buf = make([]byte, 8)
// 	binary.BigEndian.PutUint64(buf, i)
// 	return buf
// }

// func BytesToInt64(buf []byte) uint64 {
// 	return binary.BigEndian.Uint64(buf)
// }

// String .
func (b *Bitmap) String() string {
	var buffer bytes.Buffer
	start := []byte{}
	buffer.Write(start)
	var counter uint64
	i, e := b.NextSet(0)
	for e {
		counter++
		buffer.WriteString(strconv.FormatInt(int64(i), 10))
		i, e = b.NextSet(i + 1)
		if e {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

// NextSet .
func (b *Bitmap) NextSet(i uint) (uint, bool) {
	x := int(i >> logSize)
	if x >= len(b.element) {
		return 0, false
	}
	w := b.element[x]
	w = w >> (i & (Size - 1))
	if w != 0 {
		return i + trailingZeroes64(w), true
	}
	x = x + 1
	for x < len(b.element) {
		if b.element[x] != 0 {
			return uint(x)*Size + trailingZeroes64(b.element[x]), true
		}
		x = x + 1

	}
	return 0, false
}

var deBruijn = [...]byte{
	0, 1, 56, 2, 57, 49, 28, 3, 61, 58, 42, 50, 38, 29, 17, 4,
	62, 47, 59, 36, 45, 43, 51, 22, 53, 39, 33, 30, 24, 18, 12, 5,
	63, 55, 48, 27, 60, 41, 37, 16, 46, 35, 44, 21, 52, 32, 23, 11,
	54, 26, 40, 15, 34, 20, 31, 10, 25, 14, 19, 9, 13, 8, 7, 6,
}

func trailingZeroes64(v uint64) uint {
	return uint(deBruijn[((v&-v)*0x03f79d71b4ca8b09)>>58])
}
