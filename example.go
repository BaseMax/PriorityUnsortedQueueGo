/*
 * @Name: Priority Queue Go
 * @Author: Max Base
 * @Date: 2022-11-12
 * @Repository: https://github.com/basemax/PriorityQueueGo
 */

package main

import "fmt"

// Main example
func main() {
	// create a priority queue
	pq := PriorityQueue{}

	// push items to the queue
	pq.Push(1, 1)
	pq.Push(2, 2)
	pq.Push(3, 3)
	pq.Push(4, 4)
	pq.Push(5, 5)
	pq.Push(6, 1)

	// print the queue
	fmt.Println(pq.String())

	// pop the queue
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())

	// iterate the queue
	for pq.Len() > 0 {
		fmt.Println(pq.Pop())
	}

	// print as string
	fmt.Println(pq.String())
}
