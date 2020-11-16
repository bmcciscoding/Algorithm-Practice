package base

import "fmt"
import "sort"

func ShowArraySliceMap() {
	fmt.Println("show array usage")
	array_usage()
}

func use() {
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println( array[2])

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

// array 的
func array_usage() {
	// 创建数组
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("创建", arr1)

	// 遍历数组
	for i:=0; i<len(arr1); i++ {
		fmt.Println("遍历", arr1[i])
	}
	
	// 增、删、改、查
	// arr1 = append(arr1, 4); // error: only slice
	slice1 := arr1[0:]
	slice1 = append(slice1, 6)
	fmt.Println("增", slice1)

	deleteIndex := 2
	slice1 = append(slice1[:deleteIndex], slice1[deleteIndex+1:]...)
	fmt.Println("删", slice1)

	slice1[2] = 20
	fmt.Println("改", slice1)

	fmt.Println("查", slice1[2])

	// 排序
	sort.Ints(slice1)
	fmt.Println("排序", slice1)

	// 倒转
	sort.Sort(sort.Reverse(sort.IntSlice(slice1)))
	fmt.Println("倒转", slice1)
}