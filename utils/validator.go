package utils

import (
	"errors"
	"net/mail"
	"strconv"
	"strings"
)

// This Function Validates Input Email.
func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// This Function Validates Input Phone Number.
func ValidatePhone(phone string) bool {
	hasCharacter := false
	for _, digit := range phone {
		if digit < 48 || digit > 57 {
			hasCharacter = true
			break
		}
	}
	if hasCharacter {
		return false
	}
	return strings.HasPrefix(phone, "09") && len(phone) == 11
}

// This Function Parse Input String to Integer on Input Base.
func parseInt(s string, base int) int {
	n, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		return 0
	}
	return int(n)
}

// This Function Validates Input National ID.
func ValidateNationalID(id string) bool {
	l := len(id)

	if l < 8 || parseInt(id, 10) == 0 {
		return false
	}

	id = ("0000" + id)[l+4-10:]
	if parseInt(id[3:9], 10) == 0 {
		return false
	}

	c := parseInt(id[9:10], 10)
	s := 0
	for i := 0; i < 9; i++ {
		s += parseInt(id[i:i+1], 10) * (10 - i)
	}
	s = s % 11

	return (s < 2 && c == s) || (s >= 2 && c == (11-s))
}

func ValidateJsonFormat(jsonBody map[string]interface{}, fields ...string) (string, error) {
	msg := "OK"
	for _, field := range fields {
		if _, ok := jsonBody[field]; !ok {
			msg = "Input Json doesn't include " + field
			break
		}
	}
	if msg != "OK" {
		return msg, errors.New("")
	}
	return msg, nil
}
