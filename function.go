package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func getPersons(fileName string) ([]Person, error) {
	filename := fmt.Sprintf(fileName)
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf(
			"reading file %q failed: %w",
			filename,
			err,
		)
	}

	rows, err := csv.NewReader(bytes.NewBuffer(content)).ReadAll()
	if err != nil {
		return nil, fmt.Errorf(
			"parsing csv: %w",
			err,
		)
	}

	var persons []Person
	for _, row := range rows {
		if len(row) != 2 {
			return nil, fmt.Errorf("expected row of lenght 2, got %d", len(row))
		}

		caser := cases.Title(language.Polish)
		name := caser.String(row[0])

		birth, err := time.Parse("2006-01-02", row[1])
		if err != nil {
			return nil, fmt.Errorf(
				"invalid date %q: %w",
				row[1],
				err,
			)
		}

		person := Person{
			Name:  name,
			Birth: birth,
		}
		persons = append(persons, person)
	}

	return persons, nil
}
