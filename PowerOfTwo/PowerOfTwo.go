package main

import "fmt"

func main() {

	fmt.Println("\nResult is", isPowerOfTwo(256))

}

func isPowerOfTwo(n int) bool {
	return (n > 0) && ((n & (n - 1)) == 0)
}

//Long Approach
// func isPowerOfTwo(n int) bool {

// 	if n == 0 {
// 		return false
// 	}
// 	for n%2 == 0 {

// 		n = n / 2
// 		fmt.Println("N is", n)
// 	}

// 	return n == 1
// }
