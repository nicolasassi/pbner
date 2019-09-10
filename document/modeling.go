package document

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"time"
)

// DataModel represents a single entry of Prodigy's JSON Lines output.
// LabeledEntity` is a structure defined by prose that specifies where the
// entities are within the given `Text`.
type DataModel struct {
	Text   string
	Spans  []LabeledEntity
	Answer string
}

type LabeledEntity struct {
	Start int
	End   int
	Label string
}

type RawDataModel struct {
	Entities [][]interface{} `json:"entities"`
	Text     string          `json:"text"`
}

// ReadData reads our JSON Lines file line-by-line, populating a
// slice of `DataModel` structures.
func ReadData(jsonLines []byte, limit int) []DataModel {
	dec := json.NewDecoder(bytes.NewReader(jsonLines))
	entries := []DataModel{}
	for {
		raw := RawDataModel{}
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
	Shuffle(entries)
	if limit == -1 {
		return entries
	}
	return entries[:limit]
}

func (raw RawDataModel) modelData() DataModel {
	ent := DataModel{}
	ent.Text = raw.Text
	ent.Answer = "accept"
	for _, span := range raw.Entities {
		ent.Spans = append(ent.Spans,
			LabeledEntity{
				Start: getInt(span[0]),
				End:   getInt(span[1]),
				Label: span[2].(string),
			})
	}
	return ent
}

func Shuffle(dm []DataModel) {
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
func Split(data []DataModel) ([]DataModel, []DataModel) {
	cutoff := int(float64(len(data)) * 0.8)

	var train, test []DataModel

	for i, entry := range data {
		if i < cutoff {
			train = append(train, entry)
		} else {
			test = append(test, entry)
		}
	}

	return train, test
}
