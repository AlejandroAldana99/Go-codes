package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name     *string
	Age      int
	Address  AddressInfo
	IsActive bool
}

type AddressInfo struct {
	Street string
}

func main() {

	/*
		Variable examples
	*/
	// var intExample int = 1
	// var doubleExample float32 = 1.0
	// var boolExample bool = true
	// var stringExample string = "Hola"

	// name := true <-- Low type declaration

	linstOfNumbers := []int{} // make([]int,2) // [1,2] <-- Tuple
	listUser := []User{}

	name := ""
	userInfo := User{
		Address: AddressInfo{
			Street: "Nada",
		},
		IsActive: false,
	}

	userInfo.Name = &name
	listUser = append(listUser, userInfo)

	fmt.Println(userInfo)
	println(fmt.Sprintf("Name: %v, Address: %s", userInfo.Name, userInfo.Address.Street))

	println(fmt.Sprintf("Is active: %v", isActiveValidator(userInfo.IsActive)))

	maxNumber, err := validMaxNumber(linstOfNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}

	println(maxNumber)
}

// Example of function validator
func isActiveValidator(isActive bool) bool {
	return isActive
}

// Example of function with 2 returns
func validMaxNumber(arrayInts []int) (int, error) {
	if len(arrayInts) <= 0 {
		return 0, errors.New("Invalid number")
	}

	number := 0
	for i := range arrayInts {
		number = arrayInts[i]
	}

	return number, nil
}
