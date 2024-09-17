package cheapestFlightsWithinKStops

import "math"

// Bellman Ford Algorithm
// BFS (But go throught every single edges)
// time complexity: O(e * k)
func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = math.MaxUint32
	}
	cost[src] = 0

	for i := 0; i < k + 1; i++ {
		tmp := make([]int, n)
		copy(tmp, cost)

		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]
			if cost[from] != math.MaxUint32 {
				tmp[to] = min(tmp[to], cost[from] + price)
			}
		}

		cost = tmp
	}

	if cost[dst] == math.MaxUint32 {
		return -1
	}
	return cost[dst]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
