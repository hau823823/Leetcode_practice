package mergeIntervals

import "sort"

// wrong answer
func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] >= intervals[i][0] && intervals[i-1][1] <= intervals[i][1] {
			if intervals[i-1][0] <= intervals[i][0] {
				intervals[i][0] = intervals[i-1][0]
			}
			intervals = append(intervals[:i-1], intervals[i:]...)
			i--
		} else if intervals[i-1][0] >= intervals[i][0] && intervals[i-1][0] <= intervals[i][1] {
			if intervals[i-1][1] >= intervals[i][1] {
				intervals[i][1] = intervals[i-1][1]
			}
			intervals = append(intervals[:i-1], intervals[i:]...)
			i--
		} else if intervals[i-1][0] <= intervals[i][0] && intervals[i-1][1] >= intervals[i][1] {
			intervals[i][0] = intervals[i-1][0]
			intervals[i][1] = intervals[i-1][1]
			intervals = append(intervals[:i-1], intervals[i:]...)
			i--
		}
	}

	return intervals
}

func MergeOp(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	
	res := [][]int{}
	for _, interval := range intervals {
		if len(res) != 0 && res[len(res) - 1][1] >= interval[0] {
			if interval[1] > res[len(res) - 1][1] {
				res[len(res) - 1][1] = interval[1]
			}
		} else {
			res = append(res, interval)
		}
	}

	return res
}