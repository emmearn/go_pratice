package main

import "fmt"

type concactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	concact   concactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		concact: concactInfo{
			email:   "jim.party@email.com",
			zipCode: 9400,
		},
	}

	fmt.Printf("%+v", jim)
}
