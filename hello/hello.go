package main

import (
	"fmt"
	"log"

	"test.com/greetings"
)

func main() {
	log.SetPrefix("greetings")
	log.SetFlags(0)

	names := []string{"Navy", "Samantha", "Darrin"}
	messages, error := greetings.Hellos(names)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(messages["Navy"])
}
