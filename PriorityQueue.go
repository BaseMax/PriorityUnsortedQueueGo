/*
 * @Name: Priority Unsorted Queue Go
 * @Author: Max Base
 * @Date: 2022-11-12
 * @Repository: https://github.com/basemax/PriorityUnsortedQueueGo
 */

package main

import "fmt"

type QueueItem struct {
	Value    interface{}
	Priority int
}

type PriorityQueue struct {
	front int
	rear  int
	size  int
	// unsorted list
	items []QueueItem
}

// Create, O(1)
func NewPriorityQueue(size int) *PriorityQueue {
	return &PriorityQueue{
		front: -1,
		rear:  -1,
		size:  size,
		items: make([]QueueItem, size),
	}
}

// isFull, O(1)
func (q *PriorityQueue) isFull() bool {
	return q.rear == q.size-1
}

// isEmpty, O(1)
func (q *PriorityQueue) isEmpty() bool {
	return q.front == -1 || q.front > q.rear
}

// Enqueue, O(1)
func (q *PriorityQueue) Enqueue(value interface{}, priority int) {
	if q.isFull() {
		fmt.Println("Queue is full!")
		return
	}
	if q.front == -1 {
		q.front = 0
	}
	q.rear = q.rear + 1%q.size
	q.items[q.rear] = QueueItem{
		Value:    value,
		Priority: priority}
}

// Dequeue, O(n)
func (q *PriorityQueue) Dequeue() interface{} {
	if q.isEmpty() {
		fmt.Println("Queue is empty!")
		return nil
	}
	max := q.items[q.front]
	maxIndex := q.front
	i := q.front + 1%q.size
	for ; i <= q.rear; i++ {
		if q.items[i].Priority > max.Priority {
			max = q.items[i]
			maxIndex = i
		}
	}
	for maxIndex < q.rear {
		q.items[maxIndex] = q.items[maxIndex+1]
		maxIndex = maxIndex + 1%q.size
	}
	q.rear--
	return max.Value
}

// Peek, O(n)
func (q *PriorityQueue) Peek() interface{} {
	if q.isEmpty() {
		fmt.Println("Queue is empty!")
		return nil
	}
	max := q.items[q.front]
	i := q.front + 1%q.size
	for ; i <= q.rear; i++ {
		if q.items[i].Priority > max.Priority {
			max = q.items[i]
		}
	}
	return max.Value
}

// Count, O(1)
func (q *PriorityQueue) Count() int {
	return q.rear - q.front + 1
}

// Clear, O(1)
func (q *PriorityQueue) Clear() {
	q.front = -1
	q.rear = -1
}

// ToString, O(n)
func (q *PriorityQueue) ToString() string {
	if q.isEmpty() {
		return "Queue is empty!"
	}
	result := ""
	i := q.front
	for ; i <= q.rear; i++ {
		result += fmt.Sprintf("%v ", q.items[i].Value)
	}
	return result
}
