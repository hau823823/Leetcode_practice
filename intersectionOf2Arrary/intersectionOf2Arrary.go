package intersectionOf2Array

func Intersect(nums1 []int, nums2 []int) []int {

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	hashMap := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		hashMap[num] ++
	}

	result := make([]int, 0)
	for _, num := range nums2 {
		v, exists := hashMap[num];if exists && v > 0 {
			result = append(result, num)
			hashMap[num]--
		}
	}
	return result
}