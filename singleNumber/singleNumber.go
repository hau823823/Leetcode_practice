package singleNumber

func SingleNumber(nums []int) int {
	hashMap := make(map[int]int)
	res := 0
	for _, v := range nums {
		hashMap[v] += 1
		if hashMap[v] <= 1 {
			res += v
		} else {
			res -= v
		}

	}
	return res
}

func SingleNumberOP(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
		//fmt.Println(v)
		//fmt.Printf("nums: %b\n\n", v)
		//mt.Println(result)
		//fmt.Printf("result: %b\n", result)
	}
	return result
}
