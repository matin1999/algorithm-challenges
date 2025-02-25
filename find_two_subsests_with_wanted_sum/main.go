package main

import (
	"fmt"
)

//You are given an array of integers arr[] and a target sum target. Your task is to find if there
//are any two distinct elements in the array that sum up to the given target. If a pair is found, return
//the indices of the two elements; otherwise, return -1.
//Constraints:

//The array may contain both positive and negative integers.
//The length of the array can go up to 10^5.
//Goal:

//You need to solve this problem in O(n) time complexity.

func main() {

	array := []int{1, 10, 8, 6, 4, 7}

	target_sum := 14

	dict := make(map[int]int)

	for index, val := range array {
		dict[val] = index
	}

	for i := 0; i < len(array); i++ {
		if val, exist := dict[target_sum-array[i]]; exist {
			fmt.Printf("indexes are %d %d \n" , i , val)
			return
		}
	}
	fmt.Println("-1")
}
