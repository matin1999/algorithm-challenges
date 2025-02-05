package main

import (
	"fmt"
)

func main() {

	arr := []int{4,3,2,1}

	// for {
	// 	find := false
	// 	index := 0
	isGhole(arr)
	// 	if find == true {
	// 		fmt.Println(index)
	// 		break
	// 	}
	// }



}

func isGhole(arr []int) {

	
	mid := int(len(arr) / 2)
	if len(arr) <= mid+1 {
		if arr[len(arr)-1] > arr[0] {
			// return true, arr[len(arr)-1], []int{}
			fmt.Println(arr[len(arr)-1])
			return
		} else {
			// return true, arr[0], []int{}
			fmt.Println(arr[0])
			return
		}
	}

	if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
		// return true, arr[mid], []int{}
		fmt.Println(arr[mid])
		return
	} else {
		if arr[mid-1] > arr[mid+1] {
			// return false, mid, arr[:mid-1]
			isGhole(arr[:mid-1])
		} else {
			// return false, mid, arr[mid+1:]
			isGhole( arr[mid+1:])
		}
	}

}
