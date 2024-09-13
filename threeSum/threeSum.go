package threeSum

import (
	"fmt"
	"sort"
)

// result 會包含重複的答案
func ThreeSum(nums []int) [][]int {

	result := make([][]int, 0)
	hashMap := make(map[int]int, len(nums))

	for i, v := range nums {
		target := v * (-1)
		fmt.Println("target: ",target)

		for j := i + 1; j < len(nums); j++ {
			h, ok := hashMap[nums[j]]
			if ok {
				result = append(result, []int{v, nums[h], nums[j]})
				delete(hashMap, nums[j])
				fmt.Println(result)
			} else {
				hashMap[target - nums[j]] = j
			}
		}
	}
	return result
}

// time complexity: O(n^2 + nlogn)
func ThreeSum2PT(nums []int) [][]int {
	result := make([][]int, 0)	
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			left := i + 1
			right := len(nums) -1
			sum := 0 - nums[i]

			for left < right {
				if nums[left] + nums[right] == sum {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					for left < right && nums[right] == nums[right - 1] {
						right--
					}
					for left < right && nums[left] == nums[left + 1]{
						left++
					}
					right--
					left++
				} else if nums[left] + nums[right] > sum {
					right--
				} else {
					left++
				}
			}
		}
	}

	return result
}