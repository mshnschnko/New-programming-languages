package main

import (
	"fmt"
	"reflect"
)

type User struct {
	firstName string `default:"Anton"`
	lastName  string
	email     string
	age       int
}

type Adress struct {
	country string
	city    string
}

// type User struct {
// 	firstName string `default:"AAA"`
// 	lastName  string
// 	email     string
// 	age       int
// 	Adress    Adress
// }

func NewAdress(city string, street string) *Adress {
	return &Adress{
		country: city,
		city:    street,
	}
}

func print_user(u *User) {
	fmt.Printf("name: %s %s\n", u.firstName, u.lastName)
	fmt.Printf("email: %s\n", u.email)
	fmt.Printf("age: %d\n", u.age)
}

func (u *User) print_me() {
	fmt.Printf("name: %s %s\n", u.firstName, u.lastName)
	fmt.Printf("email: %s\n", u.email)
	fmt.Printf("age: %d\n", u.age)
}

func (u *User) SetName(name string) {
	u.firstName = name
}

// func NewUser(
// 	lastName string,
// 	email string,
// 	age int,
// ) *User {
// 	user := User{
// 		lastName: lastName,
// 		email:    email,
// 		age:      age,
// 	}
// 	user.firstName = "Anton"
// 	return &user
// }

func NewUser(
	lastName string,
	email string,
	age int,
) *User {
	user := User{
		lastName: lastName,
		email:    email,
		age:      age,
	}
	typ := reflect.TypeOf(user)
	if user.firstName == "" {
		f, _ := typ.FieldByName("firstName")
		user.firstName = f.Tag.Get("default")
	}
	return &user
}

// func NewUser(firstName string,
// 	lastName string,
// 	email string,
// 	age int,
// ) *User {
// 	user := User{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		email:     email,
// 		age:       age,
// 	}
// 	user.Adress = *NewAdress("Russia", "Moscow")
// 	return &user
// }

type Website struct {
	url             string
	users_per_month int
}

type Admin struct {
	IsAdmin bool
	User
	department string
	Website
}

// func NewAdmin()

// func (u *Admin) print_me() {
// 	fmt.Printf("i'm an admin and my name is %s %s", u.firstName, u.lastName)
// 	fmt.Printf("i live in %v", u.Adress)
// }

func main() {
	user := NewUser("A", "exampple@g.com", 20)
	fmt.Println(user.firstName)

	ad := Admin{true, *user, "articles", Website{"url", 200}}
	fmt.Println(ad.users_per_month)

	// admin := Admin{true, *user, "articles"}
	// fmt.Printf("Hello! I'm %s %s and I administer the %s section",
	// 	admin.firstName, admin.lastName, admin.department)

	// fmt.Printf("user's adress is %s, %s", user.Adress.country, user.Adress.city)

	// user2 := User{lastName: "B", email: "exampple@g.com", age: 22}

	// user1 := NewUser("Misha", "A", "example@g.com", 20)
	// user1.print_me()
	// user1.SetName("Borya")
	// user1.print_me()

	// user2 := User{firstName: "Boris", lastName: "B", email: "exampple@g.com", age: 22}

	// print_user(user1, user2)
	// user1.SetName("kola")
	// user1.print_me()
	// user2.print_me()
	// adm := Admin{true, user1, user2}
	// adm2 := Admin{}

	// adm2.print_me()
	// adm.print_me()
	// user1.SetName("asd")
}
