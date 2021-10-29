package main

import (
	"fmt"
	"strings"
)

func cross(a string, b string) []string {
	var c []string
	for _, i := range a {
		for _, j := range b {
			c = append(c, string(i)+string(j))
		}
	}
	return c
}

var digits = "123456789"
var rows   = "ABCDEFGHI"
var cols   = digits
var squares  = cross(rows, cols)
var unitlist [][]string
var units map[string][][]string
var peers map[string]map[string]bool


func init()  {
	// unitlist
	for _, d := range digits {
		unitlist = append(unitlist, cross(rows, string(d)))
	}
	for _, r := range rows {
		unitlist = append(unitlist, cross(string(r), cols))
	}

	for _, rs := range []string{"ABC", "DEF", "GHI"} {
		for _, cs := range []string{"123", "456", "789"} {
			unitlist = append(unitlist, cross(rs, cs))
		}
	}

	// units
	units = make(map[string][][]string)
	for _, s := range squares {
		var s_units [][]string
		for _, unit := range unitlist {
			is_in := false
			for _, u := range unit {
				if u == s {
					is_in = true
					break
				}
			}
			if is_in {
				s_units = append(s_units, unit)
			}

		}

		units[s] = s_units
	}

	// peers
	peers = make(map[string]map[string]bool)
	for _, s := range squares{
		set_map := make(map[string]bool)
		for _, unit := range units[s] {
			for _, u := range unit {
				if u != s {
					set_map[u] = true
				}
			}
		}
		peers[s] = set_map
	}
}

func gridValues(grid string) map[string]string {
	var chars []string
	for _, c := range grid {
		cString := string(c)
		if strings.Contains(digits, cString) || cString == "0" || cString == "." {
			chars = append(chars, cString)
		}
	}
	if len(chars) != 81 {
		panic(fmt.Sprintf("Invalid Grid: %s", grid))
	}
	m := make(map[string]string)
	for i:=0; i<81; i++ {
		m[squares[i]] = chars[i]
	}
	return m
}

func assign(values map[string]string, s string, d string) map[string]string {
	otherValues := strings.Replace(values[s], d, "", -1)
	for _, ov := range otherValues {
		if len(eliminate(values, s, string(ov))) == 0 {
			return make(map[string]string)
		}
	}
	return values
}

func eliminate(values map[string]string, s string, d string) map[string]string {
	if !strings.Contains(values[s], d) {
		return make(map[string]string)
	}
	values[s] = strings.Replace(values[s], d, "", -1)
	if len(values[s]) == 0 {
		return make(map[string]string)
	}
	if len(values[s]) == 1 {
		d2 := values[s]
		for s2, v_ := range peers[s] {
			if !v_ {
				continue
			}
			if len(eliminate(values, s2, d2)) == 0 {
				return make(map[string]string)
			}
		}
	}

	for _, u := range units[s] {
		if !strings.Contains(values[s], d) {
			return make(map[string]string)
		}
		if len(u) == 1 {
			if len(assign(values, u[0], d)) == 0 {
				return make(map[string]string)
			}
		}
	}
	return values
}

func parseGrid(grid string) map[string]string {
	values := make(map[string]string)
	for _, s := range squares {
		values[s] = digits
	}
	for s, d := range gridValues(grid) {
		if strings.Contains(digits, d) && len(assign(values, s, d)) == 0 {
			return make(map[string]string)  // length == 0 --> false
		}
	}
	return values
}



var grid1 = "003020600900305001001806400008102900700000008006708200002609500800203009005010300"


func main()  {
	//init()
	fmt.Println(parseGrid(grid1))
}
