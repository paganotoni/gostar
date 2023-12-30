// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg text is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <text> SVG element renders the first character at the initial current text
// position
// This position is modified by the lengthAdjust and textLength attributes.
type SVGTEXTElement struct {
	*Element
}

// Create a new SVGTEXTElement element.
// This will create a new element with the tag
// "text" during rendering.
func SVG_TEXT(children ...ElementRenderer) *SVGTEXTElement {
	e := NewElement("text", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGTEXTElement{Element: e}
}

func (e *SVGTEXTElement) Children(children ...ElementRenderer) *SVGTEXTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGTEXTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGTEXTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGTEXTElement) Attr(name string, value string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGTEXTElement) Attrs(attrs ...string) *SVGTEXTElement {
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

func (e *SVGTEXTElement) AttrsMap(attrs map[string]string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGTEXTElement) Text(text string) *SVGTEXTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGTEXTElement) TextF(format string, args ...any) *SVGTEXTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGTEXTElement) IfText(condition bool, text string) *SVGTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGTEXTElement) IfTextF(condition bool, format string, args ...any) *SVGTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGTEXTElement) Escaped(text string) *SVGTEXTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGTEXTElement) IfEscaped(condition bool, text string) *SVGTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGTEXTElement) EscapedF(format string, args ...any) *SVGTEXTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGTEXTElement) IfEscapedF(condition bool, format string, args ...any) *SVGTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGTEXTElement) CustomData(key, value string) *SVGTEXTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGTEXTElement) IfCustomData(condition bool, key, value string) *SVGTEXTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGTEXTElement) CustomDataF(key, format string, args ...any) *SVGTEXTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGTEXTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGTEXTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGTEXTElement) CustomDataRemove(key string) *SVGTEXTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The x-axis coordinate of the initial current text position.
func (e *SVGTEXTElement) X(f float64) *SVGTEXTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("x", f)
	return e
}

func (e *SVGTEXTElement) IfX(condition bool, f float64) *SVGTEXTElement {
	if condition {
		e.X(f)
	}
	return e
}

// The y-axis coordinate of the initial current text position.
func (e *SVGTEXTElement) Y(f float64) *SVGTEXTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("y", f)
	return e
}

func (e *SVGTEXTElement) IfY(condition bool, f float64) *SVGTEXTElement {
	if condition {
		e.Y(f)
	}
	return e
}

// The x-axis coordinate of the current text position.
func (e *SVGTEXTElement) DX(f float64) *SVGTEXTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("dx", f)
	return e
}

func (e *SVGTEXTElement) IfDX(condition bool, f float64) *SVGTEXTElement {
	if condition {
		e.DX(f)
	}
	return e
}

// The y-axis coordinate of the current text position.
func (e *SVGTEXTElement) DY(f float64) *SVGTEXTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("dy", f)
	return e
}

func (e *SVGTEXTElement) IfDY(condition bool, f float64) *SVGTEXTElement {
	if condition {
		e.DY(f)
	}
	return e
}

// The rotation angle about the current text position.
func (e *SVGTEXTElement) ROTATE(f float64) *SVGTEXTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("rotate", f)
	return e
}

func (e *SVGTEXTElement) IfROTATE(condition bool, f float64) *SVGTEXTElement {
	if condition {
		e.ROTATE(f)
	}
	return e
}

// The total sum of all of the advance values from rendering all of the characters
// within this element, including the advance value on the glyph (horizontal or
// vertical), the effect of properties 'kerning', 'letter-spacing' and
// 'word-spacing' and adjustments due to attributes 'x' and 'y' on the <text>
// element.
func (e *SVGTEXTElement) TEXT_LENGTH(f float64) *SVGTEXTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("textLength", f)
	return e
}

func (e *SVGTEXTElement) IfTEXT_LENGTH(condition bool, f float64) *SVGTEXTElement {
	if condition {
		e.TEXT_LENGTH(f)
	}
	return e
}

// Indicates how the text is stretched or compressed to fit the textLength.
func (e *SVGTEXTElement) LENGTH_ADJUST(c SVGTextLengthAdjustChoice) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("lengthAdjust", string(c))
	return e
}

type SVGTextLengthAdjustChoice string

const (
	// Indicates how the text is stretched or compressed to fit the textLength.
	SVGTextLengthAdjust_spacing SVGTextLengthAdjustChoice = "spacing"
	// Indicates how the text is stretched or compressed to fit the textLength.
	SVGTextLengthAdjust_spacingAndGlyphs SVGTextLengthAdjustChoice = "spacingAndGlyphs"
)

// Remove the attribute LENGTH_ADJUST from the element.
func (e *SVGTEXTElement) LENGTH_ADJUSTRemove(c SVGTextLengthAdjustChoice) *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("lengthAdjust")
	return e
}

// Specifies a unique id for an element
func (e *SVGTEXTElement) ID(s string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGTEXTElement) IfID(condition bool, s string) *SVGTEXTElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGTEXTElement) IDRemove(s string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGTEXTElement) CLASS(s ...string) *SVGTEXTElement {
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

func (e *SVGTEXTElement) IfCLASS(condition bool, s ...string) *SVGTEXTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGTEXTElement) CLASSRemove(s ...string) *SVGTEXTElement {
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
func (e *SVGTEXTElement) STYLEF(k string, format string, args ...any) *SVGTEXTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGTEXTElement) IfSTYLE(condition bool, k string, v string) *SVGTEXTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGTEXTElement) STYLE(k string, v string) *SVGTEXTElement {
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

func (e *SVGTEXTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGTEXTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGTEXTElement) STYLEMap(m map[string]string) *SVGTEXTElement {
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
func (e *SVGTEXTElement) STYLEPairs(pairs ...string) *SVGTEXTElement {
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

func (e *SVGTEXTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGTEXTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGTEXTElement) STYLERemove(keys ...string) *SVGTEXTElement {
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

func (e *SVGTEXTElement) DATASTAR_MERGE_STORE(v any) *SVGTEXTElement {
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

func (e *SVGTEXTElement) DATASTAR_REF(expression string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTElement) IfDATASTAR_REF(condition bool, expression string) *SVGTEXTElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGTEXTElement) DATASTAR_REFRemove() *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGTEXTElement) DATASTAR_BIND(key string, expression string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGTEXTElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGTEXTElement) DATASTAR_BINDRemove() *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGTEXTElement) DATASTAR_MODEL(expression string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGTEXTElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGTEXTElement) DATASTAR_MODELRemove() *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGTEXTElement) DATASTAR_TEXT(expression string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGTEXTElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGTEXTElement) DATASTAR_TEXTRemove() *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGTextDataOnMod customDataKeyModifier

// Debounces the event handler
func SVGTextDataOnModDebounce(
	d time.Duration,
) SVGTextDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGTextDataOnModThrottle(
	d time.Duration,
) SVGTextDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGTEXTElement) DATASTAR_ON(key string, expression string, modifiers ...SVGTextDataOnMod) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGTextDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGTextDataOnMod) *SVGTEXTElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGTEXTElement) DATASTAR_ONRemove() *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGTEXTElement) DATASTAR_FOCUSSet(b bool) *SVGTEXTElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGTEXTElement) DATASTAR_FOCUS() *SVGTEXTElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGTEXTElement) DATASTAR_HEADER(key string, expression string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGTEXTElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGTEXTElement) DATASTAR_HEADERRemove() *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGTEXTElement) DATASTAR_FETCH_URL(expression string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGTEXTElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGTEXTElement) DATASTAR_FETCH_URLRemove() *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGTEXTElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGTEXTElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGTEXTElement) DATASTAR_FETCH_INDICATORRemove() *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SVGTEXTElement) DATASTAR_SHOWSet(b bool) *SVGTEXTElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGTEXTElement) DATASTAR_SHOW() *SVGTEXTElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGTEXTElement) DATASTAR_INTERSECTS(expression string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGTEXTElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGTEXTElement) DATASTAR_INTERSECTSRemove() *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGTEXTElement) DATASTAR_TELEPORTSet(b bool) *SVGTEXTElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGTEXTElement) DATASTAR_TELEPORT() *SVGTEXTElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGTEXTElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGTEXTElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGTEXTElement) DATASTAR_SCROLL_INTO_VIEW() *SVGTEXTElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGTEXTElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *SVGTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-view-transition-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *SVGTEXTElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGTEXTElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
