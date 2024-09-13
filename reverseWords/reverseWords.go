package reverseWords

import (
	"strings"
)

func ReverseWordsPk(s string) string {
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func ReverseWords(s string) string {

    tmp := make([]string, 0)
    start := -1
    for i := 0; i < len(s); i++ {
        if s[i] == ' '{
            if start != -1 {
                tmp = append(tmp, s[start:i])
                start = -1
            }
        } else if start == -1 {
            start = i
        }
    }
	// 如果結尾不是 0
	if start != -1 {
		tmp = append(tmp, s[start:])
	}

	res := make([]byte, 0)
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, []byte(tmp[i])...)
		if i != 0 {
			res = append(res, ' ')
		}
	}
	return string(res)
}
