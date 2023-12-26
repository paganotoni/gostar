// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feDropShadow is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <feDropShadow> filter primitive creates a drop shadow of the input image
// It is a shorthand filter, and is defined in terms of the <feGaussianBlur> and
// <feOffset> filter primitives.
type SVGFEDROPSHADOWElement struct {
	*Element
}

// Create a new SVGFEDROPSHADOWElement element.
// This will create a new element with the tag
// "feDropShadow" during rendering.
func SVG_FEDROPSHADOW(children ...ElementRenderer) *SVGFEDROPSHADOWElement {
	e := NewElement("feDropShadow", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEDROPSHADOWElement{Element: e}
}

func (e *SVGFEDROPSHADOWElement) Children(children ...ElementRenderer) *SVGFEDROPSHADOWElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) Text(text string) *SVGFEDROPSHADOWElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEDROPSHADOWElement) TextF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) Escaped(text string) *SVGFEDROPSHADOWElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEDROPSHADOWElement) EscapedF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) CustomData(key, value string) *SVGFEDROPSHADOWElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEDROPSHADOWElement) CustomDataRemove(key string) *SVGFEDROPSHADOWElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The amount of offset in the x direction
// If the <length> is 0, the shadow is placed at the same position as the input.
func (e *SVGFEDROPSHADOWElement) DX(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("dx", f)
	return e
}

// The amount of offset in the y direction
// If the <length> is 0, the shadow is placed at the same position as the input.
func (e *SVGFEDROPSHADOWElement) DY(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("dy", f)
	return e
}

// The standard deviation for the blur operation
// If two <numbers> are provided, the first number represents a standard deviation
// value along the x-axis of the coordinate system established by attribute
// 'primitiveUnits' on the <filter> element
// The second value represents a standard deviation in Y
// If one number is provided, then that value is used for both X and Y
// Negative values are not allowed
// A value of zero disables the effect of the given filter primitive (i.e., the
// result is a transparent black image).
func (e *SVGFEDROPSHADOWElement) STDDEVIATION(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("stdDeviation", f)
	return e
}

// The flood-color attribute indicates what color to use to flood the current
// filter primitive subregion defined through the <feFlood> element
// If attribute 'flood-color' is not specified, then the effect is as if a value
// of black were specified.
func (e *SVGFEDROPSHADOWElement) FLOOD_COLOR(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("flood-color", s)
	return e
}

// Remove the attribute flood-color from the element.
func (e *SVGFEDROPSHADOWElement) FLOOD_COLORRemove(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("flood-color")
	return e
}

// The flood-opacity attribute indicates the opacity value to use across the
// current filter primitive subregion defined through the <feFlood> element.
func (e *SVGFEDROPSHADOWElement) FLOOD_OPACITY(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("flood-opacity", f)
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEDROPSHADOWElement) CLASS(s ...string) *SVGFEDROPSHADOWElement {
	if e.DelimitedStrings == nil {
		e.DelimitedStrings = treemap.New[string, *DelimitedBuilder[string]]()
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		ds = NewDelimitedBuilder[string](" ")
		e.DelimitedStrings.Set("class", ds)
	}
	ds.Add(s...)
	return e
}

// Remove the attribute class from the element.
func (e *SVGFEDROPSHADOWElement) CLASSRemove(s ...string) *SVGFEDROPSHADOWElement {
	if e.DelimitedStrings == nil {
		return e
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		return e
	}
	ds.Remove(s...)
	return e
}

// Specifies a unique id for an element
func (e *SVGFEDROPSHADOWElement) ID(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *SVGFEDROPSHADOWElement) IDRemove(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGFEDROPSHADOWElement) STYLE(k string, v string) *SVGFEDROPSHADOWElement {
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	kv.Add(k, v)
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEDROPSHADOWElement) STYLEMap(m map[string]string) *SVGFEDROPSHADOWElement {
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	for k, v := range m {
		kv.Add(k, v)
	}
	return e
}

// Add pairs of attributes to the element.
func (e *SVGFEDROPSHADOWElement) STYLEPairs(pairs ...string) *SVGFEDROPSHADOWElement {
	if len(pairs)%2 != 0 {
		panic("Must have an even number of pairs")
	}
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}

	for i := 0; i < len(pairs); i += 2 {
		kv.Add(pairs[i], pairs[i+1])
	}

	return e
}

// Remove the attribute style from the element.
func (e *SVGFEDROPSHADOWElement) STYLERemove(keys ...string) *SVGFEDROPSHADOWElement {
	if e.KVStrings == nil {
		return e
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		return e
	}
	for _, k := range keys {
		kv.Remove(k)
	}
	return e
}
