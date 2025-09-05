package builder

import (
	"fmt"
	"os"
	"text/template"

	"github.com/gobeam/stringy"
)

var enumTemplate *template.Template

func init() {
	var err error
	enumTemplate, err = template.ParseFiles("builder/enum.gotempl")
	if err != nil {
		panic(err)
	}
}

type EnumBuilder struct {
	Name     string
	FileName string
	Type     string
	Values   []string
}

// Enum generates enum type
func Enum(name string, t string) *EnumBuilder {
	return &EnumBuilder{
		Name:     name,
		FileName: "../" + stringy.New(name).SnakeCase().ToLower() + ".gen.go",
		Type:     t,
		Values:   nil,
	}
}

// Value adds enum values
func (m *EnumBuilder) Value(values ...string) *EnumBuilder {
	m.Values = append(m.Values, values...)
	return m
}

func (m *EnumBuilder) BuildName(value string) string {
	return fmt.Sprintf("%s%s", m.Name, stringy.New(value).CamelCase())
}

func (m *EnumBuilder) BuildValue(value string) string {
	switch m.Type {
	case "string":
		return `"` + value + `"`
	default:
		return value
	}
}

func (m *EnumBuilder) Build() error {
	file, err := os.Create(m.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return enumTemplate.Execute(file, m)
}
