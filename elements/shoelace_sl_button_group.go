// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package Shoelace ButtonGroup is generated from configuration file.
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

type SLBUTTONGROUPElement struct {
	*Element
}

// Create a new SLBUTTONGROUPElement element.
// This will create a new element with the tag
// "sl-button-group" during rendering.
func SL_BUTTONGROUP(children ...ElementRenderer) *SLBUTTONGROUPElement {
	e := NewElement("sl-button-group", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SLBUTTONGROUPElement{Element: e}
}

func (e *SLBUTTONGROUPElement) Children(children ...ElementRenderer) *SLBUTTONGROUPElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SLBUTTONGROUPElement) IfChildren(condition bool, children ...ElementRenderer) *SLBUTTONGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SLBUTTONGROUPElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SLBUTTONGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SLBUTTONGROUPElement) Attr(name string, value string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SLBUTTONGROUPElement) Attrs(attrs ...string) *SLBUTTONGROUPElement {
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

func (e *SLBUTTONGROUPElement) AttrsMap(attrs map[string]string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SLBUTTONGROUPElement) Text(text string) *SLBUTTONGROUPElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SLBUTTONGROUPElement) TextF(format string, args ...any) *SLBUTTONGROUPElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SLBUTTONGROUPElement) IfText(condition bool, text string) *SLBUTTONGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SLBUTTONGROUPElement) IfTextF(condition bool, format string, args ...any) *SLBUTTONGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SLBUTTONGROUPElement) Escaped(text string) *SLBUTTONGROUPElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SLBUTTONGROUPElement) IfEscaped(condition bool, text string) *SLBUTTONGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SLBUTTONGROUPElement) EscapedF(format string, args ...any) *SLBUTTONGROUPElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SLBUTTONGROUPElement) IfEscapedF(condition bool, format string, args ...any) *SLBUTTONGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SLBUTTONGROUPElement) CustomData(key, value string) *SLBUTTONGROUPElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SLBUTTONGROUPElement) IfCustomData(condition bool, key, value string) *SLBUTTONGROUPElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SLBUTTONGROUPElement) CustomDataF(key, format string, args ...any) *SLBUTTONGROUPElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SLBUTTONGROUPElement) IfCustomDataF(condition bool, key, format string, args ...any) *SLBUTTONGROUPElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SLBUTTONGROUPElement) CustomDataRemove(key string) *SLBUTTONGROUPElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

func (e *SLBUTTONGROUPElement) LABEL(s string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("label", s)
	return e
}

func (e *SLBUTTONGROUPElement) LABELF(format string, args ...any) *SLBUTTONGROUPElement {
	return e.LABEL(fmt.Sprintf(format, args...))
}

func (e *SLBUTTONGROUPElement) IfLABEL(condition bool, s string) *SLBUTTONGROUPElement {
	if condition {
		e.LABEL(s)
	}
	return e
}

func (e *SLBUTTONGROUPElement) IfLABELF(condition bool, format string, args ...any) *SLBUTTONGROUPElement {
	if condition {
		e.LABEL(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute LABEL from the element.
func (e *SLBUTTONGROUPElement) LABELRemove(s string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("label")
	return e
}

func (e *SLBUTTONGROUPElement) LABELRemoveF(format string, args ...any) *SLBUTTONGROUPElement {
	return e.LABELRemove(fmt.Sprintf(format, args...))
}

// Merges the singleton store with the given object

func (e *SLBUTTONGROUPElement) DATASTAR_STORE(v any) *SLBUTTONGROUPElement {
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

func (e *SLBUTTONGROUPElement) DATASTAR_REF(expression string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBUTTONGROUPElement) IfDATASTAR_REF(condition bool, expression string) *SLBUTTONGROUPElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SLBUTTONGROUPElement) DATASTAR_REFRemove() *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SLBUTTONGROUPElement) DATASTAR_BIND(key string, expression string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBUTTONGROUPElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SLBUTTONGROUPElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SLBUTTONGROUPElement) DATASTAR_BINDRemove() *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SLBUTTONGROUPElement) DATASTAR_MODEL(expression string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBUTTONGROUPElement) IfDATASTAR_MODEL(condition bool, expression string) *SLBUTTONGROUPElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SLBUTTONGROUPElement) DATASTAR_MODELRemove() *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SLBUTTONGROUPElement) DATASTAR_TEXT(expression string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBUTTONGROUPElement) IfDATASTAR_TEXT(condition bool, expression string) *SLBUTTONGROUPElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SLBUTTONGROUPElement) DATASTAR_TEXTRemove() *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SLButtonGroupOnMod customDataKeyModifier

// Debounces the event handler
func SLButtonGroupOnModDebounce(
	d time.Duration,
) SLButtonGroupOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SLButtonGroupOnModThrottle(
	d time.Duration,
) SLButtonGroupOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SLBUTTONGROUPElement) DATASTAR_ON(key string, expression string, modifiers ...SLButtonGroupOnMod) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SLButtonGroupOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBUTTONGROUPElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SLButtonGroupOnMod) *SLBUTTONGROUPElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SLBUTTONGROUPElement) DATASTAR_ONRemove() *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SLBUTTONGROUPElement) DATASTAR_FOCUSSet(b bool) *SLBUTTONGROUPElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLBUTTONGROUPElement) DATASTAR_FOCUS() *SLBUTTONGROUPElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SLBUTTONGROUPElement) DATASTAR_HEADER(key string, expression string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBUTTONGROUPElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SLBUTTONGROUPElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SLBUTTONGROUPElement) DATASTAR_HEADERRemove() *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SLBUTTONGROUPElement) DATASTAR_FETCH_INDICATOR(expression string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBUTTONGROUPElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SLBUTTONGROUPElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SLBUTTONGROUPElement) DATASTAR_FETCH_INDICATORRemove() *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SLBUTTONGROUPElement) DATASTAR_SHOWSet(b bool) *SLBUTTONGROUPElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLBUTTONGROUPElement) DATASTAR_SHOW() *SLBUTTONGROUPElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SLBUTTONGROUPElement) DATASTAR_INTERSECTS(expression string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBUTTONGROUPElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SLBUTTONGROUPElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SLBUTTONGROUPElement) DATASTAR_INTERSECTSRemove() *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SLBUTTONGROUPElement) DATASTAR_TELEPORTSet(b bool) *SLBUTTONGROUPElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLBUTTONGROUPElement) DATASTAR_TELEPORT() *SLBUTTONGROUPElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SLBUTTONGROUPElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SLBUTTONGROUPElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLBUTTONGROUPElement) DATASTAR_SCROLL_INTO_VIEW() *SLBUTTONGROUPElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SLBUTTONGROUPElement) DATASTAR_VIEW_TRANSITION(expression string) *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBUTTONGROUPElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SLBUTTONGROUPElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SLBUTTONGROUPElement) DATASTAR_VIEW_TRANSITIONRemove() *SLBUTTONGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
