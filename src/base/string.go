package base

import "fmt"
import "strings"
import "reflect"

func ShowString() {

 		const nihongo = "日本語"
    for index, runeValue := range nihongo {
				fmt.Printf("%d : %#U\n", index, runeValue)
        // fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }
	
	return
	// 字符串声明和赋值
	var s1 string
	const s2 = "zhangsan"
	s3 := "lisi"
	// 字符串拼接
	s4 := s1 + s2 +s3
	fmt.Println(s4)
	fmt.Println("s4[:1]\t", s4[:1])
	fmt.Println("s4[3]\t", s4[2:3])
	// 长度
	s4len := len(s4)
	fmt.Println(s4len)
	// 遍历
	// len
	fmt.Println("for use len")
	for i :=0; i<s4len; i++  {
		c := s4[i]
		fmt.Printf("%c type is %s\n", c, reflect.TypeOf(c))
	}
	// range
	fmt.Println("for use range")
	for _, c := range s4 {
		fmt.Printf("%c type is %s\n", c, reflect.TypeOf(c))
	}
	// 比较
	if s4 != s3 {	
		fmt.Println("s4 != s3")
	}
	// 包含
	if strings.Contains(s4, s3) {
		fmt.Println("s4 contain s3")
	}
}

// https://berryjam.github.io/2018/03/%E4%BB%8Egolang%E5%AD%97%E7%AC%A6%E4%B8%B2string%E9%81%8D%E5%8E%86%E8%AF%B4%E8%B5%B7/
// https://blog.golang.org/slices
// https://golang.org/pkg/strings/
