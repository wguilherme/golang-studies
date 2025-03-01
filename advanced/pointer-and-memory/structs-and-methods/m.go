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
	// nesse conntexto, o "u" é equivalente a um self ou this em outras linguagens
	// em Go, a convenção é usar o primeiro caractere de cada palavra do tipo
	// por exemplo, se o tipo fosse "UserAccount", o primeiro caractere de cada palavra seria "ua"
	// então o "u" é o primeiro caractere de "User"
	u.Name = newName
}

func main() {
	user := User{Name: "Old Name", ID: 1}
	// a variável user não é um ponteiro e mesmo assim funciona
	// isso é chamado de Pointer Indirection
	user.UpdateName("New Name")
	fmt.Println(user.Name)
}
