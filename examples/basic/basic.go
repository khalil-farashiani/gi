package main

import (
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/gimain"
)

func main() { gimain.Run(app) }

func app() {
	// gi.DebugSettings.RenderTrace = true
	// gi.DebugSettings.WinEventTrace = true
	// gi.EventTrace = true
	// gi.DebugSettings.LayoutTrace = true
	// gi.DebugSettings.LayoutTraceDetail = true
	// gi.DebugSettings.RenderTrace = true

	b := gi.NewAppBody("basic")
	b.App().AppBarConfig = nil

	gi.NewButton(b).SetType(gi.ButtonAction).SetText("Action")

	// gi.NewIcon(sc).SetIcon(icons.Add)

	// gi.NewLabel(b).SetText("Hello, World!")

	// gi.NewButton(sc).
	// 	SetText("Open Dialog").SetIcon(icons.OpenInNew).
	// 	OnClick(func(e events.Event) {
	// 		fmt.Println("button clicked")
	// 		// dialog := gi.NewScene("dialog")
	// 		// gi.NewLabel(&dialog.Frame, "dialog").SetText("Dialog!")
	// 		// gi.NewBody(dialog, but).SetModal().SetMovable().SetCloseable().Run()
	// 	})

	// gi.NewSwitch(sc)

	b.NewWindow().Run().Wait()
}
