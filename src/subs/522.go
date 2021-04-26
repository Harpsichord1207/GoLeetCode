package subs

import "fmt"

func isSub(a string, b string) bool {
	la := len(a)
	lb := len(b)
	if la > lb {return false}
	i := 0
	for j:=0;j<lb;j++{
		if b[j] == a[i]{
			i++
		}
		if i == la {
			break
		}
	}
	return i == la
}

func findLUSlength2(strs []string) int {
	l := len(strs)
	maxLength := -1
	for i:=0;i<l;i++{
		found := false
		for j:=0;j<l;j++{
			if i==j {continue}
			if isSub(strs[i], strs[j]){
				found = true
				break
			}
		}
		if !found {
			if len(strs[i]) > maxLength {
				maxLength = len(strs[i])
			}
		}
	}
	return maxLength
}


func Test522(){
	fmt.Println(findLUSlength2([]string{"aaa","aaa","aa"}))
}