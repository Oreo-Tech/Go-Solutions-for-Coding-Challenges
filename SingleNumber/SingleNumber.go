package main

import (
	"fmt"
)

func main() {

	arr := []int64{4, 1, 2, 1, 2}
	for i := 0; i < len(arr); i++ {

		//fmt.Println(arr[i], " ", strconv.FormatInt(arr[i], 2))
		fmt.Printf("%d = %04b\n", arr[i], arr[i])
	}
}
