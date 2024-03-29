package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	contactInfo
	lastName  string
	firstName string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// 	var alex person

	// 	alex.firstName = "Alex"
	// 	alex.lastName = "Anderson"

	// 	fmt.Printf("%+v", alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jimPointer := &jim

	jim.updateName("Jimmy") // pointer shortcut
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
