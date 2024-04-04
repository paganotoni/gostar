// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg script is generated from configuration file.
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

// The <script> SVG element includes scripts, which can be used to trigger user
// interface events.
type SVGSCRIPTElement struct {
	*Element
}

// Create a new SVGSCRIPTElement element.
// This will create a new element with the tag
// "script" during rendering.
func SVG_SCRIPT(children ...ElementRenderer) *SVGSCRIPTElement {
	e := NewElement("script", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSCRIPTElement{Element: e}
}

func (e *SVGSCRIPTElement) Children(children ...ElementRenderer) *SVGSCRIPTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSCRIPTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSCRIPTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSCRIPTElement) Attr(name string, value string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGSCRIPTElement) Attrs(attrs ...string) *SVGSCRIPTElement {
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

func (e *SVGSCRIPTElement) AttrsMap(attrs map[string]string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSCRIPTElement) Text(text string) *SVGSCRIPTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSCRIPTElement) TextF(format string, args ...any) *SVGSCRIPTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfText(condition bool, text string) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSCRIPTElement) IfTextF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSCRIPTElement) Escaped(text string) *SVGSCRIPTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSCRIPTElement) IfEscaped(condition bool, text string) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSCRIPTElement) EscapedF(format string, args ...any) *SVGSCRIPTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfEscapedF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSCRIPTElement) CustomData(key, value string) *SVGSCRIPTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSCRIPTElement) IfCustomData(condition bool, key, value string) *SVGSCRIPTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSCRIPTElement) CustomDataF(key, format string, args ...any) *SVGSCRIPTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSCRIPTElement) CustomDataRemove(key string) *SVGSCRIPTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The scripting language used for the given script element.
func (e *SVGSCRIPTElement) TYPE(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", s)
	return e
}

func (e *SVGSCRIPTElement) TYPEF(format string, args ...any) *SVGSCRIPTElement {
	return e.TYPE(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfTYPE(condition bool, s string) *SVGSCRIPTElement {
	if condition {
		e.TYPE(s)
	}
	return e
}

func (e *SVGSCRIPTElement) IfTYPEF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.TYPE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TYPE from the element.
func (e *SVGSCRIPTElement) TYPERemove(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

func (e *SVGSCRIPTElement) TYPERemoveF(format string, args ...any) *SVGSCRIPTElement {
	return e.TYPERemove(fmt.Sprintf(format, args...))
}

// A Uniform Resource Identifier (URI) reference that specifies the location of
// the script to execute.
func (e *SVGSCRIPTElement) HREF(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGSCRIPTElement) HREFF(format string, args ...any) *SVGSCRIPTElement {
	return e.HREF(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfHREF(condition bool, s string) *SVGSCRIPTElement {
	if condition {
		e.HREF(s)
	}
	return e
}

func (e *SVGSCRIPTElement) IfHREFF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.HREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGSCRIPTElement) HREFRemove(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("href")
	return e
}

func (e *SVGSCRIPTElement) HREFRemoveF(format string, args ...any) *SVGSCRIPTElement {
	return e.HREFRemove(fmt.Sprintf(format, args...))
}

// How the element handles crossorigin requests.
func (e *SVGSCRIPTElement) CROSSORIGIN(c SVGScriptCrossoriginChoice) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("crossorigin", string(c))
	return e
}

type SVGScriptCrossoriginChoice string

const (
	// How the element handles crossorigin requests.
	SVGScriptCrossorigin_anonymous SVGScriptCrossoriginChoice = "anonymous"
	// How the element handles crossorigin requests.
	SVGScriptCrossorigin_use_credentials SVGScriptCrossoriginChoice = "use-credentials"
)

// Remove the attribute CROSSORIGIN from the element.
func (e *SVGSCRIPTElement) CROSSORIGINRemove(c SVGScriptCrossoriginChoice) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("crossorigin")
	return e
}

// Specifies a unique id for an element
func (e *SVGSCRIPTElement) ID(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSCRIPTElement) IDF(format string, args ...any) *SVGSCRIPTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfID(condition bool, s string) *SVGSCRIPTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGSCRIPTElement) IfIDF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGSCRIPTElement) IDRemove(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGSCRIPTElement) IDRemoveF(format string, args ...any) *SVGSCRIPTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSCRIPTElement) CLASS(s ...string) *SVGSCRIPTElement {
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

func (e *SVGSCRIPTElement) IfCLASS(condition bool, s ...string) *SVGSCRIPTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGSCRIPTElement) CLASSRemove(s ...string) *SVGSCRIPTElement {
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
func (e *SVGSCRIPTElement) STYLEF(k string, format string, args ...any) *SVGSCRIPTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfSTYLE(condition bool, k string, v string) *SVGSCRIPTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSCRIPTElement) STYLE(k string, v string) *SVGSCRIPTElement {
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

func (e *SVGSCRIPTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSCRIPTElement) STYLEMap(m map[string]string) *SVGSCRIPTElement {
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
func (e *SVGSCRIPTElement) STYLEPairs(pairs ...string) *SVGSCRIPTElement {
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

func (e *SVGSCRIPTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSCRIPTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGSCRIPTElement) STYLERemove(keys ...string) *SVGSCRIPTElement {
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

func (e *SVGSCRIPTElement) DATASTAR_STORE(v any) *SVGSCRIPTElement {
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

func (e *SVGSCRIPTElement) DATASTAR_REF(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfDATASTAR_REF(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGSCRIPTElement) DATASTAR_REFRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGSCRIPTElement) DATASTAR_BIND(key string, expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGSCRIPTElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGSCRIPTElement) DATASTAR_BINDRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGSCRIPTElement) DATASTAR_MODEL(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGSCRIPTElement) DATASTAR_MODELRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGSCRIPTElement) DATASTAR_TEXT(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGSCRIPTElement) DATASTAR_TEXTRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGScriptOnMod customDataKeyModifier

// Debounces the event handler
func SVGScriptOnModDebounce(
	d time.Duration,
) SVGScriptOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGScriptOnModThrottle(
	d time.Duration,
) SVGScriptOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGSCRIPTElement) DATASTAR_ON(key string, expression string, modifiers ...SVGScriptOnMod) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGScriptOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGScriptOnMod) *SVGSCRIPTElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGSCRIPTElement) DATASTAR_ONRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGSCRIPTElement) DATASTAR_FOCUSSet(b bool) *SVGSCRIPTElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSCRIPTElement) DATASTAR_FOCUS() *SVGSCRIPTElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGSCRIPTElement) DATASTAR_HEADER(key string, expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGSCRIPTElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGSCRIPTElement) DATASTAR_HEADERRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGSCRIPTElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGSCRIPTElement) DATASTAR_FETCH_INDICATORRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGSCRIPTElement) DATASTAR_SHOWSet(b bool) *SVGSCRIPTElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSCRIPTElement) DATASTAR_SHOW() *SVGSCRIPTElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGSCRIPTElement) DATASTAR_INTERSECTS(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGSCRIPTElement) DATASTAR_INTERSECTSRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGSCRIPTElement) DATASTAR_TELEPORTSet(b bool) *SVGSCRIPTElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSCRIPTElement) DATASTAR_TELEPORT() *SVGSCRIPTElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGSCRIPTElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGSCRIPTElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSCRIPTElement) DATASTAR_SCROLL_INTO_VIEW() *SVGSCRIPTElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGSCRIPTElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGSCRIPTElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
