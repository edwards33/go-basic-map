package main

import "fmt"

func main() {
	// init
	var user map[string]string = map[string]string{
		"name":     "Vasily",
		"lastName": "Romanov",
	}

	// immediately with the right capacity
	profile := make(map[string]string, 10)

	// amount of elements
	mapLength := len(user)

	fmt.Printf("%d %+v\n", mapLength, profile)

	// if there is no key, it will return the default value for the type
	mName := user["middleName"]
	fmt.Println("mName:", mName)

	// checks for key existence
	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	// empty variabl; checks existence of the keyь
	_, mNameExist2 := user["middleName"]
	fmt.Println("mNameExist2", mNameExist2)

	delete(user, "lastName")
	fmt.Printf("%#v\n", user)
}
