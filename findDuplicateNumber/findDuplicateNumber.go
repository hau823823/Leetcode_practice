package findDuplicateNumber

import (
	"math"
)

// solution1: hashmap
// time complexity: O(n)
// space complexity: O(n)
// need linea space, not constant
func FindDuplicateHM(nums []int) int {
	dupExist := make(map[int]int)

	for _, num := range nums {
		if _, exists := dupExist[num]; exists {
			return num
		}
		dupExist[num]++
	}

	return 0
}

// solution: marking in arrary approch
// time complexity: O(n)
// space complexity: O(1)
// modify origin array, not allowed
func FindDuplicateMark(nums []int) int {

	for _, num := range nums {
		idx := int(math.Abs(float64(num)))
		if nums[idx] < 0 {
			return idx
		}
		nums[idx] = -nums[idx]
	}

	return 0
}

// solution: binary search
// time complexity: O(nlogn)
// space complexity: O(1)
func FindDuplicateBs(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}

	}

	return left
}

// solution: Fast-Slow Pointers Approach (Floyd's Cycle Detection)
// time complexity: O(n)
// space complexity: O(1)
// 由于所有数字都在 1 到 n 的范围内，使用数字作为索引将始终指向数组内的另一个有效索引，因此必然形成一个环
// 第一次循环找到环中的一个点后，第二次循环用于找到环的起始位置，也就是重复数字的位置。
// 这是因为从数组起点到环入口的距离等于从第一次相遇点到环入口的距离。
func FindDuplicateFloyd(nums []int) int {
	// 第一次循环：寻找环内相遇点
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// 第二次循环：寻找环的入口
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
