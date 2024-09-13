package findPeakElement

import (
	"math"
)

func FindPeakElement(nums []int) int {
	nums = append(nums, math.MinInt)
	nums = append([]int{math.MinInt}, nums...)

	max := nums[0]
	result := 0

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] {
			if nums[i] > nums[i+1] && nums[i] > max {
				result = i
			}
		}
	}

	return result - 1
}

func FindPeakElementMax(nums []int) int {
	max := nums[0]
	result := 0
	for i, num := range nums {
		if num > max {
			max = num
			result = i
		}
	}

	return result
}

func FindPeakElementOp(nums []int) int {
	left, right := 0, len(nums) - 1

	for left < right {
		mid := (left+right)/2
		if nums[mid] >= nums[mid + 1] {
			right = mid
		} else if nums[mid] < nums[mid + 1] {
			left = mid + 1
		} else {
			return mid
		}
	}

	return left
}
