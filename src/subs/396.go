package subs

func maxRotateFunction(nums []int) int {
	var sum int
	var f int
	for i, n := range nums {
		sum += n
		f += i * n
	}
	var res = f
	size := len(nums)
	for i:=1;i<size;i++{
		f = f + sum - size * nums[size-i]
		if f > res {
			res = f
		}
	}
	return res
}
