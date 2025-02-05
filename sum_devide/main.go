package main

import "fmt"

func main() {

	arr := []int{10, 3, 5, 6, 2}

	var new_arr []int
	for i := 0; i < len(arr); i++ {
		sum := 1
		for j := 0; j < len(arr); j++ {
			if j != i {
				sum =  sum * arr[j]
			}
		}
		new_arr = append(new_arr, sum)
	}
	fmt.Println(new_arr)

}
