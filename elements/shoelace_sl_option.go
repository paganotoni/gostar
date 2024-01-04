// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package Shoelace Option is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

type SLOPTIONElement struct {
	*Element
}

// Create a new SLOPTIONElement element.
// This will create a new element with the tag
// "sl-option" during rendering.
func SL_OPTION(children ...ElementRenderer) *SLOPTIONElement {
	e := NewElement("sl-option", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SLOPTIONElement{Element: e}
}

func (e *SLOPTIONElement) Children(children ...ElementRenderer) *SLOPTIONElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SLOPTIONElement) IfChildren(condition bool, children ...ElementRenderer) *SLOPTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SLOPTIONElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SLOPTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SLOPTIONElement) Attr(name string, value string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SLOPTIONElement) Attrs(attrs ...string) *SLOPTIONElement {
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

func (e *SLOPTIONElement) AttrsMap(attrs map[string]string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SLOPTIONElement) Text(text string) *SLOPTIONElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SLOPTIONElement) TextF(format string, args ...any) *SLOPTIONElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SLOPTIONElement) IfText(condition bool, text string) *SLOPTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SLOPTIONElement) IfTextF(condition bool, format string, args ...any) *SLOPTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SLOPTIONElement) Escaped(text string) *SLOPTIONElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SLOPTIONElement) IfEscaped(condition bool, text string) *SLOPTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SLOPTIONElement) EscapedF(format string, args ...any) *SLOPTIONElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SLOPTIONElement) IfEscapedF(condition bool, format string, args ...any) *SLOPTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SLOPTIONElement) CustomData(key, value string) *SLOPTIONElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SLOPTIONElement) IfCustomData(condition bool, key, value string) *SLOPTIONElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SLOPTIONElement) CustomDataF(key, format string, args ...any) *SLOPTIONElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SLOPTIONElement) IfCustomDataF(condition bool, key, format string, args ...any) *SLOPTIONElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SLOPTIONElement) CustomDataRemove(key string) *SLOPTIONElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

func (e *SLOPTIONElement) VALUE(s string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("value", s)
	return e
}

func (e *SLOPTIONElement) IfVALUE(condition bool, s string) *SLOPTIONElement {
	if condition {
		e.VALUE(s)
	}
	return e
}

// Remove the attribute VALUE from the element.
func (e *SLOPTIONElement) VALUERemove(s string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("value")
	return e
}

// Merges the store with the given object

func (e *SLOPTIONElement) DATASTAR_MERGE_STORE(v any) *SLOPTIONElement {
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

func (e *SLOPTIONElement) DATASTAR_REF(expression string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLOPTIONElement) IfDATASTAR_REF(condition bool, expression string) *SLOPTIONElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SLOPTIONElement) DATASTAR_REFRemove() *SLOPTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SLOPTIONElement) DATASTAR_BIND(key string, expression string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLOPTIONElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SLOPTIONElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SLOPTIONElement) DATASTAR_BINDRemove() *SLOPTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SLOPTIONElement) DATASTAR_MODEL(expression string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLOPTIONElement) IfDATASTAR_MODEL(condition bool, expression string) *SLOPTIONElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SLOPTIONElement) DATASTAR_MODELRemove() *SLOPTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SLOPTIONElement) DATASTAR_TEXT(expression string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLOPTIONElement) IfDATASTAR_TEXT(condition bool, expression string) *SLOPTIONElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SLOPTIONElement) DATASTAR_TEXTRemove() *SLOPTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SLOptionDataOnMod customDataKeyModifier

// Debounces the event handler
func SLOptionDataOnModDebounce(
	d time.Duration,
) SLOptionDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SLOptionDataOnModThrottle(
	d time.Duration,
) SLOptionDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SLOPTIONElement) DATASTAR_ON(key string, expression string, modifiers ...SLOptionDataOnMod) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SLOptionDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLOPTIONElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SLOptionDataOnMod) *SLOPTIONElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SLOPTIONElement) DATASTAR_ONRemove() *SLOPTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SLOPTIONElement) DATASTAR_FOCUSSet(b bool) *SLOPTIONElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLOPTIONElement) DATASTAR_FOCUS() *SLOPTIONElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SLOPTIONElement) DATASTAR_HEADER(key string, expression string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLOPTIONElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SLOPTIONElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SLOPTIONElement) DATASTAR_HEADERRemove() *SLOPTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SLOPTIONElement) DATASTAR_FETCH_URL(expression string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLOPTIONElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SLOPTIONElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SLOPTIONElement) DATASTAR_FETCH_URLRemove() *SLOPTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SLOPTIONElement) DATASTAR_FETCH_INDICATOR(expression string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLOPTIONElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SLOPTIONElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SLOPTIONElement) DATASTAR_FETCH_INDICATORRemove() *SLOPTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SLOPTIONElement) DATASTAR_SHOWSet(b bool) *SLOPTIONElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLOPTIONElement) DATASTAR_SHOW() *SLOPTIONElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SLOPTIONElement) DATASTAR_INTERSECTS(expression string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLOPTIONElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SLOPTIONElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SLOPTIONElement) DATASTAR_INTERSECTSRemove() *SLOPTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SLOPTIONElement) DATASTAR_TELEPORTSet(b bool) *SLOPTIONElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLOPTIONElement) DATASTAR_TELEPORT() *SLOPTIONElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SLOPTIONElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SLOPTIONElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLOPTIONElement) DATASTAR_SCROLL_INTO_VIEW() *SLOPTIONElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SLOPTIONElement) DATASTAR_VIEW_TRANSITION(expression string) *SLOPTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLOPTIONElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SLOPTIONElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SLOPTIONElement) DATASTAR_VIEW_TRANSITIONRemove() *SLOPTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
