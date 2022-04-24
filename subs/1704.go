package subs

func halvesAreAlike(s string) bool {
	h := len(s) / 2
	fp, sp := 0, 0
	m := make(map[byte]bool)
	m[byte('a')] = true
	m[byte('e')] = true
	m[byte('i')] = true
	m[byte('o')] = true
	m[byte('u')] = true
	m[byte('A')] = true
	m[byte('E')] = true
	m[byte('I')] = true
	m[byte('O')] = true
	m[byte('U')] = true
	for i := 0; i < h; i++ {
		if m[s[i]] {
			fp++
		}
		if m[s[i+h]] {
			sp++
		}
	}
	return fp == sp
}
