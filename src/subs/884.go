package subs

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	m1 := make(map[string]int)
	for _, word := range strings.Fields(s1) {
		m1[word] ++
	}

	m2 := make(map[string]int)
	for _, word := range strings.Fields(s2) {
		m2[word] ++
	}

	var r []string
	for word, count := range m1 {
		if count > 1 || m2[word] > 0 {
			delete(m2, word)
		} else {
			r = append(r, word)
		}
	}
	for word, count := range m2 {
		if count == 1 {
			r = append(r, word)
		}
	}

	return r
}
