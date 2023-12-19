// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on https://github.com/srwiley/rasterx:
// Copyright 2018 by the rasterx Authors. All rights reserved.
// Created 2018 by S.R.Wiley

package gradient

import (
	"image/color"

	"goki.dev/mat32/v2"
)

// Linear represents a linear gradient. It implements the [image.Image] interface.
type Linear struct { //gti:add -setters
	Base

	// the starting point of the gradient (x1 and y1 in SVG)
	Start mat32.Vec2

	// the ending point of the gradient (x2 and y2 in SVG)
	End mat32.Vec2
}

var _ Gradient = &Linear{}

// NewLinear returns a new downward-facing [Linear] gradient.
func NewLinear() *Linear {
	l := &Linear{
		Base: NewBase(),
		// default in CSS is "to bottom"
		End: mat32.V2(0, 1),
	}
	l.Update()
	return l
}

// AddStop adds a new stop with the given color and position to the linear gradient.
func (l *Linear) AddStop(color color.RGBA, pos float32) *Linear {
	l.Base.AddStop(color, pos)
	return l
}

// SetBox sets the [Linear.Box]
func (l *Linear) SetBox(v mat32.Box2) *Linear {
	l.Box = v
	l.Update()
	return l
}

// SetTransform sets the [Linear.Transform]
func (l *Linear) SetTransform(v mat32.Mat2) *Linear {
	l.Transform = v
	l.Update()
	return l
}

// Update updates the computed fields of the linear gradient after it has been modified.
// It should only be called by end users when they modify properties of the linear gradient
// outside of Set functions that have comments stating that they must be set using Set functions.
func (l *Linear) Update() {
	l.UpdateBase()
}

// At returns the color of the linear gradient at the given point
func (l *Linear) At(x, y int) color.Color {
	switch len(l.Stops) {
	case 0:
		return color.RGBA{}
	case 1:
		return l.Stops[0].Color
	}

	s, e := l.Start, l.End
	if l.Units == ObjectBoundingBox {
		s = l.Box.Min.Add(l.Box.Size().Mul(s))
		e = l.Box.Min.Add(l.Box.Size().Mul(e))
	} else {
		s = l.Transform.MulVec2AsPt(s)
		e = l.Transform.MulVec2AsPt(e)
	}

	d := e.Sub(s)
	dd := d.X*d.X + d.Y*d.Y // self inner prod

	pt := mat32.V2(float32(x)+0.5, float32(y)+0.5)
	if l.Units == ObjectBoundingBox {
		pt = l.ObjectMatrix.MulVec2AsPt(pt)
	}
	df := pt.Sub(s)
	pos := (d.X*df.X + d.Y*df.Y) / dd
	return l.GetColor(pos)
}
