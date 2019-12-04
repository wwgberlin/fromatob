package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadGraph(name string) (*Graph, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("error opening graph file: %w", err)
	}
	defer file.Close()

	g := &Graph{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), "-")
		if len(nodes) != 2 {
			// skip lines that don't have the format "foo-bar"
			continue
		}
		i := g.AddNode(nodes[0])
		j := g.AddNode(nodes[1])
		g.AddEdge(i, j)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error parsing graph file: %w", err)
	}

	return g, nil
}
