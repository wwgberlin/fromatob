package pkg_test

import (
	"testing"

	"github.com/fromAtoB/go-workshop/pkg"
	"fmt"
	"time"
	"reflect"
)

func TestGraph_ConnectedNoCycle(t *testing.T) {
	t.Log("Test simple directed graph with no cycles")
	var g pkg.Graph
	a := g.AddNode("A")
	b := g.AddNode("B")
	c := g.AddNode("C")
	_ = g.AddNode("D")

	g.AddEdgeDirected(a, b)
	g.AddEdgeDirected(b, c)

	tcases := []struct {
		a, b string
		res  bool
	}{
		{"A", "C", true},
		{"C", "A", false},
		{"A", "D", false},
		{"A", "B", true},
		{"B", "D", false},
	}
	for _, c := range tcases {
		done := make(chan struct{})
		go func() {
			if connected(g, c.a, c.b) != c.res {
				t.Errorf("the connection between '%s' and '%s' was expected to be: %t", c.a, c.b, c.res)
			}
			close(done)
		}()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatalf("timeout at 'Connected' between '%s' and '%s'. Possible endless loop", c.a, c.b)
		}
	}
}
func TestGraph_ConnectedNonDirected(t *testing.T) {
	t.Log("Test non directed graph (for each edge a->b there's an edge b->a)")
	var g pkg.Graph
	a := g.AddNode("A")
	b := g.AddNode("B")
	c := g.AddNode("C")
	_ = g.AddNode("D")

	g.AddEdge(a, b)
	g.AddEdge(b, c)
	g.AddEdge(b, c)

	tcases := []struct {
		a, b string
		res  bool
	}{
		{"A", "C", true},
		{"A", "D", false},
		{"D", "A", false},
		{"A", "B", true},
		{"B", "D", false},
	}
	for _, c := range tcases {
		done := make(chan struct{})
		go func() {
			if connected(g, c.a, c.b) != c.res {
				t.Errorf("the connection between '%s' and '%s' was expected to be: %t", c.a, c.b, c.res)
			}
			close(done)
		}()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatalf("timeout at 'Connected' between '%s' and '%s'. Possible endless loop", c.a, c.b)
		}
	}
}

func TestGraph_ShortestPath(t *testing.T) {
	t.Log("Test non directed graph (for each edge a->b there's an edge b->a)")
	var g pkg.Graph
	a := g.AddNode("A")
	b := g.AddNode("B")
	c := g.AddNode("C")
	_ = g.AddNode("D")

	g.AddEdge(a, b)
	g.AddEdge(b, c)
	g.AddEdge(b, c)

	tcases := []struct {
		a, b string
		res  []int
	}{
		{"A", "C", []int{0, 1, 2}},
		{"C", "A", []int{2, 1, 0}},
		{"A", "D", nil},
		{"B", "C", []int{1, 2}},
		{"C", "B", []int{2, 1}},
		{"B", "D", nil},
	}
	for _, c := range tcases {
		done := make(chan struct{})
		go func() {
			res := shortestPath(g, c.a, c.b)
			if c.res == nil {
				if len(res) != 0 {
					t.Errorf("unexpected shortest path from '%s' to '%s'. Expected an empty path (not connected), but received '%s'",
						c.a, c.b, g.PrintPath(res, "->"))
				}
			}
			if !reflect.DeepEqual(c.res, res) {
				if !g.PathExists(res) {
					t.Errorf("unexpected shortest path from %s to %s: %s\n", c.a, c.b, g.PrintPath(res, "->"))
				} else if len(c.res) != len(res) {
					t.Errorf("unexpected shortest path from %s to %s: %s has length %d, but we can do better!\n",
						c.a, c.b, g.PrintPath(res, "-"), len(res))
				}

			}
			close(done)
		}()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatalf("timeout at 'Connected' between '%s' and '%s'. Possible endless loop", c.a, c.b)
		}
	}

}

func connected(g pkg.Graph, a, b string) bool {
	ai, ok := g.LookupNode(a)
	if !ok {
		panic(fmt.Sprintf("'%s' not found in graph", a))
	}
	bi, ok := g.LookupNode(b)
	if !ok {
		panic(fmt.Sprintf("'%s' not found in graph", b))
	}
	return g.Connected(ai, bi)
}

func shortestPath(g pkg.Graph, a, b string) []int {
	ai, ok := g.LookupNode(a)
	if !ok {
		panic(fmt.Sprintf("'%s' not found in graph", a))
	}
	bi, ok := g.LookupNode(b)
	if !ok {
		panic(fmt.Sprintf("'%s' not found in graph", b))
	}
	return g.ShortestPath(ai, bi)
}
