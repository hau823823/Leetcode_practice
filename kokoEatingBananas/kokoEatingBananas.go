package kokoEatingBananas

// original
func MinEatingSpeed(piles []int, h int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	left, right := 1, max
	for left < right {
		mid := (left + right) / 2
		hours := 0

		for _, pile := range piles {
			hours += (pile + mid - 1) / mid
		}

		if hours <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
