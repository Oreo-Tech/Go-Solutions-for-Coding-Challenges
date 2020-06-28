
//////////////////////////////////////////////////////////////////

//Explanation at: https://medium.com/@sunil18031992/numbers-with-even-number-of-digits-d402be34fedb

//////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 22, 333, 4444, 55555, 666666}
	fmt.Println("Answer is", findNumbers(nums))
}

func findNumbers(nums []int) int {
	//result, res := 0, 0
	res := 0
	for _, val := range nums {
		res += int(math.Log10(float64(val))) & 1
		fmt.Println("Surprise ", res, " ", math.Log10(float64(val)), " ", int(math.Log10(float64(val)))&1)
		// 	count := getCountOFDigits(val)
		// 	if count&1 == 0 {
		// 		result++
		// 	}
	}
	return res
}

//Function to return the digit count of a given number
func getCountOFDigits(num int) int {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}
	return count
}
