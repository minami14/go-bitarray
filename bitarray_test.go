package bitarray

import (
	"errors"
	"testing"
)

func TestBitArray(t *testing.T) {
	for length := 0; length < 1000; length++ {
		bitArray, err := NewBitArray(length)
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
}

func TestBitArray_Not(t *testing.T) {
	for length := 0; length < 10000; length++ {
		bitArray, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArray.length; i += 7 {
			if err := bitArray.Set(i); err != nil {
				t.Error(err)
			}
		}

		not, err := bitArray.Not()
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArray.length; i++ {
			b, err := bitArray.Get(i)
			if err != nil {
				t.Error(err)
			}

			n, err := not.Get(i)
			if err != nil {
				t.Error(err)
			}

			if b == n {
				t.Errorf("value matches %v", i)
			}
		}
	}
}

func TestBitArray_AndNot(t *testing.T) {
	for length := 0; length < 10000; length++ {
		bitArrayA, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArrayA.length; i += 7 {
			if err := bitArrayA.Set(i); err != nil {
				t.Error(err)
			}
		}

		bitArrayB, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArrayB.length; i += 4 {
			if err := bitArrayB.Set(i); err != nil {
				t.Error(err)
			}
		}

		bitArrayAndNot, err := bitArrayA.AndNot(bitArrayB)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < length; i++ {
			a, err := bitArrayA.Get(i)
			if err != nil {
				t.Error(err)
			}

			b, err := bitArrayB.Get(i)
			if err != nil {
				t.Error(err)
			}

			andNot, err := bitArrayAndNot.Get(i)
			if err != nil {
				t.Error(err)
			}

			if andNot != ((!a) && b) {
				t.Errorf("value does not match %v %v %v %v", i, a, b, andNot)
			}
		}
	}
}

func TestAnd(t *testing.T) {
	for length := 0; length < 10000; length++ {
		bitArrayA, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArrayA.length; i += 7 {
			if err := bitArrayA.Set(i); err != nil {
				t.Error(err)
			}
		}

		bitArrayB, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArrayB.length; i += 4 {
			if err := bitArrayB.Set(i); err != nil {
				t.Error(err)
			}
		}

		bitArrayAnd, err := And(bitArrayA, bitArrayB)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < length; i++ {
			a, err := bitArrayA.Get(i)
			if err != nil {
				t.Error(err)
			}

			b, err := bitArrayB.Get(i)
			if err != nil {
				t.Error(err)
			}

			and, err := bitArrayAnd.Get(i)
			if err != nil {
				t.Error(err)
			}

			if and != (a && b) {
				t.Error("value does not match")
			}
		}
	}
}

func TestOr(t *testing.T) {
	for length := 0; length < 10000; length++ {
		bitArrayA, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArrayA.length; i += 7 {
			if err := bitArrayA.Set(i); err != nil {
				t.Error(err)
			}
		}

		bitArrayB, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArrayB.length; i += 4 {
			if err := bitArrayB.Set(i); err != nil {
				t.Error(err)
			}
		}

		bitArrayOr, err := Or(bitArrayA, bitArrayB)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < length; i++ {
			a, err := bitArrayA.Get(i)
			if err != nil {
				t.Error(err)
			}

			b, err := bitArrayB.Get(i)
			if err != nil {
				t.Error(err)
			}

			or, err := bitArrayOr.Get(i)
			if err != nil {
				t.Error(err)
			}

			if or != (a || b) {
				t.Error("value does not match")
			}
		}
	}
}

func TestXor(t *testing.T) {
	for length := 0; length < 10000; length++ {
		bitArrayA, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArrayA.length; i += 7 {
			if err := bitArrayA.Set(i); err != nil {
				t.Error(err)
			}
		}

		bitArrayB, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArrayB.length; i += 4 {
			if err := bitArrayB.Set(i); err != nil {
				t.Error(err)
			}
		}

		bitArrayXor, err := Xor(bitArrayA, bitArrayB)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < length; i++ {
			a, err := bitArrayA.Get(i)
			if err != nil {
				t.Error(err)
			}

			b, err := bitArrayB.Get(i)
			if err != nil {
				t.Error(err)
			}

			xor, err := bitArrayXor.Get(i)
			if err != nil {
				t.Error(err)
			}

			if xor != (a != b) {
				t.Error("value does not match")
			}
		}
	}
}

func TestBitArray_LeftShift(t *testing.T) {
	for length := 0; length < 1000; length++ {
		bitArray, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArray.length; i += 7 {
			if err := bitArray.Set(i); err != nil {
				t.Error(err)
			}
		}

		for n := 0; n < 128; n++ {
			bitArrayLeft, err := bitArray.LeftShift(n)
			if err != nil {
				t.Fatal(err)
			}

			for i := 0; i < bitArray.length; i++ {
				b, err := bitArray.Get(i)
				if err != nil {
					t.Error(err)
				}

				left, err := bitArrayLeft.Get(i + n)
				if err != nil {
					t.Error(err)
				}

				if b != left {
					t.Errorf("value does not match %v %v %v %v", i, n, b, left)
				}
			}
		}
	}
}

func TestBitArray_RightShift(t *testing.T) {
	for length := 0; length < 1000; length++ {
		bitArray, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArray.length; i += 7 {
			if err := bitArray.Set(i); err != nil {
				t.Error(err)
			}
		}

		for n := 0; n < 128; n++ {
			bitArrayRight, err := bitArray.RightShift(n)
			if err != nil {
				t.Fatal(err)
			}

			for i := n; i < bitArray.length; i++ {
				b, err := bitArray.Get(i)
				if err != nil {
					t.Error(err)
				}

				right, err := bitArrayRight.Get(i - n)
				if err != nil {
					t.Error(err)
				}

				if b != right {
					t.Errorf("value does not match %v %v %v %v", i, n, b, right)
				}
			}
		}
	}
}

func TestBitArray_Reverse(t *testing.T) {
	for length := 0; length < 1000; length++ {
		bitArray, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArray.length; i += 7 {
			if err := bitArray.Set(i); err != nil {
				t.Error(err)
			}
		}

		reversed, err := bitArray.ReverseBytes()
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArray.length; i++ {
			b, err := bitArray.Get(i)
			if err != nil {
				t.Error(err)
			}

			right, err := reversed.Get(bitArray.length - i - 1)
			if err != nil {
				t.Error(err)
			}

			if b != right {
				t.Errorf("value does not match %v %v %v %v", length, i, b, right)
			}
		}
	}
}

func TestBitArray_OnesCount(t *testing.T) {
	for length := 0; length < 1000; length++ {
		bitArray, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		count := 0
		for i := 0; i < bitArray.length; i += 7 {
			if err := bitArray.Set(i); err != nil {
				t.Error(err)
			}
			count++
		}

		c := bitArray.OnesCount()
		if c != count {
			t.Errorf("value does not match %v %v %v", length, c, count)
		}
	}
}

func TestBitArray_TrailingZeros(t *testing.T) {
	for length := 0; length < 1000; length++ {
		for i := 0; i < length; i++ {
			bitArray, err := NewBitArray(length)
			if err != nil {
				t.Fatal(err)
			}

			if err := bitArray.Set(i); err != nil {
				t.Error(err)
			}

			zeros := bitArray.TrailingZeros()
			if zeros != i {
				t.Errorf("value does not match %v %v %v", length, zeros, i)
			}
		}
	}
}

func TestBitArray_Slice(t *testing.T) {
	for length := 0; length < 10000; length++ {
		bitArray, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArray.length; i += 7 {
			if err := bitArray.Set(i); err != nil {
				t.Error(err)
			}
		}

		start := 200
		end := 900
		bitSlice, err := bitArray.Slice(start, end)
		if err != nil {
			t.Fatal(err)
		}

		max := end - start
		if max > bitArray.length-start {
			max = bitArray.length - start
		}

		for i := 0; i < max; i++ {
			isSetArray, err := bitArray.Get(i + start)
			if err != nil {
				t.Error(err)
			}

			isSetSlice, err := bitSlice.Get(i)
			if err != nil {
				t.Error(err)
			}

			if isSetArray != isSetSlice {
				t.Error("value does not match")
			}
		}
	}
}

func TestBitArray_Clone(t *testing.T) {
	for length := 0; length < 10000; length++ {
		bitArray, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArray.length; i += 7 {
			if err := bitArray.Set(i); err != nil {
				t.Error(err)
			}
		}

		clone, err := bitArray.Clone()
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < bitArray.length; i++ {
			isSetArray, err := bitArray.Get(i)
			if err != nil {
				t.Error(err)
			}

			isSetSlice, err := clone.Get(i)
			if err != nil {
				t.Error(err)
			}

			if isSetArray != isSetSlice {
				t.Error("value does not match")
			}
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
