package main

import (
	"github.com/gobeam/stringy"
	"os"
	"strings"
	"text/template"
)

var wrapStructTempl *template.Template

func init() {
	var err error
	wrapStructTempl, err = template.ParseFiles("wrap_struct.gotempl")
	if err != nil {
		panic(err)
	}
}

type WrapStructGen struct {
	Name         string
	JsonTag      string
	ReceiverName string
	InnerType    *NormalStructGen
}

func WrapStruct(name string, inner *NormalStructGen) *WrapStructGen {
	return &WrapStructGen{
		Name:         name,
		JsonTag:      stringy.New(name).SnakeCase().ToLower(),
		ReceiverName: strings.ToLower(name[:1]),
		InnerType:    inner,
	}
}

func (w *WrapStructGen) Build() error {
	file, err := os.Create("../" + stringy.New(w.Name).SnakeCase().ToLower() + ".go")
	if err != nil {
		return err
	}
	defer file.Close()
	return wrapStructTempl.Execute(file, w)
}
