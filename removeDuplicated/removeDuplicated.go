package removeDuplicated

func RemoveDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}
	
	j := 0 // j 表示不重複的字元，i 用來遍歷檢查
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j += 1
			nums[j] = nums[i]
		}
	}

    return j + 1
}