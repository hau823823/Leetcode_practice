package productExceptSelf

import "fmt"

// division way
func ProductExceptSelf(nums []int) []int {
	l := len(nums)
	left := make([]int, l)
	right := make([]int, l)
	res := make([]int, l)
	left[0], right[l-1] = 1, 1

	for i := 1; i < l; i++ {
		left[i] = left[i-1] * nums[i-1]
		right[l-i-1] = right[l-i] * nums[l-i]
	}
	fmt.Println("left: ",left)
	fmt.Println("right: ",right)
	for i := 0; i < l; i++ {
		res[i] = left[i] * right[i]
	}

	return res
}

func ProductExceptSelfOp(nums []int) []int {
	res := make([]int, len(nums))
	prefix := 1
	for i, num := range nums {
		res[i] = prefix
		prefix *= num
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}