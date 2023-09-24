package main

import "fmt"

type Adress struct {
	city   string
	street string
}

type User struct {
	firstName string
	lastName  string
	email     string
	age       int
	Adress    Adress
}

func (u *User) print_me() {
	fmt.Println(*u)
}

func (u *User) SetName(name string) {
	u.firstName = name
}

func print_user(u1 *User, u2 User) {
	fmt.Println(u1.firstName, "\n", u2.firstName)
}

// func NewUser(...) *User{
// 	user := User{...}

// }

type Admin struct {
	IsAdmin bool
	User
	parent User
}

// func (u *Admin) print_me() {
// 	fmt.Printf("i'm an admin and my name is %s %s", u.firstName, u.lastName)
// 	fmt.Printf("i live in %v", u.Adress)
// }

func main() {
	user1 := User{firstName: "Misha", lastName: "A", email: "exampple@g.com", age: 20}
	user2 := User{firstName: "Boris", lastName: "B", email: "exampple@g.com", age: 22}
	print_user(&user1, user2)
	// user1.print_me()
	// user2.print_me()
	// adm := Admin{true, user1, user2}
	// adm2 := Admin{}

	// adm2.print_me()
	// adm.print_me()
	// user1.SetName("asd")
}
