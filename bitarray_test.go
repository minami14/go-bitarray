package bitarray

import (
	"errors"
	"math"
	"testing"
)

func TestBitArray(t *testing.T) {
	bitArray, err := NewBitArray(math.MaxInt16)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < bitArray.Length(); i++ {
		bit, err := bitArray.Get(i)
		if err != nil {
			t.Error(err)
		}

		if bit {
			t.Error("initialize error")
		}

		if err := bitArray.Set(i); err != nil {
			t.Error(err)
		}

		bit, err = bitArray.Get(i)
		if err != nil {
			t.Error(err)
		}

		if !bit {
			t.Error("set error")
		}

		if err := bitArray.Clear(i); err != nil {
			t.Error(err)
		}

		bit, err = bitArray.Get(i)
		if err != nil {
			t.Error(err)
		}

		if bit {
			t.Error("clear error")
		}

		if err := bitArray.Set(i); err != nil {
			t.Error(err)
		}
	}

	bitArray.Reset()

	for i := 0; i < bitArray.Length(); i++ {
		bit, err := bitArray.Get(i)
		if err != nil {
			t.Error(err)
		}

		if bit {
			t.Error("reset error")
		}
	}
}

func BenchmarkBitArray_Get(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	bitArray, err := NewBitArray(b.N)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if _, err = bitArray.Get(i); err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkBitArray_Set(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	bitArray, err := NewBitArray(b.N)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if err := bitArray.Set(i); err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkBoolSlice_Get(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	boolSlice := make([]bool, b.N)

	for i := 0; i < b.N; i++ {
		_ = boolSlice[i]
	}
}

func BenchmarkBoolSlice_Set(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	boolSlice := make([]bool, b.N)

	for i := 0; i < b.N; i++ {
		boolSlice[i] = true
	}
}

func BenchmarkBitArray8_Get(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	bitArray, err := NewBitArray8(b.N)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if _, err = bitArray.Get(i); err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkBitArray8_Set(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	bitArray, err := NewBitArray8(b.N)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if err := bitArray.Set(i); err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkBitArray16_Get(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	bitArray, err := NewBitArray16(b.N)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if _, err = bitArray.Get(i); err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkBitArray16_Set(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	bitArray, err := NewBitArray16(b.N)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if err := bitArray.Set(i); err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkBitArray32_Get(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	bitArray, err := NewBitArray32(b.N)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if _, err = bitArray.Get(i); err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkBitArray32_Set(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	bitArray, err := NewBitArray32(b.N)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if err := bitArray.Set(i); err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkBitArray64_Get(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	bitArray, err := NewBitArray64(b.N)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if _, err = bitArray.Get(i); err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkBitArray64_Set(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	bitArray, err := NewBitArray64(b.N)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if err := bitArray.Set(i); err != nil {
			b.Log(err)
		}
	}
}

type BitArray8 struct {
	blocks []byte
	length int
}

func NewBitArray8(size int) (*BitArray8, error) {
	if size <= 0 {
		return nil, errors.New("size is 0 or less")
	}

	return &BitArray8{
		blocks: make([]byte, size/8+1),
		length: size,
	}, nil
}

func (b *BitArray8) Set(index int) error {
	if index < 0 || index >= b.length {
		return errors.New("index out of range")
	}

	i := index / 8
	u := b.blocks[i]
	shift := byte(index % 8)
	mask := byte(1 << shift)
	flag := u | mask
	b.blocks[i] = flag
	return nil
}

func (b *BitArray8) Get(index int) (bool, error) {
	if index < 0 || index >= b.length {
		return false, errors.New("index out of range")
	}

	i := index / 8
	u := b.blocks[i]
	shift := byte(index % 8)
	mask := byte(1 << shift)
	flag := u & mask
	return flag != 0, nil
}

type BitArray16 struct {
	blocks []uint16
	length int
}

func NewBitArray16(size int) (*BitArray16, error) {
	if size <= 0 {
		return nil, errors.New("size is 0 or less")
	}

	return &BitArray16{
		blocks: make([]uint16, size/16+1),
		length: size,
	}, nil
}

func (b *BitArray16) Set(index int) error {
	if index < 0 || index >= b.length {
		return errors.New("index out of range")
	}

	i := index / 16
	u := b.blocks[i]
	shift := uint16(index % 16)
	mask := uint16(1 << shift)
	flag := u | mask
	b.blocks[i] = flag
	return nil
}

func (b *BitArray16) Get(index int) (bool, error) {
	if index < 0 || index >= b.length {
		return false, errors.New("index out of range")
	}

	i := index / 16
	u := b.blocks[i]
	shift := uint16(index % 16)
	mask := uint16(1 << shift)
	flag := u & mask
	return flag != 0, nil
}

type BitArray32 struct {
	blocks []uint32
	length int
}

func NewBitArray32(size int) (*BitArray32, error) {
	if size <= 0 {
		return nil, errors.New("size is 0 or less")
	}

	return &BitArray32{
		blocks: make([]uint32, size/32+1),
		length: size,
	}, nil
}

func (b *BitArray32) Set(index int) error {
	if index < 0 || index >= b.length {
		return errors.New("index out of range")
	}

	i := index / 32
	u := b.blocks[i]
	shift := uint32(index % 32)
	mask := uint32(1 << shift)
	flag := u | mask
	b.blocks[i] = flag
	return nil
}

func (b *BitArray32) Get(index int) (bool, error) {
	if index < 0 || index >= b.length {
		return false, errors.New("index out of range")
	}

	i := index / 32
	u := b.blocks[i]
	shift := uint32(index % 32)
	mask := uint32(1 << shift)
	flag := u & mask
	return flag != 0, nil
}

type BitArray64 struct {
	blocks []uint64
	length int
}

func NewBitArray64(size int) (*BitArray64, error) {
	if size <= 0 {
		return nil, errors.New("size is 0 or less")
	}

	return &BitArray64{
		blocks: make([]uint64, size/64+1),
		length: size,
	}, nil
}

func (b *BitArray64) Set(index int) error {
	if index < 0 || index >= b.length {
		return errors.New("index out of range")
	}

	i := index / 64
	u := b.blocks[i]
	shift := uint64(index % 64)
	mask := uint64(1 << shift)
	flag := u | mask
	b.blocks[i] = flag
	return nil
}

func (b *BitArray64) Get(index int) (bool, error) {
	if index < 0 || index >= b.length {
		return false, errors.New("index out of range")
	}

	i := index / 64
	u := b.blocks[i]
	shift := uint64(index % 64)
	mask := uint64(1 << shift)
	flag := u & mask
	return flag != 0, nil
}
