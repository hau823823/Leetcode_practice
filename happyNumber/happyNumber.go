package happyNumber

import "fmt"

//import "fmt"

func IsHappy(n int) bool {
	visited := make(map[int]bool)

	for {
		// 計算平方和
		nextValue := 0
		for ; n > 0; n = n / 10 {
			fmt.Println("n: ", n)
			nextValue += (n % 10) * (n % 10)
		}
		fmt.Println(nextValue)

		if nextValue == 1 {
			return true
		} else if _, ok := visited[nextValue]; ok {
			return false
		} else {
			visited[nextValue] = true
		}

		n = nextValue
	}
}
