package subs

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	m := make(map[byte]int)
	index := make(map[byte]int)
	sb := []byte(s)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
		index[s[i]] = i
	}

	sort.Slice(sb, func(i, j int) bool {
		if m[sb[i]] > m[sb[j]] {
			return true
		}
		if m[sb[i]] < m[sb[j]] {
			return false
		}
		return index[sb[i]] < index[sb[j]]
	})
	return string(sb)
}

func Test451() {
	fmt.Println(frequencySort("tree"))
}
