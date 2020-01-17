Lesson content flow (potential, 20min altogether)

##Introduction to graphs (10 min Mandy)
#What are graphs?
Graphs are mathematical structures used to model pairwise relations between objects. A graph in this context is made up of vertices (also called nodes or points) which are connected by edges

https://en.wikipedia.org/wiki/Graph_theory

#What kind of problems do graphs sovle?

- Computer networks - Often nodes will represent end-systems or routers, while edges represent connections between these systems.
- Sequence of events - Prioritising projects, finding duration of each project. Use tree structures
- Pathing and Maps: Trying to find shortest or longest paths from some location to a destination makes use of graphs. This can include pathing like you see in an application like Google maps, or calculating paths for AI characters to take in a video game, and many other similar problems.
- Constraint Satisfaction: Find some goal that satisfies a list of constraints. For example,Matching professors to courses so classes don't clash
- Constraint satisfaction problems like this are often modeled and solved using graphs.
- Molecules: Model atoms and molecules for studying their interaction and structure among other things.

#How are graphs used at fromatob?
Our main search functions invovles using an undirected graph to connect cities on a map.
We use open data on public transit timetables and extract the fields we need to build network graphs. Then we create indivudual graphs for each mode of transportation. After that, we run a traversal algorithm like Dijkstra algorithm to find the shortest distance

Smaller things we do with graphs
Eg. Clean up Nodes which are not relevant stations
Eg. Consolidate cluster of stations into one node to speed up search - Alexanderplatz is seen as 10 train stations on a map so we use a cluster algorithm to represent it with a single node

Check out an example of a graph here `pkg/graph_visual.png`

#Explore mini challenges
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


##BFS (10 min Jessica)
Why is it useful?
What are Queues?
Step by step on how to queue and visit (screenshot )
https://en.wikipedia.org/wiki/Breadth-first_search
https://medium.com/basecs/going-broad-in-a-graph-bfs-traversal-959bd1a09255




