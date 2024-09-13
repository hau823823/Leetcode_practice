package jumpGame


func CanJump(nums []int) bool {
	lastReach := nums[0] + 0

	for idx, num := range nums {
		if idx > lastReach {
			return false
		}

		tmp := num + idx
		if tmp > lastReach {
			lastReach = tmp
		}
	}

	return true
}
