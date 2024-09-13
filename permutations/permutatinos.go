package permutations

func Permute(nums []int) [][]int {
	result := make([][]int, 0)

	var backtrack func([]int, []int)
	backtrack = func(source, tmp []int) {
		if len(source) == 0 {
			result = append(result, tmp)
			return
		}
		for i, num := range source {
			newSource := []int{}
			newSource = append(newSource, source[:i]...)
			newSource = append(newSource, source[i+1:]...)
			backtrack(newSource, append(tmp, num))
		}

	}

	backtrack(nums, []int{})
    return result
}