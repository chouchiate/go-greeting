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

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with message.
	messages := make(map[string]string)
	// Loop through the received slice of anmes, calling
	// the Hello Function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map , associate the retrieved message with
		// the name
		messages[name] = message

	}
	return messages, nil
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