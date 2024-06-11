package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("NOMBRE VACIO")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	mensanes := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		mensanes[name] = message
	}
	return mensanes, nil

}

func randomFormat() string {
	formats := []string{
		"Hello %v! !Bienvenido",
		"Que bueno verte, %v",
		"Saludo que bueno conocerte, %v",
	}

	return formats[rand.Intn(len(formats))]
}
