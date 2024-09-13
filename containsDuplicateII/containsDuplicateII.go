package containsDuplicateII

func ContainsNearbyDuplicate(nums []int, k int) bool {
	hashMap := make(map[int]int)
    for i, nums := range nums {
		idx, hasDup := hashMap[nums]
		if hasDup {
			tmp := idx - i
			if tmp < 0 {
				tmp *= -1
			}
			if tmp <= k {
				return true
			}
		}
		hashMap[nums] = i
	}
	return false
}