package document

import (
	"regexp"
)

var (
	upper           = regexp.MustCompile(`[\p{Lu}]+`)
	letter          = regexp.MustCompile(`[\p{L}]+`)
	digit           = regexp.MustCompile(`\d`)
	whiteSpace      = regexp.MustCompile(`[\s\t\n\r\f\v]`)
	multiWhiteSpace = regexp.MustCompile(`[\s\t\n\r\f\v]+`)
	punct           = regexp.MustCompile(`[!\\"#$%&'()*+,\-./:;<=>?@\[\]¨\^\_\x60\{\|\}\~]`)
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

func isNum(s string) bool {
	return digit.MatchString(s)
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

func beforeIsNum(tokens []string, index int) bool {
	if index == 0 {
		return false
	}
	return isNum(tokens[index-1][len(tokens[index-1])-1:])
}

func afterIsUpper(tokens []string, index int) bool {
	if index == len(tokens)-1 {
		return false
	}
	return isUpper(string(tokens[index+1]))
}

func beforeIsPunct(tokens []string, index int) bool {
	if index == 0 {
		return false
	}
	return isPunct(tokens[index-1][len(tokens[index-1])-1:])
}

func afterIsPunct(tokens []string, index int) bool {
	if index == len(tokens)-1 {
		return false
	}
	return isPunct(string(tokens[index+1]))
}

func afterIsNum(tokens []string, index int) bool {
	if index == len(tokens)-1 {
		return false
	}
	return isNum(string(tokens[index+1]))
}

func cleanDoubleSpaces(text string) string {
	return multiWhiteSpace.ReplaceAllString(text, " ")
}

func nextPunctType(tokens []string, index int) string {
	for i, token := range tokens {
		if i <= index {
			continue
		}
		if isPunct(token) {
			return token
		}
	}
	return ""
}

func nextPunct(tokens []string, index int) int {
	for i, token := range tokens {
		if i <= index {
			continue
		}
		if isPunct(token) {
			return i
		}
	}
	return -1
}

func lastPunctType(tokens []string, index int) string {
	for i, j := 0, len(tokens)-1; i < j; i, j = i+1, j-1 {
		tokens[i], tokens[j] = tokens[j], tokens[i]
	}
	index = len(tokens) - index
	for i, token := range tokens {
		if i <= index {
			continue
		}
		if isPunct(token) {
			return token
		}
	}
	return ""
}

func lastPunct(tokens []string, index int) int {
	for i, j := 0, len(tokens)-1; i < j; i, j = i+1, j-1 {
		tokens[i], tokens[j] = tokens[j], tokens[i]
	}
	index = len(tokens) - index
	for i, token := range tokens {
		if i <= index {
			continue
		}
		if isPunct(token) {
			return i - index
		}
	}
	return -1
}

func nextSimilar(tokens []string, index int) int {
	for i, token := range tokens {
		if i <= index {
			continue
		}
		if token == tokens[index] {
			return i
		}
	}
	return -1
}

func nextDotPunct(tokens []string, index int) int {
	for i, token := range tokens {
		if i <= index {
			continue
		}
		if isPunct(token) {
			return i
		}
	}
	return -1
}
