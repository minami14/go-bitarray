package bitarray

import "errors"

// BitArray is bool array with low memory usage.
type BitArray struct {
	blocks []uint64
	length int
}

// bit per block
const bits = 64

// NewBitArray is BitArray constructed.
func NewBitArray(size int) (*BitArray, error) {
	if size <= 0 {
		return nil, errors.New("size is 0 or less")
	}

	return &BitArray{
		blocks: make([]uint64, size/bits+1),
		length: size,
	}, nil
}

// Set sets the specified bit to true.
func (b *BitArray) Set(index int) error {
	if index < 0 || index >= b.length {
		return errors.New("index out of range")
	}

	i := index / bits
	u := b.blocks[i]
	shift := uint64(index % bits)
	mask := uint64(1 << shift)
	flag := u | mask
	b.blocks[i] = flag
	return nil
}

// Get gets the specified bit.
func (b *BitArray) Get(index int) (bool, error) {
	if index < 0 || index >= b.length {
		return false, errors.New("index out of range")
	}

	i := index / bits
	u := b.blocks[i]
	shift := uint64(index % bits)
	mask := uint64(1 << shift)
	flag := u & mask
	return flag != 0, nil
}

// Clear sets the specified bit to false.
func (b *BitArray) Clear(index int) error {
	if index < 0 || index >= b.length {
		return errors.New("index out of range")
	}

	i := index / bits
	u := b.blocks[i]
	shift := uint64(index % bits)
	mask := uint64(1 << shift)
	flag := u & ^mask
	b.blocks[i] = flag
	return nil
}

// Reset set all bits to false.
func (b *BitArray) Reset() {
	b.blocks = make([]uint64, b.length/bits+1)
}

// Length returns number of bits in the BitArray.
func (b *BitArray) Length() int {
	return b.length
}
