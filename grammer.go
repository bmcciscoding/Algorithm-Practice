package main

import (
    "fmt"
    "time"
)

func main() {
    sum()
    leanrSwitch()
}

func sum() {
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