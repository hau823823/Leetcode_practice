package pow

import "fmt"

// goroutine
func MyPow(x float64, n int) float64 {
	isNegative := false
	if n < 0 {
		n *= -1
		isNegative = true
	}

	s := n / 1000
	r := n % 1000
	res := float64(1)
	c := make(chan float64, s+1)

	for i := 0; i < s; i++ {
		go pow(x, 1000, c)
	}
	go pow(x, r, c)

	if !isNegative {
		for i := 0; i < s+1; i++ {
			res *= <-c
		}
	} else {
		for i := 0; i < s+1; i++ {
			res /= <-c
		}
	}
	close(c)

	return res
}

func pow(x float64, n int, c chan float64) {
	res := float64(1)
	for i := 0; i < n; i++ {
		res *= x
	}
	c <- res
}

// solution 2
// time complexity: O(logn)
func MyPowOp(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}

	result := 1.0
	for n > 0 {
		if n%2 == 1 {
			result *= x
			fmt.Println("in if")
		}
		x *= x
		n = n >> 1 // same as n /= 2
		fmt.Println("result: ", result)
		fmt.Println("x: ", x)
		fmt.Println("n: ", n)
		fmt.Printf("\n")
	}
	return result
}
