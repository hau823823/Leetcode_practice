package majorityElement

// hash map
// time complexity: O(n)
// space complexity: O(n)
func MajorityElement(nums []int) int {
    hashMap := make(map[int]int)
	res, tmp := nums[0], 0

	for _, v := range nums {
		hashMap[v] += 1

		if hashMap[v] > tmp{
			tmp = hashMap[v]
			res = v
		}
	}
	return res
}

// Boyer Moore
// time complexity: O(n)
// space complexity: O(1)
func MajorityElementOp(nums []int) int {
	count, candidate := 0, nums[0]

	for _, num := range nums{
		if num == candidate {
			count++
		} else {
			count--
		}
		if count == 0 {
			candidate = num
			count = 1
		}
	}

	return candidate
}