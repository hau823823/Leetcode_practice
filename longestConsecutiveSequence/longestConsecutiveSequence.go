package longestConsecutiveSequence

// wrong answer
/*
func LongestConsecutive(nums []int) int {
	res := 0
	set := make(map[int]int)
	for _, num := range nums {
		set[num] = 1
	}

	for _, num := range nums {
	
		if v, hasNumMinus := set[num-1]; hasNumMinus {
			set[num] += v

			if set[num] > res {
				res = set[num]
			}

			fmt.Println("Now Num: ", num, ", Value: ", set[num])
			fmt.Println("Min Num: ", num-1, ", Value: ", set[num-1])
		} 
		if v, hasNumPlus := set[num+1]; hasNumPlus {
			set[num] += v

			if set[num] > res {
				res = set[num]
			}

			fmt.Println("Now Num: ", num, ", Value: ", set[num])
			fmt.Println("Plus Num: ", num+1, ", Value: ", set[num+1])
		}
		fmt.Printf("\n")
	}

	return res
}
*/

// solution 1
func LongestConsecutiveOP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int)
	longest := 0

	for _, num := range nums {
		if _, exists := m[num]; exists {
			continue
		}

		left := 0
		if l, exists := m[num-1]; exists {
			left = l
		}

		right := 0
		if r, exists := m[num+1]; exists {
			right = r
		}

		sum := left + right + 1
		m[num] = sum
		
		if sum > longest {
			longest = sum
		}

		m[num-left] = sum
		m[num+right] = sum
	}

	return longest
}

// solution2
func LongestConsecutiveOP1(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    set := make(map[int]bool)
    for _, num := range nums {
        set[num] = true
    }

    longest := 0
    for num := range set {
        if !set[num-1] {
            length := 1
            for set[num+length] {
                length++
            }
            if length > longest {
                longest = length
            }
        }
    }
    return longest
}

// solution3
func LongestConsecutiveOP2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]bool)
	longest := 0

	for _, num := range nums {
		set[num] = true
	}

	for _, num := range nums {
		if !set[num] {
			continue
		}
		delete(set, num)
		
		pre, next := num-1, num+1
		for set[pre] {
			delete(set, pre)
			pre--
		}
		for set[next] {
			delete(set, next)
			next++
		}

		if next-pre-1 > longest {
			longest = next - pre - 1
		}
	}

	return longest
}