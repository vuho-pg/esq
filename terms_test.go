package esq

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTerms(t *testing.T) {
	q := Bool().Filter(Terms().Field("payment_type", []string{"VISA", "MC"}).Boost(1.2))
	data, err := json.Marshal(q)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}
