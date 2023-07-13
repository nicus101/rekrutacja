package main

import (
	"fmt"
	"log"
)

func main() {
	persons, err := getPersons("php_internship_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(persons)

	// TODO: policz imiona
	// TODO: policz daty urodzenia
}
