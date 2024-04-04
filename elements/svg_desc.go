// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg desc is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"html"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <desc> SVG element provides a description container for SVG content.
type SVGDESCElement struct {
	*Element
}

// Create a new SVGDESCElement element.
// This will create a new element with the tag
// "desc" during rendering.
func SVG_DESC(children ...ElementRenderer) *SVGDESCElement {
	e := NewElement("desc", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGDESCElement{Element: e}
}

func (e *SVGDESCElement) Children(children ...ElementRenderer) *SVGDESCElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGDESCElement) IfChildren(condition bool, children ...ElementRenderer) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGDESCElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGDESCElement) Attr(name string, value string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGDESCElement) Attrs(attrs ...string) *SVGDESCElement {
	if len(attrs)%2 != 0 {
		panic("attrs must be a multiple of 2")
	}
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for i := 0; i < len(attrs); i += 2 {
		k := attrs[i]
		v := attrs[i+1]
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGDESCElement) AttrsMap(attrs map[string]string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGDESCElement) Text(text string) *SVGDESCElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGDESCElement) TextF(format string, args ...any) *SVGDESCElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGDESCElement) IfText(condition bool, text string) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGDESCElement) IfTextF(condition bool, format string, args ...any) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGDESCElement) Escaped(text string) *SVGDESCElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGDESCElement) IfEscaped(condition bool, text string) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGDESCElement) EscapedF(format string, args ...any) *SVGDESCElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGDESCElement) IfEscapedF(condition bool, format string, args ...any) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGDESCElement) CustomData(key, value string) *SVGDESCElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGDESCElement) IfCustomData(condition bool, key, value string) *SVGDESCElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGDESCElement) CustomDataF(key, format string, args ...any) *SVGDESCElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGDESCElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGDESCElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGDESCElement) CustomDataRemove(key string) *SVGDESCElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Specifies a unique id for an element
func (e *SVGDESCElement) ID(s string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGDESCElement) IDF(format string, args ...any) *SVGDESCElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGDESCElement) IfID(condition bool, s string) *SVGDESCElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGDESCElement) IfIDF(condition bool, format string, args ...any) *SVGDESCElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGDESCElement) IDRemove(s string) *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGDESCElement) IDRemoveF(format string, args ...any) *SVGDESCElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGDESCElement) CLASS(s ...string) *SVGDESCElement {
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

func (e *SVGDESCElement) IfCLASS(condition bool, s ...string) *SVGDESCElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGDESCElement) CLASSRemove(s ...string) *SVGDESCElement {
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

// Specifies an inline CSS style for an element
func (e *SVGDESCElement) STYLEF(k string, format string, args ...any) *SVGDESCElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGDESCElement) IfSTYLE(condition bool, k string, v string) *SVGDESCElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGDESCElement) STYLE(k string, v string) *SVGDESCElement {
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

func (e *SVGDESCElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGDESCElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGDESCElement) STYLEMap(m map[string]string) *SVGDESCElement {
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
func (e *SVGDESCElement) STYLEPairs(pairs ...string) *SVGDESCElement {
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

func (e *SVGDESCElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGDESCElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGDESCElement) STYLERemove(keys ...string) *SVGDESCElement {
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

// Merges the singleton store with the given object

func (e *SVGDESCElement) DATASTAR_STORE(v any) *SVGDESCElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("store", html.EscapeString(string(b)))
	return e
}

// Sets the reference of the element

func (e *SVGDESCElement) DATASTAR_REF(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfDATASTAR_REF(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGDESCElement) DATASTAR_REFRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGDESCElement) DATASTAR_BIND(key string, expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGDESCElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGDESCElement) DATASTAR_BINDRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGDESCElement) DATASTAR_MODEL(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGDESCElement) DATASTAR_MODELRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGDESCElement) DATASTAR_TEXT(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGDESCElement) DATASTAR_TEXTRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGDescOnMod customDataKeyModifier

// Debounces the event handler
func SVGDescOnModDebounce(
	d time.Duration,
) SVGDescOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGDescOnModThrottle(
	d time.Duration,
) SVGDescOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGDESCElement) DATASTAR_ON(key string, expression string, modifiers ...SVGDescOnMod) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGDescOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGDescOnMod) *SVGDESCElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGDESCElement) DATASTAR_ONRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGDESCElement) DATASTAR_FOCUSSet(b bool) *SVGDESCElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGDESCElement) DATASTAR_FOCUS() *SVGDESCElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGDESCElement) DATASTAR_HEADER(key string, expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGDESCElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGDESCElement) DATASTAR_HEADERRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGDESCElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGDESCElement) DATASTAR_FETCH_INDICATORRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGDESCElement) DATASTAR_SHOWSet(b bool) *SVGDESCElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGDESCElement) DATASTAR_SHOW() *SVGDESCElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGDESCElement) DATASTAR_INTERSECTS(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGDESCElement) DATASTAR_INTERSECTSRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGDESCElement) DATASTAR_TELEPORTSet(b bool) *SVGDESCElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGDESCElement) DATASTAR_TELEPORT() *SVGDESCElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGDESCElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGDESCElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGDESCElement) DATASTAR_SCROLL_INTO_VIEW() *SVGDESCElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGDESCElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGDESCElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
