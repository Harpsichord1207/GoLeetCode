package subs

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	Data []int
	Map  map[int][]int
}

func Constructor398(nums []int) Solution {
	s := Solution{}
	s.Data = nums
	s.Map = make(map[int][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if v, ok := s.Map[n]; ok {
			v = append(v, i)
			s.Map[n] = v
		} else {
			s.Map[n] = []int{i}
		}
	}
	return s
}

func (s *Solution) Pick(target int) int {
	v := s.Map[target]
	size := len(v)
	return v[rand.Intn(size)]
}

func Test398() {
	s := Constructor398([]int{1, 2, 3, 3, 3})
	for i := 0; i < 100; i++ {
		fmt.Println(s.Pick(3))
		fmt.Println(s.Pick(2))
		fmt.Println(s.Pick(1))
	}

}
