package main

import "fmt"

func introduction(firstName string, lastName string) (string, string) {
	introFirstName := "Hello My First Name Is " + firstName
	introLastName := "Hello My Last Name Is " + lastName
	return introFirstName, introLastName
}

func main() {
	firstName, lastName := introduction("John", "Doe")
	fmt.Println(firstName, lastName)
	
	firstName2, _ := introduction("John", "Doe")
	fmt.Println(firstName2)
}
