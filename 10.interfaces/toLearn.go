package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PersonI interface {
	Add(plist *Person)
	Update(id uint64, plist *Person) error
	Delete(id uint64) error
	Get(id uint64) *Person
	Print()
}

type Person struct {
	PersonId  	uint64
	FirstName 	string
	LastName 	string
	Address 	Address
}

func NewPerson() PersonI {
	return &PersonList{}
}

type Address struct {
	Country    		string
	Region 			string
	Street 			string
	HouseNumber 	int
}

type PersonList struct {
	Persons []*Person
}

func (plist *PersonList) Add(person *Person)  {
	plist.Persons = append(plist.Persons, person)
}

func (plist *PersonList) Update(id uint64, person *Person) error {
	for _, val := range plist.Persons {
		if id == val.PersonId {
			val.FirstName = person.FirstName
			val.LastName = person.LastName
			val.Address = person.Address
		}
	}
	return nil
}

func (plist *PersonList) Delete(id uint64) error {
	for ind, val := range plist.Persons {
		if id == val.PersonId {
			plist.Persons = append(plist.Persons[:ind], plist.Persons[ind + 1:]...)
		}
	}
	return nil
}

func (plist *PersonList) Get(id uint64) *Person {
	var p Person
	for _, val := range plist.Persons {
		if id == val.PersonId {
			p.PersonId = val.PersonId
			p.FirstName = val.FirstName
			p.LastName = val.LastName
			p.Address = val.Address
			}
		}

	return &p
}

func (plist *PersonList) Print() {
	for _, val := range plist.Persons {
		fmt.Println(val)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	plist := NewPerson()

	adr1 := Address{
		Country:     "England",
		Region:      "London",
		Street:      "Lincoln",
		HouseNumber: 7,
	}

	adr2 := Address{
		Country:     "England",
		Region:      "London",
		Street:      "Lincoln",
		HouseNumber: 11,
	}

	adr3 := Address{
		Country:     "England",
		Region:      "London",
		Street:      "Lincoln",
		HouseNumber: 8,
	}

	p1 := Person{
		PersonId:  456789,
		FirstName: "Otebk",
		LastName:  "Muzropov",
		Address:   adr1,
	}

	p2 := Person{
		PersonId:  16387138,
		FirstName: "Sardor",
		LastName:  "Jo'rayev",
		Address:   adr2,
	}

	p3 := Person{
		PersonId:  1783788,
		FirstName: "John",
		LastName:  "McField",
		Address:   adr3,
	}

	plist.Add(&p1)
	plist.Add(&p2)
	plist.Add(&p3)
	plist.Print()

}
