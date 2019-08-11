package bitarray

import (
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
