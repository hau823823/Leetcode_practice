package missingNumber

func MissingNumber(nums []int) int {
	// 計算 nums 所有值的合

	res := (len(nums) + 1) * len(nums) / 2

	for _, v := range nums {
		res -= v
	}

	return res
}