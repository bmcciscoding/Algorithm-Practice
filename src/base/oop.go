
package base

import "fmt"

// 封装 
// 定一个类
type Animal struct {
	name string
	Name string
}

func NewStudent() *Animal {
	a := new(Animal)
	return a
}

// 定义方法 方法名小写开头即为私有方法
func (s Animal) sayHi() {
	s.name = "name"
}

// 定义方法 方法名小写开头即为公开方法
func (s Animal) SayHi() {
	s.name = "name"
}

func (s *Animal) realSayHi() {
	s.name = "real name"
}

// 继承 
// 子类继承父类
type Cat struct {
	Animal
	age int 
}

func (c *Cat) sleep() {
	fmt.Println("sleep")
	fmt.Println(c)
}

// 多态
type Sleep interface {
	sleep()
}

func test_Polymorphism(a Sleep) {
	a.sleep()
}

func Test() {
	cat := new(Cat)
	test_Polymorphism(cat)
}