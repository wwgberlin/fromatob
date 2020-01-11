Lesson content (potential, 45min approx)

##Introduction to graphs (20 min)
#What are graphs?
Graphs are mathematical structures used to model pairwise relations between objects. A graph in this context is made up of vertices (also called nodes or points) which are connected by edges

#non directed v directed with pictures

#Application of graphs

#What kind of problems need solving?
represent networks of communication, data organization, computational devices, the flow of computation, etc. For instance, the link structure of a website can be represented by a directed graph, in which the vertices represent web pages and directed edges represent links from one page to another.


#In context of travel - google maps?

##BFS (25 min)
Why is it useful?
What are Queues?
Step by step on how to queue and visit (screenshot )

##Breakdown individually for beginners? |

#Looking at graphs in detail - to start with the basic structure of the algorithm
- Plaing around in main.go with graph.go functions
eg.
  //showing how to use the functions in graph
  fmt.Println(g.LookupNode("Berlin"))
  fmt.Println(g.LookupNode("Leipzig"))
  fmt.Println(g.Neighbors(0))

  //Mini challenge get from Berlin to Leipzig
  var n = g.Neighbors(0)
  var j, _ = g.LookupNode("Leipzig")

  for _, a := range n {
    if a == j {
      fmt.Printf("Connection exists!")
    }

##Resources
https://medium.com/basecs/going-broad-in-a-graph-bfs-traversal-959bd1a09255

