package esq

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSortNested(t *testing.T) {
	q := Search().Query(
		Nested("parent",
			Bool().
				Must(Range("parent.age").GTE(21)).
				Filter(Nested("parent.child", Match("parent.child.name", "matt"))),
		),
	).Sort(
		SortNested("parent").
			Filter(Range("parent.age").GTE(21)).
			Nested(
				SortNested("parent.child").
					Filter(Match("parent.child.name", "matt")),
			),
	)

	j, err := json.MarshalIndent(q, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))

}
