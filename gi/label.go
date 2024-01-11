// Copyright (c) 2018, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"fmt"

	"goki.dev/cursors"
	"goki.dev/gi/v2/keyfun"
	"goki.dev/girl/abilities"
	"goki.dev/girl/paint"
	"goki.dev/girl/states"
	"goki.dev/girl/styles"
	"goki.dev/goosi"
	"goki.dev/goosi/events"
	"goki.dev/goosi/mimedata"
	"goki.dev/mat32/v2"
)

// Label is a widget for rendering text labels -- supports full widget model
// including box rendering, and full HTML styling, including links -- LinkSig
// emits link with data of URL -- opens default browser if nobody receiving
// signal.  The default white-space option is 'pre' -- set to 'normal' or
// other options to get word-wrapping etc.
type Label struct { //goki:embedder
	WidgetBase

	// label to display
	Text string

	// the type of label
	Type LabelTypes

	// render data for text label
	TextRender paint.Text `copy:"-" xml:"-" json:"-" set:"-"`
}

func (lb *Label) CopyFieldsFrom(frm any) {
	fr := frm.(*Label)
	lb.WidgetBase.CopyFieldsFrom(&fr.WidgetBase)
	lb.Text = fr.Text
}

// LabelTypes is an enum containing the different
// possible types of labels
type LabelTypes int32 //enums:enum -trim-prefix Label

const (
	// LabelDisplayLarge is a large, short, and important
	// display label with a default font size of 57dp.
	LabelDisplayLarge LabelTypes = iota
	// LabelDisplayMedium is a medium-sized, short, and important
	// display label with a default font size of 45dp.
	LabelDisplayMedium
	// LabelDisplaySmall is a small, short, and important
	// display label with a default font size of 36dp.
	LabelDisplaySmall

	// LabelHeadlineLarge is a large, high-emphasis
	// headline label with a default font size of 32dp.
	LabelHeadlineLarge
	// LabelHeadlineMedium is a medium-sized, high-emphasis
	// headline label with a default font size of 28dp.
	LabelHeadlineMedium
	// LabelHeadlineSmall is a small, high-emphasis
	// headline label with a default font size of 24dp.
	LabelHeadlineSmall

	// LabelTitleLarge is a large, medium-emphasis
	// title label with a default font size of 22dp.
	LabelTitleLarge
	// LabelTitleMedium is a medium-sized, medium-emphasis
	// title label with a default font size of 16dp.
	LabelTitleMedium
	// LabelTitleSmall is a small, medium-emphasis
	// title label with a default font size of 14dp.
	LabelTitleSmall

	// LabelBodyLarge is a large body label used for longer
	// passages of text with a default font size of 16dp.
	LabelBodyLarge
	// LabelBodyMedium is a medium-sized body label used for longer
	// passages of text with a default font size of 14dp.
	LabelBodyMedium
	// LabelBodySmall is a small body label used for longer
	// passages of text with a default font size of 12dp.
	LabelBodySmall

	// LabelLabelLarge is a large label used for label text (like a caption
	// or the text inside a button) with a default font size of 14dp.
	LabelLabelLarge
	// LabelLabelMedium is a medium-sized label used for label text (like a caption
	// or the text inside a button) with a default font size of 12dp.
	LabelLabelMedium
	// LabelLabelSmall is a small label used for label text (like a caption
	// or the text inside a button) with a default font size of 11dp.
	LabelLabelSmall
)

func (lb *Label) OnInit() {
	lb.WidgetBase.OnInit()
	lb.HandleEvents()
	lb.SetStyles()
}

func (lb *Label) SetStyles() {
	lb.Type = LabelBodyLarge
	lb.Style(func(s *styles.Style) {
		s.SetAbilities(true, abilities.Selectable, abilities.DoubleClickable)
		if len(lb.TextRender.Links) > 0 {
			s.SetAbilities(true, abilities.Pressable, abilities.LongHoverable, abilities.LongPressable)
		}
		if !lb.IsReadOnly() {
			s.Cursor = cursors.Text
		}
		s.Min.Y.Em(1)
		s.Min.X.Ch(3)
		s.Text.WhiteSpace = styles.WhiteSpaceNormal
		// s.Align.Y = styles.Center
		s.Grow.Set(0, 0)
		s.GrowWrap = true

		// Label styles based on https://m3.material.io/styles/typography/type-scale-tokens
		// TODO: maybe support brand and plain global fonts with larger labels defaulting to brand and smaller to plain
		switch lb.Type {
		case LabelLabelLarge:
			s.Text.LineHeight.Dp(20)
			s.Font.Size.Dp(14)
			s.Text.LetterSpacing.Dp(0.1)
			s.Font.Weight = styles.WeightMedium // note: excludes all fonts except Go!
		case LabelLabelMedium:
			s.Text.LineHeight.Dp(16)
			s.Font.Size.Dp(12)
			s.Text.LetterSpacing.Dp(0.5)
			s.Font.Weight = styles.WeightMedium
		case LabelLabelSmall:
			s.Text.LineHeight.Dp(16)
			s.Font.Size.Dp(11)
			s.Text.LetterSpacing.Dp(0.5)
			s.Font.Weight = styles.WeightMedium
		case LabelBodyLarge:
			s.Text.LineHeight.Dp(24)
			s.Font.Size.Dp(16)
			s.Text.LetterSpacing.Dp(0.5)
			s.Font.Weight = styles.WeightNormal
		case LabelBodyMedium:
			s.Text.LineHeight.Dp(20)
			s.Font.Size.Dp(14)
			s.Text.LetterSpacing.Dp(0.25)
			s.Font.Weight = styles.WeightNormal
		case LabelBodySmall:
			s.Text.LineHeight.Dp(16)
			s.Font.Size.Dp(12)
			s.Text.LetterSpacing.Dp(0.4)
			s.Font.Weight = styles.WeightNormal
		case LabelTitleLarge:
			s.Text.LineHeight.Dp(28)
			s.Font.Size.Dp(22)
			s.Text.LetterSpacing.Zero()
			s.Font.Weight = styles.WeightNormal
		case LabelTitleMedium:
			s.Text.LineHeight.Dp(24)
			s.Font.Size.Dp(16)
			s.Text.LetterSpacing.Dp(0.15)
			s.Font.Weight = styles.WeightMedium
		case LabelTitleSmall:
			s.Text.LineHeight.Dp(20)
			s.Font.Size.Dp(14)
			s.Text.LetterSpacing.Dp(0.1)
			s.Font.Weight = styles.WeightMedium
		case LabelHeadlineLarge:
			s.Text.LineHeight.Dp(40)
			s.Font.Size.Dp(32)
			s.Text.LetterSpacing.Zero()
			s.Font.Weight = styles.WeightNormal
		case LabelHeadlineMedium:
			s.Text.LineHeight.Dp(36)
			s.Font.Size.Dp(28)
			s.Text.LetterSpacing.Zero()
			s.Font.Weight = styles.WeightNormal
		case LabelHeadlineSmall:
			s.Text.LineHeight.Dp(32)
			s.Font.Size.Dp(24)
			s.Text.LetterSpacing.Zero()
			s.Font.Weight = styles.WeightNormal
		case LabelDisplayLarge:
			s.Text.LineHeight.Dp(64)
			s.Font.Size.Dp(57)
			s.Text.LetterSpacing.Dp(-0.25)
			s.Font.Weight = styles.WeightNormal
		case LabelDisplayMedium:
			s.Text.LineHeight.Dp(52)
			s.Font.Size.Dp(45)
			s.Text.LetterSpacing.Zero()
			s.Font.Weight = styles.WeightNormal
		case LabelDisplaySmall:
			s.Text.LineHeight.Dp(44)
			s.Font.Size.Dp(36)
			s.Text.LetterSpacing.Zero()
			s.Font.Weight = styles.WeightNormal
		}
	})
}

// SetTextUpdate sets the text and updates the underlying render
// (i.e., ConfigLabel) so it will actually show on next render.
// (including an UpdateRender flag)
func (lb *Label) SetTextUpdate(text string) *Label {
	updt := lb.UpdateStart()
	lb.Text = text
	lb.ConfigLabel()
	lb.UpdateEndRender(updt)
	return lb
}

func (lb *Label) HandleEvents() {
	// lb.HandleLabelLongHover()
	lb.HandleLabelClick(func(tl *paint.TextLink) {
		goosi.TheApp.OpenURL(tl.URL)
	})
	lb.OnKeyChord(func(e events.Event) {
		// TODO(kai): get label copying working
		fmt.Println("kc", e)
		if !lb.StateIs(states.Selected) {
			return
		}
		kf := keyfun.Of(e.KeyChord())
		if kf == keyfun.Copy {
			e.SetHandled()
			lb.Copy(true)
		}
	})
	lb.On(events.MouseMove, func(e events.Event) {
		pos := lb.Geom.Pos.Content
		inLink := false
		for _, tl := range lb.TextRender.Links {
			// TODO(kai/link): is there a better way to be safe here?
			if tl.Label == "" {
				continue
			}
			tlb := tl.Bounds(&lb.TextRender, pos)
			if e.LocalPos().In(tlb) {
				inLink = true
				if lb.StateIs(states.LongHovered) || lb.StateIs(states.LongPressed) {
					NewTooltipTextAt(lb, tl.URL, tlb.Min, tlb.Size()).Run()
					e.SetHandled()
				}
				break
			}
		}
		if inLink {
			lb.Styles.Cursor = cursors.Pointer
		} else if lb.AbilityIs(abilities.Selectable) {
			lb.Styles.Cursor = cursors.Text
		} else {
			lb.Styles.Cursor = cursors.None
		}
	})
}

// func (lb *Label) HandleLabelLongHover() {
// 	lb.On(events.LongHoverStart, func(e events.Event) {
// 		fmt.Println("lb lhs")
// 		pos := lb.Geom.Pos.Content
// 		for _, tl := range lb.TextRender.Links {
// 			tlb := tl.Bounds(&lb.TextRender, pos)
// 			fmt.Println(pos, tlb, e.LocalPos(), e.Pos())
// 			if e.LocalPos().In(tlb) {
// 				fmt.Println("ntt")
// 				NewTooltipText(lb, tl.URL, tlb.Min).Run()
// 				e.SetHandled()
// 				return
// 			}
// 		}
// 	})
// }

// HandleLabelClick handles click events such that the given function will be called
// on any links that are clicked on.
func (lb *Label) HandleLabelClick(openLink func(tl *paint.TextLink)) {
	lb.OnClick(func(e events.Event) {
		pos := lb.Geom.Pos.Content
		for _, tl := range lb.TextRender.Links {
			// TODO(kai/link): is there a better way to be safe here?
			if tl.Label == "" {
				continue
			}
			tlb := tl.Bounds(&lb.TextRender, pos)
			if e.LocalPos().In(tlb) {
				openLink(&tl)
				e.SetHandled()
				return
			}
		}
	})
}

func (lb *Label) Copy(reset bool) {
	md := mimedata.NewText(lb.Text)
	em := lb.EventMgr()
	if em != nil {
		em.Clipboard().Write(md)
	}
}

func (lb *Label) ConfigWidget() {
	lb.ConfigLabel()
}

// todo: ideally it would be possible to only call SetHTML once during config
// and then do the layout only during sizing.  However, layout starts with
// existing line breaks (which could come from <br> and <p> in HTML),
// so that is never able to undo initial word wrapping from constrained sizes.

// ConfigLabel does the HTML and Layout in TextRender for label text,
// using actual content size to constrain layout.
func (lb *Label) ConfigLabel() {
	lb.ConfigLabelSize(lb.Geom.Size.Actual.Content)
}

// ConfigLabel does the HTML and Layout in TextRender for label text,
// using given size to constrain layout.
func (lb *Label) ConfigLabelSize(sz mat32.Vec2) {
	lb.StyMu.RLock()
	defer lb.StyMu.RUnlock()

	// todo: last arg is CSSAgg.  Can synthesize that some other way?
	fs := lb.Styles.FontRender()
	txs := &lb.Styles.Text
	lb.TextRender.SetHTML(lb.Text, fs, txs, &lb.Styles.UnContext, nil)
	lb.TextRender.LayoutStdLR(txs, fs, &lb.Styles.UnContext, sz)
}

// ConfigLabelAlloc is used for determining how much space the label
// takes, using given size (typically Alloc).
// In this case, alignment factors are turned off,
// because they otherwise can absorb much more space, which should
// instead be controlled by the base Align X,Y factors.
func (lb *Label) ConfigLabelAlloc(sz mat32.Vec2) mat32.Vec2 {
	lb.StyMu.RLock()
	defer lb.StyMu.RUnlock()

	// todo: last arg is CSSAgg.  Can synthesize that some other way?
	fs := lb.Styles.FontRender()
	txs := &lb.Styles.Text
	align, alignV := txs.Align, txs.AlignV
	txs.Align, txs.AlignV = styles.Start, styles.Start
	lb.TextRender.SetHTML(lb.Text, fs, txs, &lb.Styles.UnContext, nil)
	lb.TextRender.LayoutStdLR(txs, fs, &lb.Styles.UnContext, sz)
	rsz := lb.TextRender.Size.Ceil()
	txs.Align, txs.AlignV = align, alignV
	lb.TextRender.LayoutStdLR(txs, fs, &lb.Styles.UnContext, rsz)
	return rsz
}

// SizeUpWrapSize is the size to use for layout during the SizeUp pass,
// for word wrap case, where the sizing actually matters.
// this is based on the existing styled Actual.Content aspect ratio and
// very rough estimate of total rendered size.
func (lb *Label) SizeUpWrapSize() mat32.Vec2 {
	csz := lb.Geom.Size.Actual.Content
	chars := float32(len(lb.Text))
	fht := float32(16)
	if lb.Styles.Font.Face != nil {
		fht = lb.Styles.Font.Face.Metrics.Height
	}
	area := chars * fht * fht
	ratio := float32(1.618) // default to golden
	if csz.X > 0 && csz.Y > 0 {
		ratio = csz.X / csz.Y
		// fmt.Println(lb, "content size ratio:", ratio)
	}
	// w = ratio * h
	// w^2 + h^2 = a
	// (ratio*h)^2 + h^2 = a
	h := mat32.Sqrt(area) / mat32.Sqrt(ratio+1)
	w := ratio * h
	if w < csz.X { // must be at least this
		w = csz.X
		h = area / w
		h = max(h, csz.Y)
	}
	sz := mat32.V2(w, h)
	if DebugSettings.LayoutTrace {
		fmt.Println(lb, "SizeUpWrapSize chars:", chars, "area:", area, "sz:", sz)
	}
	return sz
}

func (lb *Label) SizeUp() {
	lb.WidgetBase.SizeUp() // sets Actual size based on styles
	sz := &lb.Geom.Size
	if lb.Styles.Text.HasWordWrap() {
		lb.ConfigLabelSize(lb.SizeUpWrapSize())
	} else {
		lb.ConfigLabelSize(sz.Actual.Content)
	}
	rsz := lb.TextRender.Size.Ceil()
	sz.FitSizeMax(&sz.Actual.Content, rsz)
	sz.SetTotalFromContent(&sz.Actual)
	if DebugSettings.LayoutTrace {
		fmt.Println(lb, "Label SizeUp:", rsz, "Actual:", sz.Actual.Content)
	}
}

func (lb *Label) SizeDown(iter int) bool {
	if !lb.Styles.Text.HasWordWrap() || iter > 1 {
		return false
	}
	sz := &lb.Geom.Size
	rsz := lb.ConfigLabelAlloc(sz.Alloc.Content) // use allocation
	prevContent := sz.Actual.Content
	// start over so we don't reflect hysteresis of prior guess
	sz.SetInitContentMin(lb.Styles.Min.Dots().Ceil())
	sz.FitSizeMax(&sz.Actual.Content, rsz)
	sz.SetTotalFromContent(&sz.Actual)
	chg := prevContent != sz.Actual.Content
	if chg {
		if DebugSettings.LayoutTrace {
			fmt.Println(lb, "Label Size Changed:", sz.Actual.Content, "was:", prevContent)
		}
	}
	return chg
}

func (lb *Label) RenderLabel() {
	pc, st := lb.RenderLock()
	lb.RenderStdBox(st)
	lb.TextRender.Render(pc, lb.Geom.Pos.Content)
	lb.RenderUnlock()
}

func (lb *Label) Render() {
	if lb.PushBounds() {
		lb.RenderLabel()
		lb.RenderChildren()
		lb.PopBounds()
	}
}