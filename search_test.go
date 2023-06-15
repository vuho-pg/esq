package esq

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	q := Search().Query(Term("user.id", "kimchy"))
	j, err := q.JSON()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}
