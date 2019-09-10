package document

import (
	"strings"
)

func PunckTokenize(text string) []string {
	var tokenized []string
	chars := strings.Split(text, "")
	for _, char := range chars {
		if isSpace(char) {
			continue
		}
		tokenized = append(tokenized, char)
	}
	return tokenized
}
