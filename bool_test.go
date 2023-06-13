package esq

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBool(t *testing.T) {
	q := Bool().SetFilter(
		MatchPhrase())
	j, _ := json.MarshalIndent(q, "", "\t")
	fmt.Println(string(j))
}
