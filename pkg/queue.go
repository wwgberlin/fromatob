package pkg

// Queue implements the queue data type for ints. It's essentially a list where you can add elements
// at one end ("enqueue" an element) and take them out at the other ("dequeue" the element).
type Queue struct {
	list []int
}

// Length returns the number of elements in the list.
func (q *Queue) Length() int {
	return len(q.list)
}

// Enqueue adds i to the queue.
func (q *Queue) Enqueue(i int) {
	q.list = append(q.list, i)
}

// Dequeue takes an element out of the queue. It panics if the queue is empty.
func (q *Queue) Dequeue() int {
	i := q.list[0]
	q.list = q.list[1:]
	return i
}
