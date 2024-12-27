package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	// log.SetFlags(log.Ldate | log.Ltime)
	log.SetFlags(0)

	names := []string{"Jobe", "Júlia", "Nicholas"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
