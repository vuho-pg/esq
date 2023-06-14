package esq

type TermsSetInnerBuilder struct {
	Terms_                    []string `json:"terms"`
	MinimumShouldMatchField_  *string  `json:"minimum_should_match_field,omitempty"`
	MinimumShouldMatchScript_ *string  `json:"minimum_should_match_script,omitempty"`
}

type TermsSetBuilder struct {
	fieldName string
	TermsSet  map[string]*TermsSetInnerBuilder `json:"terms_set"`
}

func (*TermsSetBuilder) IsQuery() {}

func TermsSet(
	_fieldName string,
	_terms []string,
) *TermsSetBuilder {
	return &TermsSetBuilder{
		fieldName: _fieldName,
		TermsSet: map[string]*TermsSetInnerBuilder{
			_fieldName: {
				Terms_: _terms,
			},
		},
	}
}

func (_termsSet *TermsSetBuilder) Terms(_terms ...string) *TermsSetBuilder {
	_termsSet.TermsSet[_termsSet.fieldName].Terms_ = _terms
	return _termsSet
}
func (_termsSet *TermsSetBuilder) MinimumShouldMatchField(_minimumShouldMatchField string) *TermsSetBuilder {
	_termsSet.TermsSet[_termsSet.fieldName].MinimumShouldMatchField_ = &_minimumShouldMatchField
	return _termsSet
}
func (_termsSet *TermsSetBuilder) MinimumShouldMatchScript(_minimumShouldMatchScript string) *TermsSetBuilder {
	_termsSet.TermsSet[_termsSet.fieldName].MinimumShouldMatchScript_ = &_minimumShouldMatchScript
	return _termsSet
}
