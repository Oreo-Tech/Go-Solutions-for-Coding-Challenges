package main

import "fmt"

func main() {
	//	merge([]int{1, 2, 4, 0}, 3, []int{3}, 1)
	mergeFromEnd([]int{0}, 0, []int{4}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	for i, j := 0, 0; i < len(nums1) && j < n; i++ {

		if nums1[i] > nums2[j] {
			copy(nums1[i+1:], nums1[i:])
			nums1[i] = nums2[j]
			j++
		} else {
			m--
		}
		if m < 0 {
			nums1[i] = nums2[j]
			j++
		}

	}

	fmt.Println("Array output is", nums1)

}

func mergeFromEnd(nums1 []int, m int, nums2 []int, n int) {
	index := len(nums1) - 1
	//i, j := 0, len(nums1)-1

	for index >= 0 && m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[index] = nums1[m-1]
			m--
		} else {
			nums1[index] = nums2[n-1]
			n--
		}
		index--
	}
	if n > 0 && index >= 0 {
		for index >= 0 && n > 0 {
			nums1[index] = nums2[n-1]
			index--
			n--
		}
	}
	fmt.Println("Array output is", nums1)

}
