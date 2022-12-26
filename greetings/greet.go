package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the name
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name received")
	}

	msg := fmt.Sprintf("Hello %v", name)
	return msg, nil
}
