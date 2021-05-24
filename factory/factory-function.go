package factory

import "fmt"

func MainFactoryFunction() {
	p := NewPerson("John", 48)
	fmt.Printf("%+v\n", p)

	p = NewPerson("Grand Jim", 130)
	fmt.Printf("%+v\n", p)

	p = NewPerson("Little Marry", 12)
	fmt.Printf("%+v\n", p)
}

type Person struct {
	Name    string
	Age     int
	isAdult bool
	Live    bool
}

func NewPerson(name string, age int) *Person {
	p := Person{
		Name:    name,
		Age:     age,
		isAdult: true,
		Live:    true,
	}
	if age < 18 {
		p.isAdult = false
	}
	if age > 120 {
		p.Live = false
	}
	return &p
}
