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
	// Create a new queue
	queue := NewPriorityQueue(5)

	// Enqueue some items
	queue.Enqueue("A", 1)
	queue.Enqueue("B", 2)
	queue.Enqueue("C", 3)
	queue.Enqueue("D", 4)
	queue.Enqueue("E", 5)

	fmt.Println(queue.ToString())

	// Dequeue and print
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	queue.Enqueue("A", 10)
	queue.Enqueue("B", 20)
	queue.Enqueue("C", 3)

	fmt.Println(queue.ToString())
}
