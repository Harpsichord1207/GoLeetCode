package subs

func restoreString(s string, indices []int) string {

	cs := make([]rune, len(indices))
	for i, c := range s {
		cs[indices[i]] = c
	}
	return string(cs)
}
