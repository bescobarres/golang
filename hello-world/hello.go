package main

import (
	"fmt"
	"log"

	greeting "example.com/greetings"
)

func main() {

	log.SetPrefix("Greetings: ")
	log.SetFlags(1)

	message, err := greeting.Hello("Bryan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
