package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
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
	for _, fi := range rd {
		no := fmt.Sprintf(" [%04s](%s%s) ",strings.Replace(fi.Name(), ".go", "", -1), urlPrefix, fi.Name())
		text += no + "|"
		i += 1
		if i % 5 == 0 {text += "  \n|"}
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
	subs.Test355()
	generateReadMe()
}
