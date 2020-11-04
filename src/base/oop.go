
package base

import "fmt"

type Student struct {
	name string
}

func Test() {
	s := Student{"haha"}
	fmt.Println(s)
	//sayHi(s)
	s.sayHi()
	fmt.Println(s)
	//realSayHi(&s)
	s.realSayHi()
	fmt.Println(s)
}

func sayHi(p Student) {
	p.name = "new name"
}

func (s Student) sayHi() {
	s.name = "name"
}

func realSayHi(p *Student) {
	p.name = "real name"
}

func (s *Student) realSayHi() {
	s.name = "real name"
}
