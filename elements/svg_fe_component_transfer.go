// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feComponentTransfer is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <feComponentTransfer> SVG filter primitive performs color-component-wise
// remapping of data for each pixel
// It allows operations like brightness adjustment, contrast adjustment, color
// balance or thresholding.
type SVGFECOMPONENTTRANSFERElement struct {
	*Element
}

// Create a new SVGFECOMPONENTTRANSFERElement element.
// This will create a new element with the tag
// "feComponentTransfer" during rendering.
func SVG_FECOMPONENTTRANSFER(children ...ElementRenderer) *SVGFECOMPONENTTRANSFERElement {
	e := NewElement("feComponentTransfer", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFECOMPONENTTRANSFERElement{Element: e}
}

func (e *SVGFECOMPONENTTRANSFERElement) Children(children ...ElementRenderer) *SVGFECOMPONENTTRANSFERElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) Attr(name string, value string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) Attrs(attrs ...string) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) AttrsMap(attrs map[string]string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) Text(text string) *SVGFECOMPONENTTRANSFERElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) TextF(format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPONENTTRANSFERElement) IfText(condition bool, text string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfTextF(condition bool, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) Escaped(text string) *SVGFECOMPONENTTRANSFERElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfEscaped(condition bool, text string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) EscapedF(format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPONENTTRANSFERElement) IfEscapedF(condition bool, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) CustomData(key, value string) *SVGFECOMPONENTTRANSFERElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfCustomData(condition bool, key, value string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) CustomDataF(key, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPONENTTRANSFERElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) CustomDataRemove(key string) *SVGFECOMPONENTTRANSFERElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The input for this filter.
func (e *SVGFECOMPONENTTRANSFERElement) IN(s string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfIN(condition bool, s string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.IN(s)
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFECOMPONENTTRANSFERElement) INRemove(s string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("in")
	return e
}

// Specifies a unique id for an element
func (e *SVGFECOMPONENTTRANSFERElement) ID(s string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfID(condition bool, s string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFECOMPONENTTRANSFERElement) IDRemove(s string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFECOMPONENTTRANSFERElement) CLASS(s ...string) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) IfCLASS(condition bool, s ...string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFECOMPONENTTRANSFERElement) CLASSRemove(s ...string) *SVGFECOMPONENTTRANSFERElement {
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
func (e *SVGFECOMPONENTTRANSFERElement) STYLEF(k string, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPONENTTRANSFERElement) IfSTYLE(condition bool, k string, v string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) STYLE(k string, v string) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFECOMPONENTTRANSFERElement) STYLEMap(m map[string]string) *SVGFECOMPONENTTRANSFERElement {
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
func (e *SVGFECOMPONENTTRANSFERElement) STYLEPairs(pairs ...string) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFECOMPONENTTRANSFERElement) STYLERemove(keys ...string) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_MERGE_STORE(v any) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_REF(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfDATASTAR_REF(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_REFRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_BIND(key string, expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_BINDRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_MODEL(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_MODELRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_TEXT(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_TEXTRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGFeComponentTransferDataOnMod customDataKeyModifier

// Debounces the event handler
func SVGFeComponentTransferDataOnModDebounce(
	d time.Duration,
) SVGFeComponentTransferDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGFeComponentTransferDataOnModThrottle(
	d time.Duration,
) SVGFeComponentTransferDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeComponentTransferDataOnMod) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGFeComponentTransferDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeComponentTransferDataOnMod) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_ONRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_FOCUSSet(b bool) *SVGFECOMPONENTTRANSFERElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_FOCUS() *SVGFECOMPONENTTRANSFERElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_HEADER(key string, expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_HEADERRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_FETCH_URL(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_FETCH_URLRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_FETCH_INDICATORRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_SHOWSet(b bool) *SVGFECOMPONENTTRANSFERElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_SHOW() *SVGFECOMPONENTTRANSFERElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_INTERSECTS(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_INTERSECTSRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_TELEPORTSet(b bool) *SVGFECOMPONENTTRANSFERElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_TELEPORT() *SVGFECOMPONENTTRANSFERElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFECOMPONENTTRANSFERElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFECOMPONENTTRANSFERElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFECOMPONENTTRANSFERElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
