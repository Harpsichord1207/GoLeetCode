package subs

import "fmt"


func dfs207(graph [][]int, visit []int, i int) bool {
	if visit[i] == -1 {return false}
	if visit[i] == 1 {return true}

	visit[i] = -1
	for _, a := range graph[i] {
		if !dfs207(graph, visit, a) {return false}
	}
	visit[i] = 1
	return true
}


func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	visit := make([]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}
	for i:=0;i<numCourses;i++{
		if !dfs207(graph, visit, i) {return false}
	}
	return true
}

func Test207() {
	pre := [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}
	fmt.Println(canFinish(20, pre))
}