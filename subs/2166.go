package subs

import "strings"

type Bitset struct {
	size  int
	zeros map[int]struct{}
	ones  map[int]struct{}
}

func Constructor2166(size int) Bitset {
	b := Bitset{size: size}
	b.zeros = make(map[int]struct{}, size)
	b.ones = make(map[int]struct{}, size)
	for i := 0; i < size; i++ {
		b.zeros[i] = struct{}{}
	}
	return b
}

func (b *Bitset) Fix(idx int) {
	if _, ok := b.ones[idx]; ok {
		return
	}
	b.ones[idx] = struct{}{}
	delete(b.zeros, idx)
}

func (b *Bitset) Unfix(idx int) {
	if _, ok := b.zeros[idx]; ok {
		return
	}
	b.zeros[idx] = struct{}{}
	delete(b.ones, idx)
}

func (b *Bitset) Flip() {
	b.ones, b.zeros = b.zeros, b.ones
}

func (b *Bitset) All() bool {
	return len(b.ones) == b.size
}

func (b *Bitset) One() bool {
	return len(b.ones) > 0
}

func (b *Bitset) Count() int {
	return len(b.ones)
}

func (b *Bitset) ToString() string {
	var sb strings.Builder
	for i := 0; i < b.size; i++ {
		if _, ok := b.zeros[i]; ok {
			sb.WriteString("0")
		} else {
			sb.WriteString("1")
		}
	}
	return sb.String()
}
