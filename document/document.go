package document

type Document struct {
	RawText       string   `json:"text"`
	PuncTokenized []string `json:"punct_tokenized"`
	Puncts        []*Punct `json:"punctuation"`
}

type AnnotatedDocument struct {
	Document    Document
	Annotations Annotations
}

type Punct struct {
	Index               int    `json:"index"`
	Type                string `json:"type"`
	NTokensAfter        int    `json:"n_tokens_after"`
	BeforeIsUpper       bool   `json:"before_is_upper"`
	AfterIsUpper        bool   `json:"after_is_upper"`
	AfterIsNum          bool   `json:"after_is_n"`
	BeforeIsNum         bool   `json:"before_is_n"`
	BeforeIsPunct       bool   `json:"before_is_punct"`
	AfterIsPunct        bool   `json:"after_is_punct"`
	NToNextSimilarPunct int    `json:"n_to_next_similar_punct"`
	NToNextDotPunct     int    `json:"n_to_next_dot"`
}

func NewDocument(text string) *Document {
	tokenized := punctTokenize(text)
	doc := &Document{
		RawText:       text,
		PuncTokenized: tokenized,
	}
	for i, tok := range tokenized {
		if !isPunct(tok) {
			continue
		}
		doc.Puncts = append(doc.Puncts, &Punct{
			Index:               i,
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
