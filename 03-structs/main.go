package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

func (c contactInfo) print(params ...string) {
	prefix := ""
	if len(params) != 0 {
		prefix = params[0]
	}

	fmt.Printf("%vEmail: %v\n", prefix, c.email)
	fmt.Printf("%vZIP Code: %v\n", prefix, c.zip)
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("Name: %v\n", p.firstName)
	fmt.Printf("Surname: %v\n", p.lastName)
	fmt.Println("Contact info:")
	p.contactInfo.print(">>> ")
}

func main() {

	var myPerson person

	myPerson.firstName = "Roberto"
	myPerson.lastName = "Pastor"
	myPerson.contactInfo.email = "roberto.pastor@anywho.com"
	myPerson.contactInfo.zip = 13245

	personFromCtor := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email: "jd@gmail.com",
			zip:   12345,
		},
	}

	fmt.Printf("%+v", myPerson)
	fmt.Printf("%+v", personFromCtor)
	fmt.Println("")

	myPerson.print()
	myPerson.contactInfo.print()

	personFromCtor.updateName("Johnie")
	personFromCtor.print()
}
