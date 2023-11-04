// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package giv

import (
	"log/slog"
	"strconv"
	"time"

	"goki.dev/gi/v2/gi"
	"goki.dev/girl/styles"
	"goki.dev/girl/units"
	"goki.dev/goosi/events"
	"goki.dev/gti"
	"goki.dev/icons"
	"goki.dev/laser"
	"goki.dev/pi/v2/filecat"
)

// TimeView is a view for selecting a time
type TimeView struct {
	gi.Frame

	// the time that we are viewing
	Time time.Time `set:"-"`

	// the raw input hour
	Hour int `set:"-"`

	// whether we are in PM mode (so we have to add 12h to everything)
	PM bool `set:"-"`
}

// SetTime sets the source time and updates the view
func (tv *TimeView) SetTime(tim time.Time) *TimeView {
	updt := tv.UpdateStart()
	tv.Time = tim
	tv.UpdateEndRender(updt)
	tv.SendChange()
	return tv
}

func (tv *TimeView) ConfigWidget(sc *gi.Scene) {
	if tv.HasChildren() {
		return
	}
	updt := tv.UpdateStart()

	tv.SetLayout(gi.LayoutHoriz)

	hour := gi.NewTextField(tv, "hour")
	if gi.Prefs.Clock24 {
		hour.SetText(strconv.Itoa(tv.Time.Hour()))
	} else {
		hour.SetText(strconv.Itoa(tv.Time.Hour() % 12))
	}
	hour.Style(func(s *styles.Style) {
		s.Font.Size.Dp(57)
		s.SetFixedWidth(units.Dp(96))
	})
	hour.OnChange(func(e events.Event) {
		hr, err := strconv.Atoi(hour.Text())
		// TODO(kai/snack)
		if err != nil {
			slog.Error(err.Error())
		}
		tv.Hour = hr
		if tv.PM {
			// only add to local
			hr += 12
		}
		// we set our hour and keep everything else
		tv.Time = tv.Time.Truncate(time.Hour).Add(time.Hour * time.Duration(hr)).Add(time.Minute * time.Duration(tv.Time.Minute()))
	})

	gi.NewLabel(tv, "colon").SetType(gi.LabelDisplayLarge).SetText(":")

	minute := gi.NewTextField(tv, "minute").
		SetText(strconv.Itoa(tv.Time.Minute()))
	minute.Style(func(s *styles.Style) {
		s.Font.Size.Dp(57)
		s.SetFixedWidth(units.Dp(96))
	})
	minute.OnChange(func(e events.Event) {
		min, err := strconv.Atoi(minute.Text())
		// TODO(kai/snack)
		if err != nil {
			slog.Error(err.Error())
		}
		// we set our minute and keep everything else
		tv.Time = tv.Time.Truncate(time.Minute).Add(time.Minute * time.Duration(min))
	})

	if !gi.Prefs.Clock24 {
		sw := gi.NewSwitches(tv, "am-pm").SetMutex(true).SetType(gi.SwitchSegmentedButton).SetLayout(gi.LayoutVert).SetItems([]string{"AM", "PM"})
		if tv.Time.Hour() < 12 {
			sw.SelectItemAction(0)
		} else {
			sw.SelectItemAction(1)
		}
		sw.OnChange(func(e events.Event) {
			si := sw.SelectedItem()
			switch si {
			case "AM":
				tv.Time = tv.Time.Truncate(time.Hour).Add(time.Hour * time.Duration(tv.Hour))
			case "PM":
				tv.Time = tv.Time.Truncate(time.Hour).Add(time.Hour * time.Duration(tv.Hour+12))
			default:
				// must always have something valid selected
				sw.SelectItem(0)
			}
		})
	}

	tv.UpdateEnd(updt)
}

// TimeValue presents two text fields for editing a date and time,
// both of which can pull up corresponding picker view dialogs.
type TimeValue struct {
	ValueBase
}

func (vv *TimeValue) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.LayoutType
	return vv.WidgetTyp
}

// TimeVal decodes Value into a *time.Time value -- also handles FileTime case
func (vv *TimeValue) TimeVal() *time.Time {
	tmi := laser.PtrValue(vv.Value).Interface()
	switch v := tmi.(type) {
	case *time.Time:
		return v
	case *filecat.FileTime:
		return (*time.Time)(v)
	}
	return nil
}

func (vv *TimeValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	fr := vv.Widget.(*gi.Layout)
	tm := vv.TimeVal()

	fr.ChildByName("date").(*gi.TextField).SetText(tm.Format(time.DateOnly))
	fr.ChildByName("time").(*gi.TextField).SetText(tm.Format(gi.Prefs.TimeFormat()))
}

func (vv *TimeValue) ConfigWidget(w gi.Widget, sc *gi.Scene) {
	if vv.Widget == w {
		vv.UpdateWidget()
		return
	}
	vv.Widget = w
	vv.StdConfigWidget(w)
	ly := vv.Widget.(*gi.Layout)
	ly.SetLayout(gi.LayoutHoriz)

	if len(ly.Kids) > 0 {
		return
	}

	dt := gi.NewTextField(ly, "date").SetTooltip("The date")
	dt.SetReadOnly(vv.IsReadOnly())
	dt.OnChange(func(e events.Event) {
		d, err := time.Parse(time.DateOnly, dt.Text())
		if err != err {
			// TODO(kai/snack)
			slog.Error(err.Error())
			return
		}
		tv := vv.TimeVal()
		// new date and old time
		*tv = time.Date(d.Year(), d.Month(), d.Day(), tv.Hour(), tv.Minute(), tv.Second(), tv.Nanosecond(), tv.Location())
	})
	dt.Config(sc)

	tm := gi.NewTextField(ly, "time").SetTooltip("The time").
		SetLeadingIcon(icons.Schedule, func(e events.Event) {
			d := gi.NewDialog(w).Title("Edit time")
			NewTimeView(d).SetTime(*vv.TimeVal())
			d.Cancel().Ok().Run()
		})
	tm.SetReadOnly(vv.IsReadOnly())
	tm.OnChange(func(e events.Event) {
		t, err := time.Parse(gi.Prefs.TimeFormat(), tm.Text())
		if err != err {
			// TODO(kai/snack)
			slog.Error(err.Error())
			return
		}
		tv := vv.TimeVal()
		// old date and new time
		*tv = time.Date(tv.Year(), tv.Month(), tv.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), tv.Location())
	})
	dt.Config(sc)

	vv.UpdateWidget()
}

var durationUnits = []string{
	"nanoseconds",
	"microseconds",
	"milliseconds",
	"seconds",
	"minutes",
	"hours",
	"days",
	"weeks",
	"months",
	"years",
}

var durationUnitsMap = map[string]time.Duration{
	"nanoseconds":  time.Nanosecond,
	"microseconds": time.Microsecond,
	"milliseconds": time.Millisecond,
	"seconds":      time.Second,
	"minutes":      time.Minute,
	"hours":        time.Hour,
	"days":         24 * time.Hour,
	"weeks":        7 * 24 * time.Hour,
	"months":       30 * 24 * time.Hour,
	"years":        365 * 24 * time.Hour,
}

// DurationValue presents a spinner and unit chooser for a [time.Duration]
type DurationValue struct {
	ValueBase
}

func (vv *DurationValue) WidgetType() *gti.Type {
	vv.WidgetTyp = gi.LayoutType
	return vv.WidgetTyp
}

func (vv *DurationValue) UpdateWidget() {
	if vv.Widget == nil {
		return
	}
	npv := laser.NonPtrValue(vv.Value)
	dur := npv.Interface().(time.Duration)
	un := "seconds"
	undur := time.Duration(0)
	for _, u := range durationUnits {
		v := durationUnitsMap[u]
		if v > dur {
			break
		}
		un = u
		undur = v
	}
	adur := dur
	if undur != 0 {
		adur = dur / undur
	}

	ly := vv.Widget.(*gi.Layout)
	ly.ChildByName("value").(*gi.Spinner).SetValue(float32(adur))
	ly.ChildByName("unit").(*gi.Chooser).SetCurVal(un)
}

func (vv *DurationValue) ConfigWidget(w gi.Widget, sc *gi.Scene) {
	if vv.Widget == w {
		vv.UpdateWidget()
		return
	}
	vv.Widget = w
	vv.StdConfigWidget(w)
	fr := vv.Widget.(*gi.Layout)
	fr.SetLayout(gi.LayoutHoriz)

	if len(fr.Kids) > 0 {
		return
	}

	var ch *gi.Chooser

	sp := gi.NewSpinner(fr, "value").SetTooltip("The value of time").SetStep(1).SetPageStep(10)
	sp.OnChange(func(e events.Event) {
		vv.SetValue(sp.Value * float32(durationUnitsMap[ch.CurLabel]))
	})
	sp.Config(sc)

	units := []any{}
	for _, u := range durationUnits {
		units = append(units, u)
	}

	ch = gi.NewChooser(fr, "unit").SetTooltip("The unit of time").SetItems(units)
	ch.OnChange(func(e events.Event) {
		// we update the value to fit the unit
		npv := laser.NonPtrValue(vv.Value)
		dur := npv.Interface().(time.Duration)
		sp.SetValue(float32(dur) / float32(durationUnitsMap[ch.CurLabel]))
	})

	vv.UpdateWidget()
}
