package main

import (
	"fmt"
)

type User struct {
	Name    *string
	Age     int
	Address AddressInfo
}

type AddressInfo struct {
	Street string
}

func main() {
	println("Hello world")

	// var intExample int = 1
	// var doubleExample float32 = 1.0
	// var boolExample bool = true
	// var stringExample string = "Hola"

	// name := true

	name := ""
	userInfo := User{
		Address: AddressInfo{
			Street: "Nada",
		},
	}

	userInfo.Name = &name

	fmt.Println(userInfo)
	println(fmt.Sprintf("Name: %s, Address: %s", userInfo.Name, userInfo.Address.Street))
}
