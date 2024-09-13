package longestCommonPrefix

//import "fmt"

func LongestCommonPrefix(strs []string) string {
	for index, val := range strs[0] {
		for _, str := range strs[1:] {
			//fmt.Println("byte(val):", string(val))
			if index == len(str) || byte(val) != str[index] {
				//fmt.Println("str[index]",string(str[index]))
				return strs[0][:index]
			}
		}
	}
	return strs[0]
}
