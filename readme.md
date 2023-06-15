# Elastic search DSL query builder

ESQ Is an Elasticsearch DSL v7.7 query builder for Go. 
## Installation
```shell
go get -u github.com/vuho-pg/esq
```
## Feature
`esp` supports many query types of Elasticsearch DSL, including:
- [x] [Compound queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.17/compound-queries.html)
- [x] [Full text queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.17/full-text-queries.html)
- [ ] [Geo queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.17/geo-queries.html)
- [ ] [Shape queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.17/shape-queries.html)
- [x] [Joining queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.17/joining-queries.html)
- [x] [Match all query](https://www.elastic.co/guide/en/elasticsearch/reference/7.17/query-dsl-match-all-query.html)
- [ ] [Span queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.17/span-queries.html)
- [ ] [Specialized queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.17/specialized-queries.html)
- [x] [Term-level queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.17/term-level-queries.html)

This library is just a query builder, it does not provide any function to interact with Elasticsearch.

Please note that this library is still in development and not all query types are supported yet.

## Why?
I tried a few libraries to build Elasticsearch queries but none of them satisfied me, they are too complex and to 'Java' for me. So I decided to write my own library that is simple and easy to use, use method chaining make it easy to build complex queries.

## Example
And example of bool query [here](./example/compound.go)

From this query
```json
{
  "query": {
    "bool" : {
      "must" : {
        "term" : { "user.id" : "kimchy" }
      },
      "filter": {
        "term" : { "tags" : "production" }
      },
      "must_not" : {
        "range" : {
          "age" : { "gte" : 10, "lte" : 20 }
        }
      },
      "should" : [
        { "term" : { "tags" : "env1" } },
        { "term" : { "tags" : "deployed" } }
      ],
      "minimum_should_match" : 1,
      "boost" : 1.0
    }
  }
}
```
We can build it with this code
```go
New(Bool().
	Must(Term("user.id", "kimchy")).
	Filter(Term("tags", "production")).
	MustNot(Range("age").GTE(10).LTE(20)).
	Should(Term("tags", "env1"), Term("tags", "env2")).
	MinimumShouldMatch("1").Boost(1.0))
```

## References
[Elasticsearch DSL](https://elasticsearch-dsl.readthedocs.io/en/latest/index.html)
