package rotateArray

import "fmt"

// brute
func Rotate(nums []int, k int) {
	l := len(nums)
	k = k%l
	tmp := append(nums[l-k:], nums[:l-k]...)
	copy(nums, tmp)
}

// optimize solution 1
func RotateOP1(nums []int, k int) {
	l := len(nums) 
	k = k % l
	if k == 0 {
		return
	}

	start, idx, pre, cur := 0, 0, 0, nums[0]

	for i := 0; i < l; i++ {
		pre = cur
		idx = (idx + k) % l

		cur = nums[idx]
		nums[idx] = pre

		if idx == start {
			idx = start + 1
			start = idx
			cur = nums[idx]
		}
		fmt.Println(nums)
	}
}

// optimize solution 2
func RotateOP2(nums []int, k int) {
	n := len(nums)
	k = k%n
	for i := 0; i < (n-k)/2; i++ {
		nums[i], nums[n-k-1-i] = nums[n-k-1-i], nums[i]
	}
	fmt.Println(nums)
	for j := 0; j < k/2; j++ {
		nums[n-k+j], nums[n-j-1] = nums[n-1-j], nums[n-k+j]
	}
	fmt.Println(nums)
	for  a:= 0; a < n/2; a++ {
		nums[a], nums[n-a-1] = nums[n-a-1], nums[a]
	}
	fmt.Println(nums)
}

// optimize solution 3
func RotateOP3(nums []int, k int) {
	n, start := len(nums), 0
	k %= n 
    for k != 0 {
        for i := 0; i < k; i++ {
            nums[start+i], nums[start+n-k+i] = nums[start+n-k+i], nums[start+i]
        }
        n -= k
        start += k
        k %= n
    }
}