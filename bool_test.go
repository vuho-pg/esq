package esq

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBool(t *testing.T) {
	b := Bool().SetFilter(
		MatchPhrase().Set("register_services.parent_merchant_code", MatchPhraseField("NCCDF2GR")),
	)

	data, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
