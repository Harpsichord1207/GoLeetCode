package subs

func findClosestElements(arr []int, k int, x int) []int {
	i, j := 0, len(arr)-k
	for i < j {
		m := i + (j-i)/2
		if x-arr[m] > arr[m+k]-x {
			i = m + 1
		} else {
			j = m
		}
	}
	return arr[i : i+k]
}
