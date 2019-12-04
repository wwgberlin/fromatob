package pkg

import "strings"

// A Graph is an undirected graph.
type Graph struct {
	nodes []string
	edges [][]int
}

// Nodes returns the number of nodes (or vertices) in the graph.
func (g *Graph) Nodes() int {
	return len(g.nodes)
}

// Edges returns the number of edges in the graph.
func (g *Graph) Edges() int {
	count := 0
	for _, list := range g.edges {
		count += len(list)
	}
	return count / 2
}

// Neighbors returns the list of neighbors of a given node.
func (g *Graph) Neighbors(i int) []int {
	return g.edges[i]
}

// PathExists checks if the given path exists in the graph. The path is specified by a list of the
// node IDs in the graph.
func (g *Graph) PathExists(path []int) bool {
	for i := range path {
		if i == 0 {
			continue
		}
		a := path[i-1]
		b := path[i]
		if !g.HasEdge(a, b) {
			return false
		}
	}
	return true
}

// NodeName returns the name for a node ID.
func (g *Graph) NodeName(n int) string {
	return g.nodes[n]
}

// PrintPath prints the names of the nodes in the path, separated by sep.
func (g *Graph) PrintPath(path []int, sep string) string {
	names := make([]string, len(path))
	for i, n := range path {
		names[i] = g.NodeName(n)
	}
	return strings.Join(names, sep)
}

// AddNode adds a node to the graph and returns the node ID.
func (g *Graph) AddNode(name string) int {
	i, ok := g.LookupNode(name)
	if ok {
		return i
	}
	i = len(g.nodes)
	g.nodes = append(g.nodes, name)
	g.edges = append(g.edges, []int{})
	return i
}

// LookupNode looks up a node by name. If the node can't be found, it returns ok == false.
func (g *Graph) LookupNode(name string) (id int, ok bool) {
	for i, n := range g.nodes {
		if n == name {
			return i, true
		}
	}
	return 0, false
}

// AddEdge adds an edge between the given nodes.
func (g *Graph) AddEdge(i, j int) {
	if g.HasEdge(i, j) {
		return
	}
	g.edges[i] = append(g.edges[i], j)
	g.edges[j] = append(g.edges[j], i)
}

// HasEdge returns true if there is an edge between the given nodes.
func (g *Graph) HasEdge(i, j int) bool {
	for _, k := range g.edges[i] {
		if k == j {
			return true
		}
	}
	return false
}
