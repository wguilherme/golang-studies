package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

func (u *User) UpdateName(newName string) {
	/*
		pointer receiver
		nesse conntexto, o "u" é equivalente a um self ou this em outras linguagens
		em Go, a convenção é usar o primeiro caractere de cada palavra do tipo
		por exemplo, se o tipo fosse "UserAccount", o primeiro caractere de cada palavra seria "ua"
		então o "u" é o primeiro caractere de "User"
		aqui usamos pointer receiver porque queremos modificar o valor do campo Name
		se usássemos um receiver normal, o campo Name não seria modificado
		a convennção eem Go é que se um método modifica o valor de um campo, ele deve ser um pointer receiver
		outra convenção é que se um dos métodos de um tipo é um pointer receiver, todos os outros métodos devem ser pointer receivers
	*/
	u.Name = newName
}

func main() {
	user := User{Name: "Old Name", ID: 1}
	/*
		a variável user não é um ponteiro e mesmo assim funciona
		isso é chamado de Pointer Indirection
	*/
	user.UpdateName("New Name")
	fmt.Println("raw", user)

	res, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("json", string(res))

	/*
		output:
		raw {New Name 1}
		json {"name":"New Name","id":1}
	*/
}
