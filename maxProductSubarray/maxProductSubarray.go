package maxProductSubarray

func MaxProduct(nums []int) int {
	res := nums[0]
	curMin, curMax := 1, 1

	for _, num := range nums {

		// something will be wrong at this case
		// [0,10,10,10,10,10,10,10,10,10,-10,10,10,10,10,10,10,10,10,10,0]
		/**
		if num == 0 {
			curMin, curMax = 1, 1
			continue
		}

		tmp := curMax * num
		curMax = max(max(num*curMax, num*curMin), num)
		curMin = min(min(tmp, num*curMin), num)
		*/

		if num < 0 {
			curMin, curMax = curMax, curMin
		}

		curMin = min(num, curMin * num)
		curMax = max(num, curMax * num)

		res = max(res, curMax)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
