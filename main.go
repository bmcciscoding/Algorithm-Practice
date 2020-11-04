package main

import "base"
import "fmt"

func main() {
	base.Test()
	s := new(base.Animal)
	s.Name = "Name"
	fmt.Println(s)
}