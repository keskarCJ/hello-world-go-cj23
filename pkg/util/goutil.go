package util

import "errors"

const HelloWorld = "Hello World Constant"

//public
func GiveMessage() string  {
	return "World Hello"
}

//public
func CreateMessage(x int) (string, int, error)  {
	message := "Hello World"

	if x<0 {
		return "", x, errors.New("positives only")
	}

	return message, x, nil
}