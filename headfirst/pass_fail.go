package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	// System widget count
	fmt.Print("Enter a grade : ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
