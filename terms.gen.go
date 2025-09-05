package esq

type TermsInnerBuilder map[string]any

type TermsBuilder struct {
	Terms TermsInnerBuilder `json:"terms"`
}

func (*TermsBuilder) IsQuery() {}

func Terms() *TermsBuilder {
	return &TermsBuilder{
		Terms: TermsInnerBuilder{},
	}
}
func (_terms *TermsBuilder) Field(key_ string, value_ []string) *TermsBuilder {
	_terms.Terms[key_] = value_
	return _terms
}
func (_terms *TermsBuilder) Boost(_boost any) *TermsBuilder {
	_terms.Terms["boost"] = _boost
	return _terms
}
