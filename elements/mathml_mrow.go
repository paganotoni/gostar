// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml mrow is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// This element is used to display a sequence of expressions.
type MathMLMROWElement struct {
	*Element
}

// Create a new MathMLMROWElement element.
// This will create a new element with the tag
// "mrow" during rendering.
func MathML_MROW(children ...ElementRenderer) *MathMLMROWElement {
	e := NewElement("mrow", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMROWElement{Element: e}
}

func (e *MathMLMROWElement) Children(children ...ElementRenderer) *MathMLMROWElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMROWElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMROWElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMROWElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMROWElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMROWElement) Attr(name string, value string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *MathMLMROWElement) Attrs(attrs ...string) *MathMLMROWElement {
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

func (e *MathMLMROWElement) AttrsMap(attrs map[string]string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMROWElement) Text(text string) *MathMLMROWElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMROWElement) TextF(format string, args ...any) *MathMLMROWElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMROWElement) IfText(condition bool, text string) *MathMLMROWElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMROWElement) IfTextF(condition bool, format string, args ...any) *MathMLMROWElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMROWElement) Escaped(text string) *MathMLMROWElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMROWElement) IfEscaped(condition bool, text string) *MathMLMROWElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMROWElement) EscapedF(format string, args ...any) *MathMLMROWElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMROWElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMROWElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMROWElement) CustomData(key, value string) *MathMLMROWElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMROWElement) IfCustomData(condition bool, key, value string) *MathMLMROWElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMROWElement) CustomDataF(key, format string, args ...any) *MathMLMROWElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMROWElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMROWElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMROWElement) CustomDataRemove(key string) *MathMLMROWElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLMROWElement) CLASS(s ...string) *MathMLMROWElement {
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

func (e *MathMLMROWElement) IfCLASS(condition bool, s ...string) *MathMLMROWElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMROWElement) CLASSRemove(s ...string) *MathMLMROWElement {
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

// This attribute specifies the text directionality of the element, merely
// indicating what direction the text flows when surrounded by text with inherent
// directionality (such as Arabic or Hebrew)
// Possible values are ltr (left-to-right) and rtl (right-to-left).
func (e *MathMLMROWElement) DIR(c MathMLMrowDirChoice) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMrowDirChoice string

const (
	// left-to-right
	MathMLMrowDir_ltr MathMLMrowDirChoice = "ltr"
	// right-to-left
	MathMLMrowDir_rtl MathMLMrowDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMROWElement) DIRRemove(c MathMLMrowDirChoice) *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMROWElement) DISPLAYSTYLE(c MathMLMrowDisplaystyleChoice) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMrowDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMrowDisplaystyle_true MathMLMrowDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMrowDisplaystyle_false MathMLMrowDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMROWElement) DISPLAYSTYLERemove(c MathMLMrowDisplaystyleChoice) *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMROWElement) ID(s string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMROWElement) IfID(condition bool, s string) *MathMLMROWElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMROWElement) IDRemove(s string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMROWElement) MATHBACKGROUND(s string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMROWElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMROWElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMROWElement) MATHBACKGROUNDRemove(s string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMROWElement) MATHCOLOR(s string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMROWElement) IfMATHCOLOR(condition bool, s string) *MathMLMROWElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMROWElement) MATHCOLORRemove(s string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMROWElement) MATHSIZE_STR(s string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMROWElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMROWElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMROWElement) MATHSIZE_STRRemove(s string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathsize")
	return e
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMROWElement) NONCE(s string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMROWElement) IfNONCE(condition bool, s string) *MathMLMROWElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMROWElement) NONCERemove(s string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMROWElement) SCRIPTLEVEL(i int) *MathMLMROWElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMROWElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMROWElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMROWElement) SCRIPTLEVELRemove(i int) *MathMLMROWElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMROWElement) STYLEF(k string, format string, args ...any) *MathMLMROWElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMROWElement) IfSTYLE(condition bool, k string, v string) *MathMLMROWElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMROWElement) STYLE(k string, v string) *MathMLMROWElement {
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

func (e *MathMLMROWElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMROWElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMROWElement) STYLEMap(m map[string]string) *MathMLMROWElement {
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
func (e *MathMLMROWElement) STYLEPairs(pairs ...string) *MathMLMROWElement {
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

func (e *MathMLMROWElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMROWElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMROWElement) STYLERemove(keys ...string) *MathMLMROWElement {
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

// This attribute specifies the position of the current element in the tabbing
// order for the current document
// This value must be a number between 0 and 32767
// User agents should ignore leading zeros.
func (e *MathMLMROWElement) TABINDEX(i int) *MathMLMROWElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMROWElement) IfTABINDEX(condition bool, i int) *MathMLMROWElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMROWElement) TABINDEXRemove(i int) *MathMLMROWElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the store with the given object

func (e *MathMLMROWElement) DATASTAR_MERGE_STORE(v any) *MathMLMROWElement {
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

func (e *MathMLMROWElement) DATASTAR_REF(expression string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROWElement) IfDATASTAR_REF(condition bool, expression string) *MathMLMROWElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLMROWElement) DATASTAR_REFRemove() *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLMROWElement) DATASTAR_BIND(key string, expression string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROWElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLMROWElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLMROWElement) DATASTAR_BINDRemove() *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLMROWElement) DATASTAR_MODEL(expression string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROWElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLMROWElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLMROWElement) DATASTAR_MODELRemove() *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLMROWElement) DATASTAR_TEXT(expression string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROWElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLMROWElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLMROWElement) DATASTAR_TEXTRemove() *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLMrowDataOnMod customDataKeyModifier

// Debounces the event handler
func MathMLMrowDataOnModDebounce(
	d time.Duration,
) MathMLMrowDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func MathMLMrowDataOnModThrottle(
	d time.Duration,
) MathMLMrowDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *MathMLMROWElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLMrowDataOnMod) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m MathMLMrowDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROWElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLMrowDataOnMod) *MathMLMROWElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLMROWElement) DATASTAR_ONRemove() *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLMROWElement) DATASTAR_FOCUSSet(b bool) *MathMLMROWElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMROWElement) DATASTAR_FOCUS() *MathMLMROWElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLMROWElement) DATASTAR_HEADER(key string, expression string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROWElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLMROWElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLMROWElement) DATASTAR_HEADERRemove() *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *MathMLMROWElement) DATASTAR_FETCH_URL(expression string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROWElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLMROWElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *MathMLMROWElement) DATASTAR_FETCH_URLRemove() *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLMROWElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROWElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLMROWElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLMROWElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLMROWElement) DATASTAR_SHOWSet(b bool) *MathMLMROWElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMROWElement) DATASTAR_SHOW() *MathMLMROWElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLMROWElement) DATASTAR_INTERSECTS(expression string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROWElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *MathMLMROWElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *MathMLMROWElement) DATASTAR_INTERSECTSRemove() *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *MathMLMROWElement) DATASTAR_TELEPORTSet(b bool) *MathMLMROWElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMROWElement) DATASTAR_TELEPORT() *MathMLMROWElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLMROWElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMROWElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMROWElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMROWElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLMROWElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *MathMLMROWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-view-transition-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROWElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *MathMLMROWElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLMROWElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMROWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
