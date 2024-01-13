// Code generated by "gtigen generate -output gtigen_gen.go"; DO NOT EDIT.

package gtigen

import (
	"goki.dev/gti"
)

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/gti/gtigen.Config",
	ShortName: "gtigen.Config",
	IDName:    "config",
	Doc:       "Config contains the configuration information\nused by gtigen",
	Directives: gti.Directives{
		{Tool: "gti", Directive: "add"},
	},
	Fields:  []gti.Field{{Name: "Dir", Type: "string", LocalType: "string", Doc: "the source directory to run gtigen on (can be set to multiple through paths like ./...)"}, {Name: "Output", Type: "string", LocalType: "string", Doc: "the output file location relative to the package on which gtigen is being called"}, {Name: "AddTypes", Type: "bool", LocalType: "bool", Doc: "whether to add types to gtigen by default"}, {Name: "AddMethods", Type: "bool", LocalType: "bool", Doc: "whether to add methods to gtigen by default"}, {Name: "AddFuncs", Type: "bool", LocalType: "bool", Doc: "whether to add functions to gtigen by default"}, {Name: "InterfaceConfigs", Type: "*goki.dev/ordmap.Map", LocalType: "*ordmap.Map[string, *Config]", Doc: "An ordered map of configs keyed by fully-qualified interface type names; if a type implements the interface, the config will be applied to it.\nThe configs are applied in sequential ascending order, which means that\nthe last config overrides the other ones, so the most specific\ninterfaces should typically be put last.\nNote: the package gtigen is run on must explicitly reference this interface at some point for this to work; adding a simple\n`var _ MyInterface = (*MyType)(nil)` statement to check for interface implementation is an easy way to accomplish that.\nNote: gtigen will still succeed if it can not find one of the interfaces specified here in order to allow it to work generically across multiple directories; you can use the -v flag to get log warnings about this if you suspect that it is not finding interfaces when it should."}, {Name: "Instance", Type: "bool", LocalType: "bool", Doc: "whether to generate an instance of the type(s)"}, {Name: "TypeVar", Type: "bool", LocalType: "bool", Doc: "whether to generate a global type variable of the form 'TypeNameType'"}, {Name: "Setters", Type: "bool", LocalType: "bool", Doc: "Whether to generate chaining `Set*` methods for each field of each type (eg: \"SetText\" for field \"Text\").\nIf this is set to true, then you can add `set:\"-\"` struct tags to individual fields\nto prevent Set methods being generated for them."}, {Name: "Templates", Type: "[]*text/template.Template", LocalType: "[]*template.Template", Doc: "a slice of templates to execute on each type being added; the template data is of the type gtigen.Type"}},
	Embeds:  []gti.Field(nil),
	Methods: []gti.Method{},
})

var _ = gti.AddFunc(&gti.Func{
	Name: "goki.dev/gti/gtigen.Generate",
	Doc:  "Generate generates gti type info, using the\nconfiguration information, loading the packages from the\nconfiguration source directory, and writing the result\nto the configuration output file.\n\nIt is a simple entry point to gtigen that does all\nof the steps; for more specific functionality, create\na new [Generator] with [NewGenerator] and call methods on it.",
	Directives: gti.Directives{
		{Tool: "grease", Directive: "cmd", Args: []string{"-root"}},
		{Tool: "gti", Directive: "add"},
	},
	Args:    []gti.Field{{Name: "cfg", Type: "*goki.dev/gti/gtigen.Config", LocalType: "*Config"}},
	Returns: []gti.Field{{Name: "error", Type: "error", LocalType: "error"}},
})
