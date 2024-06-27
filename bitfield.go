package bitfield

import "strings"

const (
	BITS_PER_WORD = 32
)

type BitfieldWord_t uint32

type Bitfield struct {
	bits []BitfieldWord_t
	len  int
}

func NewBitfield(len int) *Bitfield {
	if len <= 0 {
		return nil
	}

	return &Bitfield{
		len:  len,
		bits: make([]BitfieldWord_t, (len+BITS_PER_WORD-1)/BITS_PER_WORD),
	}
}

func (bf Bitfield) Len() int {
	return bf.len
}

func (bf Bitfield) String() string {
	var sb strings.Builder
	for i := bf.len - 1; i >= 0; i-- {
		if bf.GetBit(i) {
			sb.WriteByte('1')
		} else {
			sb.WriteByte('0')
		}
	}
	return sb.String()
}

func (bf *Bitfield) GetBit(i int) bool {
	if i >= bf.len || i < 0 {
		return false
	}
	return (bf.bits[i/BITS_PER_WORD]>>(i%BITS_PER_WORD))&0b1 == 1
}

func (bf *Bitfield) SetBit(i int) {
	if i < bf.len && i >= 0 {
		bf.bits[i/BITS_PER_WORD] |= 0b1 << (i % BITS_PER_WORD)
	}
}

func (bf *Bitfield) SetBits(j []int) {
	for _, i := range j {
		bf.SetBit(i)
	}
}

func (bf *Bitfield) UnsetBit(i int) {
	if i < bf.len && i >= 0 {
		bf.bits[i/BITS_PER_WORD] &= ^(0b1 << (i % BITS_PER_WORD))
	}
}

func (bf *Bitfield) UnsetBits(j []int) {
	for _, i := range j {
		bf.UnsetBit(i)
	}
}

func Not(a *Bitfield) *Bitfield {
	b := NewBitfield(a.len)
	for i := range len(b.bits) {
		b.bits[i] = ^a.bits[i]
	}
	return b
}

func And(a, b *Bitfield) *Bitfield {
	if a.len != b.len {
		return nil
	}
	c := NewBitfield(a.len)
	for i := range len(c.bits) {
		c.bits[i] = a.bits[i] & b.bits[i]
	}
	return c
}

func Or(a, b *Bitfield) *Bitfield {
	if a.len != b.len {
		return nil
	}
	c := NewBitfield(a.len)
	for i := range len(c.bits) {
		c.bits[i] = a.bits[i] | b.bits[i]
	}
	return c
}

func Xor(a, b *Bitfield) *Bitfield {
	if a.len != b.len {
		return nil
	}
	c := NewBitfield(a.len)
	for i := range len(c.bits) {
		c.bits[i] = a.bits[i] ^ b.bits[i]
	}
	return c
}
