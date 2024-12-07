package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"example.com/hello"
)

func main() {
	// hello package
	hello.Hello()

	// greetings package
	// See error logs
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Tom", "Samantha", "Kim"}

	//Call Greeting func
	messages, err := greetings.Greetings(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
