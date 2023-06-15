package esq

import (
	"bytes"
	"encoding/json"
	"io"
)

type QueryBuilder struct {
	Query Query `json:"query"`
}

func New(q Query) *QueryBuilder {
	return &QueryBuilder{Query: q}
}

func (q *QueryBuilder) Reader() (io.Reader, error) {
	data, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func (q *QueryBuilder) JSON() (string, error) {
	data, err := json.Marshal(q)
	return string(data), err
}

func (q *QueryBuilder) JSONBeauty() (string, error) {
	data, err := json.MarshalIndent(q, "", "\t")
	return string(data), err
}
