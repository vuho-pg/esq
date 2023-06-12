package main

import (
	"github.com/gobeam/stringy"
	"os"
	"strings"
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
	FileName     string
	Name         string
	KeyType      string
	ValueType    string
	ReceiverName string
	Methods      []string
}

func MapType(name string, keyType string, valueType string) *MapTypeGen {
	return &MapTypeGen{
		Name:         name,
		FileName:     "../" + stringy.New(name).SnakeCase().ToLower() + ".go",
		KeyType:      keyType,
		ValueType:    valueType,
		ReceiverName: strings.ToLower(name[:1]),
	}
}

func (m *MapTypeGen) Method(methods ...string) *MapTypeGen {
	m.Methods = append(m.Methods, methods...)
	return m
}

func (m *MapTypeGen) Build() error {
	file, err := os.Create(m.FileName)
	if err != nil {
		return err
	}
	return mapTypeTempl.Execute(file, m)
}
