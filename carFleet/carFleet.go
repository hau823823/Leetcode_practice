package carFleet

import (
	"fmt"
	"sort"
)

func CarFleet(target int, position []int, speed []int) int {
	n := len(speed)
	v := make([][2]int, n)
	for i := 0; i < n; i++ {
		v[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(v, func(i, j int) bool {
		return v[i][0] < v[j][0]
	})
	fmt.Println("sorted pairs (position/speed): ",v)

    // 到達目的地需要的時間
	time := make([]float64, n)
	for i := 0; i < n; i++ {
		time[i] = float64(target-v[i][0]) / float64(v[i][1])
	}
    fmt.Println("times to target: ", time)

	cur := -1.0
	res := 0
	for i := n - 1; i >= 0; i-- {
		if time[i] > cur {
			cur = time[i]
			res++
		}
	}

	return res
}

// optimize
type carInfo struct {
    pos int
    spd int
}

func CarFleetOp(target int, position []int, speed []int) int {
    pairs := []carInfo{}
    stack := []float64{}

    for i := range position {
        pairs = append(pairs, carInfo{position[i], speed[i]})
    }

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].pos < pairs[j].pos
	})

	for i := len(pairs) - 1; i >= 0; i-- {
		stack = append(stack, float64(target - pairs[i].pos) / float64(pairs[i].spd))
		if len(stack) > 1 && stack[len(stack) - 1] <= stack[len(stack) - 2] {
			stack = stack[:len(stack) - 1]
		}
	}

	return len(stack)
}