package validPalindrome

import "fmt"

func IsPalindrome(s string) bool {
	res := make([]rune, 0)
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			res = append(res, v)
		} else if v >= 'A' && v <= 'Z' {
			res = append(res, v-'A'+'a')
		} else if v >= '0' && v <= '9' {
			res = append(res, v)
		}
	}
	fmt.Println(string(res))

	for i := 0; i < len(res)/2; i++ {
		if res[i] != res[len(res)-i-1] {
			return false
		}
	}

	return true
}
