package twoSum

func TwoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, n := range nums {
		h, ok := hashmap[n]; if ok {
			return []int{h, i}
		}
		hashmap[target - n] = i
	}
	 
    return nil
}