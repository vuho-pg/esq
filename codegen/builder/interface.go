package builder

import (
	"os"
	"text/template"

	"github.com/gobeam/stringy"
)

var interfaceTempl *template.Template

func init() {
	var err error
	interfaceTempl, err = template.ParseFiles("builder/interface.gotempl")
	if err != nil {
		panic(err)
	}
}

type InterfaceBuilder struct {
	Name     string
	FileName string
	Method   []string
}

func Interface(name string) *InterfaceBuilder {
	return &InterfaceBuilder{
		Name:     name,
		FileName: "../" + stringy.New(name).SnakeCase().ToLower() + ".gen.go",
	}
}

func (m *InterfaceBuilder) Func(name ...string) *InterfaceBuilder {
	m.Method = append(m.Method, name...)
	return m
}

func (m *InterfaceBuilder) Build() error {
	file, err := os.Create(m.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return interfaceTempl.Execute(file, m)
}
