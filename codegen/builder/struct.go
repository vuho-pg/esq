package builder

import (
	_ "embed"
	"github.com/gobeam/stringy"
	"os"
	"text/template"
)

type Flag int

func (f Flag) IsStruct() bool {
	return f&IsStruct != 0
}

func (f Flag) IsRequired() bool {
	return f&IsRequired != 0
}

func (f Flag) IsArr() bool {
	return f&IsArr != 0
}

func (f Flag) IsForcePtr() bool {
	return f&IsForcePtr != 0
}

const (
	IsStruct Flag = 1 << iota
	IsRequired
	IsArr
	IsForcePtr
)

var structTmp *template.Template

func init() {
	var err error
	if structTmp, err = template.ParseFiles("builder/struct.gotempl"); err != nil {
		panic(err)
	}
}

type StructBuilder struct {
	Name               string
	FileName           string
	ReceiverName       string
	Fields             []*FieldBuilder
	ExtendBuilder      []string
	ImplementInterface *InterfaceBuilder
}

func Struct(name string, fields ...*FieldBuilder) *StructBuilder {
	return &StructBuilder{
		Name:         name,
		FileName:     "../" + stringy.New(name).SnakeCase().ToLower() + ".gen.go",
		ReceiverName: "_" + stringy.New(stringy.New(name).CamelCase()).LcFirst(),
		Fields:       fields,
	}
}

func (s *StructBuilder) Implement(i *InterfaceBuilder) *StructBuilder {
	s.ImplementInterface = i
	return s
}

func (s *StructBuilder) Extend(structs ...string) *StructBuilder {
	s.ExtendBuilder = append(s.ExtendBuilder, structs...)
	return s
}

func (s *StructBuilder) ParamName(name string) string {
	return "_" + stringy.New(stringy.New(name).CamelCase()).LcFirst()
}

func (s *StructBuilder) Build() error {
	file, err := os.Create(s.FileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return structTmp.Execute(file, s)
}
