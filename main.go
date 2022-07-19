package main

import (
	"LeetCode/subs"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func checkFileLinesCount(fullPath string) {
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read %s", fullPath))
	}
	count := 0
	for _, line := range strings.Split(string(content), "\n") {
		if len(strings.Replace(line, " ", "", -1)) > 0 {
			count++
		}
	}
	if count < 20 {
		fmt.Printf("File %s has too few lines: %d\n", fullPath, count)
	}
}

func generateReadMe() {
	_, file, _, _ := runtime.Caller(1)
	dir, _ := filepath.Split(file)
	rd, _ := ioutil.ReadDir(dir + "subs")
	text := "## LeetCode By Go\n\n### Finished Subjects\n\n|"
	i := 0

	urlPrefix := "https://github.com/Harpsichord1207/GoLeetCode/blob/main/subs/"
	var subNumbers []int
	for _, fi := range rd {
		checkFileLinesCount(filepath.Join(dir, "subs", fi.Name()))
		no := strings.ReplaceAll(fi.Name(), ".go", "")
		noInt, _ := strconv.Atoi(no)
		subNumbers = append(subNumbers, noInt)
	}
	sort.Ints(subNumbers)
	total := len(subNumbers)
	rowSize := 10
	for _, subNumber := range subNumbers {
		no := fmt.Sprintf(" [%04s](%s%s.go) ", strconv.Itoa(subNumber), urlPrefix, strconv.Itoa(subNumber))
		text += no + "|"
		i += 1
		if i%rowSize == 0 {
			text += "  \n"
			if i != total {
				text += "|"
			}
		}
	}
	readMeFile := ""
	for _, p := range strings.Split(dir, "/") {
		if p == "src" {
			break
		}
		readMeFile += p + "/"
	}
	readMeFile += "ReadMe.md"
	_ = ioutil.WriteFile(readMeFile, []byte(text), 0666)

	fmt.Println("Finish to write ReadMe.md")
}

func main() {
	generateReadMe()
	fmt.Println("=====================================")
	fmt.Println("Subs output:")
	fmt.Println("-------------------------------------")
	subs.Test2341()
}
