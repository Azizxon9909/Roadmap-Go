package main

import (
	"os"
	"github.com/azizxon9909/book_keeper"
)

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int `gorm:"unique_index"`
	PersonID   int
}

func main()  {
	// Loading env variables
	dialect := os.Getenv("DIALECT")
}
