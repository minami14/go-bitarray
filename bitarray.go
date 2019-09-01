package bitarray

import "errors"

// BitArray is bool array with low memory usage.
type BitArray struct {
	blocks []byte
	length int
}

// bit per block
const bits = 8

// NewBitArray is BitArray constructed.
func NewBitArray(length int) (*BitArray, error) {
	if length < 0 {
		return nil, errors.New("negative length argument")
	}

	blockSize := length / bits
	if length%bits != 0 {
		blockSize++
	}

	return &BitArray{
		blocks: make([]byte, blockSize),
		length: length,
	}, nil
}

// Set sets the specified bit to true.
func (b *BitArray) Set(index int) error {
	if index < 0 || index >= b.length {
		return errors.New("index out of range")
	}

	i := index / bits
	u := b.blocks[i]
	shift := byte(index % bits)
	mask := byte(1 << shift)
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
	shift := byte(index % bits)
	mask := byte(1 << shift)
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
	shift := byte(index % bits)
	mask := byte(1 << shift)
	flag := u & ^mask
	b.blocks[i] = flag
	return nil
}

// Reset set all bits to false.
func (b *BitArray) Reset() {
	b.blocks = make([]byte, b.length/bits+1)
}

// Length returns number of bits in the BitArray.
func (b *BitArray) Length() int {
	return b.length
}

// Slice the BitArray
func (b *BitArray) Slice(start, end int) (*BitArray, error) {
	size := end - start
	bitArray, err := NewBitArray(size)
	if err != nil {
		return nil, err
	}

	copySize := size
	if size > b.length-start {
		copySize = b.length - start
	}

	for i := 0; i < copySize; i++ {
		isSet, err := b.Get(i + start)
		if err != nil {
			return nil, err
		}

		if isSet {
			if err := bitArray.Set(i); err != nil {
				return nil, err
			}
		}
	}

	return bitArray, nil
}

// Clone the BitArray
func (b *BitArray) Clone() (*BitArray, error) {
	clone, err := NewBitArray(b.length)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(clone.blocks); i++ {
		clone.blocks[i] = b.blocks[i]
	}

	return clone, nil
}

// Not inverts all bits
func (b *BitArray) Not() (*BitArray, error) {
	bitArray, err := NewBitArray(b.length)
	if err != nil {
		return nil, err
	}

	for i, v := range b.blocks {
		bitArray.blocks[i] = ^v
	}

	mask := byte(0xFF) >> byte(b.length%bits)
	bitArray.blocks[len(b.blocks)-1] &= mask

	return bitArray, err
}

// And is the logical AND of two BitArrays
func And(a, b *BitArray) (*BitArray, error) {
	if a.length > b.length {
		a, b = b, a
	}

	bitArray, err := NewBitArray(b.length)
	if err != nil {
		return nil, err
	}

	for i, v := range a.blocks {
		bitArray.blocks[i] = v & b.blocks[i]
	}

	mask := byte(0xFF) >> byte(a.length%bits)
	bitArray.blocks[len(a.blocks)-1] &= mask

	return bitArray, nil
}

// Or is the logical Or of two BitArrays
func Or(a, b *BitArray) (*BitArray, error) {
	if a.length > b.length {
		a, b = b, a
	}

	bitArray, err := NewBitArray(b.length)
	if err != nil {
		return nil, err
	}

	for i, v := range a.blocks {
		bitArray.blocks[i] = v | b.blocks[i]
	}

	mask := byte(0xFF) >> byte(a.length%bits)
	bitArray.blocks[len(a.blocks)-1] &= mask

	return bitArray, nil
}

// AndNot clears bits specified by argument BitArray
func (b *BitArray) AndNot(bitArray *BitArray) (*BitArray, error) {
	andNot, err := NewBitArray(b.length)
	if err != nil {
		return nil, err
	}

	if bitArray.length != b.length {
		bitArray, err = bitArray.Slice(0, b.length)
		if err != nil {
			return nil, err
		}
	}

	for i, v := range b.blocks {
		andNot.blocks[i] = bitArray.blocks[i] &^ v
	}

	mask := byte(0xFF) >> byte(andNot.length%bits)
	andNot.blocks[len(b.blocks)-1] &= mask

	return andNot, nil
}
