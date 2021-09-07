package subs

import "fmt"

type element struct {
	value  string
	number int
}

func minMutation(start string, end string, bank []string) int {

	set := make(map[string]int)
	for _, s := range bank {
		set[s] += 1
	}
	queue := []element{{start, 1}}
	for {
		if len(queue) == 0 {
			break
		}
		e := queue[0]
		queue = append(queue[1:])
		for i := 0; i < len(e.value); i++ {
			tmp1 := string(e.value[:i])
			tmp2 := string(e.value[i+1:])
			for _, c := range "ACGT" {
				if uint8(c) == e.value[i] {
					continue
				}
				tmp := tmp1 + string(c) + tmp2
				if set[tmp] > 0 {
					if tmp == end {
						return e.number
					}
					set[tmp] = 0
					queue = append(queue, element{tmp, e.number + 1})
				}
			}
		}
	}
	return -1
}

func Test433() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
}
