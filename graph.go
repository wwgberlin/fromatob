package main

type Graph struct {
	nodes []string
	edges [][]int
}

func (g *Graph) Nodes() int {
	return len(g.nodes)
}

func (g *Graph) Edges() int {
	count := 0
	for _, list := range g.edges {
		count += len(list)
	}
	return count / 2
}

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

func (g *Graph) LookupNode(name string) (int, bool) {
	for i, n := range g.nodes {
		if n == name {
			return i, true
		}
	}
	return 0, false
}

func (g *Graph) AddEdge(i, j int) {
	if g.HasEdge(i, j) {
		return
	}
	g.edges[i] = append(g.edges[i], j)
	g.edges[j] = append(g.edges[j], i)
}

func (g *Graph) HasEdge(i, j int) bool {
	for _, k := range g.edges[i] {
		if k == j {
			return true
		}
	}
	return false
}
