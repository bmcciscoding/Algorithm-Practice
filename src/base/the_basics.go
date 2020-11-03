package base

import "fmt"
import "math"

func ConstantsAndVariables() {
	const PI = 3.1415926
	var day int
	age := 18
	fmt.Println("const PI =", PI)
	fmt.Println("var day int", day)
	fmt.Println("age :=", age)
}

func TypeAnnotatios() {
	var msg string
	msg = "hello world"
	var hour, min, sec int
	fmt.Println(msg, hour, min, sec)
}

func NamingConstantsAndVariables() {

}

func Comments() {
	// single line comment

	/*
			mutil line comment
	*/
}

func IntAndFloat() {
	const MaxUInt32 = math.MaxUint32
	const MinInt32 = math.MinInt32
} 

func TypeSafetyandTypeInference() {
	var age int
	// age = 18.4
	fmt.Println(age)
}

func NumericLiterals() {
	const decimal = 17
	const binary = 0b10101
	const octal = 0o10101
	const hexadecimal = 0x10101
}