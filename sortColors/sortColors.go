package sortColors

// bubble sort
func SortColors(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j ++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

// two pointer
func SortColorsTP(nums []int) {
	red, blue := 0, 2
	idxRed, idxBlue := 0, len(nums) - 1

	for i := 0; i <= idxBlue; i++ {
		if nums[i] == red {
			nums[i], nums[idxRed] = nums[idxRed], nums[i]
			idxRed ++
		} else if nums[i] == blue {
			nums[i], nums[idxBlue] = nums[idxBlue], nums[i]
			idxBlue --
			i--
		}
	}
}