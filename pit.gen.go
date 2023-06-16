package esq

import "encoding/json"

type PITBuilder struct {
	ID_        string  `json:"id"`
	KeepAlive_ *string `json:"keep_alive,omitempty"`
}

func PIT(
	_iD string,
) *PITBuilder {
	return &PITBuilder{
		ID_: _iD,
	}
}

func (_pITBuilder *PITBuilder) ID(_iD string) *PITBuilder {
	_pITBuilder.ID_ = _iD
	return _pITBuilder
}
func (_pITBuilder *PITBuilder) KeepAlive(_keepAlive string) *PITBuilder {
	_pITBuilder.KeepAlive_ = &_keepAlive
	return _pITBuilder
}

func (_pITBuilder *PITBuilder) JSON() ([]byte, error) {
	return json.Marshal(_pITBuilder)
}

func (_pITBuilder *PITBuilder) JSONIndent() ([]byte, error) {
	return json.MarshalIndent(_pITBuilder, "", "\t")
}
