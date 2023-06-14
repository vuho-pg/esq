package esq

import (
	"encoding/json"
	"testing"
)

func TestNested(t *testing.T) {
	n := Nested("register_services", Bool().Filter(MatchPhrase("register_services.parent_merchant_code", "NCCDF2GR"))).ScoreMode("avg")
	j, _ := json.MarshalIndent(n, "", "\t")
	t.Log(string(j))
}
