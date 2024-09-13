package removeElement

func RemoveElement(nums []int, val int) int {
    left := 0              // 用來記錄目前檢查到的元素
	right := len(nums) - 1 // 用來記錄與 val 相同的元素往後交換的位置

	for right > left {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left] // 將相同的直丟到後面
			right--
		} else {
			left++
		}
	}
	
	count := 0
	for _,v := range nums {
		if v != val {
			count++
		}
	}

	return count
}