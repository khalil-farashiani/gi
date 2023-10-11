// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package giv

//go:generate goki generate

import (
	"fmt"
	"image/color"
	"log"
	"log/slog"
	"reflect"
	"strings"
	"time"

	"goki.dev/enums"
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/histyle"
	"goki.dev/girl/states"
	"goki.dev/girl/styles"
	"goki.dev/girl/units"
	"goki.dev/goosi/events"
	"goki.dev/goosi/events/key"
	"goki.dev/gti"
	"goki.dev/icons"
	"goki.dev/ki/v2"
	"goki.dev/laser"
)

func init() {
	gi.TheViewIFace = &ViewIFace{}
	ValueViewMapAdd(laser.LongTypeName(reflect.TypeOf(icons.Icon(""))), func() ValueView {
		vv := &IconValueView{}
		ki.InitNode(vv)
		return vv
	})
	ValueViewMapAdd(laser.LongTypeName(reflect.TypeOf(gi.FontName(""))), func() ValueView {
		vv := &FontValueView{}
		ki.InitNode(vv)
		return vv
	})
	ValueViewMapAdd(laser.LongTypeName(reflect.TypeOf(gi.FileName(""))), func() ValueView {
		vv := &FileValueView{}
		ki.InitNode(vv)
		return vv
	})
	ValueViewMapAdd(laser.LongTypeName(reflect.TypeOf(gi.KeyMapName(""))), func() ValueView {
		vv := &KeyMapValueView{}
		ki.InitNode(vv)
		return vv
	})
	ValueViewMapAdd(laser.LongTypeName(reflect.TypeOf(gi.ColorName(""))), func() ValueView {
		vv := &ColorNameValueView{}
		ki.InitNode(vv)
		return vv
	})
	ValueViewMapAdd(laser.LongTypeName(reflect.TypeOf(key.Chord(""))), func() ValueView {
		vv := &KeyChordValueView{}
		ki.InitNode(vv)
		return vv
	})
	ValueViewMapAdd(laser.LongTypeName(reflect.TypeOf(gi.HiStyleName(""))), func() ValueView {
		vv := &HiStyleValueView{}
		ki.InitNode(vv)
		return vv
	})
	ValueViewMapAdd(laser.LongTypeName(reflect.TypeOf(time.Time{})), func() ValueView {
		vv := &TimeValueView{}
		ki.InitNode(vv)
		return vv
	})
	ValueViewMapAdd(laser.LongTypeName(reflect.TypeOf(FileTime{})), func() ValueView {
		vv := &TimeValueView{}
		ki.InitNode(vv)
		return vv
	})
}

var (
	// MapInlineLen is the number of map elements at or below which an inline
	// representation of the map will be presented -- more convenient for small
	// #'s of props
	MapInlineLen = 3

	// StructInlineLen is the number of elemental struct fields at or below which an inline
	// representation of the struct will be presented -- more convenient for small structs
	StructInlineLen = 6

	// SliceInlineLen is the number of slice elements below which inline will be used
	SliceInlineLen = 6
)

////////////////////////////////////////////////////////////////////////////////////////
//  ValueViewer -- an interface for selecting ValueView GUI representation of types

// ValueViewer interface supplies the appropriate type of ValueView -- called
// on a given receiver item if defined for that receiver type (tries both
// pointer and non-pointer receivers) -- can use this for custom types to
// provide alternative custom interfaces -- must call Init on ValueView before
// returning it
type ValueViewer interface {
	ValueView() ValueView
}

// example implementation of ValueViewer interface -- can't implement on
// non-local types, so all the basic types are handled separately:
//
// func (s string) ValueView() ValueView {
// 	vv := &ValueViewBase{}
// 	ki.InitNode(vv)
// 	return vv
// }

// FieldValueViewer interface supplies the appropriate type of ValueView for a
// given field name and current field value on the receiver parent struct --
// called on a given receiver struct if defined for that receiver type (tries
// both pointer and non-pointer receivers) -- if a struct implements this
// interface, then it is used first for structs -- return nil to fall back on
// the default ToValueView result
type FieldValueViewer interface {
	FieldValueView(field string, fval any) ValueView
}

////////////////////////////////////////////////////////////////////////////////////////
//  ValueViewMap -- alternative way to connect value view with type

// ValueViewFunc is a function that returns a new initialized ValueView
// of an appropriate type as registered in the ValueViewMap
type ValueViewFunc func() ValueView

// The ValueViewMap is used to connect type names with corresponding ValueView
// representations of those types -- this can be used when it is not possible
// to use the ValueViewer interface (e.g., interface methods can only be
// defined within the package that defines the type -- so we need this for
// all types in gi which don't know about giv).
// You must use laser.LongTypeName (full package name + "." . type name) for
// the type name, as that is how it will be looked up.
var ValueViewMap map[string]ValueViewFunc

// ValueViewMapAdd adds a ValueViewFunc for a given type name.
// You must use laser.LongTypeName (full package name + "." . type name) for
// the type name, as that is how it will be looked up.
func ValueViewMapAdd(typeNm string, fun ValueViewFunc) {
	if ValueViewMap == nil {
		ValueViewMap = make(map[string]ValueViewFunc)
	}
	ValueViewMap[typeNm] = fun
}

// StructTagVal returns the value for given key in given struct tag string
// uses reflect.StructTag Lookup method -- just a wrapper for external
// use (e.g., in Python code)
func StructTagVal(key, tags string) string {
	stag := reflect.StructTag(tags)
	val, _ := stag.Lookup(key)
	return val
}

// ToValueView returns the appropriate ValueView for given item, based only on
// its type -- attempts to get the ValueViewer interface and failing that,
// falls back on default Kind-based options.  tags are optional tags, e.g.,
// from the field in a struct, that control the view properties -- see the gi wiki
// for details on supported tags -- these are NOT set for the view element, only
// used for options that affect what kind of view to create.
// See FieldToValueView for version that takes into account the properties of the owner.
// gopy:interface=handle
func ToValueView(it any, tags string) ValueView {
	if it == nil {
		vv := &ValueViewBase{}
		ki.InitNode(vv)
		return vv
	}
	if vv, ok := it.(ValueViewer); ok {
		vvo := vv.ValueView()
		if vvo != nil {
			return vvo
		}
	}
	// try pointer version..
	if vv, ok := laser.PtrInterface(it).(ValueViewer); ok {
		vvo := vv.ValueView()
		if vvo != nil {
			return vvo
		}
	}

	if _, ok := it.(enums.BitFlag); ok {
		vv := &BitFlagView{}
		ki.InitNode(vv)
		return vv
	}
	if _, ok := it.(enums.Enum); ok {
		vv := &EnumValueView{}
		ki.InitNode(vv)
		return vv
	}

	typ := reflect.TypeOf(it)
	nptyp := laser.NonPtrType(typ)
	vk := typ.Kind()
	// fmt.Printf("vv val %v: typ: %v nptyp: %v kind: %v\n", it, typ.String(), nptyp.String(), vk)

	nptypnm := laser.LongTypeName(nptyp)
	if vvf, has := ValueViewMap[nptypnm]; has {
		vv := vvf()
		return vv
	}

	forceInline := false
	forceNoInline := false

	/*
		tprops := kit.Types.Properties(typ, false) // don't make
		if tprops != nil {
			if inprop, ok := kit.TypeProp(*tprops, "inline"); ok {
				forceInline, ok = kit.ToBool(inprop)
			}
			if inprop, ok := kit.TypeProp(*tprops, "no-inline"); ok {
				forceNoInline, ok = kit.ToBool(inprop)
			}
		}
	*/

	if tags != "" {
		stag := reflect.StructTag(tags)
		if vwtag, ok := stag.Lookup("view"); ok {
			switch vwtag {
			case "inline":
				forceInline = true
			case "no-inline":
				forceNoInline = true
			}
		}
	}

	switch {
	case vk >= reflect.Int && vk <= reflect.Uint64:
		if _, ok := it.(fmt.Stringer); ok { // use stringer
			vv := &ValueViewBase{}
			ki.InitNode(vv)
			return vv
		} else {
			vv := &IntValueView{}
			ki.InitNode(vv)
			return vv
		}
	case vk == reflect.Bool:
		vv := &BoolValueView{}
		ki.InitNode(vv)
		return vv
	case vk >= reflect.Float32 && vk <= reflect.Float64:
		vv := &FloatValueView{} // handles step, min / max etc
		ki.InitNode(vv)
		return vv
	case vk >= reflect.Complex64 && vk <= reflect.Complex128:
		// todo: special edit with 2 fields..
		vv := &ValueViewBase{}
		ki.InitNode(vv)
		return vv
	case vk == reflect.Ptr:
		if ki.IsKi(nptyp) {
			vv := &KiPtrValueView{}
			ki.InitNode(vv)
			return vv
		}
		if laser.AnyIsNil(it) {
			vv := &NilValueView{}
			ki.InitNode(vv)
			return vv
		}
		v := reflect.ValueOf(it)
		if !laser.ValueIsZero(v) {
			// note: interfaces go here:
			// fmt.Printf("vv indirecting on pointer: %v type: %v\n", it, nptyp.String())
			return ToValueView(v.Elem().Interface(), tags)
		}
	case vk == reflect.Array:
		fallthrough
	case vk == reflect.Slice:
		v := reflect.ValueOf(it)
		sz := v.Len()
		eltyp := laser.SliceElType(it)
		if _, ok := it.([]byte); ok {
			vv := &ByteSliceValueView{}
			ki.InitNode(vv)
			return vv
		}
		if _, ok := it.([]rune); ok {
			vv := &RuneSliceValueView{}
			ki.InitNode(vv)
			return vv
		}
		isstru := (laser.NonPtrType(eltyp).Kind() == reflect.Struct)
		if !forceNoInline && (forceInline || (!isstru && sz <= SliceInlineLen && !ki.IsKi(eltyp))) {
			vv := &SliceInlineValueView{}
			ki.InitNode(vv)
			return vv
		} else {
			vv := &SliceValueView{}
			ki.InitNode(vv)
			return vv
		}
	case vk == reflect.Map:
		v := reflect.ValueOf(it)
		sz := v.Len()
		sz = laser.MapStructElsN(it)
		if !forceNoInline && (forceInline || sz <= MapInlineLen) {
			vv := &MapInlineValueView{}
			ki.InitNode(vv)
			return vv
		} else {
			vv := &MapValueView{}
			ki.InitNode(vv)
			return vv
		}
	case vk == reflect.Struct:
		// note: we need to handle these here b/c cannot define new methods for gi types
		if nptyp == laser.TypeFor[color.RGBA]() {
			vv := &ColorValueView{}
			ki.InitNode(vv)
			return vv
		}
		nfld := laser.AllFieldsN(nptyp)
		if nfld > 0 && !forceNoInline && (forceInline || nfld <= StructInlineLen) {
			vv := &StructInlineValueView{}
			ki.InitNode(vv)
			return vv
		} else {
			vv := &StructValueView{}
			ki.InitNode(vv)
			return vv
		}
	case vk == reflect.Interface:
		// note: we never get here -- all interfaces are captured by pointer kind above
		// apparently (because the non-ptr vk indirection does that I guess?)
		fmt.Printf("interface kind: %v %v %v\n", nptyp, nptyp.Name(), nptyp.String())
		switch {
		case nptyp == laser.TypeFor[reflect.Type]():
			vv := &TypeValueView{}
			ki.InitNode(vv)
			return vv
		}
	}
	// fallback.
	vv := &ValueViewBase{}
	ki.InitNode(vv)
	return vv
}

// FieldToValueView returns the appropriate ValueView for given field on a
// struct -- attempts to get the FieldValueViewer interface, and falls back on
// ToValueView otherwise, using field value (fval)
// gopy:interface=handle
func FieldToValueView(it any, field string, fval any) ValueView {
	if it == nil || field == "" {
		return ToValueView(fval, "")
	}
	if vv, ok := it.(FieldValueViewer); ok {
		vvo := vv.FieldValueView(field, fval)
		if vvo != nil {
			return vvo
		}
	}
	// try pointer version..
	if vv, ok := laser.PtrInterface(it).(FieldValueViewer); ok {
		vvo := vv.FieldValueView(field, fval)
		if vvo != nil {
			return vvo
		}
	}

	typ := reflect.TypeOf(it)
	nptyp := laser.NonPtrType(typ)

	/*
		if pv, has := kit.Types.Prop(nptyp, "EnumType:"+field); has {
			et := pv.(reflect.Type)
			if kit.Enums.IsBitFlag(et) {
				vv := &BitFlagView{}
				vv.AltType = et
				ki.InitNode(vv)
				return vv
			} else {
				vv := &EnumValueView{}
				vv.AltType = et
				ki.InitNode(vv)
				return vv
			}
		}
	*/

	ftyp, ok := nptyp.FieldByName(field)
	if ok {
		return ToValueView(fval, string(ftyp.Tag))
	}
	return ToValueView(fval, "")
}

// ValueView is an interface for managing the GUI representation of values
// (e.g., fields, map values, slice values) in Views (StructView, MapView,
// etc).  The different types of ValueView are for different Kinds of values
// (bool, float, etc) -- which can have different Kinds of owners.  The
// ValueViewBase class supports all the basic fields for managing the owner
// kinds.
type ValueView interface {
	ki.Ki

	// AsValueViewBase gives access to the basic data fields so that the
	// interface doesn't need to provide accessors for them.
	AsValueViewBase() *ValueViewBase

	// SetStructValue sets the value, owner and field information for a struct field.
	SetStructValue(val reflect.Value, owner any, field *reflect.StructField, tmpSave ValueView, viewPath string)

	// SetMapKey sets the key value and owner for a map key.
	SetMapKey(val reflect.Value, owner any, tmpSave ValueView)

	// SetMapValue sets the value, owner and map key information for a map
	// element -- needs pointer to ValueView representation of key to track
	// current key value.
	SetMapValue(val reflect.Value, owner any, key any, keyView ValueView, tmpSave ValueView, viewPath string)

	// SetSliceValue sets the value, owner and index information for a slice element.
	SetSliceValue(val reflect.Value, owner any, idx int, tmpSave ValueView, viewPath string)

	// SetSoloValue sets the value for a singleton standalone value
	// (e.g., for arg values).
	SetSoloValue(val reflect.Value)

	// OwnerKind returns the reflect.Kind of the owner: Struct, Map, or Slice
	// (or Invalid for standalone values such as args).
	OwnerKind() reflect.Kind

	// IsInactive returns whether the value is inactive -- e.g., Map owners
	// have Inactive values, and some fields can be marked as Inactive using a
	// struct tag.
	IsInactive() bool

	// WidgetType returns an appropriate type of widget to represent the
	// current value.
	WidgetType() *gti.Type

	// UpdateWidget updates the widget representation to reflect the current
	// value.  Must first check for a nil widget -- can be called in a
	// no-widget context (e.g., for single-argument values with actions).
	UpdateWidget()

	// ConfigWidget configures a widget of WidgetType for representing the
	// value, including setting up the signal connections to set the value
	// when the user edits it (values are always set immediately when the
	// widget is updated).
	ConfigWidget(widg gi.Widget)

	// HasAction returns true if this value has an associated action, such as
	// pulling up a dialog or chooser for this value.  Activate method will
	// trigger this action.
	HasDialog() bool

	// OpenDialog opens a dialog or chooser for this value.
	// This is called by default for single-argument methods that have value
	// representations with actions.  The ctx Widget provides a context
	// for opening the dialog, and function is called with the the relevant dialog,
	// so that the caller can execute its own actions based on the user
	// hitting Ok or Cancel.
	OpenDialog(ctx gi.Widget, fun func(dlg *gi.Dialog))

	// Val returns the reflect.Value representation for this item.
	Val() reflect.Value

	// SetValue assigns given value to this item (if not Inactive), using
	// Ki.SetField for Ki types and laser.SetRobust otherwise -- emits a ViewSig
	// signal when set.
	SetValue(val any) bool

	// SendChange sends events.Change event to all listeners registered on this view.
	// This is the primary notification event for all ValueView elements.
	SendChange()

	// OnChange registers given listener function for Change events on ValueView.
	// This is the primary notification event for all ValueView elements.
	OnChange(fun func(e events.Event))

	// SetTags sets tags for this valueview, for non-struct values, to
	// influence interface for this value -- see
	// https://goki.dev/gi/v2/wiki/Tags for valid options.  Adds to
	// existing tags if some are already set.
	SetTags(tags map[string]string)

	// SetTag sets given tag to given value for this valueview, for non-struct
	// values, to influence interface for this value -- see
	// https://goki.dev/gi/v2/wiki/Tags for valid options.
	SetTag(tag, value string)

	// Tag returns value for given tag -- looks first at tags set by
	// SetTag(s) methods, and then at field tags if this is a field in a
	// struct -- returns false if tag was not set.
	Tag(tag string) (string, bool)

	// AllTags returns all the tags for this value view, from structfield or set
	// specifically using SetTag* methods
	AllTags() map[string]string

	// SaveTmp saves a temporary copy of a struct to a map -- map values must
	// be explicitly re-saved and cannot be directly written to by the value
	// elements -- each ValueView has a pointer to any parent ValueView that
	// might need to be saved after SetValue -- SaveTmp called automatically
	// in SetValue but other cases that use something different need to call
	// it explicitly.
	SaveTmp()
}

// note: could have a more efficient way to represent the different owner type
// data (Key vs. Field vs. Idx), instead of just having everything for
// everything.  However, ValueView itself gets customized for different target
// value types, and those are orthogonal to the owner type, so need a separate
// ValueViewOwner class that encodes these options more efficiently -- but
// that introduces another struct alloc and pointer -- not clear if it is
// worth it..

// ValueViewBase provides the basis for implementations of the ValueView
// interface, representing values in the interface -- it implements a generic
// TextField representation of the string value, and provides the generic
// fallback for everything that doesn't provide a specific ValueViewer type.
type ValueViewBase struct {
	ki.Node

	// the reflect.Value representation of the value
	Value reflect.Value `desc:"the reflect.Value representation of the value"`

	// kind of owner that we have -- reflect.Struct, .Map, .Slice are supported
	OwnKind reflect.Kind `desc:"kind of owner that we have -- reflect.Struct, .Map, .Slice are supported"`

	// for OwnKind = Map, this value represents the Key -- otherwise the Value
	IsMapKey bool `desc:"for OwnKind = Map, this value represents the Key -- otherwise the Value"`

	// a record of parent View names that have led up to this view -- displayed as extra contextual information in view dialog windows
	ViewPath string `desc:"a record of parent View names that have led up to this view -- displayed as extra contextual information in view dialog windows"`

	// the object that owns this value, either a struct, slice, or map, if non-nil -- if a Ki Node, then SetField is used to set value, to provide proper updating
	Owner any `desc:"the object that owns this value, either a struct, slice, or map, if non-nil -- if a Ki Node, then SetField is used to set value, to provide proper updating"`

	// if Owner is a struct, this is the reflect.StructField associated with the value
	Field *reflect.StructField `desc:"if Owner is a struct, this is the reflect.StructField associated with the value"`

	// set of tags that can be set to customize interface for different types of values -- only source for non-structfield values
	Tags map[string]string `desc:"set of tags that can be set to customize interface for different types of values -- only source for non-structfield values"`

	// if Owner is a map, and this is a value, this is the key for this value in the map
	Key any `desc:"if Owner is a map, and this is a value, this is the key for this value in the map"`

	// if Owner is a map, and this is a value, this is the value view representing the key -- its value has the *current* value of the key, which can be edited
	KeyView ValueView `desc:"if Owner is a map, and this is a value, this is the value view representing the key -- its value has the *current* value of the key, which can be edited"`

	// if Owner is a slice, this is the index for the value in the slice
	Idx int `desc:"if Owner is a slice, this is the index for the value in the slice"`

	// type of widget to create -- cached during WidgetType method -- chosen based on the ValueView type and reflect.Value type -- see ValueViewer interface
	WidgetTyp *gti.Type `desc:"type of widget to create -- cached during WidgetType method -- chosen based on the ValueView type and reflect.Value type -- see ValueViewer interface"`

	// the widget used to display and edit the value in the interface -- this is created for us externally and we cache it during ConfigWidget
	Widget gi.Widget `desc:"the widget used to display and edit the value in the interface -- this is created for us externally and we cache it during ConfigWidget"`

	// Listeners are event listener functions for processing events on this widget.
	// type specific Listeners are added in OnInit when the widget is initialized.
	Listeners events.Listeners

	// value view that needs to have SaveTmp called on it whenever a change is made to one of the underlying values -- pass this down to any sub-views created from a parent
	TmpSave ValueView `desc:"value view that needs to have SaveTmp called on it whenever a change is made to one of the underlying values -- pass this down to any sub-views created from a parent"`
}

func (vv *ValueViewBase) AsValueViewBase() *ValueViewBase {
	return vv
}

func (vv *ValueViewBase) SetStructValue(val reflect.Value, owner any, field *reflect.StructField, tmpSave ValueView, viewPath string) {
	vv.OwnKind = reflect.Struct
	vv.Value = val
	vv.Owner = owner
	vv.Field = field
	vv.TmpSave = tmpSave
	vv.ViewPath = viewPath + "." + field.Name
	vv.SetName(field.Name)
}

func (vv *ValueViewBase) SetMapKey(key reflect.Value, owner any, tmpSave ValueView) {
	vv.OwnKind = reflect.Map
	vv.IsMapKey = true
	vv.Value = key
	vv.Owner = owner
	vv.TmpSave = tmpSave
	vv.SetName(laser.ToString(key.Interface()))
}

func (vv *ValueViewBase) SetMapValue(val reflect.Value, owner any, key any, keyView ValueView, tmpSave ValueView, viewPath string) {
	vv.OwnKind = reflect.Map
	vv.Value = val
	vv.Owner = owner
	vv.Key = key
	vv.KeyView = keyView
	vv.TmpSave = tmpSave
	keystr := laser.ToString(key)
	vv.ViewPath = viewPath + "." + keystr
	vv.SetName(keystr)
}

func (vv *ValueViewBase) SetSliceValue(val reflect.Value, owner any, idx int, tmpSave ValueView, viewPath string) {
	vv.OwnKind = reflect.Slice
	vv.Value = val
	vv.Owner = owner
	vv.Idx = idx
	vv.TmpSave = tmpSave
	idxstr := fmt.Sprintf("%v", idx)
	vpath := "[" + idxstr + "]"
	if vv.Owner != nil {
		if lblr, ok := vv.Owner.(gi.SliceLabeler); ok {
			slbl := lblr.ElemLabel(idx)
			if slbl != "" {
				vpath = "[" + slbl + "]"
			}
		}
	}
	vv.ViewPath = viewPath + vpath
	vv.SetName(idxstr)
}

// SetSoloValue sets the value for a singleton standalone value
// (e.g., for arg values).
func (vv *ValueViewBase) SetSoloValue(val reflect.Value) {
	vv.OwnKind = reflect.Invalid
	vv.Value = val
}

// SetSoloValueIface sets the value for a singleton standalone value
// using an interface for the value -- you must pass a pointer.
// for now, this cannot be a method because gopy doesn't find the
// key comment below that tells it what to do with the interface
// gopy:interface=handle
func SetSoloValueIface(vv *ValueViewBase, val any) {
	vv.OwnKind = reflect.Invalid
	vv.Value = reflect.ValueOf(val)
}

// we have this one accessor b/c it is more useful for outside consumers vs. internal usage
func (vv *ValueViewBase) OwnerKind() reflect.Kind {
	return vv.OwnKind
}

func (vv *ValueViewBase) IsInactive() bool {
	if vv.OwnKind == reflect.Struct {
		if _, ok := vv.Tag("inactive"); ok {
			return true
		}
	}
	npv := laser.NonPtrValue(vv.Value)
	if npv.Kind() == reflect.Interface && laser.ValueIsZero(npv) {
		return true
	}
	return false
}

func (vv *ValueViewBase) HasDialog() bool {
	return false
}

func (vv *ValueViewBase) OpenDialog(ctx gi.Widget, fun func(dlg *gi.Dialog)) {
}

func (vv *ValueViewBase) Val() reflect.Value {
	return vv.Value
}

func (vv *ValueViewBase) SetValue(val any) bool {
	if vv.This().(ValueView).IsInactive() {
		return false
	}
	var err error
	wasSet := false
	if vv.Owner != nil {
		switch vv.OwnKind {
		case reflect.Struct:
			err = laser.SetRobust(laser.PtrValue(vv.Value).Interface(), val)
			wasSet = true
		case reflect.Map:
			wasSet, err = vv.SetValueMap(val)
		case reflect.Slice:
			err = laser.SetRobust(laser.PtrValue(vv.Value).Interface(), val)
		}
		if updtr, ok := vv.Owner.(gi.Updater); ok {
			// fmt.Printf("updating: %v\n", updtr)
			updtr.Update()
		}
	} else {
		err = laser.SetRobust(laser.PtrValue(vv.Value).Interface(), val)
		wasSet = true
	}
	if wasSet {
		vv.This().(ValueView).SaveTmp()
	}
	// fmt.Printf("value view: %T sending for setting val %v\n", vv.This(), val)
	vv.SendChange()
	if err != nil {
		// todo: snackbar for error?
		slog.Error(err.Error())
	}
	return wasSet
}

func (vv *ValueViewBase) SetValueMap(val any) (bool, error) {
	ov := laser.NonPtrValue(reflect.ValueOf(vv.Owner))
	wasSet := false
	var err error
	if vv.IsMapKey {
		nv := laser.NonPtrValue(reflect.ValueOf(val)) // new key value
		kv := laser.NonPtrValue(vv.Value)
		cv := ov.MapIndex(kv)    // get current value
		curnv := ov.MapIndex(nv) // see if new value there already
		if val != kv.Interface() && !laser.ValueIsZero(curnv) {
			// actually new key and current exists
			gi.ChoiceDialog(vv.Widget, gi.DlgOpts{Title: "Map Key Conflict", Prompt: fmt.Sprintf("The map key value: %v already exists in the map -- are you sure you want to overwrite the current value?", val)},
				[]string{"Cancel Change", "Overwrite"}, func(dlg *gi.Dialog) {
					switch dlg.Data.(int) {
					case 1:
						cv := ov.MapIndex(kv)               // get current value
						ov.SetMapIndex(kv, reflect.Value{}) // delete old key
						ov.SetMapIndex(nv, cv)              // set new key to current value
						vv.Value = nv                       // update value to new key
						vv.This().(ValueView).SaveTmp()
						vv.SendChange()
					}
				})
			return false, nil // abort this action right now
		}
		ov.SetMapIndex(kv, reflect.Value{}) // delete old key
		ov.SetMapIndex(nv, cv)              // set new key to current value
		vv.Value = nv                       // update value to new key
		wasSet = true
	} else {
		vv.Value = laser.NonPtrValue(reflect.ValueOf(val))
		if vv.KeyView != nil {
			ck := laser.NonPtrValue(vv.KeyView.Val())                 // current key value
			wasSet = laser.SetMapRobust(ov, ck, reflect.ValueOf(val)) // todo: error
		} else { // static, key not editable?
			wasSet = laser.SetMapRobust(ov, laser.NonPtrValue(reflect.ValueOf(vv.Key)), vv.Value) // todo: error
		}
		// wasSet = true
	}
	return wasSet, err
}

// OnChange registers given listener function for Change events on ValueView.
// This is the primary notification event for all ValueView elements.
func (vv *ValueViewBase) OnChange(fun func(e events.Event)) {
	vv.On(events.Change, fun)
}

// On adds an event listener function for the given event type
func (vv *ValueViewBase) On(etype events.Types, fun func(e events.Event)) {
	vv.Listeners.Add(etype, fun)
}

// SendChange sends events.Change event to all listeners registered on this view.
// This is the primary notification event for all ValueView elements.
func (vv *ValueViewBase) SendChange() {
	vv.Send(events.Change, nil)
}

// Send sends an NEW event of given type to this widget,
// optionally starting from values in the given original event
// (recommended to include where possible).
// Do NOT send an existing event using this method if you
// want the Handled state to persist throughout the call chain;
// call HandleEvent directly for any existing events.
func (vv *ValueViewBase) Send(typ events.Types, orig events.Event) {
	var e events.Event
	if orig != nil {
		e = orig.Clone()
		e.AsBase().Typ = typ
	} else {
		e = &events.Base{Typ: typ}
	}
	vv.HandleEvent(e)
}

// HandleEvent sends the given event to all Listeners for that event type.
// It also checks if the State has changed and calls ApplyStyle if so.
// If more significant Config level changes are needed due to an event,
// the event handler must do this itself.
func (vv *ValueViewBase) HandleEvent(ev events.Event) {
	if gi.EventTrace {
		fmt.Println("Event to ValueView:", vv.String(), ev.String())
	}
	vv.Listeners.Call(ev)
}

func (vv *ValueViewBase) SaveTmp() {
	if vv.TmpSave == nil {
		return
	}
	if vv.TmpSave == vv.This().(ValueView) {
		// if we are a map value, of a struct value, we save our value
		if vv.Owner != nil && vv.OwnKind == reflect.Map && !vv.IsMapKey {
			if laser.NonPtrValue(vv.Value).Kind() == reflect.Struct {
				ov := laser.NonPtrValue(reflect.ValueOf(vv.Owner))
				if vv.KeyView != nil {
					ck := laser.NonPtrValue(vv.KeyView.Val())
					laser.SetMapRobust(ov, ck, laser.NonPtrValue(vv.Value))
				} else {
					laser.SetMapRobust(ov, laser.NonPtrValue(reflect.ValueOf(vv.Key)), laser.NonPtrValue(vv.Value))
					// fmt.Printf("save tmp of struct value in key: %v\n", vv.Key)
				}
			}
		}
	} else {
		vv.TmpSave.SaveTmp()
	}
}

func (vv *ValueViewBase) CreateTempIfNotPtr() bool {
	if vv.Value.Kind() != reflect.Ptr { // we create a temp variable -- SaveTmp will save it!
		vv.TmpSave = vv.This().(ValueView) // we are it!
		vtyp := reflect.TypeOf(vv.Value.Interface())
		vtp := reflect.New(vtyp)
		// fmt.Printf("vtyp: %v %v %v, vtp: %v %v %T\n", vtyp, vtyp.Name(), vtyp.String(), vtp, vtp.Type(), vtp.Interface())
		laser.SetRobust(vtp.Interface(), vv.Value.Interface())
		vv.Value = vtp // use this instead
		return true
	}
	return false
}

func (vv *ValueViewBase) SetTags(tags map[string]string) {
	if vv.Tags == nil {
		vv.Tags = make(map[string]string, len(tags))
	}
	for tag, val := range tags {
		vv.Tags[tag] = val
	}
}

func (vv *ValueViewBase) SetTag(tag, value string) {
	if vv.Tags == nil {
		vv.Tags = make(map[string]string, 10)
	}
	vv.Tags[tag] = value
}

func (vv *ValueViewBase) Tag(tag string) (string, bool) {
	if vv.Tags != nil {
		if tv, ok := vv.Tags[tag]; ok {
			return tv, ok
		}
	}
	if !(vv.Owner != nil && vv.OwnKind == reflect.Struct) {
		return "", false
	}
	return vv.Field.Tag.Lookup(tag)
}

func (vv *ValueViewBase) AllTags() map[string]string {
	rvt := make(map[string]string)
	if vv.Tags != nil {
		for key, val := range vv.Tags {
			rvt[key] = val
		}
	}
	if !(vv.Owner != nil && vv.OwnKind == reflect.Struct) {
		return rvt
	}
	smap := laser.StructTags(vv.Field.Tag)
	for key, val := range smap {
		rvt[key] = val
	}
	return rvt
}

// OwnerLabel returns some extra info about the owner of this value view
// which is useful to put in title of our object
func (vv *ValueViewBase) OwnerLabel() string {
	if vv.Owner == nil {
		return ""
	}
	olbl, _ := gi.ToLabeler(vv.Owner)
	switch vv.OwnKind {
	case reflect.Struct:
		if olbl != "" {
			return olbl + "." + vv.Field.Name
		}
		return vv.Field.Name
	case reflect.Map:
		kystr := ""
		if vv.IsMapKey {
			kv := laser.NonPtrValue(vv.Value)
			kystr = laser.ToString(kv.Interface())
		} else {
			if vv.KeyView != nil {
				ck := laser.NonPtrValue(vv.KeyView.Val()) // current key value
				kystr = laser.ToString(ck.Interface())
			} else {
				kystr = laser.ToString(vv.Key)
			}
		}
		if kystr != "" {
			return olbl + "[" + kystr + "]"
		}
		return olbl
	case reflect.Slice:
		if lblr, ok := vv.Owner.(gi.SliceLabeler); ok {
			slbl := lblr.ElemLabel(vv.Idx)
			if slbl != "" {
				return fmt.Sprintf("%s[%s]", olbl, slbl)
			}
		}
		return fmt.Sprintf("%s[%d]", olbl, vv.Idx)
	}
	return olbl
}

// Label returns a label for this item suitable for a window title etc.
// Based on the underlying value type name, owner label, and ViewPath.
// newPath returns just what should be added to a ViewPath
// also includes zero value check reported in the isZero bool, which
// can be used for not proceeding in case of non-value-based types.
func (vv *ValueViewBase) Label() (label, newPath string, isZero bool) {
	lbl := ""
	var npt reflect.Type
	if laser.ValueIsZero(vv.Value) || laser.ValueIsZero(laser.NonPtrValue(vv.Value)) {
		npt = laser.NonPtrType(vv.Value.Type())
		isZero = true
	} else {
		opv := laser.OnePtrUnderlyingValue(vv.Value)
		npt = laser.NonPtrType(opv.Type())
	}
	lbl += npt.String()
	newPath = lbl
	olbl := vv.OwnerLabel()
	if olbl != "" {
		lbl += ": " + olbl
	}
	if vv.ViewPath != "" {
		lbl += " [" + vv.ViewPath + "]"
	}
	return lbl, newPath, isZero
}

////////////////////////////////////////////////////////////////////////////////////////
//   Base Widget Functions -- these are typically redefined in ValueView subtypes

func (vv *ValueViewBase) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.TextFieldType
	return vv.WidgetTyp
}

func (vv *ValueViewBase) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	tf := vv.Widget.(*gi.TextField)
	npv := laser.NonPtrValue(vv.Value)
	// fmt.Printf("vvb val: %v  type: %v  kind: %v\n", npv.Interface(), npv.Type().String(), npv.Kind())
	if npv.Kind() == reflect.Interface && laser.ValueIsZero(npv) {
		tf.SetText("nil")
	} else {
		txt := laser.ToString(vv.Value.Interface())
		tf.SetText(txt)
	}
}

func (vv *ValueViewBase) ConfigWidget(widg gi.Widget) {
	vv.Widget = widg
	vv.StdConfigWidget(widg)
	tf, ok := vv.Widget.(*gi.TextField)
	if !ok {
		return
	}
	tf.SetStretchMaxWidth()
	tf.Tooltip, _ = vv.Tag("desc")
	tf.SetState(vv.This().(ValueView).IsInactive(), states.Disabled)
	// STYTODO: need better solution to value view style configuration (this will add too many stylers)
	tf.AddStyles(func(s *styles.Style) {
		s.MinWidth.SetCh(16)
	})
	if completetag, ok := vv.Tag("complete"); ok {
		// todo: this does not seem to be up-to-date and should use Completer interface..
		in := []reflect.Value{reflect.ValueOf(tf)}
		in = append(in, reflect.ValueOf(completetag)) // pass tag value - object may doing completion on multiple fields
		cmpfv := reflect.ValueOf(vv.Owner).MethodByName("SetCompleter")
		if laser.ValueIsZero(cmpfv) {
			log.Printf("giv.ValueViewBase: programmer error -- SetCompleter method not found in type: %T\n", vv.Owner)
		} else {
			cmpfv.Call(in)
		}
	}

	tf.OnChange(func(e events.Event) {
		if vv.SetValue(tf.Text()) {
			vv.UpdateWidget() // always update after setting value..
		}
	})
	vv.UpdateWidget()
}

// StdConfigWidget does all of the standard widget configuration tag options
func (vv *ValueViewBase) StdConfigWidget(widg gi.Widget) {
	// STYTODO: get rid of this
	nb := widg.AsWidget()
	if widthtag, ok := vv.Tag("width"); ok {
		width, err := laser.ToFloat32(widthtag)
		if err == nil {
			nb.SetMinPrefWidth(units.Ch(width))
		}
	}
	if maxwidthtag, ok := vv.Tag("max-width"); ok {
		width, err := laser.ToFloat32(maxwidthtag)
		if err == nil {
			nb.SetProp("max-width", units.Ch(width))
		}
	}
	if heighttag, ok := vv.Tag("height"); ok {
		height, err := laser.ToFloat32(heighttag)
		if err == nil {
			nb.SetMinPrefHeight(units.Em(height))
		}
	}
	if maxheighttag, ok := vv.Tag("max-height"); ok {
		height, err := laser.ToFloat32(maxheighttag)
		if err == nil {
			nb.SetProp("max-height", units.Em(height))
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////
// ViewIFace

// giv.ViewIFace is THE implementation of the gi.ViewIFace interface
type ViewIFace struct {
}

func (vi *ViewIFace) CtxtMenuView(val any, inactive bool, sc *gi.Scene, menu *gi.Menu) bool {
	// return CtxtMenuView(val, inactive, sc, menu)
	return false
}

func (vi *ViewIFace) GoGiEditor(obj ki.Ki) {
	GoGiEditorDialog(obj)
}

func (vi *ViewIFace) PrefsView(prefs *gi.Preferences) {
	PrefsView(prefs)
}

func (vi *ViewIFace) KeyMapsView(maps *gi.KeyMaps) {
	KeyMapsView(maps)
}

func (vi *ViewIFace) PrefsDetView(prefs *gi.PrefsDetailed) {
	PrefsDetView(prefs)
}

func (vi *ViewIFace) HiStylesView(std bool) {
	if std {
		HiStylesView(&histyle.StdStyles)
	} else {
		HiStylesView(&histyle.CustomStyles)
	}
}

func (vi *ViewIFace) HiStyleInit() {
	histyle.Init()
}

func (vi *ViewIFace) SetHiStyleDefault(hsty gi.HiStyleName) {
	histyle.StyleDefault = hsty
}

func (vi *ViewIFace) PrefsDetDefaults(pf *gi.PrefsDetailed) {
	// pf.TextViewClipHistMax = TextViewClipHistMax
	pf.TextBufMaxScopeLines = TextBufMaxScopeLines
	pf.TextBufDiffRevertLines = TextBufDiffRevertLines
	pf.TextBufDiffRevertDiffs = TextBufDiffRevertDiffs
	pf.TextBufMarkupDelayMSec = TextBufMarkupDelayMSec
	pf.MapInlineLen = MapInlineLen
	pf.StructInlineLen = StructInlineLen
	pf.SliceInlineLen = SliceInlineLen
}

func (vi *ViewIFace) PrefsDetApply(pf *gi.PrefsDetailed) {
	// TextViewClipHistMax = pf.TextViewClipHistMax
	TextBufMaxScopeLines = pf.TextBufMaxScopeLines
	TextBufDiffRevertLines = pf.TextBufDiffRevertLines
	TextBufDiffRevertDiffs = pf.TextBufDiffRevertDiffs
	TextBufMarkupDelayMSec = pf.TextBufMarkupDelayMSec
	MapInlineLen = pf.MapInlineLen
	StructInlineLen = pf.StructInlineLen
	SliceInlineLen = pf.SliceInlineLen
}

func (vi *ViewIFace) PrefsDbgView(prefs *gi.PrefsDebug) {
	PrefsDbgView(prefs)
}

////////////////////////////////////////////////////////////////////////////////////////
//  VersCtrlValueView

// VersCtrlSystems is a list of supported Version Control Systems.
// These must match the VCS Types from goki/pi/vci which in turn
// is based on masterminds/vcs
var VersCtrlSystems = []string{"git", "svn", "bzr", "hg"}

// IsVersCtrlSystem returns true if the given string matches one of the
// standard VersCtrlSystems -- uses lowercase version of str.
func IsVersCtrlSystem(str string) bool {
	stl := strings.ToLower(str)
	for _, vcn := range VersCtrlSystems {
		if stl == vcn {
			return true
		}
	}
	return false
}

// VersCtrlName is the name of a version control system
type VersCtrlName string

func VersCtrlNameProper(vc string) VersCtrlName {
	vcl := strings.ToLower(vc)
	for _, vcnp := range VersCtrlSystems {
		vcnpl := strings.ToLower(vcnp)
		if strings.Compare(vcl, vcnpl) == 0 {
			return VersCtrlName(vcnp)
		}
	}
	return ""
}

// ValueView registers VersCtrlValueView as the viewer of VersCtrlName
func (kn VersCtrlName) ValueView() ValueView {
	vv := &VersCtrlValueView{}
	ki.InitNode(vv)
	return vv
}

// VersCtrlValueView presents an action for displaying an VersCtrlName and selecting
// from StringPopup
type VersCtrlValueView struct {
	ValueViewBase
}

func (vv *VersCtrlValueView) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.ButtonType
	return vv.WidgetTyp
}

func (vv *VersCtrlValueView) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	ac := vv.Widget.(*gi.Button)
	txt := laser.ToString(vv.Value.Interface())
	if txt == "" {
		txt = "(none)"
	}
	ac.SetText(txt)
}

func (vv *VersCtrlValueView) ConfigWidget(widg gi.Widget) {
	vv.Widget = widg
	ac := vv.Widget.(*gi.Button)
	ac.OnClick(func(e events.Event) {
		vv.OpenDialog(vv.Widget, nil)
	})
	vv.UpdateWidget()
}

func (vv *VersCtrlValueView) HasDialog() bool {
	return true
}

func (vv *VersCtrlValueView) OpenDialog(ctx gi.Widget, fun func(dlg *gi.Dialog)) {
	if vv.IsInactive() {
		return
	}
	cur := laser.ToString(vv.Value.Interface())
	gi.StringsChooserPopup(VersCtrlSystems, cur, ctx, func(ac *gi.Button) {
		vv.SetValue(ac.Text)
		vv.UpdateWidget()
	})
}
