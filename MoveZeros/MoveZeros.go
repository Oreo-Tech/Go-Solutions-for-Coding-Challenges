package main

import "fmt"

func main() {

	nums := []int{0, 2, 3, 0, 5, 6, 4, 0, 0, 9, 1, 0}

	moveZeroes(nums)
	fmt.Println("Changed", nums)
}

func moveZeroes(nums []int) {
	var count int

	for i := 0; i < len(nums); i++ {

		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}

	for count < len(nums) {
		nums[count] = 0
		count++
	}
}
