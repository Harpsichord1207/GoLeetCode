package subs

import (
	"fmt"
	"sort"
)

func customSortString(order string, str string) string {
	strByte := []byte(str)
	m := make(map[byte]int)
	for i:=0;i<len(order);i++{
		m[order[i]] = i+1
	}
	
	sort.Slice(strByte, func(i, j int) bool {
		return m[strByte[i]] < m[strByte[j]]
	})
	
	return string(strByte)
}

func Test791()  {
	fmt.Println(customSortString("cbd", "abcd"))
}
