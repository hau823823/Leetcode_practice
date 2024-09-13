package firstBadVersion

var Bad int

func isBadVersion(version int) bool {
	return version == Bad
}

func FirstBadVersion(n int) int {
	left, right := 1, n
	for left + 1 < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	if isBadVersion(right) {
		return right
	} else {
		return left
	}
}
