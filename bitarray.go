package bitarray

import (
	"errors"
	"math"
	"math/bits"
)

// BitArray is bool array with low memory usage.
type BitArray struct {
	blocks []uint64
	length int
}

const bitPerBlock = 64
const max = math.MaxUint64

// NewBitArray is BitArray constructed.
func NewBitArray(length int) (*BitArray, error) {
	if length < 0 {
		return nil, errors.New("negative length argument")
	}

	blockSize := length / bitPerBlock
	if length%bitPerBlock != 0 {
		blockSize++
	}

	return &BitArray{
		blocks: make([]uint64, blockSize),
		length: length,
	}, nil
}

// Set sets the specified bit to true.
func (b *BitArray) Set(index int) error {
	if index < 0 || index >= b.length {
		return errors.New("index out of range")
	}

	i := index / bitPerBlock
	u := b.blocks[i]
	shift := uint64(index % bitPerBlock)
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

	i := index / bitPerBlock
	u := b.blocks[i]
	shift := uint64(index % bitPerBlock)
	mask := uint64(1 << shift)
	flag := u & mask
	return flag != 0, nil
}

// Clear sets the specified bit to false.
func (b *BitArray) Clear(index int) error {
	if index < 0 || index >= b.length {
		return errors.New("index out of range")
	}

	i := index / bitPerBlock
	u := b.blocks[i]
	shift := uint64(index % bitPerBlock)
	mask := uint64(1 << shift)
	flag := u & ^mask
	b.blocks[i] = flag
	return nil
}

// Reset set all bitPerBlock to false.
func (b *BitArray) Reset() {
	b.blocks = make([]uint64, b.length/bitPerBlock+1)
}

// Length returns number of bitPerBlock in the BitArray.
func (b *BitArray) Length() int {
	return b.length
}

// Slice the BitArray.
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

	i := 0
	if start < 0 {
		i = -start
	}

	for ; i < copySize; i++ {
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

// Clone the BitArray.
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

// Not inverts all bitPerBlock.
func (b *BitArray) Not() (*BitArray, error) {
	bitArray, err := NewBitArray(b.length)
	if err != nil {
		return nil, err
	}

	for i, v := range b.blocks {
		bitArray.blocks[i] = ^v
	}

	if len(b.blocks) > 0 && b.length%bitPerBlock != 0 {
		mask := uint64(max) >> uint64(bitPerBlock-b.length%bitPerBlock)
		bitArray.blocks[len(b.blocks)-1] &= mask
	}

	return bitArray, err
}

// And is the logical AND of two BitArrays.
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

	if len(a.blocks) > 0 && a.length%bitPerBlock != 0 {
		mask := uint64(max) >> uint64(bitPerBlock-a.length%bitPerBlock)
		bitArray.blocks[len(a.blocks)-1] &= mask
	}

	return bitArray, nil
}

// Or is the logical OR of two BitArrays.
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

	if len(a.blocks) > 0 && a.length%bitPerBlock != 0 {
		mask := uint64(max) >> uint64(bitPerBlock-a.length%bitPerBlock)
		bitArray.blocks[len(a.blocks)-1] &= mask
	}

	return bitArray, nil
}

// Xor is the Exclusive OR of two BitArrays.
func Xor(a, b *BitArray) (*BitArray, error) {
	if a.length > b.length {
		a, b = b, a
	}

	bitArray, err := NewBitArray(b.length)
	if err != nil {
		return nil, err
	}

	for i, v := range a.blocks {
		bitArray.blocks[i] = v ^ b.blocks[i]
	}

	if len(a.blocks) > 0 && a.length%bitPerBlock != 0 {
		mask := uint64(max) >> uint64(bitPerBlock-a.length%bitPerBlock)
		bitArray.blocks[len(a.blocks)-1] &= mask
	}

	return bitArray, nil
}

// AndNot clears bitPerBlock specified by argument BitArray.
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

	if len(andNot.blocks) > 0 && andNot.length%bitPerBlock != 0 {
		mask := uint64(max) >> uint64(bitPerBlock-b.length%bitPerBlock)
		andNot.blocks[len(andNot.blocks)-1] &= mask
	}

	return andNot, nil
}

// LeftShift shifts the BitArray to the left.
func (b *BitArray) LeftShift(n int) (*BitArray, error) {
	if n < 0 {
		return b.RightShift(-n)
	}

	if n == 0 {
		return b.Clone()
	}

	length := b.length + n

	bitArray, err := NewBitArray(length)
	if err != nil {
		return nil, err
	}

	div := n / bitPerBlock
	mod := uint64(n % bitPerBlock)
	shift := bitPerBlock - mod
	for i := 1; i < len(b.blocks); i++ {
		bitArray.blocks[i+div] = (b.blocks[i] << mod) | (b.blocks[i-1] >> shift)
	}

	if len(b.blocks) > 0 {
		bitArray.blocks[div] = b.blocks[0] << mod
		if len(b.blocks)+div < len(bitArray.blocks) {
			bitArray.blocks[len(bitArray.blocks)-1] = b.blocks[len(b.blocks)-1] >> shift
		}
	}

	return bitArray, nil
}

// RightShift shifts the BitArray to the right.
func (b *BitArray) RightShift(n int) (*BitArray, error) {
	if n < 0 {
		return b.LeftShift(-n)
	}

	if n == 0 {
		return b.Clone()
	}

	bitArray, err := NewBitArray(b.length)
	if err != nil {
		return nil, err
	}

	div := n / bitPerBlock
	mod := uint64(n % bitPerBlock)
	shift := bitPerBlock - mod
	for i := div; i < len(b.blocks)-1; i++ {
		bitArray.blocks[i-div] = (b.blocks[i] >> mod) | (b.blocks[i+1] << shift)
	}

	if len(b.blocks)-div > 0 {
		bitArray.blocks[len(b.blocks)-1-div] = b.blocks[len(b.blocks)-1] >> mod
	}

	return bitArray, nil
}

// Reverse reverses bit order
func (b *BitArray) Reverse() (*BitArray, error) {
	reversed, err := NewBitArray(b.length)
	if err != nil {
		return nil, err
	}

	lastIndex := len(b.blocks) - 1
	mod := uint64(b.length % bitPerBlock)
	if mod == 0 {
		mod = bitPerBlock
	}
	shift := bitPerBlock - mod

	for i := 0; i < lastIndex; i++ {
		reversed.blocks[i] = (bits.Reverse64(b.blocks[lastIndex-i]) >> shift) | bits.Reverse64(b.blocks[lastIndex-i-1]>>mod)
	}

	if len(b.blocks) > 0 {
		reversed.blocks[lastIndex] = bits.Reverse64(b.blocks[0]) >> (bitPerBlock - mod)
	}

	return reversed, nil
}

// OnesCount returns the number of true bits in BitArray
func (b *BitArray) OnesCount() int {
	count := 0
	for i := 0; i < len(b.blocks); i++ {
		count += bits.OnesCount64(b.blocks[i])
	}

	return count
}
