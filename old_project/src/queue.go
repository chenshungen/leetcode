package main

import "fmt"

// ArrayQueue 基于数组实现的队列
type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

// NewArrayQueue ...
func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

// EnQueue  ...
func (queue *ArrayQueue) EnQueue(v interface{}) bool {
	if queue.tail == queue.capacity {
		return false
	}
	queue.q[queue.tail] = v
	queue.tail++
	return true
}

// DeQueue ...
func (queue *ArrayQueue) DeQueue() interface{} {
	if queue.head == queue.tail {
		return nil
	}
	v := queue.q[queue.head]
	queue.head++
	return v
}

//  String ....
func (queue *ArrayQueue) String() string {
	if queue.head == queue.tail {
		return "empty queue"
	}
	result := "head"
	for i := queue.head; i <= queue.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", queue.q[i])
	}
	result += "<-tail"
	return result
}
