package esq

import (
	"bytes"
	"encoding/json"
	"io"
)

func (_searchBuilder *SearchBuilder) Reader() (io.Reader, error) {
	data, err := json.Marshal(_searchBuilder)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
