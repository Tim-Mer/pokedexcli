package main

import (
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	output := strings.Fields(text)

	return output
}
