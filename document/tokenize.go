package document

import (
	"strings"
)

func PunckTokenize(text string) []string {
	var puctSensitive string
	chars := strings.Split(text, "")
	for i, s := range chars {
		if isPunct(s) {
			if i != 0 {
				if !isSpace(chars[i-1]) {
					puctSensitive += ` `
				}
			}
			puctSensitive += s
			if i+1 != len(chars) {
				if !isSpace(chars[i+1]) {
					puctSensitive += ` `
				}
			}
			continue
		}
		puctSensitive += s
	}
	return whiteSpace.Split(puctSensitive, -1)
}
