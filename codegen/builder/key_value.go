package builder

import (
	"github.com/gobeam/stringy"
	"os"
	"text/template"
)

var kvTemplate *template.Template

func init() {
	var err error
	kvTemplate, err = template.ParseFiles("builder/key_value.gotempl")
	if err != nil {
		panic(err)
	}
}

type KeyValueBuilder struct {
	Name      string
	FileName  string
	KeyType   string
	ValueType string
	IsSingle  bool
}

func KeyValue(name, keyType, valueType string) *KeyValueBuilder {
	return &KeyValueBuilder{
		Name:      name,
		KeyType:   keyType,
		FileName:  "../" + stringy.New(name).SnakeCase().ToLower() + ".gen.go",
		ValueType: valueType,
		IsSingle:  true,
	}
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
