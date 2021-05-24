package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func MainProtoSerial() {
	john := &Person{
		Name: "John",
		Address: &Address{
			StreetAddress: "123 London Rd",
			City:          "London",
			Country:       "UK",
		},
		Friends: []string{"Chris", "Matt"},
	}

	jane := john.DeepCopy2()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker str"
	jane.Friends[0] = "Angela"
	fmt.Printf("%+v\t%+v\n", john, john.Address)
	fmt.Printf("%+v\t%+v\n", jane, jane.Address)
}

func (p *Person) DeepCopy2() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(*p)

	fmt.Println(b.String())

	q := Person{}
	d := gob.NewDecoder(&b)
	_ = d.Decode(&q)

	return &q
}
