package subs

func isAnagram(s string, t string) bool {
	m := make([]int, 26)

	if len(t) != len(s) {
		return false
	}

	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, n := range m {
		if n != 0 {
			return false
		}
	}
	return true
}

func Test242() {
	isAnagram("abcde", "abcde")
}
