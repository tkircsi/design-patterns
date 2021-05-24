package factory

import "fmt"

func MainProtoFactory() {
	d := NewEmployee(Developer)
	d.Name = "Bogi"
	fmt.Printf("%+v\n", *d)

	b := NewEmployee(Boss)
	b.Name = "Boci"
	fmt.Printf("%+v\n", *b)

	t := NewEmployee(5)
	t.Name = "Tibi"
	fmt.Printf("%+v\n", *t)
}

type Employee2 struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
	Boss
)

func NewEmployee(role int) *Employee2 {
	switch role {
	case Developer:
		return &Employee2{
			Position:     "developer",
			AnnualIncome: 60000,
		}
	case Manager:
		return &Employee2{
			Position:     "manager",
			AnnualIncome: 80000,
		}
	case Boss:
		return &Employee2{
			Position:     "CEO",
			AnnualIncome: 110000,
		}
	default:
		panic("unknown position")
	}
}
