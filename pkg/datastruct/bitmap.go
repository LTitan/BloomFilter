package datastruct

import "sync"

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
