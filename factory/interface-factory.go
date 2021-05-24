package factory

import "fmt"

func MainInterfaceFacory() {
	p := NewPerson2("Matt", 23)
	// can not access to private members
	// p.age++
	p.SayHello()

	p = NewPerson2("Jim", 102)
	p.SayHello()
}

type Person2 interface {
	SayHello()
}

type person2 struct {
	name string
	age  int
}

func (p *person2) SayHello() {
	fmt.Printf("Hi, my name is %s, I'm %d years old.\n", p.name, p.age)
}

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Hi, my name is %s, I'm %d years old and soooo tired!\n", p.name, p.age)
}

func NewPerson2(name string, age int) Person2 {
	if age > 100 {
		return &tiredPerson{
			name: name,
			age:  age,
		}
	}
	return &person2{
		name: name,
		age:  age,
	}
}
