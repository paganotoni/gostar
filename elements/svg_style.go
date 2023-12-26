// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg style is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <style> SVG element allows style sheets to be embedded directly within SVG
// content
// SVG's style element has the same attributes as the corresponding element in
// HTML (see HTML's <style> element).
type SVGSTYLEElement struct {
	*Element
}

// Create a new SVGSTYLEElement element.
// This will create a new element with the tag
// "style" during rendering.
func SVG_STYLE(children ...ElementRenderer) *SVGSTYLEElement {
	e := NewElement("style", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSTYLEElement{Element: e}
}

func (e *SVGSTYLEElement) Children(children ...ElementRenderer) *SVGSTYLEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSTYLEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSTYLEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSTYLEElement) Text(text string) *SVGSTYLEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSTYLEElement) TextF(format string, args ...any) *SVGSTYLEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) Escaped(text string) *SVGSTYLEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSTYLEElement) EscapedF(format string, args ...any) *SVGSTYLEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) CustomData(key, value string) *SVGSTYLEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSTYLEElement) CustomDataRemove(key string) *SVGSTYLEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The style sheet language.
func (e *SVGSTYLEElement) TYPE(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", s)
	return e
}

// Remove the attribute type from the element.
func (e *SVGSTYLEElement) TYPERemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

// The intended destination medium for style information.
func (e *SVGSTYLEElement) MEDIA(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("media", s)
	return e
}

// Remove the attribute media from the element.
func (e *SVGSTYLEElement) MEDIARemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("media")
	return e
}

// The advisory title.
func (e *SVGSTYLEElement) TITLE(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("title", s)
	return e
}

// Remove the attribute title from the element.
func (e *SVGSTYLEElement) TITLERemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("title")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSTYLEElement) CLASS(s ...string) *SVGSTYLEElement {
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
func (e *SVGSTYLEElement) CLASSRemove(s ...string) *SVGSTYLEElement {
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
func (e *SVGSTYLEElement) ID(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *SVGSTYLEElement) IDRemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGSTYLEElement) STYLE(k string, v string) *SVGSTYLEElement {
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
func (e *SVGSTYLEElement) STYLEMap(m map[string]string) *SVGSTYLEElement {
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
func (e *SVGSTYLEElement) STYLEPairs(pairs ...string) *SVGSTYLEElement {
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
func (e *SVGSTYLEElement) STYLERemove(keys ...string) *SVGSTYLEElement {
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
