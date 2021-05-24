package builder

import "fmt"

func MainBuilderFacets() {
	pb := NewPersonBuilder()

	pb.
		Lives().
		At("Széchenyi út 76.").
		In("Budapest").
		WithPostcode("1211").
		Works().
		At("NASA").
		AsA("developer").
		Earning(65000)

	p := pb.Build()
	fmt.Printf("%+v\n", p)
}

type Person struct {
	Street, Postcode, City string

	Company, Position string
	AnnualIncome      int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

func (b *PersonAddressBuilder) At(street string) *PersonAddressBuilder {
	b.person.Street = street
	return b
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonJobBuilder) At(company string) *PersonJobBuilder {
	b.person.Company = company
	return b
}

func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome
	return b
}
