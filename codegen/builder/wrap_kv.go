package builder

import (
	"os"
	"text/template"

	"github.com/gobeam/stringy"
)

var wrapKVTempl *template.Template

func init() {
	var err error
	wrapKVTempl, err = template.ParseFiles("builder/wrap_kv.gotempl")
	if err != nil {
		panic(err)
	}
}

type WrapKVSetter struct {
	Name      string
	ValueType string
}

type WrapKVBuilder struct {
	SetterFn           []*WrapKVSetter
	Name               string
	FileName           string
	ReceiverName       string
	JsonTag            string
	ImplementInterface *InterfaceBuilder
	InnerType          *KeyValueBuilder
}

func WrapKV(name string, inner *KeyValueBuilder) *WrapKVBuilder {
	return &WrapKVBuilder{
		Name:         name,
		FileName:     "../" + stringy.New(name).SnakeCase().ToLower() + ".gen.go",
		ReceiverName: "_" + stringy.New(stringy.New(name).CamelCase()).LcFirst(),
		JsonTag:      stringy.New(name).SnakeCase().ToLower(),
		InnerType:    inner,
	}
}

func (m *WrapKVBuilder) Setter(name string, valueType string) *WrapKVBuilder {
	if valueType == "" {
		valueType = m.InnerType.ValueType
	}
	m.SetterFn = append(m.SetterFn, &WrapKVSetter{
		Name:      name,
		ValueType: valueType,
	})
	return m
}

func (m *WrapKVBuilder) Implement(i *InterfaceBuilder) *WrapKVBuilder {
	m.ImplementInterface = i
	return m
}

func (m *WrapKVBuilder) Build() error {
	file, err := os.Create(m.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return wrapKVTempl.Execute(file, m)
}
