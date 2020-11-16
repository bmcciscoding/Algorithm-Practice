# Slice 切片

在 Objc 和 Swift 中有可变数组和不可变数组的概念。在 GO 里，则分为数组和切片。数组即为不可变数组，切片即为可变数组

## 基操

```go
// 创建数组
arr1 := [3]int{1, 2, 3} 

// 遍历数组
for i:=0; i<len(arr1); i++ {
  fmt.Println(arr1[i])
}

// 增
append(arr1, 4) // 这样会报错，append 只有 slice 才能使用

slice1 := arr1[0:]
slice1 = append(slice1, 4)


```


