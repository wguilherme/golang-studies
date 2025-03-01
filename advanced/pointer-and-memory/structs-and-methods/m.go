package main

import (
	"fmt"
)

type User struct {
	Name string
	ID   uint64
}

func (u *User) UpdateName(newName string) {
	// pointer receiver
	u.Name = newName
}

func main() {
	user := User{Name: "Old Name", ID: 1}
	// a variável user não é um ponteiro e mesmo assim funciona
	// isso é chamado de Pointer Indirection
	user.UpdateName("New Name")
	fmt.Println(user.Name)
}
