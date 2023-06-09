package builder

import (
	"github.com/gobeam/stringy"
	"os"
	"text/template"
)

var mapSingleTempl *template.Template

func init() {
	var err error
	mapSingleTempl, err = template.ParseFiles("builder/map_single.gotempl")
	if err != nil {
		panic(err)
	}
}

type MapBuilder struct {
	Name               string
	FileName           string
	ReceiverName       string
	JsonTag            string
	ImplementInterface *InterfaceBuilder
	InnerType          *StructBuilder
}

func Map(name string, inner *StructBuilder) *MapBuilder {
	return &MapBuilder{
		Name:         name,
		FileName:     "../" + stringy.New(name).SnakeCase().ToLower() + ".gen.go",
		ReceiverName: "_" + stringy.New(stringy.New(name).CamelCase()).LcFirst(),
		JsonTag:      stringy.New(name).SnakeCase().ToLower(),
		InnerType:    inner,
	}
}

func (m *MapBuilder) Implement(i *InterfaceBuilder) *MapBuilder {
	m.ImplementInterface = i
	return m
}

func (m *MapBuilder) Build() error {
	file, err := os.Create(m.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return mapSingleTempl.Execute(file, m)
}
