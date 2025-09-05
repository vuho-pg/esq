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

// Field generates a field in struct
func Field(name string) *FieldBuilder {
	return &FieldBuilder{
		Name:      name,
		JsonTag:   stringy.New(name).SnakeCase().ToLower(),
		ParamName: "_" + stringy.New(stringy.New(name).CamelCase()).LcFirst(),
	}
}

// String make the field type to string
func (f *FieldBuilder) String() *FieldBuilder {
	return f.Type("string")
}

// Float64 make the field type to float64
func (f *FieldBuilder) Float64() *FieldBuilder {
	return f.Type("float64")
}

// Bool make the field type to bool
func (f *FieldBuilder) Bool() *FieldBuilder {
	return f.Type("bool")
}

// Int make the field type to int
func (f *FieldBuilder) Int() *FieldBuilder {
	return f.Type("int")
}

// Any make the field type to any
func (f *FieldBuilder) Any() *FieldBuilder {
	return f.Type("any")
}

// Type sets field to custom type
func (f *FieldBuilder) Type(t string) *FieldBuilder {
	f.FieldType = t
	return f
}

// Enum sets field to enum type
func (f *FieldBuilder) Enum() *FieldBuilder {
	f.FieldType = f.Name
	return f
}

// Struct sets field to struct type
func (f *FieldBuilder) Struct(s *StructBuilder) *FieldBuilder {
	f.Flag |= IsStruct
	f.StructType = s
	return f
}

// Array marks the field type as array
func (f *FieldBuilder) Array() *FieldBuilder {
	f.Flag |= IsArr
	return f
}

// Required marks the field as required
func (f *FieldBuilder) Required() *FieldBuilder {
	f.Flag |= IsRequired
	return f
}

// ForcePtr forces the field to be a pointer type
func (f *FieldBuilder) ForcePtr() *FieldBuilder {
	f.Flag |= IsForcePtr
	return f
}
