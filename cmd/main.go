package main

import (
	"fmt"

	"github.com/nicolasassi/pbner/document"
)

func main() {
	text := "PRATI DIDOMIZIO-GAUNTNER, Wan B.; GUTHEIL KREB-SWISTON, Zelina G.; IGLESIAS LANGÓNBRUNN GILLANI, Laetitia Elianiv; RÜMIR ENSTROM TAKAO, Selin M.; POLLIARD RENFER BEACHER, Franciszek do I. Self-knowledge: reflecting on the influence of IHC publications on its own event."
	doc := document.NewDocument(text)
	for _, p := range doc.Puncts {
		fmt.Printf("%+v\n", p)
	}
}
