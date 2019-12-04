package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "paths.md"
	g, err := LoadGraph(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Loaded %s (%d nodes, %d edges)\n", filename, g.Nodes(), g.Edges())
}
