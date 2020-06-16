package main

import "fmt"

func main() {

	fmt.Println("Output is", isHappy(19))

}

func isHappy(n int) bool {

	var slow, fast = n, n

	for {
		slow = getNumSquare(slow)
		fast = getNumSquare(fast)

		if slow == fast {
			break
		}
	}
	return slow == 1
}

func getNumSquare(num int) int {
	var sumsq int
	for num > 0 {
		digit := num % 10
		num = num / 10
		sumsq += digit * digit
	}
	return sumsq
}
