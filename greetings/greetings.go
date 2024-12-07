package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Uppercase is used for exporting
func Greeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Greetings(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Greeting(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

// Lowercase is not exported
// message is selected at random
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Will met!",
	}

	return formats[rand.Intn(len(formats))]
}
