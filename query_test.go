package esq

import (
	"testing"
)

func TestQuery(t *testing.T) {
	q := New(Nested("register_services",
		Bool().
			Should(
				Bool().MustNot(Wildcard("register_services.parent_merchant_code", "*")),
				Bool().MustNot(Exists("register_services.parent_merchant_code")),
			).MinimumShouldMatch("1"),
	),
	)
	data, err := q.JSONBeauty()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
