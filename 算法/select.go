package main

import "fmt"

func main() {
	arr := []int{10, 33, 45, 88, 76, 23, 32, 55, 66, 92}
	fmt.Println(arr)
	fmt.Println(select_sort(arr))
}

func select_sort(arr []int) []int {
	count := len(arr)
	for i :=0; i<count-1; i++ {
		min := i
		for j :=i+1; j<count-1; j++ {
			if (arr[j] < arr[min]) {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

