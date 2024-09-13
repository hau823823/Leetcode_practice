package lengthOfLastWord

//import "fmt"

func LengthOfLastWord(s string) int {
	b := []byte(s)
	l := len(b)
	res := 0

	for i := l - 1; i >= 0; i-- {
		//fmt.Println(b[i])
		if b[i] != byte(' ') {
			res++
		} else if res != 0 && b[i] == byte(' ') {
			break
		}
	}

	return res
}
