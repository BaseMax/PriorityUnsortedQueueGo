/*
 * @Name: Priority Queue Go
 * @Author: Max Base
 * @Date: 2022-11-12
 * @Repository: https://github.com/basemax/PriorityQueueGo
 */

package main

import "fmt"

type PriorityQueue struct {
	front int
	rear  int
	items []interface{}
}

func (pq *PriorityQueue) Len() int {
	return pq.rear - pq.front
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.items[i].(int) < pq.items[j].(int)
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue) Push(x interface{}, priority int) {
	// priority is priority of x
	pq.items = append(pq.items, x)
	pq.rear++

	// move up
	for i := pq.rear - 1; i > pq.front && pq.Less(i, i-1); i-- {
		pq.Swap(i, i-1)
	}
}

func (pq *PriorityQueue) Pop() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	pq.rear--
	return pq.items[pq.rear]
}

func (pq *PriorityQueue) Peek() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	return pq.items[pq.front]
}

func (pq *PriorityQueue) Remove(i int) interface{} {
	if pq.Len() == 0 {
		return nil
	}
	pq.rear--
	pq.Swap(i, pq.rear)
	return pq.items[pq.rear]
}

func (pq *PriorityQueue) Update(i int, x interface{}, priority int) {
	// priority is priority of x
	pq.items[i] = x

	// move up
	for i > pq.front && pq.Less(i, i-1) {
		pq.Swap(i, i-1)
		i--
	}

	// move down
	for i < pq.rear-1 && pq.Less(i+1, i) {
		pq.Swap(i, i+1)
		i++
	}
}

func (pq *PriorityQueue) String() string {
	return fmt.Sprintf("%v", pq.items[pq.front:pq.rear])
}
