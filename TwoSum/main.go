package main

import "fmt"

func main() {
	arr := []int{3, -2, 7, 11, 15, 5}
	fmt.Println("Output is", twoSum(arr, 9))
}

func twoSum(nums []int, target int) []int {
	res := []int{}
	reminderMap := map[int]int{}
	// Map to store index of number againt the key of target - number
	for i := 0; i < len(nums); i++ {
		value, ok := reminderMap[nums[i]]
		if ok {// If target - number already exists in the map means we would solution as key=target-number at index of value
			res = append(res, value, i) // add index of first operand and index of key operand and return the solution array
			break
		}
		reminderMap[target-nums[i]] = i // Store the key as key=target-number at index of value

	}
	fmt.Println("Reminder map is", reminderMap)
	return res
}
