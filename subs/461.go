package subs

func hammingDistance(x int, y int) int {
	z := x ^ y
	c := 0

	for z > 0 {
		c += 1
		z = z & (z - 1)
	}
	return c
}

func Test461() {
	println(hammingDistance(1, 3))
	println(hammingDistance(1, 4))
}
