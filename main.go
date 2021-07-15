package graph

import "fmt"

//matrix is a 2d slice for crate structure of graph
var matrix [][]bool
var Node int

//Create is function for create 2d graph by size of node
func Create(node int) {
	Node = node
	matrix = make([][]bool, node)
	for i := range matrix {
		matrix[i] = make([]bool, node)
	}
}

//AddEdge is function for add edge to graph
func AddEdge(i, j int8) {
	matrix[i][j] = true
	//matrix[j][i] = true
}

//DFS for depth first search for graph
func DFS(start int) {
	visited := make([]bool, Node)
	dfsUtil(start, visited)
}

//dfsUtil is a recursive function that used in DFS function
func dfsUtil(start int, visited []bool) {
	visited[start] = true
	fmt.Print(start, " ")
	for i := range matrix[start] {
		if matrix[start][i] && !visited[i] {
			dfsUtil(i, visited)
		}
	}
}
//BFS for  breadth search for graph
func BFS(start int) {
	visited := make([]bool, Node)
	queue := make([]int, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		node := queue[0]
		visited[node] = true
		fmt.Print(node, " ")
		queue = queue[1:]
		for i := range matrix[node] {
			if matrix[node][i] && !visited[i] {
				queue = append(queue, i)
			}
		}
	}
}

