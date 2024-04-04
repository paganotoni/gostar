// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package Shoelace RadioGroup is generated from configuration file.
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

type SLRADIOGROUPElement struct {
	*Element
}

// Create a new SLRADIOGROUPElement element.
// This will create a new element with the tag
// "sl-radio-group" during rendering.
func SL_RADIOGROUP(children ...ElementRenderer) *SLRADIOGROUPElement {
	e := NewElement("sl-radio-group", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SLRADIOGROUPElement{Element: e}
}

func (e *SLRADIOGROUPElement) Children(children ...ElementRenderer) *SLRADIOGROUPElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SLRADIOGROUPElement) IfChildren(condition bool, children ...ElementRenderer) *SLRADIOGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SLRADIOGROUPElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SLRADIOGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SLRADIOGROUPElement) Attr(name string, value string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SLRADIOGROUPElement) Attrs(attrs ...string) *SLRADIOGROUPElement {
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

func (e *SLRADIOGROUPElement) AttrsMap(attrs map[string]string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SLRADIOGROUPElement) Text(text string) *SLRADIOGROUPElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SLRADIOGROUPElement) TextF(format string, args ...any) *SLRADIOGROUPElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SLRADIOGROUPElement) IfText(condition bool, text string) *SLRADIOGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SLRADIOGROUPElement) IfTextF(condition bool, format string, args ...any) *SLRADIOGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SLRADIOGROUPElement) Escaped(text string) *SLRADIOGROUPElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SLRADIOGROUPElement) IfEscaped(condition bool, text string) *SLRADIOGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SLRADIOGROUPElement) EscapedF(format string, args ...any) *SLRADIOGROUPElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SLRADIOGROUPElement) IfEscapedF(condition bool, format string, args ...any) *SLRADIOGROUPElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SLRADIOGROUPElement) CustomData(key, value string) *SLRADIOGROUPElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SLRADIOGROUPElement) IfCustomData(condition bool, key, value string) *SLRADIOGROUPElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SLRADIOGROUPElement) CustomDataF(key, format string, args ...any) *SLRADIOGROUPElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SLRADIOGROUPElement) IfCustomDataF(condition bool, key, format string, args ...any) *SLRADIOGROUPElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SLRADIOGROUPElement) CustomDataRemove(key string) *SLRADIOGROUPElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

func (e *SLRADIOGROUPElement) SIZE(c SLRadioGroupSizeChoice) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("size", string(c))
	return e
}

type SLRadioGroupSizeChoice string

const (
	// Small
	SLRadioGroupSize_small SLRadioGroupSizeChoice = "small"
	// Medium
	SLRadioGroupSize_medium SLRadioGroupSizeChoice = "medium"
	// Large
	SLRadioGroupSize_large SLRadioGroupSizeChoice = "large"
)

// Remove the attribute SIZE from the element.
func (e *SLRADIOGROUPElement) SIZERemove(c SLRadioGroupSizeChoice) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("size")
	return e
}

func (e *SLRADIOGROUPElement) LABEL(s string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("label", s)
	return e
}

func (e *SLRADIOGROUPElement) LABELF(format string, args ...any) *SLRADIOGROUPElement {
	return e.LABEL(fmt.Sprintf(format, args...))
}

func (e *SLRADIOGROUPElement) IfLABEL(condition bool, s string) *SLRADIOGROUPElement {
	if condition {
		e.LABEL(s)
	}
	return e
}

func (e *SLRADIOGROUPElement) IfLABELF(condition bool, format string, args ...any) *SLRADIOGROUPElement {
	if condition {
		e.LABEL(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute LABEL from the element.
func (e *SLRADIOGROUPElement) LABELRemove(s string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("label")
	return e
}

func (e *SLRADIOGROUPElement) LABELRemoveF(format string, args ...any) *SLRADIOGROUPElement {
	return e.LABELRemove(fmt.Sprintf(format, args...))
}

func (e *SLRADIOGROUPElement) HELP_TEXT(s string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("help-text", s)
	return e
}

func (e *SLRADIOGROUPElement) HELP_TEXTF(format string, args ...any) *SLRADIOGROUPElement {
	return e.HELP_TEXT(fmt.Sprintf(format, args...))
}

func (e *SLRADIOGROUPElement) IfHELP_TEXT(condition bool, s string) *SLRADIOGROUPElement {
	if condition {
		e.HELP_TEXT(s)
	}
	return e
}

func (e *SLRADIOGROUPElement) IfHELP_TEXTF(condition bool, format string, args ...any) *SLRADIOGROUPElement {
	if condition {
		e.HELP_TEXT(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HELP_TEXT from the element.
func (e *SLRADIOGROUPElement) HELP_TEXTRemove(s string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("help-text")
	return e
}

func (e *SLRADIOGROUPElement) HELP_TEXTRemoveF(format string, args ...any) *SLRADIOGROUPElement {
	return e.HELP_TEXTRemove(fmt.Sprintf(format, args...))
}

func (e *SLRADIOGROUPElement) NAME(s string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("name", s)
	return e
}

func (e *SLRADIOGROUPElement) NAMEF(format string, args ...any) *SLRADIOGROUPElement {
	return e.NAME(fmt.Sprintf(format, args...))
}

func (e *SLRADIOGROUPElement) IfNAME(condition bool, s string) *SLRADIOGROUPElement {
	if condition {
		e.NAME(s)
	}
	return e
}

func (e *SLRADIOGROUPElement) IfNAMEF(condition bool, format string, args ...any) *SLRADIOGROUPElement {
	if condition {
		e.NAME(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NAME from the element.
func (e *SLRADIOGROUPElement) NAMERemove(s string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("name")
	return e
}

func (e *SLRADIOGROUPElement) NAMERemoveF(format string, args ...any) *SLRADIOGROUPElement {
	return e.NAMERemove(fmt.Sprintf(format, args...))
}

func (e *SLRADIOGROUPElement) VALUE(s string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("value", s)
	return e
}

func (e *SLRADIOGROUPElement) VALUEF(format string, args ...any) *SLRADIOGROUPElement {
	return e.VALUE(fmt.Sprintf(format, args...))
}

func (e *SLRADIOGROUPElement) IfVALUE(condition bool, s string) *SLRADIOGROUPElement {
	if condition {
		e.VALUE(s)
	}
	return e
}

func (e *SLRADIOGROUPElement) IfVALUEF(condition bool, format string, args ...any) *SLRADIOGROUPElement {
	if condition {
		e.VALUE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute VALUE from the element.
func (e *SLRADIOGROUPElement) VALUERemove(s string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("value")
	return e
}

func (e *SLRADIOGROUPElement) VALUERemoveF(format string, args ...any) *SLRADIOGROUPElement {
	return e.VALUERemove(fmt.Sprintf(format, args...))
}

func (e *SLRADIOGROUPElement) REQUIRED() *SLRADIOGROUPElement {
	e.REQUIREDSet(true)
	return e
}

func (e *SLRADIOGROUPElement) IfREQUIRED(condition bool) *SLRADIOGROUPElement {
	if condition {
		e.REQUIREDSet(true)
	}
	return e
}

// Set the attribute REQUIRED to the value b explicitly.
func (e *SLRADIOGROUPElement) REQUIREDSet(b bool) *SLRADIOGROUPElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("required", b)
	return e
}

func (e *SLRADIOGROUPElement) IfSetREQUIRED(condition bool, b bool) *SLRADIOGROUPElement {
	if condition {
		e.REQUIREDSet(b)
	}
	return e
}

// Remove the attribute REQUIRED from the element.
func (e *SLRADIOGROUPElement) REQUIREDRemove(b bool) *SLRADIOGROUPElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("required")
	return e
}

// Merges the singleton store with the given object

func (e *SLRADIOGROUPElement) DATASTAR_STORE(v any) *SLRADIOGROUPElement {
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

func (e *SLRADIOGROUPElement) DATASTAR_REF(expression string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOGROUPElement) IfDATASTAR_REF(condition bool, expression string) *SLRADIOGROUPElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SLRADIOGROUPElement) DATASTAR_REFRemove() *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SLRADIOGROUPElement) DATASTAR_BIND(key string, expression string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOGROUPElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SLRADIOGROUPElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SLRADIOGROUPElement) DATASTAR_BINDRemove() *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SLRADIOGROUPElement) DATASTAR_MODEL(expression string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOGROUPElement) IfDATASTAR_MODEL(condition bool, expression string) *SLRADIOGROUPElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SLRADIOGROUPElement) DATASTAR_MODELRemove() *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SLRADIOGROUPElement) DATASTAR_TEXT(expression string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOGROUPElement) IfDATASTAR_TEXT(condition bool, expression string) *SLRADIOGROUPElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SLRADIOGROUPElement) DATASTAR_TEXTRemove() *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SLRadioGroupOnMod customDataKeyModifier

// Debounces the event handler
func SLRadioGroupOnModDebounce(
	d time.Duration,
) SLRadioGroupOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SLRadioGroupOnModThrottle(
	d time.Duration,
) SLRadioGroupOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SLRADIOGROUPElement) DATASTAR_ON(key string, expression string, modifiers ...SLRadioGroupOnMod) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SLRadioGroupOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOGROUPElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SLRadioGroupOnMod) *SLRADIOGROUPElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SLRADIOGROUPElement) DATASTAR_ONRemove() *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SLRADIOGROUPElement) DATASTAR_FOCUSSet(b bool) *SLRADIOGROUPElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLRADIOGROUPElement) DATASTAR_FOCUS() *SLRADIOGROUPElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SLRADIOGROUPElement) DATASTAR_HEADER(key string, expression string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOGROUPElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SLRADIOGROUPElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SLRADIOGROUPElement) DATASTAR_HEADERRemove() *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SLRADIOGROUPElement) DATASTAR_FETCH_INDICATOR(expression string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOGROUPElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SLRADIOGROUPElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SLRADIOGROUPElement) DATASTAR_FETCH_INDICATORRemove() *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SLRADIOGROUPElement) DATASTAR_SHOWSet(b bool) *SLRADIOGROUPElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLRADIOGROUPElement) DATASTAR_SHOW() *SLRADIOGROUPElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SLRADIOGROUPElement) DATASTAR_INTERSECTS(expression string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOGROUPElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SLRADIOGROUPElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SLRADIOGROUPElement) DATASTAR_INTERSECTSRemove() *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SLRADIOGROUPElement) DATASTAR_TELEPORTSet(b bool) *SLRADIOGROUPElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLRADIOGROUPElement) DATASTAR_TELEPORT() *SLRADIOGROUPElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SLRADIOGROUPElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SLRADIOGROUPElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLRADIOGROUPElement) DATASTAR_SCROLL_INTO_VIEW() *SLRADIOGROUPElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SLRADIOGROUPElement) DATASTAR_VIEW_TRANSITION(expression string) *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOGROUPElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SLRADIOGROUPElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SLRADIOGROUPElement) DATASTAR_VIEW_TRANSITIONRemove() *SLRADIOGROUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
