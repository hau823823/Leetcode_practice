package increasingTripletSubSequence

import "math"

/* origin solution false
func IncreasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, second, third := 0, 1, 2
	
	for third < len(nums) && first < second && second < third {
		if nums[third] > nums[second] && nums[second] > nums[first] {
			return true
		} else if nums[third] < nums[second]{
			third ++
		} else if nums[second] < nums[first] {
			second ++
		} else {
			first ++
		}

	}

    return false
}
*/

func IncreasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt

	for _, num := range nums {
		if num <= first {
			first = num
		} else if  num <= second {
			second = num
		} else {
			return true
		}
	}

	return false
}