package containsDuplicate

func ContainsDuplicate(nums []int) bool {
	hashMap := make(map[int]int)

	for _, v := range nums {
		hashMap[v] += 1
		if hashMap[v] > 1 {
			return true
		}
	}

	return false
}
