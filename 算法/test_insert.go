package main

import "fmt"

func main() {
	arr := []int{10, 33, 45, 88, 76, 23, 32, 55, 66, 92}
	fmt.Println(arr)
	fmt.Println(insert_sort(arr))
}

func insert_sort(arr []int) []int {
	count := len(arr)
	for i :=1; i<count-1; i++ {
		now := arr[i]
		j := i-1
		for j > 0 && arr[j] > now {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = now
	}
	return arr
}