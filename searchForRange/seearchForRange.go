package searchForRange

func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	// first search 
	left, right := 0 , len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left < len(nums) && nums[left] == target {
		result[0] = left
	}

	// second search
	left, right = 0, len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 && nums[right] == target {
		result[1] = right
	}

	return result
}
