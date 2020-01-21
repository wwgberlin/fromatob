package pkg

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := &Queue{}
	if q.Length() != 0 {
		t.Fatalf("Queue length not zero")
	}
	q.Enqueue(4)
	if q.Length() != 1 {
		t.Fatalf("queue length not 1")
	}
	q.Enqueue(5)
	if q.Length() != 2 {
		t.Fatalf("queue length not 2")
	}
	if q.Dequeue() != 4 {
		t.Fatal("value expected to be 4")
	}
	if q.Length() != 1 {
		t.Fatalf("queue length not 1")
	}
	if q.Dequeue() != 5 {
		if q.Length() != 0 {
			t.Fatalf("queue length should be 0")
		}
	}
}
