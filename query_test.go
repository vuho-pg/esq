package esq

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	q := Search().Query(Match("user.id", "kimchy"))
	j, err := q.JSONIndent()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}

func TestSimpleTimeout(t *testing.T) {
	q := Search().Timeout("2s").Query(Match("user.id", "kimchy"))
	j, err := q.JSONIndent()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}

func TestSimpleTrackTotalHits(t *testing.T) {
	q := Search().TrackTotalHits(true).Query(Match("user.id", "elkbee"))
	j, err := q.JSONIndent()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}

// https://www.elastic.co/guide/en/elasticsearch/reference/current/collapse-search-results.html#collapse-search-results
func TestCollapseSearch(t *testing.T) {
	q := Search().
		Query(Match("message", "GET /search")).
		Collapse(Collapse("user.id")).
		Sort(Sort("http.response.bytes", SortValue().Order("desc"))).
		From(0)
	j, err := q.JSONIndent()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}

// https://www.elastic.co/guide/en/elasticsearch/reference/current/collapse-search-results.html#expand-collapse-results
func TestExpandCollapse(t *testing.T) {
	q := Search().
		Query(Match("message", "GET /search")).
		Collapse(
			Collapse("user.id").
				InnerHits(InnerHits().
					Name("most_recent").
					Size(5).
					Sort(Sort("@timestamp", SortValue().Order("desc"))),
				).
				MaxConcurrentGroupSearches(4),
		).
		Sort(Sort("http.response.bytes", SortValue().Order("desc")))
	j, err := q.JSONIndent()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}

// https://www.elastic.co/guide/en/elasticsearch/reference/current/collapse-search-results.html#collapsing-with-search-after
func TestCollapsingWithSearchAfter(t *testing.T) {
	q := Search().
		Query(Match("message", "GET /search")).
		Collapse(Collapse("user.id")).
		Sort(Sort("user.id", SortValue())).
		SearchAfter("dd5ce1ad")
	j, err := q.JSONIndent()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}

// https://www.elastic.co/guide/en/elasticsearch/reference/current/collapse-search-results.html#second-level-of-collapsing
func TestSecondLevelOfCollapsing(t *testing.T) {
	q := Search().
		Query(Match("message", "GET /search")).
		Collapse(
			Collapse("geo.country_name").
				InnerHits(InnerHits().
					Name("by_location").
					Collapse(Collapse("user.id")).
					Size(3),
				),
		)
	j, err := q.JSONIndent()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}
