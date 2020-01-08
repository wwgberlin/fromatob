package pkg

// Connected checks if the given nodes are connected by a path.
func (g *Graph) Connected(a, b int) bool {
	var q Queue                           //define q as a slice
	q.Enqueue(a)                          //enqueue root element a
	discovered := make([]bool, g.Nodes()) //define discovered as slice bool
	discovered[a] = true                  //make root element 'discovered'

	for q.Length() > 0 { //check if an element exist?
		v := q.Dequeue() //initalize current element as v
		if v == b {      //if current element is equivalent to the last element
			return true //we found the connection
		}
		for _, n := range g.Neighbors(v) { //look at neighbours for current element v
			if !discovered[n] { //if neighbours are not discovered
				discovered[n] = true //mark it as 'discoverd'
				q.Enqueue(n)         //add it to the queue
			}
		}
	}
	return false //otherwise, return 'not connected'
}

//Steps
//1. Make the root element 'discovered'
//2. Loop through to check if any elements are left in the slice
//3. Loop through to check if neighbours of each element is discovered or not?

type Queue struct {
	list []int
}

func (q *Queue) Length() int {
	return len(q.list)
}

func (q *Queue) Enqueue(i int) {
	q.list = append(q.list, i)
}

func (q *Queue) Dequeue() int {
	i := q.list[0]
	q.list = q.list[1:]
	return i
}

// ShortestPath returns a path between the nodes with the minimum number of edges.
func (g *Graph) ShortestPath(a, b int) []int {
	// mission 2: find the shortest path from a to b
	return []int{}
}
