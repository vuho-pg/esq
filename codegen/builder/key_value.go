package builder

import (
	"os"
	"text/template"

	"github.com/gobeam/stringy"
)

var kvTemplate *template.Template

func init() {
	var err error
	kvTemplate, err = template.ParseFiles("builder/key_value.gotempl")
	if err != nil {
		panic(err)
	}
}

type KeyValueField struct {
	Name      string
	ParamName string
	Key       string
	ValueType string
}

type KeyValueBuilder struct {
	ReceiverName string
	Name         string
	FileName     string
	KeyType      string
	ValueType    string
	IsSingle     bool
	Fields       []*KeyValueField
}

func KeyValue(name, keyType, valueType string) *KeyValueBuilder {
	return &KeyValueBuilder{
		Name:         name,
		KeyType:      keyType,
		FileName:     "../" + stringy.New(name).SnakeCase().ToLower() + ".gen.go",
		ReceiverName: "_" + stringy.New(stringy.New(name).CamelCase()).LcFirst(),
		ValueType:    valueType,
		IsSingle:     true,
	}
}

func (k *KeyValueBuilder) Field(name, key, valueType string) *KeyValueBuilder {
	if valueType == "" {
		valueType = k.ValueType
	}
	k.Fields = append(k.Fields, &KeyValueField{
		Name:      name,
		ParamName: "_" + stringy.New(stringy.New(name).CamelCase()).LcFirst(),
		Key:       key,
		ValueType: valueType,
	})

	return k
}

func (k *KeyValueBuilder) Multiple() *KeyValueBuilder {
	k.IsSingle = false
	return k
}

func (k *KeyValueBuilder) Build() error {
	file, err := os.Create(k.FileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return kvTemplate.Execute(file, k)
}
