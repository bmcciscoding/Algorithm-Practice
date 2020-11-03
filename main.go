package main

import "fmt"
// import "base"

func main() {
	const len = 1
	arr := []int{}
	slice := make([]int, 0)
	arr = append(arr, 1)
	arr[0] = 1
	//slice[0] = 1

	// s := make([]int, len)
	fmt.Println(arr)
	fmt.Println(slice)
}