package pkg

// Connected checks if the given nodes are connected by a path.
func (g *Graph) Connected(a, b int) bool {
	// mission 1: find out if there's a path from a to b

	var q Queue
	q.Enqueue(a)
	discovered := make([]bool, g.Nodes())
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

	var q Queue
	q.Enqueue(a)
	parent := make([]int, g.Nodes())
	discovered := make([]bool, g.Nodes())
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

	reversePath := []int{b}
	next := b
	for next != a {
		next = parent[next]
		reversePath = append(reversePath, next)
	}

	l := len(reversePath)
	path := make([]int, len(reversePath))
	for i, x := range reversePath {
		path[l-i-1] = x
	}

	return path
}

type Queue struct {
	list []int
}

func (q *Queue) Length() int {
	return len(q.list)
}

func (q *Queue) Enqueue(i int) {
	q.list = append(q.list, i)
}

func (q *Queue) Dequeue() int {
	i := q.list[0]
	q.list = q.list[1:]
	return i
}
