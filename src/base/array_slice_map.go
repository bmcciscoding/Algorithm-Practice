package base

import "fmt"

func ShowArraySliceMap() {
	sum := 0
	for i:=0; i<100; i++ {
		if i%2 != 0 {
			continue
		}
		sum += 1
	}
	fmt.Println("the sum is", sum)
	use()
}

func use() {
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[2])

	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr2)

	for i, v:=range arr2 {
		fmt.Println("arr", i, v)
	}

	slice1 := array[2:5]
	fmt.Println(slice1)
	fmt.Println(array[:5])
	fmt.Println(array[1:])
	slice1[1] = "11"
	fmt.Println(slice1)

	makeSlice := make([]string, 4)
	fmt.Println(makeSlice)

	ageMap := map[string]int{"winter":26}
	ageMap["show"] = 18
	fmt.Println(ageMap)
	delete(ageMap, "show")
	fmt.Println(ageMap)
	ageMap["show1"] = 18
	ageMap["show2"] = 18
	for k, v := range ageMap {
		fmt.Println("key is", k, "val is", v)
	}

	s := "show string"
	slice_s := []byte(s)
	fmt.Println(s)
	fmt.Println(slice_s)
}