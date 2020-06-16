package main

import "fmt"

func main() {
	fmt.Println("Output is", isHappy(16))
}

func isHappy(n int) bool {
	numMap := make(map[int]int)
	for {
		sumsq := getNumSquare(n)
		if sumsq == 1 { // if at any point sum of digit square is 1 return true
			return true
		}
		_, ok := numMap[sumsq]
		if ok { // if number already exists in map then loop is observed and  break
			break
		} else { // add an entry of number in sequence of already traversed digits
			numMap[sumsq] = 1
			n = sumsq
		}
	}

	return false
}

func getNumSquare(num int) int { // calculate sum of suware of each digit in a number
	var sumsq int
	for num > 0 {
		digit := num % 10
		num = num / 10
		sumsq += digit * digit
	}
	return sumsq
}
