package subs

import "strings"

func licenseKeyFormatting(s string, k int) string {

	var arr []rune

	for _, c := range s {
		if c == '-' {
			continue
		}
		arr = append(arr, c)
	}

	length := len(arr)
	group := length / k
	firstPart := length % k

	var strs []string

	if firstPart > 0 {
		strs = append(strs, string(arr[:firstPart]))
	}
	for i := 0; i < group; i++ {
		strs = append(strs, string(arr[i*k+firstPart:i*k+k+firstPart]))
	}

	return strings.ToUpper(strings.Join(strs, "-"))
}
