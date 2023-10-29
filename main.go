package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	// Open a new database connection
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// CREATE operation
	_, err = db.Exec("INSERT INTO users(name, age) VALUES (?, ?)", "John Doe", 30)
	if err != nil {
		log.Fatal(err)
	}

	// READ operation
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, age)
	}

	// UPDATE operation
	_, err = db.Exec("UPDATE users SET age = ? WHERE name = ?", 31, "John Doe")
	if err != nil {
		log.Fatal(err)
	}

	// DELETE operation
	_, err = db.Exec("DELETE FROM users WHERE name = ?", "John Doe")
	if err != nil {
		log.Fatal(err)
	}
}
