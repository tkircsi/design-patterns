package factory

import "fmt"

func MainFactoryGenerator() {
	developerFacory := NewEmployeeFacotry("developer", 60000)
	managerFactory := NewEmployeeFacotry("manager", 80000)

	developer := developerFacory("Bogi")
	manager := managerFactory("Boci")

	fmt.Printf("%+v\n", *developer)
	fmt.Printf("%+v\n", *manager)

	bossFacotry := NewEmployeeFacotry2("CEO", 120000)
	// can modify attributes after creation
	bossFacotry.AnnualIncome = 130000
	boss := bossFacotry.Create("Tibike")
	fmt.Printf("%+v\n", *boss)
}

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional
func NewEmployeeFacotry(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:         name,
			Position:     position,
			AnnualIncome: annualIncome,
		}
	}
}

// structural
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{
		Name:         name,
		Position:     f.Position,
		AnnualIncome: f.AnnualIncome,
	}
}

func NewEmployeeFacotry2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		Position:     position,
		AnnualIncome: annualIncome,
	}
}
