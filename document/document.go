package document

type Document struct {
	RawText       string
	PunkTokenized []string
	Puncts        []*Punct
}

type Punct struct {
	Index         int
	Type          string
	NWordsAfter   int
	BeforeIsUpper bool
	AfterIsUpper  bool
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
			Index:         i,
			Type:          tok,
			NWordsAfter:   len(tokenized) - i,
			BeforeIsUpper: beforeIsUpper(tokenized, i),
			AfterIsUpper:  afterIsUpper(tokenized, i),
		})
	}
	return doc
}
