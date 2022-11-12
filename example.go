/*
 * @Name: Priority Queue Go
 * @Author: Max Base
 * @Date: 2022-11-12
 * @Repository: https://github.com/basemax/PriorityQueueGo
 */

package main

import (
	"container/heap"
	"fmt"
)

// Main example
func main() {
	// create a priority queue
	pq := PriorityQueue{}
	heap.Init(&pq)

	// add items
	heap.Push(&pq, 1)
	heap.Push(&pq, 2)
	heap.Push(&pq, 3)

	// pop items
	fmt.Println(heap.Pop(&pq))
	fmt.Println(heap.Pop(&pq))
	fmt.Println(heap.Pop(&pq))
}
