package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"subs"
)

func generateReadMe() {
	_, file, _, _ := runtime.Caller(1)
	dir, _ := filepath.Split(file)
	rd, _ := ioutil.ReadDir(dir + "subs")
	text := "## LeetCode By Go\n\n### Finished Subjects\n\n|"
	i := 0
	urlPrefix := "https://github.com/Harpsichord1207/GoLeetCode/blob/main/src/subs/"
	var subNumbers []int
	for _, fi := range rd {
		no := strings.Replace(fi.Name(), ".go", "", -1)
		noInt, _ :=  strconv.Atoi(no)
		subNumbers = append(subNumbers, noInt)
	}
	sort.Ints(subNumbers)
	total := len(subNumbers)
	for _, subNumber := range subNumbers {
		no := fmt.Sprintf(" [%04s](%s%s.go) ", strconv.Itoa(subNumber), urlPrefix, strconv.Itoa(subNumber))
		text += no + "|"
		i += 1
		if i % 5 == 0 {
			text += "  \n"
			if i != total {
				text += "|"
			}
		}
	}
	readMeFile := ""
	for _, p := range strings.Split(dir, "/") {
		if p == "src" { break }
		readMeFile += p + "/"
	}
	readMeFile += "ReadMe.md"
	_ = ioutil.WriteFile(readMeFile, []byte(text), 0666)

	fmt.Println("Finish to write ReadMe.md")
}

func main(){
	subs.Test274()
	generateReadMe()
}
