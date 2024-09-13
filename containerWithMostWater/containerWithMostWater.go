package containerWithMostWater

func MaxArea(height []int) int {
	l := len(height) - 1
	leftIdx, rightIdx := 0, l
	max := 0

	for leftIdx < rightIdx {
		if height[leftIdx] < height[rightIdx]{
			max = maxInt(max, height[leftIdx] * l)
			leftIdx++
		} else {
			max = maxInt(max, height[rightIdx] * l)
			rightIdx--
		}
		l--
	}
	
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}