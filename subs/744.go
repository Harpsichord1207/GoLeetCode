package subs

func nextGreatestLetter(letters []byte, target byte) byte {
	i, j := 0, len(letters)-1

	for i <= j {
		m := (i + j) / 2
		v := letters[m]
		if v > target {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return letters[i%len(letters)]
}
