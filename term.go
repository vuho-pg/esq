package esq

type TermBuilder struct {
    Value string `json:"value"`
    Boost *float64 `json:"boost,omitempty"`
    CaseInsensitive *bool `json:"case_insensitive,omitempty"`
}
func (t *TermBuilder) IsQuery() {}


func Term(
_value string,
) *TermBuilder {
    return &TermBuilder{
       Value: _value,
    }
}
func (t *TermBuilder) SetValue(_value string) *TermBuilder {
    t.Value = _value
    return t
}
func (t *TermBuilder) SetBoost(_boost float64) *TermBuilder {
    t.Boost = &_boost
    return t
}
func (t *TermBuilder) SetCaseInsensitive(_caseInsensitive bool) *TermBuilder {
    t.CaseInsensitive = &_caseInsensitive
    return t
}