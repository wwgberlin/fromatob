package main

import (
	"fmt"
	"os"

	"github.com/fromatob/go-workshop/pkg"
)

func main() {
	filename := "trains.txt"
	g, err := pkg.LoadGraph(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Loaded %s (%d nodes)\n", filename, g.Nodes())

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
		{"Jena", "Aachen", true},
		{"Hamburg", "Augsburg", true},
		{"Hiroshima", "Tokyo", true},
		{"Kyoto", "Dresden", false},
		{"Köln", "Tokyo", false},
	}
	for _, c := range cases {
		fmt.Printf("%-9s -> %-9s ", c.a, c.b)
		a, _ := g.LookupNode(c.a)
		b, _ := g.LookupNode(c.b)
		got := g.Connected(a, b)
		if got == c.want {
			fmt.Println("ok")
		} else if got == false {
			fmt.Println("marked as “not connected”, but they are!")
		} else {
			fmt.Println("marked as “connected”, but they are not!")
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
		{"Jena", "Aachen", 5},
		{"Hamburg", "Augsburg", 5},
		{"Hiroshima", "Tokyo", 2},
	}
	for _, c := range cases {
		fmt.Printf("%-9s -> %-9s ", c.a, c.b)
		a, _ := g.LookupNode(c.a)
		b, _ := g.LookupNode(c.b)
		got := g.ShortestPath(a, b)
		if len(got) == 0 {
			fmt.Println("no result")
			continue
		}
		gotLength := len(got) - 1
		if !(len(got) >= 2 && got[0] == a && got[len(got)-1] == b && g.PathExists(got)) {
			fmt.Println()
			fmt.Printf("    not a correct path from %s to %s: %s\n",
				c.a, c.b, g.PrintPath(got, "-"))
		} else if gotLength > c.wantLength {
			fmt.Printf("%s has length %d, but we can do better!\n",
				g.PrintPath(got, "-"), gotLength)
		} else {
			fmt.Println("ok")
			fmt.Printf("    %s\n", g.PrintPath(got, "-"))
		}
	}
	fmt.Println()
}
