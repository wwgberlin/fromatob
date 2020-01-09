# go-workshop

Women Who Go workshop in January challenge

This challenge is all about graphs. A graph consists of vertices and edges, where an edge connects
two vertices. In our case, each vertex represents a city. An edge between two vertices means there’s
a train connection between those two cities.

The graph is loaded from a file (see `paths.md`) into a data structure called `Graph` (see
`pkg/graph.go`). Your challenge is to use the `Graph` structure to find out if and how we can use
trains to go from one city to another.

`go run cmd/go-workshop/main.go` runs your code with some examples. When you’re done, it should
print “ok” for all of them!


## 1st mission

For the first mission, we just want to find out if two cities are connected at all. In
`pkg/algorithms.go`, fill in the `Connected` function so it returns true if there’s a train
connection from `a` to `b`.


## 2nd mission

For the second mission, we want to find the “shortest path”: the path that connects two cities with
the smallest number of stops in-between In `pkg/algorithms.go`, fill in the `ShortestPath` function.

A path is represented by a slice of int with the node IDs. For example, a path from node 4 to node 6
and then to node 2 would be represented by `[]int{4, 6, 2}`.
