package builder

import "github.com/gobeam/stringy"

type FieldBuilder struct {
	Flag       Flag
	Name       string
	JsonTag    string
	FieldType  string
	ParamName  string
	StructType *StructBuilder
}

func Field(name string) *FieldBuilder {
	return &FieldBuilder{
		Name:      name,
		JsonTag:   stringy.New(name).SnakeCase().ToLower(),
		ParamName: "_" + stringy.New(stringy.New(name).CamelCase()).LcFirst(),
	}
}

func (f *FieldBuilder) String() *FieldBuilder {
	return f.Type("string")
}

func (f *FieldBuilder) Float64() *FieldBuilder {
	return f.Type("float64")
}

func (f *FieldBuilder) Bool() *FieldBuilder {
	return f.Type("bool")
}

func (f *FieldBuilder) Int() *FieldBuilder {
	return f.Type("int")
}

func (f *FieldBuilder) Type(t string) *FieldBuilder {
	f.FieldType = t
	return f
}

func (f *FieldBuilder) Struct(s *StructBuilder) *FieldBuilder {
	f.Flag |= IsStruct
	f.StructType = s
	return f
}

func (f *FieldBuilder) Array() *FieldBuilder {
	f.Flag |= IsArr
	return f
}

func (f *FieldBuilder) Required() *FieldBuilder {
	f.Flag |= IsRequired
	return f
}
