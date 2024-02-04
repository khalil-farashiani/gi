// Code generated by "core generate"; DO NOT EDIT.

package webcore

import (
	"io/fs"

	"cogentcore.org/core/gti"
	"cogentcore.org/core/ki"
)

// PageType is the [gti.Type] for [Page]
var PageType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/webcore.Page", IDName: "page", Doc: "Page represents one site page", Embeds: []gti.Field{{Name: "Frame"}}, Fields: []gti.Field{{Name: "Source", Doc: "Source is the filesystem in which the content is located."}, {Name: "Context", Doc: "Context is the page's [coredom.Context]."}, {Name: "History", Doc: "The history of URLs that have been visited. The oldest page is first."}, {Name: "HistoryIndex", Doc: "HistoryIndex is the current place we are at in the History"}, {Name: "PagePath", Doc: "PagePath is the fs path of the current page in [Page.Source]"}, {Name: "URLToPagePath", Doc: "URLToPagePath is a map between user-facing page URLs and underlying\nFS page paths."}}, Instance: &Page{}})

// NewPage adds a new [Page] with the given name to the given parent:
// Page represents one site page
func NewPage(par ki.Ki, name ...string) *Page {
	return par.NewChild(PageType, name...).(*Page)
}

// KiType returns the [*gti.Type] of [Page]
func (t *Page) KiType() *gti.Type { return PageType }

// New returns a new [*Page] value
func (t *Page) New() ki.Ki { return &Page{} }

// SetSource sets the [Page.Source]:
// Source is the filesystem in which the content is located.
func (t *Page) SetSource(v fs.FS) *Page { t.Source = v; return t }

// SetTooltip sets the [Page.Tooltip]
func (t *Page) SetTooltip(v string) *Page { t.Tooltip = v; return t }

// SetStackTop sets the [Page.StackTop]
func (t *Page) SetStackTop(v int) *Page { t.StackTop = v; return t }
