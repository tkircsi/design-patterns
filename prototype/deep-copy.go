package prototype

import "fmt"

func MainProtoCopy() {
	john := Person{
		Name: "John",
		Address: &Address{
			StreetAddress: "123 London Rd",
			City:          "London",
			Country:       "UK",
		},
	}

	// jane := john
	// jane.Name = "Jane"                               // OK
	// jane.Address.StreetAddress = "321 Baker str 45." // Not OK

	jane := john
	jane.Name = "Jane" // OK
	jane.Address = &Address{
		StreetAddress: john.Address.StreetAddress,
		City:          john.Address.City,
		Country:       john.Address.Country,
	}
	jane.Address.StreetAddress = "321 Baker str 45." // OK
	fmt.Printf("%+v, %+v\n", john, john.Address)
	fmt.Printf("%+v %+v\n", jane, jane.Address)
}

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}
