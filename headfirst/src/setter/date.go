package main

import (
	"errors"
	"fmt"
	"log"
)

// Date Structure
type Date struct {
	Year  int
	Month int
	Day   int
}

// SetYear Year Setter Method
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func main() {
	date := Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year)
}
