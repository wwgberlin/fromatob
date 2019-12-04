package main

import (
	"fmt"
	"os"

	"github.com/fromatob/go-workshop/pkg"
)

func main() {
	filename := "paths.md"
	g, err := pkg.LoadGraph(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Loaded %s (%d nodes, %d edges)\n", filename, g.Nodes(), g.Edges())

	mission1(g)
	mission2(g)
}

func mission1(g *pkg.Graph) {
	fmt.Print("MISSION 1   Are cities connected?\n")
	fmt.Println()

	cases := []struct {
		a, b string
		want bool
	}{
		{"Berlin", "München", true},
	}
	for _, c := range cases {
		fmt.Printf("%s -- %s ", c.a, c.b)
		a, _ := g.LookupNode(c.a)
		b, _ := g.LookupNode(c.b)
		got := g.Connected(a, b)
		if got == c.want {
			fmt.Println("✓")
		} else {
			fmt.Println("❌")
		}
	}
	fmt.Println()
}

func mission2(g *pkg.Graph) {
	fmt.Print("MISSION 2   Find the shortest path!\n")
	fmt.Println()

	cases := []struct {
		a, b       string
		wantLength int
	}{
		{"Berlin", "München", 4},
	}
	for _, c := range cases {
		fmt.Printf("%s -- %s ", c.a, c.b)
		a, _ := g.LookupNode(c.a)
		b, _ := g.LookupNode(c.b)
		got := g.ShortestPath(a, b)
		gotLength := len(got) - 1
		if g.PathExists(got) && gotLength == c.wantLength {
			fmt.Println("✓")
		} else {
			fmt.Println("❌")
		}
	}
	fmt.Println()
}
