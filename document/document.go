package document

import "strings"

type Document struct {
	RawText       string   `json:"text"`
	PuncTokenized []string `json:"punct_tokenized"`
	Puncts        []*Punct `json:"punctuation"`
}

type AnnotatedDocument struct {
	Document    *Document
	Annotations Annotations
}

func NewAnnotatedDocument(doc *Document, annot Annotations) *AnnotatedDocument {
	annotDoc := new(AnnotatedDocument)
	for _, punctObj := range doc.Puncts {
		for _, annotObj := range annot.Spans {
			if punctObj.IndexText == annotObj.Start || punctObj.IndexText == annotObj.End {
				punctObj.IsClassDelimiter = true
			}
		}
	}
	annotDoc.Document = doc
	annotDoc.Annotations = annot
	return annotDoc
}

type Punct struct {
	IndexText           int    `json:"index"`
	Type                string `json:"type"`
	IndexToken          int    `json:"index_token"`
	NTokensAfter        int    `json:"n_tokens_after"`
	BeforeIsUpper       bool   `json:"before_is_upper"`
	AfterIsUpper        bool   `json:"after_is_upper"`
	AfterIsNum          bool   `json:"after_is_n"`
	BeforeIsNum         bool   `json:"before_is_n"`
	BeforeIsPunct       bool   `json:"before_is_punct"`
	AfterIsPunct        bool   `json:"after_is_punct"`
	NToNextSimilarPunct int    `json:"n_to_next_similar_punct"`
	NToNextDotPunct     int    `json:"n_to_next_dot"`
	IsClassDelimiter    bool   `json:"is_class_delimiter,omitempty"`
}

type indexCounter struct {
	lastIndex   int
	splitedText []string
}

func newIndexCounter(text string) *indexCounter {
	return &indexCounter{lastIndex: 0, splitedText: strings.Split(text, "")}
}

func NewDocument(text string) *Document {
	tokenized := punctTokenize(text)
	doc := &Document{
		RawText:       text,
		PuncTokenized: tokenized,
	}
	indexCounter := newIndexCounter(text)
	for i, tok := range tokenized {
		if !isPunct(tok) {
			continue
		}
		index := 0
		for textIndex, punct := range indexCounter.splitedText {
			if textIndex <= indexCounter.lastIndex {
				continue
			}
			if punct == tok {
				indexCounter.lastIndex = textIndex
				index = textIndex
				break
			}
		}
		doc.Puncts = append(doc.Puncts, &Punct{
			IndexToken:          i,
			IndexText:           index,
			Type:                tok,
			NTokensAfter:        len(tokenized) - i,
			BeforeIsUpper:       beforeIsUpper(tokenized, i),
			AfterIsUpper:        afterIsUpper(tokenized, i),
			AfterIsNum:          afterIsNum(tokenized, i),
			BeforeIsNum:         beforeIsNum(tokenized, i),
			BeforeIsPunct:       beforeIsPunct(tokenized, i),
			AfterIsPunct:        afterIsPunct(tokenized, i),
			NToNextSimilarPunct: nextSimilar(tokenized, i),
			NToNextDotPunct:     nextDotPunct(tokenized, i),
		})
	}
	return doc
}
