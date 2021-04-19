package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var persons = []Person{}
var person = Person{}

type Person struct {
	id         int
	first_name string
	last_name  string
	phone      string
}

func (p *Person) add(db *sqlx.DB) {
	err := db.QueryRowx("INSERT INTO contactList(id, first_name, last_name, phone) values($1, $2, $3, $4)", p.id, p.first_name, p.last_name, p.phone)
	if err != nil {
		log.Println(err, " Bla Bla2")
		return
	}
	fmt.Println("person saved")
}

// func (p *Person) add(db *sqlx.DB) {
// 	tx := db.MustBegin()
// 	tx.MustExec("INSERT INTO contact_list(id, first_name, last_name, phone) values($1, $2, $3, $4)", p.id, p.first_name, p.last_name, p.phone)
// 	tx.Commit()
// }

func (p *Person) delete(db *sqlx.DB, id int) {
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM contact_list WHERE id = $1", id)
	tx.Commit()
}

func (p *Person) get(db *sqlx.DB) {
	row, err := db.Queryx("SELECT id, first_name, last_name, phone FROM contact_list WHERE id = $1 LIMIT 1;", p.id)
	if err == sql.ErrNoRows {
		fmt.Println("No obj is found!")
		return
	} else if err != nil {
		fmt.Println(err)
	}

	if row.Next() {
		err := row.Scan(&p.id, &p.first_name, &p.last_name, &p.phone)
		if err != nil {
			log.Println(err, " Bla Bla2")
		}
	}
	fmt.Println(p)
}

func (p *Person) list(db *sqlx.DB) {
	rows, _ := db.Queryx("SELECT * FROM contact_list")
	for rows.Next() {
		var person Person
		err := rows.Scan(&person.id, &person.first_name, &person.last_name, &person.phone)
		if err != nil {
			log.Println(err, " Bla Bla2")
		}
		persons = append(persons, person)
	}
	fmt.Println(persons)
}

func main() {
	psqlUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		5432,
		"postgres",
		"postgres",
		"psql",
	)

	db, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Whooho connected")

	person := Person{id: 3, first_name: "Jalili", last_name: "Soliyev", phone: "+9784120952"}

	var str string
	fmt.Print("Command: ")
	fmt.Scan(&str)
	if str == "add" {
		person.add(db)
	} else if str == "posts" {
		person.list(db)
	} else if str == "get" {
		p1 := Person{id: 2}

		p1.get(db)
	} else if str == "delete" {
		var i int
		fmt.Print("Person id: ")
		fmt.Scan(&i)
		person.delete(db, i)
		person.list(db)
	}

}

// func person_list(db *sqlx.DB) {
// 	rows, _ := db.Queryx("SELECT * FROM contactList")
// 	for rows.Next() {
// 		var person Person
// 		err := rows.Scan(&person.id, &person.first_name, &person.last_name, &person.phone)
// 		if err != nil {
// 			log.Println(err, " Bla Bla2")
// 		}
// 		persons = append(persons, person)
// 	}
// 	fmt.Println(persons)
// }

// func add_person(db *sqlx.DB) {
// 	var str string
// 	fmt.Print("First name: ")
// 	input, _ := fmt.Scan(&str)
// 	if input == 1 {
// 		person.first_name = str
// 		fmt.Println("Ok")
// 		last_name(db)
// 	}
// }

// func last_name(db *sqlx.DB) {
// 	var str string
// 	fmt.Print("Last name: ")
// 	input, _ := fmt.Scan(&str)
// 	if input == 1 {
// 		person.last_name = str
// 		fmt.Println("Ok")
// 		phone(db)
// 	}
// }

// func phone(db *sqlx.DB) {
// 	var str string
// 	fmt.Print("Phone: ")
// 	input, _ := fmt.Scan(&str)
// 	if input == 1 {
// 		person.phone = str
// 		fmt.Println("Ok")
// 		// Save person
// 		save(db)
// 	}
// }
// func save(db *sqlx.DB) {
// 	tx := db.MustBegin()
// 	tx.MustExec("INSERT INTO contactList(first_name, last_name, phone) values($1, $2, $3)", person.first_name, person.last_name, person.phone)
// 	tx.Commit()
// }
