package bitfield_test

import (
	"testing"

	"github.com/49pctber/bitfield"
)

func TestBitfield0(t *testing.T) {
	bf := bitfield.NewBitfield(4)

	if have, want := bf.Len(), 4; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	bf.SetBit(1)
	bf.SetBit(3)

	if have, want := bf.GetBit(0), false; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
	if have, want := bf.GetBit(1), true; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
	if have, want := bf.GetBit(2), false; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
	if have, want := bf.GetBit(3), true; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	if have, want := bf.String(), "1010"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	bf.UnsetBit(1)

	if have, want := bf.GetBit(0), false; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
	if have, want := bf.GetBit(1), false; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
	if have, want := bf.GetBit(2), false; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
	if have, want := bf.GetBit(3), true; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	if have, want := bf.String(), "1000"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
}

func TestBitfield1(t *testing.T) {
	bf := bitfield.NewBitfield(47)

	if have, want := bf.Len(), 47; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	bf.SetBit(46)

	for i := range 46 {
		if have, want := bf.GetBit(i), false; have != want {
			t.Fatalf("have %v, want %v", have, want)
		}
	}
	if have, want := bf.GetBit(46), true; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
}

func TestBitfield2(t *testing.T) {
	bf := bitfield.NewBitfield(8)

	if have, want := bf.Len(), 8; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	bf.SetBits([]int{0, 3, 7, 2})

	if have, want := bf.String(), "10001101"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	bf.SetBits([]int{0, 3, 7, 2})

	if have, want := bf.String(), "10001101"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	bf.SetBits([]int{5, 4, 7, 2})
	bf.UnsetBits([]int{0, 1, 2})

	if have, want := bf.String(), "10111000"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	bf.SetBits([]int{10, -1})
	bf.UnsetBits([]int{11, -3})

	if have, want := bf.String(), "10111000"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
}

func TestBitfield3(t *testing.T) {
	bf := bitfield.NewBitfield(0)
	if bf != nil {
		t.Fatalf("cannot have length-0 bitfield")
	}

	bf = bitfield.NewBitfield(-1)
	if bf != nil {
		t.Fatalf("cannot have bitfield with negative length")
	}
}

func TestBitfieldNot(t *testing.T) {
	a := bitfield.NewBitfield(12)
	b := bitfield.NewBitfield(12)

	a.SetBits([]int{0, 1, 4, 5, 11})
	b.SetBits([]int{0, 2, 6, 11, 9})

	c := bitfield.Not(a)
	d := bitfield.Not(b)

	if have, want := c.String(), "011111001100"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	if have, want := d.String(), "010110111010"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
}

func TestBitfieldAnd(t *testing.T) {
	a := bitfield.NewBitfield(12)
	b := bitfield.NewBitfield(12)

	a.SetBits([]int{0, 1, 4, 5, 11})
	b.SetBits([]int{0, 2, 6, 11, 9})

	c := bitfield.And(a, b)
	d := bitfield.And(c, b)

	if have, want := c.String(), "100000000001"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	if have, want := d.String(), c.String(); have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
}

func TestBitfieldOr(t *testing.T) {
	a := bitfield.NewBitfield(12)
	b := bitfield.NewBitfield(12)

	a.SetBits([]int{0, 1, 4, 5, 11})
	b.SetBits([]int{0, 2, 6, 11, 9})

	c := bitfield.Or(a, b)
	d := bitfield.Or(c, b)

	if have, want := c.String(), "101001110111"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	if have, want := d.String(), c.String(); have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
}

func TestBitfieldXor(t *testing.T) {
	a := bitfield.NewBitfield(12)
	b := bitfield.NewBitfield(12)

	a.SetBits([]int{0, 1, 4, 5, 11})
	b.SetBits([]int{0, 2, 6, 11, 9})

	c := bitfield.Xor(a, b)
	d := bitfield.Xor(c, b)

	if have, want := c.String(), "001001110110"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	if have, want := d.String(), a.String(); have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
}
