// Code generated by "gtigen.test -test.paniconexit0 -test.timeout=10m0s"; DO NOT EDIT.

package testdata

import (
	"goki.dev/gti"
)

// PersonType is the [gti.Type] for [Person]
var PersonType = gti.AddType(&gti.Type{
	Name:      "goki.dev/gti/gtigen/testdata.Person",
	ShortName: "testdata.Person",
	IDName:    "person",
	Doc:       "Person represents a person and their attributes.\nThe zero value of a Person is not valid.",
	Directives: gti.Directives{
		{Tool: "ki", Directive: "flagtype", Args: []string{"NodeFlags", "-field", "Flag"}},
		{Tool: "goki", Directive: "embedder"},
	},
	Fields: []gti.Field{{Name: "Name", Type: "string", LocalType: "string", Doc: "Name is the name of the person"}, {Name: "Age", Type: "int", LocalType: "int", Doc: "Age is the age of the person"}, {Name: "Type", Type: "*goki.dev/gti.Type", LocalType: "*gti.Type", Doc: "Type is the type of the person"}, {Name: "Nicknames", Type: "[]string", LocalType: "[]string", Doc: "Nicknames are the nicknames of the person"}},
	Embeds: []gti.Field{{Name: "RGBA", Type: "image/color.RGBA", LocalType: "color.RGBA"}},
	Methods: []gti.Method{{Name: "Introduction", Doc: "Introduction returns an introduction for the person.\nIt contains the name of the person and their age.", Directives: gti.Directives{
		{Tool: "gi", Directive: "toolbar", Args: []string{"-name", "ShowIntroduction", "-icon", "play", "-show-result", "-confirm"}},
		{Tool: "gti", Directive: "add"},
	}, Returns: []gti.Field{{Name: "string", Type: "string", LocalType: "string"}}}},
	Instance: &Person{},
})

func (t *Person) MyCustomFuncForStringers(a any) error {
	return nil
}

// SetName sets the [Person.Name]:
// Name is the name of the person
func (t *Person) SetName(v string) *Person {
	t.Name = v
	return t
}

// SetAge sets the [Person.Age]:
// Age is the age of the person
func (t *Person) SetAge(v int) *Person {
	t.Age = v
	return t
}

// SetType sets the [Person.Type]:
// Type is the type of the person
func (t *Person) SetType(v *gti.Type) *Person {
	t.Type = v
	return t
}

// SetNicknames sets the [Person.Nicknames]:
// Nicknames are the nicknames of the person
func (t *Person) SetNicknames(v ...string) *Person {
	t.Nicknames = v
	return t
}

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.Alert",
	Doc:        "Alert prints an alert with the given message",
	Directives: gti.Directives{},
	Args:       []gti.Field{{Name: "msg", Type: "string", LocalType: "string"}},
	Returns:    []gti.Field(nil),
})

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.TypeOmittedArgs0",
	Doc:        "",
	Directives: gti.Directives{},
	Args:       []gti.Field{{Name: "x", Type: "float32", LocalType: "float32"}, {Name: "y", Type: "float32", LocalType: "float32"}},
	Returns:    []gti.Field(nil),
})

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.TypeOmittedArgs1",
	Doc:        "",
	Directives: gti.Directives{},
	Args:       []gti.Field{{Name: "x", Type: "int", LocalType: "int"}, {Name: "y", Type: "struct{}", LocalType: "struct{}"}, {Name: "z", Type: "struct{}", LocalType: "struct{}"}},
	Returns:    []gti.Field(nil),
})

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.TypeOmittedArgs2",
	Doc:        "",
	Directives: gti.Directives{},
	Args:       []gti.Field{{Name: "x", Type: "int", LocalType: "int"}, {Name: "y", Type: "int", LocalType: "int"}, {Name: "z", Type: "int", LocalType: "int"}},
	Returns:    []gti.Field(nil),
})

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.TypeOmittedArgs3",
	Doc:        "",
	Directives: gti.Directives{},
	Args:       []gti.Field{{Name: "x", Type: "int", LocalType: "int"}, {Name: "y", Type: "bool", LocalType: "bool"}, {Name: "z", Type: "bool", LocalType: "bool"}, {Name: "w", Type: "float32", LocalType: "float32"}},
	Returns:    []gti.Field(nil),
})

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.TypeOmittedArgs4",
	Doc:        "",
	Directives: gti.Directives{},
	Args:       []gti.Field{{Name: "x", Type: "string", LocalType: "string"}, {Name: "y", Type: "string", LocalType: "string"}, {Name: "z", Type: "string", LocalType: "string"}, {Name: "w", Type: "bool", LocalType: "bool"}},
	Returns:    []gti.Field(nil),
})
