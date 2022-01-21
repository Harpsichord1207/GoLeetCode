package subs

import (
	"fmt"
	"unicode"
)

func detectCapitalUse(word string) bool {
	firstRuneIsUpper := false
	secondRuneIsUpper := false
	for i, r := range word {
		if i == 0 {
			if unicode.IsUpper(r) {
				firstRuneIsUpper = true
			}
		} else if i == 1 {
			if unicode.IsUpper(r) {
				secondRuneIsUpper = true
			}
			if !firstRuneIsUpper && secondRuneIsUpper {
				return false
			}
		} else {
			if secondRuneIsUpper != unicode.IsUpper(r) {
				return false
			}
		}
	}
	return true
}

func Test520() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("FlaG"))
	fmt.Println(detectCapitalUse("leetcode"))
	fmt.Println(detectCapitalUse("Google"))
}
