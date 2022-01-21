package subs

type Iterator struct{}

func (itr *Iterator) hasNext() bool {
	return true
}

func (itr *Iterator) next() int {
	return 1
}

type PeekingIterator struct {
	it       *Iterator
	value    int
	hasValue bool
}

func Constructor284(iter *Iterator) *PeekingIterator {
	pi := PeekingIterator{}
	pi.it = iter
	return &pi
}

func (pki *PeekingIterator) hasNext() bool {
	return pki.hasValue || pki.it.hasNext()
}

func (pki *PeekingIterator) next() int {
	if pki.hasValue {
		num := pki.value
		pki.hasValue = false
		return num
	}
	return pki.it.next()
}

func (pki *PeekingIterator) peek() int {
	if !pki.hasValue {
		pki.value = pki.it.next()
		pki.hasValue = true
	}
	return pki.value
}

func Test284() {
	Constructor284(&Iterator{})
}
