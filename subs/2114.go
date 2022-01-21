package subs

func mostWordsFound(sentences []string) int {
	maxLength := 0
	for _, sentence := range sentences {
		wordCount := getWordCount((sentence))
		if wordCount > maxLength {
			maxLength = wordCount
		}
	}
	return maxLength
}

func getWordCount(sentence string) int {
	res := 1
	for _, c := range sentence {
		if c == ' ' {
			res++
		}
	}
	return res
}
