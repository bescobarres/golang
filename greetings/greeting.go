package greeting

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// If no name was given, the return
	if name == "" {
		return "", errors.New("empty name")
	}

	//Hello return greeting for the named person.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v.!",
		"Hail, %v.! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
