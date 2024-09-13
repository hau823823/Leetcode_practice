package mergeSortedArrary

func Merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m == 0 || nums1[m-1] < nums2[n-1] {
			nums1[n+m-1] = nums2[n-1]
			n--
		} else {
			nums1[n+m-1] = nums1[m-1]
			m--
		}
	}
}
