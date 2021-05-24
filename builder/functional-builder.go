package builder

import "fmt"

func MainFunctionalBuilder() {
	eb := EmployeeBuilder{}
	e := eb.Called("Tibor").Build()
	fmt.Printf("%+v\n", *e)

	e = eb.Called("Bogi").WorksAsA("senior developer").Build()
	fmt.Printf("%+v\n", *e)
}

type Employee struct {
	name, position string
}

type employeeMod func(*Employee)
type EmployeeBuilder struct {
	actions []employeeMod
}

func (b *EmployeeBuilder) Called(name string) *EmployeeBuilder {
	b.actions = append(b.actions, func(e *Employee) {
		e.name = name
	})
	return b
}

func (b *EmployeeBuilder) WorksAsA(position string) *EmployeeBuilder {
	b.actions = append(b.actions, func(e *Employee) {
		e.position = position
	})
	return b
}

func (b *EmployeeBuilder) Build() *Employee {
	e := Employee{}
	for _, action := range b.actions {
		action(&e)
	}
	return &e
}
