package person

import (
	"main/address"
	"main/myerrors"
)

type Person struct {
	FirstName string
	LastName  string
	age       int
	Address   address.Address
}

func NewPerson(firstName string, lastName string, age int, address *address.Address) (*Person, error) {
	if age < 0 {
		return nil, &myerrors.AgeError{Message: "Возраст не может быть отрицательным"}
	}
	return &Person{FirstName: firstName, LastName: lastName, age: age, Address: *address}, nil
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) SetAge(age int) error {
	if age < 0 {
		return &myerrors.AgeError{Message: "Возраст не может быть отрицательным"}
	}
	p.age = age
	return nil
}
