package taskScheduler

import (
	"container/heap"
)

func LeastInterval(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, task := range tasks {
		count[task-'A']++
	}

	maxCount := count[0]
	for _, c := range count {
		maxCount = max(c, maxCount)
	}

	i := 0
	for _, c := range count {
		if c == maxCount {
			i++
		}
	}

	// maxCount - 1 group of an action and idle time
	// Each group takes n + 1 time
	// the last task executes without any idle time need i time
	return max(len(tasks), (maxCount-1)*(n+1)+i)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// priority queue
type Myheap []int

func (h Myheap) Len() int           { return len(h) }
func (h Myheap) Less(i, j int) bool { return h[i] > h[j] }
func (h Myheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Myheap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *Myheap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Popleft(q *[][2]int) [2]int {
	old := *q
	x := old[0]
	*q = old[1:]
	return x
}

func LeastIntervalPQ(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, task := range tasks {
		count[task-'A']++
	}

	h := &Myheap{}
	heap.Init(h)
	for _, cnt := range count {
		if cnt > 0 {
			heap.Push(h, cnt)
		}
	}

	time := 0
	queue := [][2]int{}
	for h.Len() > 0 || len(queue) > 0 {
		time++

		if h.Len() > 0 {
			taskCnt := heap.Pop(h).(int)
			if taskCnt-1 > 0 {
				queue = append(queue, [2]int{taskCnt - 1, time + n})
			}
		}

		if len(queue) > 0 && queue[0][1] == time {
			// pull tasks which finished idle time
			taskCnt := Popleft(&queue)[0]
			if taskCnt > 0 {
				heap.Push(h, taskCnt)
			}
		}
	}

	return time
}
