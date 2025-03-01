package main

import (
	"errors"
	"fmt"
)

func main() {

	user, err := NewUser(true)

	if err != nil {
		fmt.Println("Error on create user")
		return
	}

	user.Foo()
}

type User struct {
	foo string
}

func (u *User) Foo() {
	fmt.Println(u.foo)
}

func NewUser(wantErr bool) (*User, error) {
	if wantErr {
		return nil, errors.New("error")
	}
	return &User{}, nil
}
