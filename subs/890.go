package subs

func findAndReplacePattern(words []string, pattern string) []string {

	isMatch := func(a string, b string) bool {
		s := len(a)
		m := make(map[uint8]uint8)
		for i := 0; i < s; i++ {
			v, ok := m[a[i]]
			if !ok {
				m[a[i]] = b[i]
			} else if v != b[i] {
				return false
			}
		}
		return true
	}

	var res []string
	for _, word := range words {
		if isMatch(word, pattern) && isMatch(pattern, word) {
			res = append(res, word)
		}
	}
	return res
}
