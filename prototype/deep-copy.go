package prototype

import "fmt"

func MainProtoCopy() {
	john := &Person{
		Name: "John",
		Address: &Address{
			StreetAddress: "123 London Rd",
			City:          "London",
			Country:       "UK",
		},
		Friends: []string{"Chris", "Matt"},
	}

	// jane := john
	// jane.Name = "Jane"                               // OK
	// jane.Address.StreetAddress = "321 Baker str 45." // Not OK

	// jane := john
	// jane.Name = "Jane" // OK
	// jane.Address = &Address{
	// 	StreetAddress: john.Address.StreetAddress,
	// 	City:          john.Address.City,
	// 	Country:       john.Address.Country,
	// }
	// jane.Address.StreetAddress = "321 Baker str 45." // OK

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker str"
	jane.Friends[0] = "Angela"
	fmt.Printf("%+v, %+v\n", john, john.Address)
	fmt.Printf("%+v %+v\n", jane, jane.Address)
}

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	q.Friends = make([]string, len(p.Friends))
	copy(q.Friends, p.Friends)
	return &q
}
