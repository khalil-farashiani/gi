// Code generated by "goki generate ./..."; DO NOT EDIT.

package keyfun

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/gi/v2/keyfun.MapsItem",
	ShortName: "keyfun.MapsItem",
	IDName:    "maps-item",
	Doc:       "MapsItem is an entry in a Maps list",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{"-setters"}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Name", &gti.Field{Name: "Name", Type: "string", LocalType: "string", Doc: "name of keymap", Directives: gti.Directives{}, Tag: "width:\"20\""}},
		{"Desc", &gti.Field{Name: "Desc", Type: "string", LocalType: "string", Doc: "description of keymap -- good idea to include source it was derived from", Directives: gti.Directives{}, Tag: ""}},
		{"Map", &gti.Field{Name: "Map", Type: "goki.dev/gi/v2/keyfun.Map", LocalType: "Map", Doc: "to edit key sequence click button and type new key combination; to edit function mapped to key sequence choose from menu", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

// SetName sets the [MapsItem.Name]:
// name of keymap
func (t *MapsItem) SetName(v string) *MapsItem {
	t.Name = v
	return t
}

// SetDesc sets the [MapsItem.Desc]:
// description of keymap -- good idea to include source it was derived from
func (t *MapsItem) SetDesc(v string) *MapsItem {
	t.Desc = v
	return t
}

// SetMap sets the [MapsItem.Map]:
// to edit key sequence click button and type new key combination; to edit function mapped to key sequence choose from menu
func (t *MapsItem) SetMap(v Map) *MapsItem {
	t.Map = v
	return t
}

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/gi/v2/keyfun.Maps",
	ShortName: "keyfun.Maps",
	IDName:    "maps",
	Doc:       "Maps is a list of KeyMap's -- users can edit these in Prefs -- to create\na custom one, just duplicate an existing map, rename, and customize",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{
		{"Open", &gti.Method{Name: "Open", Doc: "Open opens keymaps from a json-formatted file.\nYou can save and open key maps to / from files to share, experiment, transfer, etc", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"filename", &gti.Field{Name: "filename", Type: "string", LocalType: "string", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		})}},
		{"Save", &gti.Method{Name: "Save", Doc: "Save saves keymaps to a json-formatted file.\nYou can save and open key maps to / from files to share, experiment, transfer, etc", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"filename", &gti.Field{Name: "filename", Type: "string", LocalType: "string", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		})}},
		{"OpenPrefs", &gti.Method{Name: "OpenPrefs", Doc: "OpenPrefs opens KeyMaps from GoGi standard prefs directory, in file key_maps_prefs.json.\nThis is called automatically, so calling it manually should not be necessary in most cases.", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		})}},
		{"SavePrefs", &gti.Method{Name: "SavePrefs", Doc: "SavePrefs saves KeyMaps to GoGi standard prefs directory, in file key_maps_prefs.json,\nwhich will be loaded automatically at startup if prefs SaveKeyMaps is checked\n(should be if you're using custom keymaps)", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		})}},
		{"RevertToStd", &gti.Method{Name: "RevertToStd", Doc: "RevertToStd reverts the keymaps to using the StdKeyMaps that are compiled into the program\nand have all the lastest key functions defined.  If you have edited your maps, and are finding\nthings not working, it is a good idea to save your current maps and try this, or at least do\nViewStdMaps to see the current standards. Your current map edits will be lost if you proceed!", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"MarkdownDoc", &gti.Method{Name: "MarkdownDoc", Doc: "MarkdownDoc generates a markdown table of all the key mappings", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"string", &gti.Field{Name: "string", Type: "string", LocalType: "string", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		})}},
	}),
})
