package main

import (
	"fmt"
	"log"
	"os"

	greeting "example.com/greetings"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(1)

	// Verificar si se proporcionó un argumento
	if len(os.Args) < 3 {
		log.Fatal("Por favor proporciona un nombre como argumento")
	}

	// Obtener el nombre del primer argumento
	nombre := os.Args[1]
	apellido := os.Args[2]
	nombreCompleto := nombre + " " + apellido

	message, err := greeting.Hello(nombreCompleto)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
