package pkg

// Connected checks if the given nodes are connected by a path.
func (g *Graph) Connected(a, b int) bool {
	// mission 1: find out if there's a path from a to b
	return false
}

// ShortestPath returns a path between the nodes with the minimum number of edges.
func (g *Graph) ShortestPath(a, b int) []int {
	// mission 2: find the shortest path from a to b
	return []int{}
}

// reversePath will reverse the path, e.g. [5, 7, 2] becomes [2, 7, 5].
func reversePath(path []int) {
	for i := 0; i < len(path); i++ {
		j := len(path) - 1 - i
		path[i], path[j] = path[j], path[i]
	}
}
