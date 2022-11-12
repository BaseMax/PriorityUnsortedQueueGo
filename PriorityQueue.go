/*
 * @Name: Priority Queue Go
 * @Author: Max Base
 * @Date: 2022-11-12
 * @Repository: https://github.com/basemax/PriorityQueueGo
 */

package main

import "fmt"

type QueueItem struct {
	Value    interface{}
	Priority int
	Index    int
}

type PriorityQueue struct {
	front int
	rear  int
	items []QueueItem
}

func (pq *PriorityQueue) swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].Index = i
	pq.items[j].Index = j
}

func (pq *PriorityQueue) Push(value interface{}, priority int) {
	item := QueueItem{value, priority, pq.rear}
	pq.items = append(pq.items, item)
	pq.rear++
}

func (pq *PriorityQueue) Pop() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	item := pq.items[pq.front]
	pq.front++
	return item.Value
}

func (pq *PriorityQueue) Len() int {
	return pq.rear - pq.front
}

func (pq *PriorityQueue) String() string {
	return fmt.Sprintf("%v", pq.items)
}
