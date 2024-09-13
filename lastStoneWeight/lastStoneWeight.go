package lastStoneWeight

import (
	"container/heap"
	"sort"
)

// sort
func LastStoneWeightSort(stones []int) int {
	sort.Ints(stones)

	for len(stones) > 1 {
		stone := stones[len(stones) - 1] - stones[len(stones) - 2]
		stones = stones[:len(stones)-2]
		if stone > 0 {
			idx := sort.SearchInts(stones, stone)
			if idx > len(stones) {
				stones = append(stones, stone)
			} else {
				stones = append(stones[:idx], append([]int{stone}, stones[idx:]...)...)
			}
		}
	}

	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}

// priority queue
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func LastStoneWeight(stones []int) int {
	pq := IntHeap(stones)
	heap.Init(&pq)

	for pq.Len() > 1 {
		x := heap.Pop(&pq).(int)
		y := heap.Pop(&pq).(int)
		if x != y {
			heap.Push(&pq, x-y)
		}
	}

	if pq.Len() == 0 {
		return 0
	}

	return heap.Pop(&pq).(int)
}
