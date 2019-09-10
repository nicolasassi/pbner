package document

import (
	"fmt"
	"regexp"
)

var (
	upper      = regexp.MustCompile(`[\p{Lu}]+`)
	letter     = regexp.MustCompile(`[\p{L}]+`)
	digit      = regexp.MustCompile(`\d`)
	whiteSpace = regexp.MustCompile(`[\s\t\n\r\f\v]`)
	punct      = regexp.MustCompile(`[!\\"#$%&'()*+,\-./:;<=>?@\[\]¨\^\_\x60\{\|\}\~]`)
)

// IsPunct determines if a character is a punctuation symbol.
func isPunct(s string) bool {
	return punct.MatchString(s)
}

// IsSpace determines if a character is a whitespace character.
func isSpace(s string) bool {
	return whiteSpace.MatchString(s)
}

// IsLetter determines if a character is letter.
func isLetter(c string) bool {
	return letter.MatchString(c)
}

// IsAlnum determines if a character is a letter or a digit.
func isAlnum(s string) bool {
	return digit.MatchString(s) || isLetter(s)
}

func isUpper(s string) bool {
	return upper.MatchString(s)
}

func beforeIsUpper(tokens []string, index int) bool {
	if index == 0 {
		return false
	}
	return isUpper(tokens[index-1][len(tokens[index-1])-1:])
}

func afterIsUpper(tokens []string, index int) bool {
	fmt.Println(index, len(tokens))
	if index == len(tokens)-1 {
		return false
	}
	fmt.Println(tokens[index+1])
	return isUpper(string(tokens[index+1]))
}
