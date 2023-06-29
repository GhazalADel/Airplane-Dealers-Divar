package utils

import "errors"

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
