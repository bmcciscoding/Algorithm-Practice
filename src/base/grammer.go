package base

import (
    "fmt"
    "time"
)

// func main() {
//     leanrFor()
//     leanrSwitch()
//     leanrDefer()
//     learnPoint()
//     learnStruct()
//     learnArray()
//     learnDict()
//     learnClosure()
// }

func LearnFor() {
    sum := 0
    fmt.Println("begin for")
	for i:=0; i<10; i++ {
		sum += i
	}
    fmt.Println(sum)
    fmt.Println("begin if")
    if sum > 20 {
        fmt.Println(sum)
    }
}

func leanrSwitch() {
    t := time.Now().Hour()
    switch t { 
    case 1:
        fmt.Println("<12")
    case 2:
        fmt.Println(">12")        
    default:
        fmt.Println(t)
    }

    switch {
    case t < 12:
        fmt.Println("<12")
    case t > 12:
        fmt.Println(">12")        
    }
}

// count done 3 2 1 
func leanrDefer() {
    fmt.Println("count")
    for i := 1; i <= 3; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("done")
}

func learnPoint() {
    var p *int 
    i := 42
    p = &i
    fmt.Println(*p)
    *p = 21
    fmt.Println(*p)
    fmt.Println(p)
}

type Point struct {
    X, Y int
}
func learnStruct() {
    var (
        v1 = Point{1, 2}
        v2 = Point{X: 1}
        v3 = Point{}
    )
    fmt.Println(v1, v2, v3)
}

func learnArray() {
    p := []int{1, 2, 3, 4, 5}
    fmt.Println(p)
    fmt.Println("length:", len(p))
    fmt.Println(p[2:4])
    fmt.Println(p[:3])
    fmt.Println(p[3:])
}

func learnDict() {
    var m = map[string]Point {
        "a": {1, 2},
        "b": {2, 3},
    }
    fmt.Println(m)
}

func learnClosure() {
    temp := 5
    block := func(x, y int) int {
        temp = x * y * temp
        return temp
    }
    fmt.Println(block(1,2))
    fmt.Println(block(2,4))
    fmt.Println(temp)
}