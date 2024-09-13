package twoSumII

// hashMap
func TwoSum(numbers []int, target int) []int {
	idxMap := make(map[int]int)
	res := make([]int, 2)
	for i, num := range numbers {
		idx, exists := idxMap[num]
		if exists {
			res[0] = idx
			res[1] = i + 1
			break
		}

		idxMap[target-num] = i + 1

	}

	return res
}

// optimize
func TwoSumOP(numbers []int, target int) []int {
	s := 0
	l := len(numbers) - 1

	for s < l {
		currentSum := numbers[s] + numbers[l]

		if currentSum > target {
			l--
		} else if currentSum < target {
			s++
		} else {
			return []int{s+1, l+1}
		}
	}

	return []int{}
}