package dailyTemperatures

func DailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0)

	for idx, temp := range temperatures {

		for len(stack) > 0 && temperatures[stack[len(stack) - 1]] < temp {
			index := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			result[index] = idx - index
		}

		stack = append(stack, idx)
	}

	return result
}