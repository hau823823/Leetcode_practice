package maximumSubarray

func MaxSubArraryGreedy(nums []int) int {

	max := nums[0]
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > max {
            max = sum
        }
		if sum < 0 {
			sum = 0
		}
	}

	return max
}

func MaxSubArraryDpKandane(nums []int) int {
	l := len(nums)
    if l == 0 {
        return 0
    }
    
    sub := make([]int, l)
    sub[0] = nums[0]
    max := sub[0]
    
    for i := 1; i < l; i++ {
        sub[i] = maxInt(sub[i-1] + nums[i], nums[i])
        max = maxInt(max, sub[i])
    }
    
    return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}