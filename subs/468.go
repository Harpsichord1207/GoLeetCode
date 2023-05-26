package subs

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if strings.Contains(queryIP, ".") {
		parts := strings.Split(queryIP, ".")
		if len(parts) != 4 {
			return "Neither"
		}
		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil || (len(part) > 1 && part[0] == '0') || number < 0 || number > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}

	if strings.Contains(queryIP, ":") {
		parts := strings.Split(queryIP, ":")
		if len(parts) != 8 {
			return "Neither"
		}
		for _, part := range parts {
			_, err := strconv.ParseInt(part, 16, 32)
			if err != nil || len(part) > 4 || len(part) == 0 {
				return "Neither"
			}
		}
		return "IPv6"
	}

	return "Neither"
}

func Test468() {
	fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("192.0.0.1"))
}
