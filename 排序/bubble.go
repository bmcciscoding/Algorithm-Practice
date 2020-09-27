package main

import "fmt"

func main() {
	arr := []int{10, 33, 45, 88, 76, 23, 32, 55, 66, 92}
	fmt.Println(arr)
	fmt.Println(bubble(arr))
}

func bubble(arr []int) []int {
	count := len(arr)
	var i, j int
	for i=0; i<count-1; i++ {
		for j=0; j<count-1-i; j++ {
			if (arr[j] > arr[j+1]) {
				t := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = t
			}
		}
	}
	return arr
}

