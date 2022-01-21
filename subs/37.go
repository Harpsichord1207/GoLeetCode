package subs

import (
	"fmt"
	"strings"
)

// http://norvig.com/top95solutions.html

var Digits = "123456789"
var Rows = "ABCDEFGHI"
var Cols = Digits
var Squares []string
var UnitList [][]string
var UnitMap map[string][][]string
var PeerMap map[string]map[string]bool
var EmptyMap = map[string]string{}

type Result37 struct {
	OK   bool
	Data map[string]string
}

func initConstants() {
	cross := func(a string, b string) []string {
		var c []string
		for _, i := range a {
			for _, j := range b {
				c = append(c, string(i)+string(j))
			}
		}
		return c
	}

	Squares = cross(Rows, Cols)

	// UnitList
	for _, d := range Digits {
		UnitList = append(UnitList, cross(Rows, string(d)))
	}
	for _, r := range Rows {
		UnitList = append(UnitList, cross(string(r), Cols))
	}
	for _, rs := range []string{"ABC", "DEF", "GHI"} {
		for _, cs := range []string{"123", "456", "789"} {
			UnitList = append(UnitList, cross(rs, cs))
		}
	}

	// UnitMap
	UnitMap = make(map[string][][]string)
	for _, s := range Squares {
		var sUnits [][]string
		for _, unit := range UnitList {
			isIn := false
			for _, u := range unit {
				if u == s {
					isIn = true
					break
				}
			}
			if isIn {
				sUnits = append(sUnits, unit)
			}

		}
		UnitMap[s] = sUnits
	}

	// PeerMap
	PeerMap = make(map[string]map[string]bool)
	for _, s := range Squares {
		setMap := make(map[string]bool)
		for _, unit := range UnitMap[s] {
			for _, u := range unit {
				if u != s {
					setMap[u] = true
				}
			}
		}
		PeerMap[s] = setMap
	}

}

func assign(values map[string]string, s string, d string) map[string]string {
	otherValues := strings.Replace(values[s], d, "", -1)
	for _, ov := range otherValues {
		if !eliminate(values, s, string(ov)).OK {
			return make(map[string]string)
		}
	}
	return values
}

func eliminate(values map[string]string, s string, d string) Result37 {
	// 消除values中位置为s的值中的d

	// 已经没有d了
	if !strings.Contains(values[s], d) {
		return Result37{true, values}
	}
	values[s] = strings.Replace(values[s], d, "", -1)
	if len(values[s]) == 0 {
		return Result37{false, EmptyMap}
	}
	if len(values[s]) == 1 {
		d2 := values[s]
		for s2, v_ := range PeerMap[s] {
			if !v_ {
				continue
			}
			if !eliminate(values, s2, d2).OK {
				return Result37{false, EmptyMap}
			}
		}
	}

	for _, u := range UnitMap[s] {
		var dPlaces []string
		for _, s_ := range u {
			if strings.Contains(values[s_], d) {
				dPlaces = append(dPlaces, s_)
			}
		}

		if len(dPlaces) == 0 {
			return Result37{false, EmptyMap}
		}
		if len(dPlaces) == 1 {
			if len(assign(values, dPlaces[0], d)) == 0 {
				return Result37{false, EmptyMap}
			}
		}
	}
	return Result37{true, values}
}

func copyMap(values map[string]string) map[string]string {
	copyValues := make(map[string]string)
	for k, v := range values {
		copyValues[k] = v
	}
	return copyValues
}

func gridValues(gridBoard [][]byte) map[string]string {
	var chars []string
	for _, row := range gridBoard {
		for _, val := range row {
			chars = append(chars, string(val))
		}
	}

	m := make(map[string]string)
	for i := 0; i < 81; i++ {
		m[Squares[i]] = chars[i]
	}
	return m
}

func parseGrid(gridBoard [][]byte) map[string]string {
	values := make(map[string]string)
	for _, s := range Squares {
		values[s] = Digits
	}
	for s, d := range gridValues(gridBoard) {
		if strings.Contains(Digits, d) && len(assign(values, s, d)) == 0 {
			return make(map[string]string) // length == 0 --> false
		}
	}
	return values
}

func search(values map[string]string) Result37 {
	if len(values) == 0 {
		return Result37{false, EmptyMap}
	}

	isSolved := true
	for _, v := range values {
		if len(v) != 1 {
			isSolved = false
		}
	}

	if isSolved {
		return Result37{true, values}
	}

	minLen := 10
	var targetS string
	for _, s := range Squares {
		if len(values[s]) > 1 && len(values[s]) < minLen {
			minLen = len(values[s])
			targetS = s
		}
	}

	for _, d := range values[targetS] {
		res := search(assign(copyMap(values), targetS, string(d)))
		if res.OK {
			return res
		}
	}

	return Result37{false, EmptyMap}
}

func display(values map[string]string) {
	var width = 0
	for _, s := range Squares {
		if len(values[s]) > width {
			width = len(values[s])
		}
	}
	width++

	linePart1 := strings.Repeat("-", width*3)
	line := linePart1 + "+" + linePart1 + "+" + linePart1

	for _, r := range Rows {
		var row string
		for _, c := range Cols {
			row += values[string(r)+string(c)]
			if strings.Contains("36", string(c)) {
				row += "|"
			}

		}
		println(row)
		if strings.Contains("CF", string(r)) {
			println(line)
		}
	}

}

func solveSudoku(board [][]byte) {
	initConstants()
	result := search(parseGrid(board))
	for i, r := range Rows {
		for j, c := range Cols {
			board[i][j] = []byte(result.Data[string(r)+string(c)])[0]
		}
	}

}

func Test37() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	fmt.Printf("%#v", board)
}
