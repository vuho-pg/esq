package main

import (
	"github.com/gobeam/stringy"
	"os"
	"text/template"
)

var mapTypeTempl *template.Template

func init() {
	var err error
	mapTypeTempl, err = template.ParseFiles("map_struct.gotempl")
	if err != nil {
		panic(err)
	}
}

type MapTypeGen struct {
	FileName  string
	Name      string
	KeyType   string
	ValueType string
}

func MapType(name string, keyType string, valueType string) *MapTypeGen {
	return &MapTypeGen{
		Name:      name,
		FileName:  "../" + stringy.New(name).SnakeCase().ToLower() + ".go",
		KeyType:   keyType,
		ValueType: valueType,
	}
}

func (m *MapTypeGen) Build() error {
	file, err := os.Create(m.FileName)
	if err != nil {
		return err
	}
	return mapTypeTempl.Execute(file, m)
}
