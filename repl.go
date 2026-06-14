package main

import "strings"

func cleanInput(text string) []string {
	cleaned_text := strings.TrimSpace(text)
	cleaned_text = strings.ReplaceAll(cleaned_text, "  ", " ")
	lowercase_text := strings.ToLower(cleaned_text)
	return strings.Split(lowercase_text, " ")
}
