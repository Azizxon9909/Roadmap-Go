package main

import (
	"fmt"
	"strconv"
)

type Contact struct {
	firstName   string
	lastName    string
	phoneNumber int
}

func (c *Contact) updateContact(firstName, lastName string, phoneNumber int) {
	c.firstName = firstName
	c.lastName = lastName
	c.phoneNumber = phoneNumber
}

func (c *Contact) getContact() string {
	return c.firstName + " " + c.lastName + " " + strconv.Itoa(c.phoneNumber)
}

func main() {

	c1 := Contact{"Azizxon", "Jalilov", 998971009909}
	c2 := Contact{"Jasur", "Ismoilov", 998957774455}

	fmt.Println(c1.getContact())

	c2.updateContact("Alisher", "Abdullaev", 88899000)
	fmt.Println(c2.getContact())

}
