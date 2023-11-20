package person

import (
	"main/address"
	"main/myerrors"
)

type User struct {
	Person
	Email    string
	password string
}

func NewUser(
	firstName string, lastName string,
	age int, address address.Address,
	email string, password string) (*User, error) {
	if age < 0 {
		return nil, &myerrors.AgeError{Message: "Возраст не может быть отрицательным"}
	}
	return &User{Person: Person{FirstName: firstName, LastName: lastName, age: age, Address: address}, Email: email, password: password}, nil
}

func (u *User) ChangePassword(oldPassword string, newPassword string) error {
	if oldPassword != u.password {
		return &myerrors.IncorrectPasswordError{Message: "Неверный старый пароль"}
	}
	if len(newPassword) < 8 {
		return &myerrors.InvalidPasswordError{Message: "Пароль должен удовлетворять условиям безопасности"}
	}
	u.password = newPassword
	return nil
}

func (u *User) Login(email string, password string) error {
	if email == u.Email && password == u.password {
		return nil
	}
	return &myerrors.InvalidLoginDataError{Message: "Неверный email или пароль"}
}
