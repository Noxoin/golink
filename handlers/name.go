package handlers

import (
	"errors"
	"regexp"
)

var nameRegex = "^[A-Za-z0-9-_]+$"

func validateLinkName(name string) (bool, error) {
	return regexp.MatchString(nameRegex, name)
}

func getLinkName(path string) (string, error) {
	if len(path) == 0 {
		return "", errors.New("Empty path string")
	}
	if path[0] == '/' {
		path = path[1:]
	}
	matched, err := validateLinkName(path)
	if !matched || err != nil {
		return "", errors.New("Invalid Link")
	}
	return path, nil
}

