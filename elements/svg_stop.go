// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg stop is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <stop> SVG element defines the ramp of colors to use on a gradient, which
// is a child element to either the <linearGradient> or the <radialGradient>
// element.
type SVGSTOPElement struct {
	*Element
}

// Create a new SVGSTOPElement element.
// This will create a new element with the tag
// "stop" during rendering.
func SVG_STOP(children ...ElementRenderer) *SVGSTOPElement {
	e := NewElement("stop", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSTOPElement{Element: e}
}

func (e *SVGSTOPElement) Children(children ...ElementRenderer) *SVGSTOPElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSTOPElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSTOPElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSTOPElement) Attr(name string, value string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGSTOPElement) Attrs(attrs ...string) *SVGSTOPElement {
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

func (e *SVGSTOPElement) AttrsMap(attrs map[string]string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSTOPElement) Text(text string) *SVGSTOPElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSTOPElement) TextF(format string, args ...any) *SVGSTOPElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSTOPElement) IfText(condition bool, text string) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSTOPElement) IfTextF(condition bool, format string, args ...any) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSTOPElement) Escaped(text string) *SVGSTOPElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSTOPElement) IfEscaped(condition bool, text string) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSTOPElement) EscapedF(format string, args ...any) *SVGSTOPElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSTOPElement) IfEscapedF(condition bool, format string, args ...any) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSTOPElement) CustomData(key, value string) *SVGSTOPElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSTOPElement) IfCustomData(condition bool, key, value string) *SVGSTOPElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSTOPElement) CustomDataF(key, format string, args ...any) *SVGSTOPElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSTOPElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSTOPElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSTOPElement) CustomDataRemove(key string) *SVGSTOPElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The offset from the start of the gradient where the color first takes effect.
func (e *SVGSTOPElement) OFFSET(f float64) *SVGSTOPElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("offset", f)
	return e
}

func (e *SVGSTOPElement) IfOFFSET(condition bool, f float64) *SVGSTOPElement {
	if condition {
		e.OFFSET(f)
	}
	return e
}

// The color of the gradient stop.
func (e *SVGSTOPElement) STOP_COLOR(s string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("stop-color", s)
	return e
}

func (e *SVGSTOPElement) IfSTOP_COLOR(condition bool, s string) *SVGSTOPElement {
	if condition {
		e.STOP_COLOR(s)
	}
	return e
}

// Remove the attribute STOP_COLOR from the element.
func (e *SVGSTOPElement) STOP_COLORRemove(s string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("stop-color")
	return e
}

// Specifies a unique id for an element
func (e *SVGSTOPElement) ID(s string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSTOPElement) IfID(condition bool, s string) *SVGSTOPElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGSTOPElement) IDRemove(s string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSTOPElement) CLASS(s ...string) *SVGSTOPElement {
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

func (e *SVGSTOPElement) IfCLASS(condition bool, s ...string) *SVGSTOPElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGSTOPElement) CLASSRemove(s ...string) *SVGSTOPElement {
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
func (e *SVGSTOPElement) STYLEF(k string, format string, args ...any) *SVGSTOPElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSTOPElement) IfSTYLE(condition bool, k string, v string) *SVGSTOPElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSTOPElement) STYLE(k string, v string) *SVGSTOPElement {
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

func (e *SVGSTOPElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSTOPElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSTOPElement) STYLEMap(m map[string]string) *SVGSTOPElement {
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
func (e *SVGSTOPElement) STYLEPairs(pairs ...string) *SVGSTOPElement {
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

func (e *SVGSTOPElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSTOPElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGSTOPElement) STYLERemove(keys ...string) *SVGSTOPElement {
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

// Merges the store with the given object

func (e *SVGSTOPElement) DATASTAR_MERGE_STORE(v any) *SVGSTOPElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("data-merge-store", string(b))
	return e
}

// Sets the reference of the element

func (e *SVGSTOPElement) DATASTAR_REF(expression string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTOPElement) IfDATASTAR_REF(condition bool, expression string) *SVGSTOPElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGSTOPElement) DATASTAR_REFRemove() *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGSTOPElement) DATASTAR_BIND(key string, expression string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTOPElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGSTOPElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGSTOPElement) DATASTAR_BINDRemove() *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGSTOPElement) DATASTAR_MODEL(expression string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTOPElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGSTOPElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGSTOPElement) DATASTAR_MODELRemove() *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGSTOPElement) DATASTAR_TEXT(expression string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTOPElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGSTOPElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGSTOPElement) DATASTAR_TEXTRemove() *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGStopDataOnMod customDataKeyModifier

// Debounces the event handler
func SVGStopDataOnModDebounce(
	d time.Duration,
) SVGStopDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGStopDataOnModThrottle(
	d time.Duration,
) SVGStopDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGSTOPElement) DATASTAR_ON(key string, expression string, modifiers ...SVGStopDataOnMod) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGStopDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTOPElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGStopDataOnMod) *SVGSTOPElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGSTOPElement) DATASTAR_ONRemove() *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGSTOPElement) DATASTAR_FOCUSSet(b bool) *SVGSTOPElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSTOPElement) DATASTAR_FOCUS() *SVGSTOPElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGSTOPElement) DATASTAR_HEADER(key string, expression string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTOPElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGSTOPElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGSTOPElement) DATASTAR_HEADERRemove() *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGSTOPElement) DATASTAR_FETCH_URL(expression string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTOPElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGSTOPElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGSTOPElement) DATASTAR_FETCH_URLRemove() *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGSTOPElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTOPElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGSTOPElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGSTOPElement) DATASTAR_FETCH_INDICATORRemove() *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SVGSTOPElement) DATASTAR_SHOWSet(b bool) *SVGSTOPElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSTOPElement) DATASTAR_SHOW() *SVGSTOPElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGSTOPElement) DATASTAR_INTERSECTS(expression string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTOPElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGSTOPElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGSTOPElement) DATASTAR_INTERSECTSRemove() *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGSTOPElement) DATASTAR_TELEPORTSet(b bool) *SVGSTOPElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSTOPElement) DATASTAR_TELEPORT() *SVGSTOPElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGSTOPElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGSTOPElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSTOPElement) DATASTAR_SCROLL_INTO_VIEW() *SVGSTOPElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGSTOPElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-view-transition-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTOPElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *SVGSTOPElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGSTOPElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
