package subs

import "fmt"

const hexString string = "0123456789abcdef"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	res := ""
	for cnt := 0; cnt < 8 && num != 0; cnt++ {
		res = string(hexString[num&0xf]) + res
		num >>= 4
	}
	return res
}

func Test405() {
	fmt.Println(toHex(-714))
}
