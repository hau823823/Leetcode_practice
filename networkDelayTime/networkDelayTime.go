package networkDelayTime

import "container/heap"

// min heap
type minHeap []heapNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(heapNode)) }
func (h *minHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// bfs
// Dijkstra's algorithm

type neighbour struct {
	destination int
	weight      int
}

type heapNode struct {
	distance  int
	nodeIndex int
}

func NetworkDelayTime(times [][]int, n int, k int) int {
	edgeMap := make(map[int][]neighbour)
	for _, time := range times {
		edgeMap[time[0]] = append(edgeMap[time[0]], neighbour{time[1], time[2]})
	}

	h := &minHeap{heapNode{0, k}}
	heap.Init(h)

	visited := make(map[int]bool)
	res := 0

	for h.Len() > 0 {
		node := heap.Pop(h).(heapNode)

		if exist := visited[node.nodeIndex]; exist {
			continue
		}

		res = max(res, node.distance)
		visited[node.nodeIndex] = true

		neighbours := edgeMap[node.nodeIndex]
		for _, neigh := range neighbours {
			if exist := visited[neigh.destination]; !exist {
				heap.Push(h, heapNode{
					distance:  neigh.weight + node.distance,
					nodeIndex: neigh.destination,
				})
			}
		}
	}

	if n == len(visited) {
		return res
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// test