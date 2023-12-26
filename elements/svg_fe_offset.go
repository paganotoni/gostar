// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feOffset is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <feOffset> SVG filter primitive allows to offset the input image
// The amount of offset can be controlled by attributes dx and dy.
type SVGFEOFFSETElement struct {
	*Element
}

// Create a new SVGFEOFFSETElement element.
// This will create a new element with the tag
// "feOffset" during rendering.
func SVG_FEOFFSET(children ...ElementRenderer) *SVGFEOFFSETElement {
	e := NewElement("feOffset", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEOFFSETElement{Element: e}
}

func (e *SVGFEOFFSETElement) Children(children ...ElementRenderer) *SVGFEOFFSETElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEOFFSETElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEOFFSETElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEOFFSETElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEOFFSETElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEOFFSETElement) Text(text string) *SVGFEOFFSETElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEOFFSETElement) TextF(format string, args ...any) *SVGFEOFFSETElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEOFFSETElement) Escaped(text string) *SVGFEOFFSETElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEOFFSETElement) EscapedF(format string, args ...any) *SVGFEOFFSETElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEOFFSETElement) CustomData(key, value string) *SVGFEOFFSETElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEOFFSETElement) CustomDataRemove(key string) *SVGFEOFFSETElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The dx attribute indicates a shift along the x-axis on the kernel matrix.
func (e *SVGFEOFFSETElement) DX(f float64) *SVGFEOFFSETElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("dx", f)
	return e
}

// The dy attribute indicates a shift along the y-axis on the kernel matrix.
func (e *SVGFEOFFSETElement) DY(f float64) *SVGFEOFFSETElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("dy", f)
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEOFFSETElement) CLASS(s ...string) *SVGFEOFFSETElement {
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
func (e *SVGFEOFFSETElement) CLASSRemove(s ...string) *SVGFEOFFSETElement {
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
func (e *SVGFEOFFSETElement) ID(s string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *SVGFEOFFSETElement) IDRemove(s string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGFEOFFSETElement) STYLE(k string, v string) *SVGFEOFFSETElement {
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
func (e *SVGFEOFFSETElement) STYLEMap(m map[string]string) *SVGFEOFFSETElement {
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
func (e *SVGFEOFFSETElement) STYLEPairs(pairs ...string) *SVGFEOFFSETElement {
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
func (e *SVGFEOFFSETElement) STYLERemove(keys ...string) *SVGFEOFFSETElement {
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
