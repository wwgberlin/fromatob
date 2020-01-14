Lesson content flow (potential, 20min altogether)

##Introduction to graphs (10 min Mandy)
#What are graphs?
Graphs are mathematical structures used to model pairwise relations between objects. A graph in this context is made up of vertices (also called nodes or points) which are connected by edges

https://en.wikipedia.org/wiki/Graph_theory

- Show Graph visual representation

#How are graphs used at fromatob?
- google maps itro
- dijkstra + dfs + bfs (find the right solution)
- What are these algorithms good for?
(find out more info from Medusa? Chris? Nikita?)

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




