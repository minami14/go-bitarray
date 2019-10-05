package bitarray

import (
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

func TestAdd(t *testing.T) {
	for length := 0; length < 1000; length++ {
		a, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		b, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		c, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < length; i += 7 {
			if err := a.Set(i); err != nil {
				t.Error(err)
			}

			if err := c.Set(i); err != nil {
				t.Error(err)
			}
		}

		carry := false
		for i := 0; i < length; i += 11 {
			if err := b.Set(i); err != nil {
				t.Error(err)
			}

			v, err := c.Get(i)
			if err != nil {
				t.Error(err)
			}

			if v == carry {
				if err := c.Set(i); err != nil {
					t.Error(err)
				}
			} else {
				if err := c.Clear(i); err != nil {
					t.Error(err)
				}
			}

			if v || carry {
				carry = true
			} else {
				carry = false
			}

			for j := i + 1; j < length; j++ {
				if !carry {
					break
				}

				v, err := c.Get(i)
				if err != nil {
					t.Error(err)
				}

				if carry == v {
					carry = true
					if err := c.Clear(j); err != nil {
						t.Error(err)
					}
				} else {
					carry = false
					if err := c.Set(j); err != nil {
						t.Error(err)
					}
				}
			}
		}

		sum, carryOut, err := Add(a, b, false)
		if err != nil {
			t.Error(err)
		}

		if carry != carryOut {
			t.Errorf("value does not match %v %v %v", length, carry, carryOut)
		}

		for i := 0; i < length; i++ {
			x, err := c.Get(i)
			if err != nil {
				t.Error(err)
			}

			y, err := sum.Get(i)
			if err != nil {
				t.Error(err)
			}

			if x != y {
				t.Errorf("value does not match %v %v %v %v", length, i, x, y)
			}
		}
	}
}

func TestBitArray_Append(t *testing.T) {
	for length := 0; length < 1000; length++ {
		x, err := NewBitArray(length)
		if err != nil {
			t.Fatal(err)
		}

		interval := 7
		for i := 0; i < length; i += interval {
			if err := x.Set(i); err != nil {
				t.Error(err)
			}
		}

		for i := 0; i < 1000; i++ {
			y, err := NewBitArray(i)
			if err != nil {
				t.Fatal(err)
			}

			z, err := NewBitArray(length + i)
			if err != nil {
				t.Fatal(err)
			}

			for j := 0; j < length; j += interval {
				if err := z.Set(j); err != nil {
					t.Error(err)
				}
			}

			for j := 0; j < i; j += 11 {
				if err := y.Set(j); err != nil {
					t.Error(err)
				}

				if err := z.Set(j + length); err != nil {
					t.Error(err)
				}
			}

			appended, err := x.Append(y)
			if err != nil {
				t.Fatal(err)
			}

			for j := 0; j < appended.length; j++ {
				a, err := appended.Get(j)
				if err != nil {
					t.Error(err)
				}

				b, err := z.Get(j)
				if err != nil {
					t.Error(err)
				}

				if a != b {
					t.Errorf("value does not match %v %v %v %v %v", length, i, j, a, b)
				}
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
