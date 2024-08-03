package main

import (
	"strings"
)

// Complete the method/function so that it converts dash/underscore
// delimited words into camel casing. The first word within the output
// should be capitalized only if the original word was capitalized (known
// as Upper Camel Case, also often referred to as Pascal case). The next
// words should be always capitalized.
func ToCamelCase(s string) string {
	if s == "" {
		return s
	}

	s = strings.ReplaceAll(s, "_", "-")

	c, s := string(s[0]), s[1:]
	s = strings.ToLower(s)

	words := strings.Split(s, "-")
	first := c + words[0]

	var newWords []string
	for _, w := range words[1:] {
		newWords = append(
			newWords,
			strings.ToUpper(string(w[0]))+w[1:],
		)
	}

	return first + strings.Join(newWords, "")
}
