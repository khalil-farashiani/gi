// Code generated by "gtigen -add-types -add-methods"; DO NOT EDIT.

package main

import (
	"goki.dev/goki/gti"
	"goki.dev/goki/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:       "main.Config",
	ShortName:  "main.Config",
	IDName:     "config",
	Doc:        "",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Name", &gti.Field{Name: "Name", Type: "string", Doc: "the name of the user", Directives: gti.Directives{}, Tag: ""}},
		{"Age", &gti.Field{Name: "Age", Type: "int", Doc: "the age of the user", Directives: gti.Directives{}, Tag: ""}},
		{"LikesGo", &gti.Field{Name: "LikesGo", Type: "bool", Doc: "whether the user likes Go", Directives: gti.Directives{}, Tag: ""}},
		{"BuildTarget", &gti.Field{Name: "BuildTarget", Type: "string", Doc: "the target platform to build for", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
