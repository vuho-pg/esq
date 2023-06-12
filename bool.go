package esq

type BoolBuilder struct {
    Must *[]Query `json:"must,omitempty"`
    Filter *[]Query `json:"filter,omitempty"`
    MustNot *[]Query `json:"must_not,omitempty"`
    Should *[]Query `json:"should,omitempty"`
    MinimumShouldMatch *int `json:"minimum_should_match,omitempty"`
    Boost *float64 `json:"boost,omitempty"`
}
func (b *BoolBuilder) IsQuery() {}


func Bool(
) *BoolBuilder {
    return &BoolBuilder{
    }
}
func (b *BoolBuilder) SetMust(_must ...Query) *BoolBuilder {
    b.Must = &_must
    return b
}
func (b *BoolBuilder) SetFilter(_filter ...Query) *BoolBuilder {
    b.Filter = &_filter
    return b
}
func (b *BoolBuilder) SetMustNot(_mustNot ...Query) *BoolBuilder {
    b.MustNot = &_mustNot
    return b
}
func (b *BoolBuilder) SetShould(_should ...Query) *BoolBuilder {
    b.Should = &_should
    return b
}
func (b *BoolBuilder) SetMinimumShouldMatch(_minimumShouldMatch int) *BoolBuilder {
    b.MinimumShouldMatch = &_minimumShouldMatch
    return b
}
func (b *BoolBuilder) SetBoost(_boost float64) *BoolBuilder {
    b.Boost = &_boost
    return b
}