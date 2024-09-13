package kClosestPoints2Origin

import (
	"sort"
)

// solution1
type m struct {
	val   int
	point []int
}

func KClosest(points [][]int, k int) [][]int {
	tmp := make([]m, len(points))
	for i, point := range points {
		tmp[i].val = square(point[0]) + square(point[1])
		tmp[i].point = point
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].val < tmp[j].val
	})

	result := [][]int{}
	for i := 0; i < k; i++ {
		result = append(result, tmp[i].point)
	}

	return result
}

// solution2
func KClosestOP(points [][]int, k int) [][]int {

	for i := 1; i < len(points); i++ {
		if i >= 1 {
			distanceOfIMinusOne := square(points[i-1][0]) + square(points[i-1][1])
            distanceOfI := square(points[i][0]) + square(points[i][1])
			if distanceOfIMinusOne > distanceOfI {
                points[i-1], points[i] = points[i], points[i-1]
                i -= 2
            }
		}
	}

	return points[:k]
}

func square(n int) int {
    return n * n
}