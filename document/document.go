package document

type Document struct {
	RawText       string   `json:"text"`
	PunkTokenized []string `json:"punkt_tokenized"`
	Puncts        []*Punct `json:"punctuation"`
}

type AnnotatedDocument struct {
	Document    Document
	Annotations Annotations
}

type Punct struct {
	Index               int
	Type                string
	NWordsAfter         int
	BeforeIsUpper       bool
	AfterIsUpper        bool
	AfterIsNum          bool
	BeforeIsNum         bool
	BeforeIsPunkt       bool
	AfterIsPunkt        bool
	NToNextSimilarPunkt int
	NToNextDotPunkt     int
}

func NewDocument(text string) *Document {
	tokenized := PunckTokenize(text)
	doc := &Document{
		RawText:       text,
		PunkTokenized: tokenized,
	}
	for i, tok := range tokenized {
		if !isPunct(tok) {
			continue
		}
		doc.Puncts = append(doc.Puncts, &Punct{
			Index:               i,
			Type:                tok,
			NWordsAfter:         len(tokenized) - i,
			BeforeIsUpper:       beforeIsUpper(tokenized, i),
			AfterIsUpper:        afterIsUpper(tokenized, i),
			AfterIsNum:          afterIsNum(tokenized, i),
			BeforeIsNum:         beforeIsNum(tokenized, i),
			BeforeIsPunkt:       beforeIsPunkt(tokenized, i),
			AfterIsPunkt:        afterIsPunkt(tokenized, i),
			NToNextSimilarPunkt: nextSimilar(tokenized, i),
			NToNextDotPunkt:     nextPunkt(tokenized, i),
		})
	}
	return doc
}
