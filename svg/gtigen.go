// Code generated by "core generate"; DO NOT EDIT.

package svg

import (
	"image"

	"cogentcore.org/core/colors/gradient"
	"cogentcore.org/core/gti"
	"cogentcore.org/core/ki"
	"cogentcore.org/core/mat32"
	"cogentcore.org/core/paint"
	"github.com/aymerick/douceur/css"
)

// CircleType is the [gti.Type] for [Circle]
var CircleType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Circle", IDName: "circle", Doc: "Circle is a SVG circle", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "Pos", Doc: "position of the center of the circle"}, {Name: "Radius", Doc: "radius of the circle"}}, Instance: &Circle{}})

// NewCircle adds a new [Circle] with the given name to the given parent:
// Circle is a SVG circle
func NewCircle(par ki.Ki, name ...string) *Circle {
	return par.NewChild(CircleType, name...).(*Circle)
}

// KiType returns the [*gti.Type] of [Circle]
func (t *Circle) KiType() *gti.Type { return CircleType }

// New returns a new [*Circle] value
func (t *Circle) New() ki.Ki { return &Circle{} }

// SetPos sets the [Circle.Pos]:
// position of the center of the circle
func (t *Circle) SetPos(v mat32.Vec2) *Circle { t.Pos = v; return t }

// SetRadius sets the [Circle.Radius]:
// radius of the circle
func (t *Circle) SetRadius(v float32) *Circle { t.Radius = v; return t }

// SetClass sets the [Circle.Class]
func (t *Circle) SetClass(v string) *Circle { t.Class = v; return t }

// ClipPathType is the [gti.Type] for [ClipPath]
var ClipPathType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.ClipPath", IDName: "clip-path", Doc: "ClipPath is used for holding a path that renders as a clip path", Embeds: []gti.Field{{Name: "NodeBase"}}, Instance: &ClipPath{}})

// NewClipPath adds a new [ClipPath] with the given name to the given parent:
// ClipPath is used for holding a path that renders as a clip path
func NewClipPath(par ki.Ki, name ...string) *ClipPath {
	return par.NewChild(ClipPathType, name...).(*ClipPath)
}

// KiType returns the [*gti.Type] of [ClipPath]
func (t *ClipPath) KiType() *gti.Type { return ClipPathType }

// New returns a new [*ClipPath] value
func (t *ClipPath) New() ki.Ki { return &ClipPath{} }

// SetClass sets the [ClipPath.Class]
func (t *ClipPath) SetClass(v string) *ClipPath { t.Class = v; return t }

// StyleSheetType is the [gti.Type] for [StyleSheet]
var StyleSheetType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.StyleSheet", IDName: "style-sheet", Doc: "StyleSheet is a Node2D node that contains a stylesheet -- property values\ncontained in this sheet can be transformed into ki.Props and set in CSS\nfield of appropriate node", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "Sheet"}}, Instance: &StyleSheet{}})

// NewStyleSheet adds a new [StyleSheet] with the given name to the given parent:
// StyleSheet is a Node2D node that contains a stylesheet -- property values
// contained in this sheet can be transformed into ki.Props and set in CSS
// field of appropriate node
func NewStyleSheet(par ki.Ki, name ...string) *StyleSheet {
	return par.NewChild(StyleSheetType, name...).(*StyleSheet)
}

// KiType returns the [*gti.Type] of [StyleSheet]
func (t *StyleSheet) KiType() *gti.Type { return StyleSheetType }

// New returns a new [*StyleSheet] value
func (t *StyleSheet) New() ki.Ki { return &StyleSheet{} }

// SetSheet sets the [StyleSheet.Sheet]
func (t *StyleSheet) SetSheet(v *css.Stylesheet) *StyleSheet { t.Sheet = v; return t }

// SetClass sets the [StyleSheet.Class]
func (t *StyleSheet) SetClass(v string) *StyleSheet { t.Class = v; return t }

// MetaDataType is the [gti.Type] for [MetaData]
var MetaDataType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.MetaData", IDName: "meta-data", Doc: "MetaData is used for holding meta data info", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "MetaData"}}, Instance: &MetaData{}})

// NewMetaData adds a new [MetaData] with the given name to the given parent:
// MetaData is used for holding meta data info
func NewMetaData(par ki.Ki, name ...string) *MetaData {
	return par.NewChild(MetaDataType, name...).(*MetaData)
}

// KiType returns the [*gti.Type] of [MetaData]
func (t *MetaData) KiType() *gti.Type { return MetaDataType }

// New returns a new [*MetaData] value
func (t *MetaData) New() ki.Ki { return &MetaData{} }

// SetMetaData sets the [MetaData.MetaData]
func (t *MetaData) SetMetaData(v string) *MetaData { t.MetaData = v; return t }

// SetClass sets the [MetaData.Class]
func (t *MetaData) SetClass(v string) *MetaData { t.Class = v; return t }

// EllipseType is the [gti.Type] for [Ellipse]
var EllipseType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Ellipse", IDName: "ellipse", Doc: "Ellipse is a SVG ellipse", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "Pos", Doc: "position of the center of the ellipse"}, {Name: "Radii", Doc: "radii of the ellipse in the horizontal, vertical axes"}}, Instance: &Ellipse{}})

// NewEllipse adds a new [Ellipse] with the given name to the given parent:
// Ellipse is a SVG ellipse
func NewEllipse(par ki.Ki, name ...string) *Ellipse {
	return par.NewChild(EllipseType, name...).(*Ellipse)
}

// KiType returns the [*gti.Type] of [Ellipse]
func (t *Ellipse) KiType() *gti.Type { return EllipseType }

// New returns a new [*Ellipse] value
func (t *Ellipse) New() ki.Ki { return &Ellipse{} }

// SetPos sets the [Ellipse.Pos]:
// position of the center of the ellipse
func (t *Ellipse) SetPos(v mat32.Vec2) *Ellipse { t.Pos = v; return t }

// SetRadii sets the [Ellipse.Radii]:
// radii of the ellipse in the horizontal, vertical axes
func (t *Ellipse) SetRadii(v mat32.Vec2) *Ellipse { t.Radii = v; return t }

// SetClass sets the [Ellipse.Class]
func (t *Ellipse) SetClass(v string) *Ellipse { t.Class = v; return t }

// FilterType is the [gti.Type] for [Filter]
var FilterType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Filter", IDName: "filter", Doc: "Filter represents SVG filter* elements", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "FilterType"}}, Instance: &Filter{}})

// NewFilter adds a new [Filter] with the given name to the given parent:
// Filter represents SVG filter* elements
func NewFilter(par ki.Ki, name ...string) *Filter {
	return par.NewChild(FilterType, name...).(*Filter)
}

// KiType returns the [*gti.Type] of [Filter]
func (t *Filter) KiType() *gti.Type { return FilterType }

// New returns a new [*Filter] value
func (t *Filter) New() ki.Ki { return &Filter{} }

// SetFilterType sets the [Filter.FilterType]
func (t *Filter) SetFilterType(v string) *Filter { t.FilterType = v; return t }

// SetClass sets the [Filter.Class]
func (t *Filter) SetClass(v string) *Filter { t.Class = v; return t }

// FlowType is the [gti.Type] for [Flow]
var FlowType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Flow", IDName: "flow", Doc: "Flow represents SVG flow* elements", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "FlowType"}}, Instance: &Flow{}})

// NewFlow adds a new [Flow] with the given name to the given parent:
// Flow represents SVG flow* elements
func NewFlow(par ki.Ki, name ...string) *Flow {
	return par.NewChild(FlowType, name...).(*Flow)
}

// KiType returns the [*gti.Type] of [Flow]
func (t *Flow) KiType() *gti.Type { return FlowType }

// New returns a new [*Flow] value
func (t *Flow) New() ki.Ki { return &Flow{} }

// SetFlowType sets the [Flow.FlowType]
func (t *Flow) SetFlowType(v string) *Flow { t.FlowType = v; return t }

// SetClass sets the [Flow.Class]
func (t *Flow) SetClass(v string) *Flow { t.Class = v; return t }

// GradientType is the [gti.Type] for [Gradient]
var GradientType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Gradient", IDName: "gradient", Doc: "Gradient is used for holding a specified color gradient.\nThe name is the id for lookup in url", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "Grad", Doc: "the color gradient"}, {Name: "StopsName", Doc: "name of another gradient to get stops from"}}, Instance: &Gradient{}})

// NewGradient adds a new [Gradient] with the given name to the given parent:
// Gradient is used for holding a specified color gradient.
// The name is the id for lookup in url
func NewGradient(par ki.Ki, name ...string) *Gradient {
	return par.NewChild(GradientType, name...).(*Gradient)
}

// KiType returns the [*gti.Type] of [Gradient]
func (t *Gradient) KiType() *gti.Type { return GradientType }

// New returns a new [*Gradient] value
func (t *Gradient) New() ki.Ki { return &Gradient{} }

// SetGrad sets the [Gradient.Grad]:
// the color gradient
func (t *Gradient) SetGrad(v gradient.Gradient) *Gradient { t.Grad = v; return t }

// SetStopsName sets the [Gradient.StopsName]:
// name of another gradient to get stops from
func (t *Gradient) SetStopsName(v string) *Gradient { t.StopsName = v; return t }

// SetClass sets the [Gradient.Class]
func (t *Gradient) SetClass(v string) *Gradient { t.Class = v; return t }

// GroupType is the [gti.Type] for [Group]
var GroupType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Group", IDName: "group", Doc: "Group groups together SVG elements.\nProvides a common transform for all group elements\nand shared style properties.", Embeds: []gti.Field{{Name: "NodeBase"}}, Instance: &Group{}})

// NewGroup adds a new [Group] with the given name to the given parent:
// Group groups together SVG elements.
// Provides a common transform for all group elements
// and shared style properties.
func NewGroup(par ki.Ki, name ...string) *Group {
	return par.NewChild(GroupType, name...).(*Group)
}

// KiType returns the [*gti.Type] of [Group]
func (t *Group) KiType() *gti.Type { return GroupType }

// New returns a new [*Group] value
func (t *Group) New() ki.Ki { return &Group{} }

// SetClass sets the [Group.Class]
func (t *Group) SetClass(v string) *Group { t.Class = v; return t }

// ImageType is the [gti.Type] for [Image]
var ImageType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Image", IDName: "image", Doc: "Image is an SVG image (bitmap)", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "Pos", Doc: "position of the top-left of the image"}, {Name: "Size", Doc: "rendered size of the image (imposes a scaling on image when it is rendered)"}, {Name: "Filename", Doc: "file name of image loaded -- set by OpenImage"}, {Name: "ViewBox", Doc: "how to scale and align the image"}, {Name: "Pixels", Doc: "the image pixels"}}, Instance: &Image{}})

// NewImage adds a new [Image] with the given name to the given parent:
// Image is an SVG image (bitmap)
func NewImage(par ki.Ki, name ...string) *Image {
	return par.NewChild(ImageType, name...).(*Image)
}

// KiType returns the [*gti.Type] of [Image]
func (t *Image) KiType() *gti.Type { return ImageType }

// New returns a new [*Image] value
func (t *Image) New() ki.Ki { return &Image{} }

// SetPos sets the [Image.Pos]:
// position of the top-left of the image
func (t *Image) SetPos(v mat32.Vec2) *Image { t.Pos = v; return t }

// SetSize sets the [Image.Size]:
// rendered size of the image (imposes a scaling on image when it is rendered)
func (t *Image) SetSize(v mat32.Vec2) *Image { t.Size = v; return t }

// SetFilename sets the [Image.Filename]:
// file name of image loaded -- set by OpenImage
func (t *Image) SetFilename(v string) *Image { t.Filename = v; return t }

// SetViewBox sets the [Image.ViewBox]:
// how to scale and align the image
func (t *Image) SetViewBox(v ViewBox) *Image { t.ViewBox = v; return t }

// SetPixels sets the [Image.Pixels]:
// the image pixels
func (t *Image) SetPixels(v *image.RGBA) *Image { t.Pixels = v; return t }

// SetClass sets the [Image.Class]
func (t *Image) SetClass(v string) *Image { t.Class = v; return t }

// LineType is the [gti.Type] for [Line]
var LineType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Line", IDName: "line", Doc: "Line is a SVG line", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "Start", Doc: "position of the start of the line"}, {Name: "End", Doc: "position of the end of the line"}}, Instance: &Line{}})

// NewLine adds a new [Line] with the given name to the given parent:
// Line is a SVG line
func NewLine(par ki.Ki, name ...string) *Line {
	return par.NewChild(LineType, name...).(*Line)
}

// KiType returns the [*gti.Type] of [Line]
func (t *Line) KiType() *gti.Type { return LineType }

// New returns a new [*Line] value
func (t *Line) New() ki.Ki { return &Line{} }

// SetStart sets the [Line.Start]:
// position of the start of the line
func (t *Line) SetStart(v mat32.Vec2) *Line { t.Start = v; return t }

// SetEnd sets the [Line.End]:
// position of the end of the line
func (t *Line) SetEnd(v mat32.Vec2) *Line { t.End = v; return t }

// SetClass sets the [Line.Class]
func (t *Line) SetClass(v string) *Line { t.Class = v; return t }

// MarkerType is the [gti.Type] for [Marker]
var MarkerType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Marker", IDName: "marker", Doc: "Marker represents marker elements that can be drawn along paths (arrow heads, etc)", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "RefPos", Doc: "reference position to align the vertex position with, specified in ViewBox coordinates"}, {Name: "Size", Doc: "size of marker to render, in Units units"}, {Name: "Units", Doc: "units to use"}, {Name: "ViewBox", Doc: "viewbox defines the internal coordinate system for the drawing elements within the marker"}, {Name: "Orient", Doc: "orientation of the marker -- either 'auto' or an angle"}, {Name: "VertexPos", Doc: "current vertex position"}, {Name: "VertexAngle", Doc: "current vertex angle in radians"}, {Name: "StrokeWidth", Doc: "current stroke width"}, {Name: "Transform", Doc: "net transform computed from settings and current values -- applied prior to rendering"}, {Name: "EffSize", Doc: "effective size for actual rendering"}}, Instance: &Marker{}})

// NewMarker adds a new [Marker] with the given name to the given parent:
// Marker represents marker elements that can be drawn along paths (arrow heads, etc)
func NewMarker(par ki.Ki, name ...string) *Marker {
	return par.NewChild(MarkerType, name...).(*Marker)
}

// KiType returns the [*gti.Type] of [Marker]
func (t *Marker) KiType() *gti.Type { return MarkerType }

// New returns a new [*Marker] value
func (t *Marker) New() ki.Ki { return &Marker{} }

// SetRefPos sets the [Marker.RefPos]:
// reference position to align the vertex position with, specified in ViewBox coordinates
func (t *Marker) SetRefPos(v mat32.Vec2) *Marker { t.RefPos = v; return t }

// SetSize sets the [Marker.Size]:
// size of marker to render, in Units units
func (t *Marker) SetSize(v mat32.Vec2) *Marker { t.Size = v; return t }

// SetUnits sets the [Marker.Units]:
// units to use
func (t *Marker) SetUnits(v MarkerUnits) *Marker { t.Units = v; return t }

// SetViewBox sets the [Marker.ViewBox]:
// viewbox defines the internal coordinate system for the drawing elements within the marker
func (t *Marker) SetViewBox(v ViewBox) *Marker { t.ViewBox = v; return t }

// SetOrient sets the [Marker.Orient]:
// orientation of the marker -- either 'auto' or an angle
func (t *Marker) SetOrient(v string) *Marker { t.Orient = v; return t }

// SetVertexPos sets the [Marker.VertexPos]:
// current vertex position
func (t *Marker) SetVertexPos(v mat32.Vec2) *Marker { t.VertexPos = v; return t }

// SetVertexAngle sets the [Marker.VertexAngle]:
// current vertex angle in radians
func (t *Marker) SetVertexAngle(v float32) *Marker { t.VertexAngle = v; return t }

// SetStrokeWidth sets the [Marker.StrokeWidth]:
// current stroke width
func (t *Marker) SetStrokeWidth(v float32) *Marker { t.StrokeWidth = v; return t }

// SetTransform sets the [Marker.Transform]:
// net transform computed from settings and current values -- applied prior to rendering
func (t *Marker) SetTransform(v mat32.Mat2) *Marker { t.Transform = v; return t }

// SetEffSize sets the [Marker.EffSize]:
// effective size for actual rendering
func (t *Marker) SetEffSize(v mat32.Vec2) *Marker { t.EffSize = v; return t }

// SetClass sets the [Marker.Class]
func (t *Marker) SetClass(v string) *Marker { t.Class = v; return t }

// NodeBaseType is the [gti.Type] for [NodeBase]
var NodeBaseType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.NodeBase", IDName: "node-base", Doc: "svg.NodeBase is the base type for elements within the SVG scenegraph", Embeds: []gti.Field{{Name: "Node"}}, Fields: []gti.Field{{Name: "Class", Doc: "user-defined class name(s) used primarily for attaching\nCSS styles to different display elements.\nMultiple class names can be used to combine properties:\nuse spaces to separate per css standard."}, {Name: "CSS", Doc: "cascading style sheet at this level.\nThese styles apply here and to everything below, until superceded.\nUse .class and #name Props elements to apply entire styles\nto given elements, and type for element type."}, {Name: "CSSAgg", Doc: "aggregated css properties from all higher nodes down to me"}, {Name: "BBox", Doc: "bounding box for the node within the SVG Pixels image.\nThis one can be outside the visible range of the SVG image.\nVisBBox is intersected and only shows visible portion."}, {Name: "VisBBox", Doc: "visible bounding box for the node intersected with the SVG image geometry"}, {Name: "Paint", Doc: "paint style information for this node"}}, Instance: &NodeBase{}})

// NewNodeBase adds a new [NodeBase] with the given name to the given parent:
// svg.NodeBase is the base type for elements within the SVG scenegraph
func NewNodeBase(par ki.Ki, name ...string) *NodeBase {
	return par.NewChild(NodeBaseType, name...).(*NodeBase)
}

// KiType returns the [*gti.Type] of [NodeBase]
func (t *NodeBase) KiType() *gti.Type { return NodeBaseType }

// New returns a new [*NodeBase] value
func (t *NodeBase) New() ki.Ki { return &NodeBase{} }

// SetClass sets the [NodeBase.Class]:
// user-defined class name(s) used primarily for attaching
// CSS styles to different display elements.
// Multiple class names can be used to combine properties:
// use spaces to separate per css standard.
func (t *NodeBase) SetClass(v string) *NodeBase { t.Class = v; return t }

// PathType is the [gti.Type] for [Path]
var PathType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Path", IDName: "path", Doc: "Path renders SVG data sequences that can render just about anything", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "Data", Doc: "the path data to render -- path commands and numbers are serialized, with each command specifying the number of floating-point coord data points that follow"}, {Name: "DataStr", Doc: "string version of the path data"}}, Instance: &Path{}})

// NewPath adds a new [Path] with the given name to the given parent:
// Path renders SVG data sequences that can render just about anything
func NewPath(par ki.Ki, name ...string) *Path {
	return par.NewChild(PathType, name...).(*Path)
}

// KiType returns the [*gti.Type] of [Path]
func (t *Path) KiType() *gti.Type { return PathType }

// New returns a new [*Path] value
func (t *Path) New() ki.Ki { return &Path{} }

// SetDataStr sets the [Path.DataStr]:
// string version of the path data
func (t *Path) SetDataStr(v string) *Path { t.DataStr = v; return t }

// SetClass sets the [Path.Class]
func (t *Path) SetClass(v string) *Path { t.Class = v; return t }

// PolygonType is the [gti.Type] for [Polygon]
var PolygonType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Polygon", IDName: "polygon", Doc: "Polygon is a SVG polygon", Embeds: []gti.Field{{Name: "Polyline"}}, Instance: &Polygon{}})

// NewPolygon adds a new [Polygon] with the given name to the given parent:
// Polygon is a SVG polygon
func NewPolygon(par ki.Ki, name ...string) *Polygon {
	return par.NewChild(PolygonType, name...).(*Polygon)
}

// KiType returns the [*gti.Type] of [Polygon]
func (t *Polygon) KiType() *gti.Type { return PolygonType }

// New returns a new [*Polygon] value
func (t *Polygon) New() ki.Ki { return &Polygon{} }

// SetClass sets the [Polygon.Class]
func (t *Polygon) SetClass(v string) *Polygon { t.Class = v; return t }

// SetPoints sets the [Polygon.Points]
func (t *Polygon) SetPoints(v ...mat32.Vec2) *Polygon { t.Points = v; return t }

// PolylineType is the [gti.Type] for [Polyline]
var PolylineType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Polyline", IDName: "polyline", Doc: "Polyline is a SVG multi-line shape", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "Points", Doc: "the coordinates to draw -- does a moveto on the first, then lineto for all the rest"}}, Instance: &Polyline{}})

// NewPolyline adds a new [Polyline] with the given name to the given parent:
// Polyline is a SVG multi-line shape
func NewPolyline(par ki.Ki, name ...string) *Polyline {
	return par.NewChild(PolylineType, name...).(*Polyline)
}

// KiType returns the [*gti.Type] of [Polyline]
func (t *Polyline) KiType() *gti.Type { return PolylineType }

// New returns a new [*Polyline] value
func (t *Polyline) New() ki.Ki { return &Polyline{} }

// SetPoints sets the [Polyline.Points]:
// the coordinates to draw -- does a moveto on the first, then lineto for all the rest
func (t *Polyline) SetPoints(v ...mat32.Vec2) *Polyline { t.Points = v; return t }

// SetClass sets the [Polyline.Class]
func (t *Polyline) SetClass(v string) *Polyline { t.Class = v; return t }

// RectType is the [gti.Type] for [Rect]
var RectType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Rect", IDName: "rect", Doc: "Rect is a SVG rectangle, optionally with rounded corners", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "Pos", Doc: "position of the top-left of the rectangle"}, {Name: "Size", Doc: "size of the rectangle"}, {Name: "Radius", Doc: "radii for curved corners, as a proportion of width, height"}}, Instance: &Rect{}})

// NewRect adds a new [Rect] with the given name to the given parent:
// Rect is a SVG rectangle, optionally with rounded corners
func NewRect(par ki.Ki, name ...string) *Rect {
	return par.NewChild(RectType, name...).(*Rect)
}

// KiType returns the [*gti.Type] of [Rect]
func (t *Rect) KiType() *gti.Type { return RectType }

// New returns a new [*Rect] value
func (t *Rect) New() ki.Ki { return &Rect{} }

// SetPos sets the [Rect.Pos]:
// position of the top-left of the rectangle
func (t *Rect) SetPos(v mat32.Vec2) *Rect { t.Pos = v; return t }

// SetSize sets the [Rect.Size]:
// size of the rectangle
func (t *Rect) SetSize(v mat32.Vec2) *Rect { t.Size = v; return t }

// SetRadius sets the [Rect.Radius]:
// radii for curved corners, as a proportion of width, height
func (t *Rect) SetRadius(v mat32.Vec2) *Rect { t.Radius = v; return t }

// SetClass sets the [Rect.Class]
func (t *Rect) SetClass(v string) *Rect { t.Class = v; return t }

// SVGNodeType is the [gti.Type] for [SVGNode]
var SVGNodeType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.SVGNode", IDName: "svg-node", Doc: "SVGNode represents the root of an SVG tree", Embeds: []gti.Field{{Name: "Group"}}, Fields: []gti.Field{{Name: "ViewBox", Doc: "viewbox defines the coordinate system for the drawing.\nThese units are mapped into the screen space allocated\nfor the SVG during rendering"}}, Instance: &SVGNode{}})

// NewSVGNode adds a new [SVGNode] with the given name to the given parent:
// SVGNode represents the root of an SVG tree
func NewSVGNode(par ki.Ki, name ...string) *SVGNode {
	return par.NewChild(SVGNodeType, name...).(*SVGNode)
}

// KiType returns the [*gti.Type] of [SVGNode]
func (t *SVGNode) KiType() *gti.Type { return SVGNodeType }

// New returns a new [*SVGNode] value
func (t *SVGNode) New() ki.Ki { return &SVGNode{} }

// SetViewBox sets the [SVGNode.ViewBox]:
// viewbox defines the coordinate system for the drawing.
// These units are mapped into the screen space allocated
// for the SVG during rendering
func (t *SVGNode) SetViewBox(v ViewBox) *SVGNode { t.ViewBox = v; return t }

// SetClass sets the [SVGNode.Class]
func (t *SVGNode) SetClass(v string) *SVGNode { t.Class = v; return t }

// TextType is the [gti.Type] for [Text]
var TextType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/svg.Text", IDName: "text", Doc: "Text renders SVG text, handling both text and tspan elements.\ntspan is nested under a parent text -- text has empty Text string.", Embeds: []gti.Field{{Name: "NodeBase"}}, Fields: []gti.Field{{Name: "Pos", Doc: "position of the left, baseline of the text"}, {Name: "Width", Doc: "width of text to render if using word-wrapping"}, {Name: "Text", Doc: "text string to render"}, {Name: "TextRender", Doc: "render version of text"}, {Name: "CharPosX", Doc: "character positions along X axis, if specified"}, {Name: "CharPosY", Doc: "character positions along Y axis, if specified"}, {Name: "CharPosDX", Doc: "character delta-positions along X axis, if specified"}, {Name: "CharPosDY", Doc: "character delta-positions along Y axis, if specified"}, {Name: "CharRots", Doc: "character rotations, if specified"}, {Name: "TextLength", Doc: "author's computed text length, if specified -- we attempt to match"}, {Name: "AdjustGlyphs", Doc: "in attempting to match TextLength, should we adjust glyphs in addition to spacing?"}, {Name: "LastPos", Doc: "last text render position -- lower-left baseline of start"}, {Name: "LastBBox", Doc: "last actual bounding box in display units (dots)"}}, Instance: &Text{}})

// NewText adds a new [Text] with the given name to the given parent:
// Text renders SVG text, handling both text and tspan elements.
// tspan is nested under a parent text -- text has empty Text string.
func NewText(par ki.Ki, name ...string) *Text {
	return par.NewChild(TextType, name...).(*Text)
}

// KiType returns the [*gti.Type] of [Text]
func (t *Text) KiType() *gti.Type { return TextType }

// New returns a new [*Text] value
func (t *Text) New() ki.Ki { return &Text{} }

// SetPos sets the [Text.Pos]:
// position of the left, baseline of the text
func (t *Text) SetPos(v mat32.Vec2) *Text { t.Pos = v; return t }

// SetWidth sets the [Text.Width]:
// width of text to render if using word-wrapping
func (t *Text) SetWidth(v float32) *Text { t.Width = v; return t }

// SetText sets the [Text.Text]:
// text string to render
func (t *Text) SetText(v string) *Text { t.Text = v; return t }

// SetTextRender sets the [Text.TextRender]:
// render version of text
func (t *Text) SetTextRender(v paint.Text) *Text { t.TextRender = v; return t }

// SetCharPosX sets the [Text.CharPosX]:
// character positions along X axis, if specified
func (t *Text) SetCharPosX(v ...float32) *Text { t.CharPosX = v; return t }

// SetCharPosY sets the [Text.CharPosY]:
// character positions along Y axis, if specified
func (t *Text) SetCharPosY(v ...float32) *Text { t.CharPosY = v; return t }

// SetCharPosDX sets the [Text.CharPosDX]:
// character delta-positions along X axis, if specified
func (t *Text) SetCharPosDX(v ...float32) *Text { t.CharPosDX = v; return t }

// SetCharPosDY sets the [Text.CharPosDY]:
// character delta-positions along Y axis, if specified
func (t *Text) SetCharPosDY(v ...float32) *Text { t.CharPosDY = v; return t }

// SetCharRots sets the [Text.CharRots]:
// character rotations, if specified
func (t *Text) SetCharRots(v ...float32) *Text { t.CharRots = v; return t }

// SetTextLength sets the [Text.TextLength]:
// author's computed text length, if specified -- we attempt to match
func (t *Text) SetTextLength(v float32) *Text { t.TextLength = v; return t }

// SetAdjustGlyphs sets the [Text.AdjustGlyphs]:
// in attempting to match TextLength, should we adjust glyphs in addition to spacing?
func (t *Text) SetAdjustGlyphs(v bool) *Text { t.AdjustGlyphs = v; return t }

// SetLastPos sets the [Text.LastPos]:
// last text render position -- lower-left baseline of start
func (t *Text) SetLastPos(v mat32.Vec2) *Text { t.LastPos = v; return t }

// SetLastBBox sets the [Text.LastBBox]:
// last actual bounding box in display units (dots)
func (t *Text) SetLastBBox(v mat32.Box2) *Text { t.LastBBox = v; return t }

// SetClass sets the [Text.Class]
func (t *Text) SetClass(v string) *Text { t.Class = v; return t }
