package longestSubStrWithoutRptChts

import "fmt"

// wrong answer
func LengthOfLongestSubString(s string) int {
	hashMap := make(map[rune]int)
	result, tmp := 0, 0

	for _, v := range s {
		_, hasDuplicate := hashMap[v]
		hashMap[v]++
		if hasDuplicate {
			hashMap = make(map[rune]int)
			tmp = 1
		} else {
			tmp++
		}
		if tmp > result {
			result = tmp

		}
		fmt.Println(hashMap)
		fmt.Println("res:", result)
		fmt.Println("tmp:", tmp)
	}

	return result
}

func LengthOfLongestSubStringOP(s string) int {
    hashMap := make(map[rune]int)
    result, left := 0, 0

    for i, v := range s {
		fmt.Println("v: ",string(v))
        if index, found := hashMap[v]; found && index >= left {
            left = index + 1
			fmt.Println("left:", left)
        }
        hashMap[v] = i
        if result < i-left+1 {
            result = i - left + 1
        }
    }

    return result
}
