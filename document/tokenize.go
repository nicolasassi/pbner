package document

import (
	"strings"
)

func PunckTokenize(text string) []string {
	allChars := strings.Split(text, "")
	var allCharPlusSpace []string
	for i, char := range allChars {
		if i == 0 {
			if isPunct(char) {
				allCharPlusSpace = append(allCharPlusSpace, char, " ")
				continue
			}
		}
		if i == len(allChars)-1 {
			if isPunct(char) {
				allCharPlusSpace = append(allCharPlusSpace, " ", char)
				continue
			}
		}
		if isPunct(char) {
			if !isSpace(allChars[i-1]) {
				allCharPlusSpace = append(allCharPlusSpace, " ")
			}
			allCharPlusSpace = append(allCharPlusSpace, char)
			if !isSpace(allChars[i+1]) {
				allCharPlusSpace = append(allCharPlusSpace, " ")
			}
			continue
		}
		allCharPlusSpace = append(allCharPlusSpace, char)
	}
	text = strings.Join(allCharPlusSpace, "")
	text = cleanDoubleSpaces(text)
	return whiteSpace.Split(strings.TrimSpace(text), -1)
}
