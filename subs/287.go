package subs

func findDuplicate(nums []int) int {

	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	res := 0
	for {
		res = nums[res]
		slow = nums[slow]
		if res == slow {
			break
		}
	}
	return res
}
