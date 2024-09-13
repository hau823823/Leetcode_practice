package validParentheses

func IsValid(s string) bool {

	barcket := map[byte]byte {
		byte(')'): byte('('),
		byte(']'): byte('['),
		byte('}'): byte('{'),
	}
	b := []byte(s)
	stack := make([]byte, 0)

	for i := 0; i < len(b); i++ {
		v, ok := barcket[b[i]]
		if ok {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, b[i])
		}
	}

	return len(stack) == 0
}