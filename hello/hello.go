package main

import (
    "fmt"

    "greetings"
	"log"
)

func main() {
    // Get a greeting message and print it.

	log.SetPrefix("greetings: ")
	log.SetFlags(0)	

	names:=[]string{"Gladys","Samantha","Darrin"}
    messages, err := greetings.Hellos(names)
	if err!=nil {
		log.Fatal(err);
	}
    fmt.Println(messages)
}