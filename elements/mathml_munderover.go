// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml munderover is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// This element is used to display an expression with both an underbar and an
// overbar.
type MathMLMUNDEROVERElement struct {
	*Element
}

// Create a new MathMLMUNDEROVERElement element.
// This will create a new element with the tag
// "munderover" during rendering.
func MathML_MUNDEROVER(children ...ElementRenderer) *MathMLMUNDEROVERElement {
	e := NewElement("munderover", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMUNDEROVERElement{Element: e}
}

func (e *MathMLMUNDEROVERElement) Children(children ...ElementRenderer) *MathMLMUNDEROVERElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMUNDEROVERElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMUNDEROVERElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMUNDEROVERElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMUNDEROVERElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMUNDEROVERElement) Attr(name string, value string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *MathMLMUNDEROVERElement) Attrs(attrs ...string) *MathMLMUNDEROVERElement {
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

func (e *MathMLMUNDEROVERElement) AttrsMap(attrs map[string]string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMUNDEROVERElement) Text(text string) *MathMLMUNDEROVERElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMUNDEROVERElement) TextF(format string, args ...any) *MathMLMUNDEROVERElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMUNDEROVERElement) IfText(condition bool, text string) *MathMLMUNDEROVERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMUNDEROVERElement) IfTextF(condition bool, format string, args ...any) *MathMLMUNDEROVERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMUNDEROVERElement) Escaped(text string) *MathMLMUNDEROVERElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMUNDEROVERElement) IfEscaped(condition bool, text string) *MathMLMUNDEROVERElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMUNDEROVERElement) EscapedF(format string, args ...any) *MathMLMUNDEROVERElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMUNDEROVERElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMUNDEROVERElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMUNDEROVERElement) CustomData(key, value string) *MathMLMUNDEROVERElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMUNDEROVERElement) IfCustomData(condition bool, key, value string) *MathMLMUNDEROVERElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMUNDEROVERElement) CustomDataF(key, format string, args ...any) *MathMLMUNDEROVERElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMUNDEROVERElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMUNDEROVERElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMUNDEROVERElement) CustomDataRemove(key string) *MathMLMUNDEROVERElement {
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
func (e *MathMLMUNDEROVERElement) CLASS(s ...string) *MathMLMUNDEROVERElement {
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

func (e *MathMLMUNDEROVERElement) IfCLASS(condition bool, s ...string) *MathMLMUNDEROVERElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMUNDEROVERElement) CLASSRemove(s ...string) *MathMLMUNDEROVERElement {
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
func (e *MathMLMUNDEROVERElement) DIR(c MathMLMunderoverDirChoice) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMunderoverDirChoice string

const (
	// left-to-right
	MathMLMunderoverDir_ltr MathMLMunderoverDirChoice = "ltr"
	// right-to-left
	MathMLMunderoverDir_rtl MathMLMunderoverDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMUNDEROVERElement) DIRRemove(c MathMLMunderoverDirChoice) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMUNDEROVERElement) DISPLAYSTYLE(c MathMLMunderoverDisplaystyleChoice) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMunderoverDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMunderoverDisplaystyle_true MathMLMunderoverDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMunderoverDisplaystyle_false MathMLMunderoverDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMUNDEROVERElement) DISPLAYSTYLERemove(c MathMLMunderoverDisplaystyleChoice) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMUNDEROVERElement) ID(s string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMUNDEROVERElement) IfID(condition bool, s string) *MathMLMUNDEROVERElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMUNDEROVERElement) IDRemove(s string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMUNDEROVERElement) MATHBACKGROUND(s string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMUNDEROVERElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMUNDEROVERElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMUNDEROVERElement) MATHBACKGROUNDRemove(s string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMUNDEROVERElement) MATHCOLOR(s string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMUNDEROVERElement) IfMATHCOLOR(condition bool, s string) *MathMLMUNDEROVERElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMUNDEROVERElement) MATHCOLORRemove(s string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMUNDEROVERElement) MATHSIZE_STR(s string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMUNDEROVERElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMUNDEROVERElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMUNDEROVERElement) MATHSIZE_STRRemove(s string) *MathMLMUNDEROVERElement {
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
func (e *MathMLMUNDEROVERElement) NONCE(s string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMUNDEROVERElement) IfNONCE(condition bool, s string) *MathMLMUNDEROVERElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMUNDEROVERElement) NONCERemove(s string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMUNDEROVERElement) SCRIPTLEVEL(i int) *MathMLMUNDEROVERElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMUNDEROVERElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMUNDEROVERElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMUNDEROVERElement) SCRIPTLEVELRemove(i int) *MathMLMUNDEROVERElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMUNDEROVERElement) STYLEF(k string, format string, args ...any) *MathMLMUNDEROVERElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMUNDEROVERElement) IfSTYLE(condition bool, k string, v string) *MathMLMUNDEROVERElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMUNDEROVERElement) STYLE(k string, v string) *MathMLMUNDEROVERElement {
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

func (e *MathMLMUNDEROVERElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMUNDEROVERElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMUNDEROVERElement) STYLEMap(m map[string]string) *MathMLMUNDEROVERElement {
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
func (e *MathMLMUNDEROVERElement) STYLEPairs(pairs ...string) *MathMLMUNDEROVERElement {
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

func (e *MathMLMUNDEROVERElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMUNDEROVERElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMUNDEROVERElement) STYLERemove(keys ...string) *MathMLMUNDEROVERElement {
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
func (e *MathMLMUNDEROVERElement) TABINDEX(i int) *MathMLMUNDEROVERElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMUNDEROVERElement) IfTABINDEX(condition bool, i int) *MathMLMUNDEROVERElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMUNDEROVERElement) TABINDEXRemove(i int) *MathMLMUNDEROVERElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the store with the given object

func (e *MathMLMUNDEROVERElement) DATASTAR_MERGE_STORE(v any) *MathMLMUNDEROVERElement {
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

func (e *MathMLMUNDEROVERElement) DATASTAR_REF(expression string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMUNDEROVERElement) IfDATASTAR_REF(condition bool, expression string) *MathMLMUNDEROVERElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLMUNDEROVERElement) DATASTAR_REFRemove() *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLMUNDEROVERElement) DATASTAR_BIND(key string, expression string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMUNDEROVERElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLMUNDEROVERElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLMUNDEROVERElement) DATASTAR_BINDRemove() *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLMUNDEROVERElement) DATASTAR_MODEL(expression string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMUNDEROVERElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLMUNDEROVERElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLMUNDEROVERElement) DATASTAR_MODELRemove() *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLMUNDEROVERElement) DATASTAR_TEXT(expression string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMUNDEROVERElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLMUNDEROVERElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLMUNDEROVERElement) DATASTAR_TEXTRemove() *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLMunderoverDataOnMod customDataKeyModifier

// Debounces the event handler
func MathMLMunderoverDataOnModDebounce(
	d time.Duration,
) MathMLMunderoverDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func MathMLMunderoverDataOnModThrottle(
	d time.Duration,
) MathMLMunderoverDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *MathMLMUNDEROVERElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLMunderoverDataOnMod) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m MathMLMunderoverDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMUNDEROVERElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLMunderoverDataOnMod) *MathMLMUNDEROVERElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLMUNDEROVERElement) DATASTAR_ONRemove() *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLMUNDEROVERElement) DATASTAR_FOCUSSet(b bool) *MathMLMUNDEROVERElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMUNDEROVERElement) DATASTAR_FOCUS() *MathMLMUNDEROVERElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLMUNDEROVERElement) DATASTAR_HEADER(key string, expression string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMUNDEROVERElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLMUNDEROVERElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLMUNDEROVERElement) DATASTAR_HEADERRemove() *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *MathMLMUNDEROVERElement) DATASTAR_FETCH_URL(expression string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMUNDEROVERElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLMUNDEROVERElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *MathMLMUNDEROVERElement) DATASTAR_FETCH_URLRemove() *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLMUNDEROVERElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMUNDEROVERElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLMUNDEROVERElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLMUNDEROVERElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLMUNDEROVERElement) DATASTAR_SHOWSet(b bool) *MathMLMUNDEROVERElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMUNDEROVERElement) DATASTAR_SHOW() *MathMLMUNDEROVERElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLMUNDEROVERElement) DATASTAR_INTERSECTS(expression string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMUNDEROVERElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *MathMLMUNDEROVERElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *MathMLMUNDEROVERElement) DATASTAR_INTERSECTSRemove() *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *MathMLMUNDEROVERElement) DATASTAR_TELEPORTSet(b bool) *MathMLMUNDEROVERElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMUNDEROVERElement) DATASTAR_TELEPORT() *MathMLMUNDEROVERElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLMUNDEROVERElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMUNDEROVERElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMUNDEROVERElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMUNDEROVERElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLMUNDEROVERElement) DATASTAR_VIEW_TRANSITION(expression string) *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMUNDEROVERElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *MathMLMUNDEROVERElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLMUNDEROVERElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMUNDEROVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
