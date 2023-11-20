package main

import (
	"fmt"
	"main/address"
	"main/person"
)

func main() {
	address1 := address.NewAddress("Russia", "Saint-P", "Vavilovikh", "10c2")
	person1, err := person.NewPerson("Misha", "Shisha", 7, address1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*person1)
		user1, err := person.NewUser("Misha", "Shisha", 7, address1, "misha@gmail.com", "123456789")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(*user1)
			fmt.Println(user1.Login("misha@gmail.com", "123456789"))
			user1.ChangePassword("123456789", "zxcasdqwe")
		}
	}
}
