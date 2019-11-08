package datastruct

const (
	Size    uint = 64
	logSize uint = 6
)

type Bitmap struct {
	element []uint64
	length  uint
}

func New(length uint) (bmp *Bitmap) {
	defer func() {
		if r := recover(); r != nil {
			bmp = &Bitmap{
				make([]uint64, 0),
				0,
			}
		}
	}()

	bmp = &Bitmap{
		make([]uint64, NeedSize(length)),
		length,
	}
	return
}

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
func (b *Bitmap) Add(n uint) {
	b.extend(n)
	b.element[n>>logSize] |= 1 << (n & (Size - 1))
}
func (b *Bitmap) Has(n uint) bool {
	if n >= b.length {
		return false
	}
	return b.element[n>>logSize]&(1<<(n&(Size-1))) != 0
}

func (b *Bitmap)Len()uint{
	return b.length
}