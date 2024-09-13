package numberOf1Bits

import "fmt"

func HammingWeight(num uint32) int {
	w := 0

	for {
		if num == 0 {
			break
		}
		if num&1 == 1 {
			w ++
		}
		fmt.Printf("num(binary): %b \n", num)
		fmt.Printf("num&1: %b \n", num&1)

		num = num >> 1
		fmt.Printf("num >> 1: %b \n", num)
	}

	return w
}