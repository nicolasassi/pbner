package document

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"time"
)

type Annotations struct {
	Text  string          `json:"-"`
	Spans []LabeledEntity `json:"spans"`
}

type LabeledEntity struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Label string `json:"label"`
}

type RawAnnotations struct {
	Entities [][]interface{} `json:"entities"`
	Text     string          `json:"text"`
}

// ReadData reads our JSON Lines file line-by-line, populating a
// slice of `DataModel` structures.
func ReadData(jsonLines []byte, args ...interface{}) []Annotations {
	dec := json.NewDecoder(bytes.NewReader(jsonLines))
	entries := []Annotations{}
	for {
		raw := RawAnnotations{}
		err := dec.Decode(&raw)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		ent := raw.modelData()
		entries = append(entries, ent)
	}
	if args != nil {
		if args[0].(int) != -1 && args[0].(int) != 0 {
			return entries[:args[0].(int)]
		}
	}
	return entries
}

func (raw RawAnnotations) modelData() Annotations {
	ent := Annotations{}
	for _, span := range raw.Entities {
		ent.Text = raw.Text
		ent.Spans = append(ent.Spans,
			LabeledEntity{
				Start: getInt(span[0]),
				End:   getInt(span[1]),
				Label: span[2].(string),
			})
	}
	ent = fix(ent)
	return ent
}

func Shuffle(dm []Annotations) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(dm); n > 0; n-- {
		randIndex := r.Intn(n)
		dm[n-1], dm[randIndex] = dm[randIndex], dm[n-1]
	}
}

func getInt(n interface{}) int {
	return int(n.(float64))
}

// Split divides our synthetic data set into two groups: one for training
// our model and one for testing it.
// We're using an 80-20 split here, although you may want to use a different
// split.
func Split(data []Annotations) ([]Annotations, []Annotations) {
	cutoff := int(float64(len(data)) * 0.8)

	var train, test []Annotations

	for i, entry := range data {
		if i < cutoff {
			train = append(train, entry)
		} else {
			test = append(test, entry)
		}
	}

	return train, test
}
