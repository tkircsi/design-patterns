package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func MainProtoFactory() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 20)

	fmt.Printf("%+v\n", john)
	fmt.Printf("%+v\n", jane)
}

type Address2 struct {
	Suite        int
	Street, City string
}

type Employee struct {
	Name   string
	Office Address2
}

func (e *Employee) DeepCopy() *Employee {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	_ = enc.Encode(e)

	var ne Employee
	dec := gob.NewDecoder(&buff)
	_ = dec.Decode(&ne)
	return &ne
}

var mainOffice = Employee{
	Name: "",
	Office: Address2{
		Street: "123 East Dr",
		Suite:  0,
		City:   "London",
	},
}

var auxOffice = Employee{
	Name: "",
	Office: Address2{
		Street: "66 West Dr",
		Suite:  0,
		City:   "London",
	},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}
