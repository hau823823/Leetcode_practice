package minCost2ConnectAllPoints

import (
	"container/heap"
)

// heap
type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *IntHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// minimun spanning tree
// prim's algorithm, time complexity: O(n^2 * logn)
func MinCostConnectPoints(points [][]int) int {

	n := len(points)
	adj := make([][][]int, n)
	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dist := abs(x1-x2) + abs(y1-y2)
			adj[i] = append(adj[i], []int{dist, j})
			adj[j] = append(adj[j], []int{dist, i})
		}
	}

	res := 0
	visited := make(map[int]bool)
	h := &IntHeap{{0, 0}}

	for len(visited) < n {
		pntObj := heap.Pop(h).([2]int)
		cost, pnt := pntObj[0], pntObj[1]

		if visited[pnt] {
			continue
		}

		res += cost
		visited[pnt] = true

		for _, neighbour := range adj[pnt] {
			nCost, nPoint := neighbour[0], neighbour[1]
			if !visited[nPoint] {
				heap.Push(h, [2]int{nCost, nPoint})
			}
		}
	}

	return res
}

// optimize
// time complexity: o(n^2)
func MinCostConnectPointsOp(points [][]int) int {
	nPoint := len(points)
	if nPoint < 2 {
		return 0
	}

	total := 0                   // 用于累积所有选取的边的总成本
	dists := make([]int, nPoint) // 用于存储其他未访问点到当前访问点的最短曼哈顿距离

	// start explore from points[0]
	curPoint := 0
	for curPoint >= 0 {
		dists[curPoint] = -1 // mark as visted
		nextPoint := -1
		minEdge := -1

		for i := range points {
			if dists[i] >= 0 {
				d := getManhattenDistance(points[curPoint][0], points[curPoint][1], points[i][0], points[i][1])

				if dists[i] == 0 || dists[i] > d {
					dists[i] = d
				}

				if minEdge == -1 || minEdge > dists[i] {
					minEdge = dists[i]
					nextPoint = i
				}
			}
		}

		if minEdge != -1 {
			total += minEdge
		}
		curPoint = nextPoint
	}

	return total
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func getManhattenDistance(x1, y1, x2, y2 int) int {
	return abs(x1 - x2) + abs(y1 - y2)
}
