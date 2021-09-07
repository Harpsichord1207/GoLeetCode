package subs

import (
	"fmt"
	"strings"
)

func _isSameRow(word string, row string) bool {
	for _, w := range word {
		if !strings.ContainsRune(row, w) {
			return false
		}
	}
	return true
}

func isSameRow(word string) bool {
	return _isSameRow(word, "qwertyuiop") || _isSameRow(word, "asdfghjkl") || _isSameRow(word, "zxcvbnm")
}

func findWords(words []string) []string {
	var res []string
	for _, word := range words {
		lWord := strings.ToLower(word)
		if isSameRow(lWord) {
			res = append(res, word)
		}
	}
	return res
}

func Test500() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	res := findWords(words)
	fmt.Println(res)
}
