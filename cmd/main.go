package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/nicolasassi/pbner/document"
)

var docs []document.AnnotatedDocument

func main() {
	// t()
	data, err := ioutil.ReadFile("train-data/super_noice.jsonl")
	if err != nil {
		panic(err)
	}
	annotations := document.ReadData(data)
	for _, a := range annotations {
		doc := document.NewDocument(a.Text)
		docs = append(docs, document.AnnotatedDocument{
			Document:    *doc,
			Annotations: a,
		})
	}
	f, err := os.OpenFile("train-data/new_noice.jsonl", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range docs {
		b, err := json.Marshal(d)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write(b); err != nil {
			log.Fatal(err)
		}
		if _, err := f.WriteString("\n"); err != nil {
			log.Fatal(err)
		}
	}
}

func t() {
	text := ".PRATI DIDOMIZIO-GAUNTNER, Wan B.; GUTHEIL KREB-SWISTON, Zelina G.; IGLESIAS LANGÓNBRUNN GILLANI, Laetitia Elianiv; RÜMIR ENSTROM TAKAO, Selin M.; POLLIARD RENFER BEACHER, Franciszek do I. Self-knowledge: reflecting on the influence of IHC publications on its own event."
	doc := document.NewDocument(text)
	for index, p := range doc.Puncts {
		fmt.Printf("%v: %+v\n", index, p)
	}
}
