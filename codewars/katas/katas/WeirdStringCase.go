package kata

// Write a function that accepts a string, and returns the same string
// with all even indexed characters in each word upper cased, and all
// odd indexed characters in each word lower cased. The indexing just
// explained is zero based, so the zero-ith index is even, therefore
// that character should be upper cased and you need to start over for
// each word.

// The passed in string will only consist of alphabetical characters and
// spaces(' '). Spaces will only be present if there are multiple words.
// Words will be separated by a single space(' '). Examples:

// "String" => "StRiNg" "Weird string case" => "WeIrD StRiNg CaSe"

import (
	"strings"
)

func ToWeirdCase(str string) (result string) {
	up := true
	for _, rune := range str {
		char := strings.ToUpper(string(rune))
		if char == " " {
			up = true
		} else {
            if !up {
                char = strings.ToLower(char)
            }
            up = !up
        }
		result += char
	}
	return
}
