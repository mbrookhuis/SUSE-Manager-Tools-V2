// Package contains - check if a string contains a substring
package contains

import (
	"errors"
	"os"
	"strings"
)

// Contains - list of strings contain a string
//
// param: s
// param: str
// return: bool
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// SubInString - list of strings contain a substring
//
// param: s
// param: str
// return: bool
func SubInString(s []string, str string) bool {
	for _, v := range s {
		if strings.Contains(v, str) {
			return true
		}
	}
	return false
}

// PartOff - substring is part of string
//
// param: s
// param: str
// return: bool
func PartOff(s string, str string) bool {
	return strings.Contains(s, str)
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, errors.New("present")
	}
	if os.IsNotExist(err) {
		return false, errors.New("not exist")
	}
	return false, err
}
