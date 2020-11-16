package base

import "fmt"



func ShowPrint() {
	num1 := 3
	fmt.Println(num1)
	fmt.Printf("%#v", num1)
	fmt.Printf("%T", num1)
	fmt.Println()

	arr := []int{1, 2, 3}
	fmt.Println(arr)
	fmt.Printf("%T", arr)
	fmt.Println()

	slice := make([]int, 10)
	//slice := []int{1, 2}
	fmt.Println("len", len(slice))
	fmt.Println("cap", cap(slice))
	slice = append(slice, 3)
	for _, v:= range slice {
		fmt.Println(v)
	}
	fmt.Printf("%T", slice)
}