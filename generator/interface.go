package main

import (
	"github.com/gobeam/stringy"
	"os"
	"text/template"
)

var interfaceTypeTempl *template.Template

func init() {
	var err error
	interfaceTypeTempl, err = template.ParseFiles("interface.gotempl")
	if err != nil {
		panic(err)
	}
}

type InterfaceGen struct {
	Name     string
	FileName string
	Methods  []string
}

func Interface(name string) *InterfaceGen {
	return &InterfaceGen{
		Name:     name,
		FileName: "../" + stringy.New(name).SnakeCase().ToLower() + ".go",
		Methods:  nil,
	}
}

func (i *InterfaceGen) Method(method ...string) *InterfaceGen {
	i.Methods = append(i.Methods, method...)
	return i
}

func (m *InterfaceGen) Build() error {
	file, err := os.Create(m.FileName)
	if err != nil {
		return err
	}
	return interfaceTypeTempl.Execute(file, m)
}
