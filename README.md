# go-workshop

Welcome to the Women Who Go workshop in January challenge with fromatob

## What does fromatob do?
fromAtoB's vision is to combine flights, trains, long distance buses and car-pooling towards your perfect trip from A to B within one booking, one payment, within less than 30 seconds.

## What are we doing today?

1. Brief explanation on graphs at fromatob
2. Mini challenge to familiarise ourselves with our graph `pkg/graph.go`
3. Breadth-first-search algorithm
4. Bring in the mission

1.
## What are graphs?

Graphs are mathematical structures used to model pairwise relations between objects. A graph in this context is made up of vertices (also called nodes or points) which are connected by edges

https://en.wikipedia.org/wiki/Graph_theory

## What kind of problems do graphs sovle?

- Computer networks - Often nodes will represent end-systems or routers, while edges represent connections between these systems.
- Sequence of events - Prioritising projects, finding duration of each project. Use tree structures
- Pathing and Maps: Trying to find shortest or longest paths from some location to a destination makes use of graphs. This can include pathing like you see in an application like Google maps, or calculating paths for AI characters to take in a video game, and many other similar problems.
- Constraint Satisfaction: Find some goal that satisfies a list of constraints. For example,Matching professors to courses so classes don't clash
- Constraint satisfaction problems like this are often modeled and solved using graphs.
- Molecules: Model atoms and molecules for studying their interaction and structure among other things.

## How are graphs used at fromatob?
Our main search functions invovles using an undirected graph to connect cities on a map.
We use open data on public transit timetables and extract the fields we need to build network graphs. Then we create indivudual graphs for each mode of transportation. After that, we run a traversal algorithm like Dijkstra algorithm to find the shortest distance
2.
## Explore mini challenges
  Print the Node ID for Berlin
  Print the Nodename for NodeID 7
  Print the nodes connected to Berin

Eg.
  fmt.Println(g.LookupNode("Berlin"))
  fmt.Println(g.NodeName(6))
  fmt.Println(g.Neighbors(0))

  Write forloop to get from Berlin to Leipzig

  var n = g.Neighbors(0)
  var j, _ = g.LookupNode("Leipzig")

  for _, a := range n {
    if a == j {
      fmt.Printf("Connection exists!")
    }

3.
## BFS (10 min Jessica)

What is queueing?
Step by step on how to queue and visit

https://en.wikipedia.org/wiki/Breadth-first_search
https://medium.com/basecs/going-broad-in-a-graph-bfs-traversal-959bd1a09255

4.
## The mission
Part 1 - Connected or not?

Use the `Graph` structure to find out if cities a and b are connected or not
Tips on how to start:
- In `pkg/algorithms.go`fill in the `Connected` function so it returns true
- Try runnig`go run cmd/main/main.go` to see the Graph structure. if your algorthm works, it should print “ok” for cities that are connected

Part 2 - Shortest path

Well done on completing part 1! Now that you know if cities a to b are connected or not, let's find the shortest path (shortest number of steps in between
Tips on how to start:
- Fill in the ShortestPath function in `pkg/algorithms.go`
- A path is represented by a slice of int with the node IDs. For example, a path from node 4 to node 6 and then to node 2 would be represented by `[]int{4, 6, 2}`
