// Code generated by "goki generate ./..."; DO NOT EDIT.

package android

import (
	"goki.dev/goki/gti"
	"goki.dev/goki/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/goki/goosi/driver/android.App",
	ShortName: "android.App",
	IDName:    "app",
	Doc:       "App is the [goosi.App] implementation for the Android platform",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"GPU", &gti.Field{Name: "GPU", Type: "*goki.dev/goki/vgpu.GPU", LocalType: "*vgpu.GPU", Doc: "GPU is the system GPU used for the app", Directives: gti.Directives{}, Tag: ""}},
		{"Winptr", &gti.Field{Name: "Winptr", Type: "uintptr", LocalType: "uintptr", Doc: "Winptr is the pointer to the underlying system window", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"AppSingle", &gti.Field{Name: "AppSingle", Type: "goki.dev/goki/goosi/driver/base.AppSingle", LocalType: "base.AppSingle[*vdraw.Drawer, *Window]", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/goki/goosi/driver/android.Window",
	ShortName: "android.Window",
	IDName:    "window",
	Doc:       "Window is the implementation of [goosi.Window] for the Android platform.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"WindowSingle", &gti.Field{Name: "WindowSingle", Type: "goki.dev/goki/goosi/driver/base.WindowSingle", LocalType: "base.WindowSingle[*App]", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
