# Priority Unsorted Queue Go (PriorityUnsortedQueueGo)

This is a priority unsorted queue written in Go (Golang).

## Features

-  Unsorted list
-  Support for custom priority
-  Support for custom data type

## Installation

```bash
$ go get github.com/basemax/PriorityUnsortedQueueGo
```

## Usage

```go
package main

import "fmt"

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
```

© Copyright 2022, Max Base
