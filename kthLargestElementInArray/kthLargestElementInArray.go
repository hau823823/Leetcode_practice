package kthLargestElementInArray

// quick select
// divide and conquer
func FindKthLargest(nums []int, k int) int {
	k = k - 1

	partition := func(left, right int) int {
		pivot := nums[right]
		i := left

		for j := left; j < right; j++ {
			if nums[j] > pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		nums[i], nums[right] = nums[right], nums[i]
		return i	
	}

	var quickSelect func(int, int) int
	quickSelect = func(left, right int) int {
		if left == right {
			return nums[left]
		}
		
		pivotIdx := partition(left, right)
		if k == pivotIdx {
			return nums[pivotIdx]
		} else if k < pivotIdx {
			return quickSelect(left, pivotIdx-1)
		} else {
			return quickSelect(pivotIdx+1, right)
		}
	}

	return quickSelect(0, len(nums)-1)
}

func FindKthLargest2(nums []int, k int) int {
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	freq := make([]int, max - min + 1)
	for _, num := range nums {
		freq[num - min]++
	}

	counter, idx := 0, len(freq) - 1
	for counter < k {
		for freq[idx] == 0 {
			idx--
		}
		freq[idx]--
		counter++
	}

	return idx + min
}