package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// Add an init function to seed the rand package with the current time. 
// Go executes init functions automatically at program startup, 
// after global variables have been initialized.
// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}
// Dive more at https://golang.org/doc/effective_go#init

// randomFormat returns on of the set of greeting messages.
// The returned message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string {
		"Hi, %v. Welcome!",
		"Greate to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected emssage format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}