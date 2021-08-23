package main

import (
	"container/list"
	"fmt"
)

// Graph 无向图
type Graph struct {
	Adj []*list.List
	V   int
}

// NewGraph ...
func NewGraph(v int) *Graph {
	graph := &Graph{}
	graph.V = v
	graph.Adj = make([]*list.List, v)
	for i := range graph.Adj {
		graph.Adj[i] = list.New()
	}
	return graph
}

// AddEdge ...
func (graph *Graph) AddEdge(s int, t int) {
	graph.Adj[s].PushBack(t)
	graph.Adj[t].PushBack(s)
}

// BFS 广度优先搜索
func (graph *Graph) BFS(s int, t int) {
	if s == t {
		return
	}

	// init prev
	prev := make([]int, graph.V)
	for index := range prev {
		prev[index] = -1
	}

	// search by queue
	var queue []int
	visited := make([]bool, graph.V)
	queue = append(queue, s)
	visited[s] = true
	isFound := false

	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linkedList := graph.Adj[top]
		for e := linkedList.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}
	if isFound {
		printPrev(prev, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}
}

// DFS 深度优先搜索
func (graph *Graph) DFS(s int, t int) {
	// init prev
	prev := make([]int, graph.V)
	for index := range prev {
		prev[index] = -1
	}

	visited := make([]bool, graph.V)

	visited[s] = true

	isFound := false
	graph.recurse(s, t, prev, visited, isFound)

	printPrev(prev, s, t)
}

//recursivly find path
func (graph *Graph) recurse(s int, t int, prev []int, visited []bool, isFound bool) {

	if isFound {
		return
	}

	visited[s] = true

	if s == t {
		isFound = true
		return
	}

	linkedlist := graph.Adj[s]
	for e := linkedlist.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			graph.recurse(k, t, prev, visited, false)
		}
	}

}

//print path recursively
func printPrev(prev []int, s int, t int) {

	if t == s || prev[t] == -1 {
		fmt.Printf("%d ", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf("%d ", t)
	}

}
