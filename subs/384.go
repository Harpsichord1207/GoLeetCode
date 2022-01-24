package subs

import "math/rand"

type Solution384 struct {
	data0  []int
	data1  []int
	length int
}

func Constructor384(nums []int) Solution384 {
	solution := Solution384{}
	solution.data0 = nums
	solution.data1 = make([]int, len(nums))
	copy(solution.data1, nums)
	solution.length = len(nums)
	return solution
}

func (this *Solution384) Reset() []int {
	return this.data0
}

func (this *Solution384) Shuffle() []int {
	for i := 0; i < this.length; i++ {
		j := rand.Intn(this.length-i) + 1
		this.data1[i], this.data1[j] = this.data1[j], this.data1[i]
	}
	return this.data1
}
