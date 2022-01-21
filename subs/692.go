package subs

import (
	"container/heap"
	"fmt"
)

type WordFreq struct {
	Word string
	Freq int
}

type WordFreqHeap []WordFreq

func (w WordFreqHeap) Len() int {
	return len(w)
}

func (w WordFreqHeap) Less(i, j int) bool {
	wf1 := w[i]
	wf2 := w[j]
	if wf1.Freq > wf2.Freq {
		return true
	} else if wf1.Freq < wf2.Freq {
		return false
	} else {
		return wf1.Word < wf2.Word
	}
}

func (w WordFreqHeap) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w *WordFreqHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*w = append(*w, x.(WordFreq))
}

func (w *WordFreqHeap) Pop() interface{} {
	old := *w
	n := len(old)
	x := old[n-1]
	*w = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	wfh := &WordFreqHeap{}

	for word, freq := range m {
		*wfh = append(*wfh, WordFreq{Word: word, Freq: freq})
	}

	heap.Init(wfh)

	var res []string

	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(wfh).(WordFreq).Word)
	}
	return res
}

func Test692() {
	words := []string{"i", "love", "leetcode", "i", "love", "code"}
	fmt.Println(topKFrequent(words, 2))
}