package builder

import (
	"github.com/gobeam/stringy"
	"os"
	"text/template"
)

var wrapStructTempl *template.Template

func init() {
	var err error
	wrapStructTempl, err = template.ParseFiles("builder/wrap_struct.gotempl")
	if err != nil {
		panic(err)
	}
}

type WrapStructBuilder struct {
	Name               string
	FileName           string
	ReceiverName       string
	JsonTag            string
	ImplementInterface *InterfaceBuilder
	InnerType          *StructBuilder
}

func Wrap(name string, inner *StructBuilder) *WrapStructBuilder {
	return &WrapStructBuilder{
		Name:         name,
		FileName:     "../" + stringy.New(name).SnakeCase().ToLower() + ".go",
		ReceiverName: stringy.New(stringy.New(name).CamelCase()).LcFirst(),
		JsonTag:      stringy.New(name).SnakeCase().ToLower(),
		InnerType:    inner,
	}
}

func (m *WrapStructBuilder) Implement(i *InterfaceBuilder) *WrapStructBuilder {
	m.ImplementInterface = i
	return m
}

func (m *WrapStructBuilder) Build() error {
	file, err := os.Create(m.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return wrapStructTempl.Execute(file, m)
}
