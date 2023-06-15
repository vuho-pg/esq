package example

import (
	"fmt"
	. "github.com/vuho-pg/esq"
)

func Compound() {
	q, err := New(
		Bool().
			Must(Term("user.id", "kimchy")).
			Filter(Term("tags", "production")).
			MustNot(Range("age").GTE(10).LTE(20)).
			Should(Term("tags", "env1"), Term("tags", "env2")).
			MinimumShouldMatch("1").Boost(1.0),
	).JSONBeauty()
	if err != nil {
		panic(err)
	}
	fmt.Println(q)
}
