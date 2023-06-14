package esq

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBool(t *testing.T) {
	b := Bool().Should(
		MatchPhrase("tcb_merchant_code", "TCB108"),
		MatchPhrase("tcb_merchant_id", "123456"),
	).MinimumShouldMatch("10")

	j, _ := json.MarshalIndent(b, "", "\t")
	fmt.Println(string(j))
}
