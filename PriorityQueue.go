/*
 * @Name: Priority Queue Go
 * @Author: Max Base
 * @Date: 2022-11-12
 * @Repository: https://github.com/basemax/PriorityQueueGo
 */

package main

type PriorityQueue struct {
	items []interface{}
}

func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.items[i].(int) < pq.items[j].(int)
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.items = append(pq.items, x)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(pq.items)
	x := pq.items[n-1]
	pq.items = pq.items[:n-1]
	return x
}
