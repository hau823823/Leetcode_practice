package houseRobber

func Rob(nums []int) int {
    l := len(nums)
    if l == 1 {
        return nums[0]
    } else if l == 2 {
        return maxInt(nums[0], nums[1])
    }

    max := maxInt(nums[0], nums[1])
    nums[1] = max

    for i := 2; i < l; i++ {
        nums[i] = maxInt(nums[i] + nums[i-2], nums[i-1])
        max = maxInt(max, nums[i])
    }

    return max
}

func maxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}