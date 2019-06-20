package main

import "fmt"

type concactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	concactInfo // it's the same of declaring a concact field of type concactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		concactInfo: concactInfo{
			email:   "jim.party@email.com",
			zipCode: 9400,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// shortcut to a pointer turn a variable into a pointer. It works because there is a pointer at a receiver
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
