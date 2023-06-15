package esq

import (
	"testing"
)

func TestQuery(t *testing.T) {
	q := New(
		Bool().
			Must(Term("user.id", "kimchy")).
			Filter(Term("tags", "production")).
			MustNot(Range("age").GTE(10).LTE(20)).
			Should(Term("tags", "env1"), Term("tags", "env2")).
			MinimumShouldMatch("1").Boost(1.0),
	)

	data, err := q.JSONBeauty()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
	/* Output:
	   	"query": {
	   		"bool": {
	   			"must": [
	   				{
	   					"term": {
	   						"user.id": {
	   							"value": "kimchy"
	   						}
	   					}
	   				}
	   			],
	   			"must_not": [
	   				{
	   					"range": {
	   						"age": {
	   							"gte": 10,
	   							"lte": 20
	   						}
	   					}
	   				}
	   			],
	   			"should": [
	   				{
	   					"term": {
	   						"tags": {
	   							"value": "env1"
	   						}
	   					}
	   				},
	   				{
	   					"term": {
	   						"tags": {
	   							"value": "env2"
	   						}
	   					}
	   				}
	   			],
	   			"filter": [
	   				{
	   					"term": {
	   						"tags": {
	   							"value": "production"
	   						}
	   					}
	   				}
	   			],
	   			"minimum_should_match": "1",
	   			"boost": 1
	   		}
	   	}
	   }
	*/
}
