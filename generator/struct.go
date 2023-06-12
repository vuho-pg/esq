package main

import (
	"github.com/gobeam/stringy"
	"os"
	"strings"
	"text/template"
)

var typeTempl *template.Template

func init() {
	var err error
	typeTempl, err = template.ParseFiles("struct.gotempl")
	if err != nil {
		panic(err)
	}
}

type FieldGen struct {
	Name         string
	ParamName    string
	SingularType string
	JsonTag      string
	Required     bool
	IsArr        bool
	HasDefault   bool
	DefaultValue string
	Type         string
}

func Field(name string, t string) *FieldGen {
	f := &FieldGen{
		Name:         name,
		SingularType: strings.TrimPrefix(t, "[]"),
		ParamName:    "_" + stringy.New(stringy.New(name).CamelCase()).LcFirst(),
		JsonTag:      stringy.New(name).SnakeCase().ToLower(),
		Type:         t,
	}
	f.IsArr = strings.HasPrefix(f.Type, "[]")

	return f
}

func (f *FieldGen) Default(value string) *FieldGen {
	f.HasDefault = true
	f.DefaultValue = value
	return f
}

func (f *FieldGen) Require() *FieldGen {
	f.Required = true
	f.JsonTag = stringy.New(f.Name).SnakeCase().ToLower()
	return f
}

type NormalStructGen struct {
	Name         string
	ReceiverName string
	FileName     string
	Fields       []*FieldGen
	Methods      []string
}

func Type(name string, fields ...*FieldGen) *NormalStructGen {
	return &NormalStructGen{
		Name:         name,
		FileName:     "../" + stringy.New(name).SnakeCase().ToLower() + ".go",
		Fields:       fields,
		ReceiverName: strings.ToLower(name[:1]),
	}
}

func (t *NormalStructGen) Method(methods ...string) *NormalStructGen {
	t.Methods = append(t.Methods, methods...)
	return t
}

func (t *NormalStructGen) Build() error {
	file, err := os.Create(t.FileName)
	if err != nil {
		return err
	}
	return typeTempl.Execute(file, t)
}

var typeGenerators = make([]Builder, 0)

func CreateType(gen ...Builder) {
	typeGenerators = append(typeGenerators, gen...)
}
