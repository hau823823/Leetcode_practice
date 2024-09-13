package topKFrequentElements

func TopKFrequent(nums []int, k int) []int {
	uniq := make([]int, 0)
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
		if count[num] == 1 {
			uniq = append(uniq, num)
		}
	}

	for i := 0; i < len(uniq); i++ {
		for j := 1; j < len(uniq)-i; j++ {
			if count[uniq[j-1]] < count[uniq[j]] {
				uniq[j-1], uniq[j] = uniq[j], uniq[j-1]
			}
		}
	}

	return uniq[:k]
}

func TopKFrequentOp(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	freqList := make([][]int, len(nums)+1)
	for key, val := range freqMap {
		freqList[val] = append(freqList[val], key)
	}

	res := make([]int, 0)
	for i := len(nums); i >= 0; i-- {
		res = append(res, freqList[i]...)
		if len(res) == k {
			break
		}
	}

	return res
}
