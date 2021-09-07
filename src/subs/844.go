package subs

func backspaceCompare(s string, t string) bool {
	var st1 []byte
	var st2 []byte

	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			st1 = append(st1, s[i])
		} else if len(st1) > 0 {
			st1 = st1[:len(st1)-1]
		}
	}

	for i := 0; i < len(t); i++ {
		if t[i] != '#' {
			st2 = append(st2, t[i])
		} else if len(st2) > 0 {
			st2 = st2[:len(st2)-1]
		}
	}

	return string(st1) == string(st2)
}
