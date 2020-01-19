package pkg

// Connected checks if the given nodes are connected by a path.
func (g *Graph) Connected(a, b int) bool {
	// mission 1: find out if there's a path from a to b
	return g.bfs(a, b)
}

func (g *Graph) dfs(a, b int) bool {
	// solution for part 1 using DFS (depth-first search)
	discovered := make([]bool, g.NumNodes())
	var f func(int) bool
	f = func(v int) bool {
		discovered[v] = true
		if v == b {
			return true
		}
		for _, n := range g.Neighbors(v) {
			if !discovered[n] {
				if f(n) {
					return true
				}
			}
		}
		return false
	}
	return f(a)
}

func (g *Graph) bfs(a, b int) bool {
	// solution for part 1 using BFS (breadth-first search)
	var q Queue
	q.Enqueue(a)
	discovered := make([]bool, g.NumNodes())
	discovered[a] = true
	for q.Length() > 0 {
		v := q.Dequeue()
		if v == b {
			return true
		}
		for _, n := range g.Neighbors(v) {
			if !discovered[n] {
				discovered[n] = true
				q.Enqueue(n)
			}
		}
	}
	return false
}

// ShortestPath returns a path between the nodes with the minimum number of edges.
func (g *Graph) ShortestPath(a, b int) []int {
	// mission 2: find the shortest path from a to b

	// depth-first search, using parent to store how we got to each node
	var q Queue
	q.Enqueue(a)
	parent := make([]int, g.NumNodes())
	discovered := make([]bool, g.NumNodes())
	discovered[a] = true
	for q.Length() > 0 {
		v := q.Dequeue()
		if v == b {
			break
		}
		for _, w := range g.Neighbors(v) {
			if !discovered[w] {
				parent[w] = v
				discovered[w] = true
				q.Enqueue(w)
			}
		}
	}

	// follow "parent" back from b to a
	reversePath := []int{b}
	next := b
	for next != a {
		next = parent[next]
		reversePath = append(reversePath, next)
	}

	// reverse the slice to get the path from a to b
	l := len(reversePath)
	path := make([]int, len(reversePath))
	for i, x := range reversePath {
		path[l-i-1] = x
	}

	return path
}
