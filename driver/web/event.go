// Copyright 2023 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js

package web

import (
	"image"
	"syscall/js"

	"goki.dev/goosi/events"
	"goki.dev/goosi/events/key"
)

func (app *App) addEventListeners() {
	g := js.Global()
	g.Call("addEventListener", "mousedown", js.FuncOf(app.onMouseDown))
	g.Call("addEventListener", "touchstart", js.FuncOf(app.onTouchStart))
	g.Call("addEventListener", "mouseup", js.FuncOf(app.onMouseUp))
	g.Call("addEventListener", "touchend", js.FuncOf(app.onTouchEnd))
	g.Call("addEventListener", "mousemove", js.FuncOf(app.onMouseMove))
	g.Call("addEventListener", "touchmove", js.FuncOf(app.onTouchMove))
	g.Call("addEventListener", "contextmenu", js.FuncOf(app.onContextMenu))
	g.Call("addEventListener", "keydown", js.FuncOf(app.onKeyDown))
	g.Call("addEventListener", "keyup", js.FuncOf(app.onKeyUp))
	g.Call("addEventListener", "beforeinput", js.FuncOf(app.onBeforeInput))
	g.Call("addEventListener", "resize", js.FuncOf(app.onResize))
}

// eventPos returns the appropriate position for the given event,
// multiplying the x and y components by the device pixel ratio
// so that they line up correctly with the canvas.
func (app *App) eventPos(e js.Value) image.Point {
	xi, yi := e.Get("clientX").Int(), e.Get("clientY").Int()
	xi = int(float32(xi) * app.screen.DevicePixelRatio)
	yi = int(float32(yi) * app.screen.DevicePixelRatio)
	return image.Pt(xi, yi)
}

func (app *App) onMouseDown(this js.Value, args []js.Value) any {
	e := args[0]
	but := e.Get("button").Int()
	var ebut events.Buttons
	switch but {
	case 0:
		ebut = events.Left
	case 1:
		ebut = events.Middle
	case 2:
		ebut = events.Right
	}
	where := app.eventPos(e)
	app.window.EvMgr.MouseButton(events.MouseDown, ebut, where, app.keyMods)
	e.Call("preventDefault")
	return nil
}

func (app *App) onTouchStart(this js.Value, args []js.Value) any {
	e := args[0]
	touches := e.Get("changedTouches")
	for i := 0; i < touches.Length(); i++ {
		touch := touches.Index(i)
		where := app.eventPos(touch)
		app.window.EvMgr.MouseButton(events.MouseDown, events.Left, where, 0)
	}
	e.Call("preventDefault")
	return nil
}

func (app *App) onMouseUp(this js.Value, args []js.Value) any {
	e := args[0]
	but := e.Get("button").Int()
	var ebut events.Buttons
	switch but {
	case 0:
		ebut = events.Left
	case 1:
		ebut = events.Middle
	case 2:
		ebut = events.Right
	}
	where := app.eventPos(e)
	app.window.EvMgr.MouseButton(events.MouseUp, ebut, where, app.keyMods)
	e.Call("preventDefault")
	return nil
}

func (app *App) onTouchEnd(this js.Value, args []js.Value) any {
	e := args[0]
	touches := e.Get("changedTouches")
	for i := 0; i < touches.Length(); i++ {
		touch := touches.Index(i)
		where := app.eventPos(touch)
		app.window.EvMgr.MouseButton(events.MouseUp, events.Left, where, 0)
	}
	e.Call("preventDefault")
	return nil
}

func (app *App) onMouseMove(this js.Value, args []js.Value) any {
	e := args[0]
	where := app.eventPos(e)
	app.window.EvMgr.MouseMove(where)
	e.Call("preventDefault")
	return nil
}

func (app *App) onTouchMove(this js.Value, args []js.Value) any {
	e := args[0]
	touches := e.Get("changedTouches")
	for i := 0; i < touches.Length(); i++ {
		touch := touches.Index(i)
		where := app.eventPos(touch)
		app.window.EvMgr.MouseMove(where)
	}
	e.Call("preventDefault")
	return nil
}

func (app *App) onContextMenu(this js.Value, args []js.Value) any {
	// no-op (we handle elsewhere), but needed to prevent browser
	// from making its own context menus on right clicks
	e := args[0]
	e.Call("preventDefault")
	return nil
}

// runeAndCodeFromKey returns the rune and key code corresponding to the given key string.
// down is whether this is from a keyDown event (as opposed to a keyUp one)
func (app *App) runeAndCodeFromKey(k string, down bool) (rune, key.Codes) {
	switch k {
	case "Shift":
		app.keyMods.SetFlag(down, key.Shift)
		return 0, key.CodeLeftShift
	case "Control":
		app.keyMods.SetFlag(down, key.Control)
		return 0, key.CodeLeftControl
	case "Alt":
		app.keyMods.SetFlag(down, key.Alt)
		return 0, key.CodeLeftAlt
	case "Meta":
		app.keyMods.SetFlag(down, key.Meta)
		return 0, key.CodeLeftMeta
	case "Backspace":
		return 0, key.CodeBackspace
	case "Delete":
		return 0, key.CodeDelete
	case "Enter":
		return 0, key.CodeReturnEnter
	case "Tab":
		return 0, key.CodeTab
	case "ArrowDown":
		return 0, key.CodeDownArrow
	case "ArrowLeft":
		return 0, key.CodeLeftArrow
	case "ArrowRight":
		return 0, key.CodeRightArrow
	case "ArrowUp":
		return 0, key.CodeUpArrow
	case "Spacebar":
		return ' ', 0
	default:
		return []rune(k)[0], 0
	}
}

func (app *App) onKeyDown(this js.Value, args []js.Value) any {
	e := args[0]
	k := e.Get("key").String()
	if k == "Unidentified" {
		return nil
	}
	r, c := app.runeAndCodeFromKey(k, true)
	app.window.EvMgr.Key(events.KeyDown, r, c, app.keyMods)
	e.Call("preventDefault")
	return nil
}

func (app *App) onKeyUp(this js.Value, args []js.Value) any {
	e := args[0]
	k := e.Get("key").String()
	if k == "Unidentified" {
		return nil
	}
	r, c := app.runeAndCodeFromKey(k, false)
	app.window.EvMgr.Key(events.KeyUp, r, c, app.keyMods)
	e.Call("preventDefault")
	return nil
}

func (app *App) onBeforeInput(this js.Value, args []js.Value) any {
	e := args[0]
	data := e.Get("data").String()
	if data == "" {
		return nil
	}
	for _, r := range data {
		app.window.EvMgr.KeyChord(r, 0, app.keyMods)
	}
	e.Call("preventDefault")
	return nil
}

func (app *App) onResize(this js.Value, args []js.Value) any {
	app.resize()
	return nil
}
