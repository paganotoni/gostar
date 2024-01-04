// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml math is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// This element is the root element of a MathML expression
// It is used to identify the document as a MathML document, and to specify the
// namespaces used in the document.
type MathMLMATHElement struct {
	*Element
}

// Create a new MathMLMATHElement element.
// This will create a new element with the tag
// "math" during rendering.
func MathML_MATH(children ...ElementRenderer) *MathMLMATHElement {
	e := NewElement("math", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMATHElement{Element: e}
}

func (e *MathMLMATHElement) Children(children ...ElementRenderer) *MathMLMATHElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMATHElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMATHElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMATHElement) Attr(name string, value string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *MathMLMATHElement) Attrs(attrs ...string) *MathMLMATHElement {
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

func (e *MathMLMATHElement) AttrsMap(attrs map[string]string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMATHElement) Text(text string) *MathMLMATHElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMATHElement) TextF(format string, args ...any) *MathMLMATHElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfText(condition bool, text string) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMATHElement) IfTextF(condition bool, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMATHElement) Escaped(text string) *MathMLMATHElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMATHElement) IfEscaped(condition bool, text string) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMATHElement) EscapedF(format string, args ...any) *MathMLMATHElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMATHElement) CustomData(key, value string) *MathMLMATHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMATHElement) IfCustomData(condition bool, key, value string) *MathMLMATHElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMATHElement) CustomDataF(key, format string, args ...any) *MathMLMATHElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMATHElement) CustomDataRemove(key string) *MathMLMATHElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// This attribute specifies the default namespace for elements and attributes in
// the document
// Possible values are http://www.w3.org/1998/Math/MathML and
// http://www.w3.org/1999/xhtml.
func (e *MathMLMATHElement) XMLNS(c MathMLMathXmlnsChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("xmlns", string(c))
	return e
}

type MathMLMathXmlnsChoice string

const (
	MathMLMathXmlns_http___www_w3_org_1998_Math_MathML MathMLMathXmlnsChoice = "http://www.w3.org/1998/Math/MathML"

	MathMLMathXmlns_http___www_w3_org_1999_xhtml MathMLMathXmlnsChoice = "http://www.w3.org/1999/xhtml"
)

// Remove the attribute XMLNS from the element.
func (e *MathMLMATHElement) XMLNSRemove(c MathMLMathXmlnsChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("xmlns")
	return e
}

// This attribute specifies the namespace for elements and attributes in the
// document whose names start with the letter m
// Possible values are http://www.w3.org/1998/Math/MathML and
// http://www.w3.org/1999/xhtml.
func (e *MathMLMATHElement) XMLNS_M(c MathMLMathXmlnsmChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("xmlns:m", string(c))
	return e
}

type MathMLMathXmlnsmChoice string

const (
	MathMLMathXmlnsm_http___www_w3_org_1998_Math_MathML MathMLMathXmlnsmChoice = "http://www.w3.org/1998/Math/MathML"

	MathMLMathXmlnsm_http___www_w3_org_1999_xhtml MathMLMathXmlnsmChoice = "http://www.w3.org/1999/xhtml"
)

// Remove the attribute XMLNS_M from the element.
func (e *MathMLMATHElement) XMLNS_MRemove(c MathMLMathXmlnsmChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("xmlns:m")
	return e
}

// This attribute specifies the namespace for elements and attributes in the
// document whose names start with the letters xlink
// Possible values are http://www.w3.org/1999/xlink and
// http://www.w3.org/1999/xhtml.
func (e *MathMLMATHElement) XMLNS_XLINK(c MathMLMathXmlnsxlinkChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("xmlns:xlink", string(c))
	return e
}

type MathMLMathXmlnsxlinkChoice string

const (
	MathMLMathXmlnsxlink_http___www_w3_org_1999_xlink MathMLMathXmlnsxlinkChoice = "http://www.w3.org/1999/xlink"

	MathMLMathXmlnsxlink_http___www_w3_org_1999_xhtml MathMLMathXmlnsxlinkChoice = "http://www.w3.org/1999/xhtml"
)

// Remove the attribute XMLNS_XLINK from the element.
func (e *MathMLMATHElement) XMLNS_XLINKRemove(c MathMLMathXmlnsxlinkChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("xmlns:xlink")
	return e
}

// This attribute specifies the namespace for elements and attributes in the
// document whose names start with the letters xml
// Possible values are http://www.w3.org/XML/1998/namespace and
// http://www.w3.org/1999/xhtml.
func (e *MathMLMATHElement) XMLNS_XML(c MathMLMathXmlnsxmlChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("xmlns:xml", string(c))
	return e
}

type MathMLMathXmlnsxmlChoice string

const (
	MathMLMathXmlnsxml_http___www_w3_org_XML_1998_namespace MathMLMathXmlnsxmlChoice = "http://www.w3.org/XML/1998/namespace"

	MathMLMathXmlnsxml_http___www_w3_org_1999_xhtml MathMLMathXmlnsxmlChoice = "http://www.w3.org/1999/xhtml"
)

// Remove the attribute XMLNS_XML from the element.
func (e *MathMLMATHElement) XMLNS_XMLRemove(c MathMLMathXmlnsxmlChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("xmlns:xml")
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLMATHElement) CLASS(s ...string) *MathMLMATHElement {
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

func (e *MathMLMATHElement) IfCLASS(condition bool, s ...string) *MathMLMATHElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMATHElement) CLASSRemove(s ...string) *MathMLMATHElement {
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
func (e *MathMLMATHElement) DIR(c MathMLMathDirChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMathDirChoice string

const (
	// left-to-right
	MathMLMathDir_ltr MathMLMathDirChoice = "ltr"
	// right-to-left
	MathMLMathDir_rtl MathMLMathDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMATHElement) DIRRemove(c MathMLMathDirChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMATHElement) DISPLAYSTYLE(c MathMLMathDisplaystyleChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMathDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMathDisplaystyle_true MathMLMathDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMathDisplaystyle_false MathMLMathDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMATHElement) DISPLAYSTYLERemove(c MathMLMathDisplaystyleChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMATHElement) ID(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMATHElement) IfID(condition bool, s string) *MathMLMATHElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMATHElement) IDRemove(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMATHElement) MATHBACKGROUND(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMATHElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMATHElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMATHElement) MATHBACKGROUNDRemove(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMATHElement) MATHCOLOR(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMATHElement) IfMATHCOLOR(condition bool, s string) *MathMLMATHElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMATHElement) MATHCOLORRemove(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMATHElement) MATHSIZE_STR(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMATHElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMATHElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMATHElement) MATHSIZE_STRRemove(s string) *MathMLMATHElement {
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
func (e *MathMLMATHElement) NONCE(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMATHElement) IfNONCE(condition bool, s string) *MathMLMATHElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMATHElement) NONCERemove(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMATHElement) SCRIPTLEVEL(i int) *MathMLMATHElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMATHElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMATHElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMATHElement) SCRIPTLEVELRemove(i int) *MathMLMATHElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMATHElement) STYLEF(k string, format string, args ...any) *MathMLMATHElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfSTYLE(condition bool, k string, v string) *MathMLMATHElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMATHElement) STYLE(k string, v string) *MathMLMATHElement {
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

func (e *MathMLMATHElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMATHElement) STYLEMap(m map[string]string) *MathMLMATHElement {
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
func (e *MathMLMATHElement) STYLEPairs(pairs ...string) *MathMLMATHElement {
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

func (e *MathMLMATHElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMATHElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMATHElement) STYLERemove(keys ...string) *MathMLMATHElement {
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
func (e *MathMLMATHElement) TABINDEX(i int) *MathMLMATHElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMATHElement) IfTABINDEX(condition bool, i int) *MathMLMATHElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMATHElement) TABINDEXRemove(i int) *MathMLMATHElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the store with the given object

func (e *MathMLMATHElement) DATASTAR_MERGE_STORE(v any) *MathMLMATHElement {
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

func (e *MathMLMATHElement) DATASTAR_REF(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfDATASTAR_REF(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLMATHElement) DATASTAR_REFRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLMATHElement) DATASTAR_BIND(key string, expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLMATHElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLMATHElement) DATASTAR_BINDRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLMATHElement) DATASTAR_MODEL(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLMATHElement) DATASTAR_MODELRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLMATHElement) DATASTAR_TEXT(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLMATHElement) DATASTAR_TEXTRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLMathDataOnMod customDataKeyModifier

// Debounces the event handler
func MathMLMathDataOnModDebounce(
	d time.Duration,
) MathMLMathDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func MathMLMathDataOnModThrottle(
	d time.Duration,
) MathMLMathDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *MathMLMATHElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLMathDataOnMod) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m MathMLMathDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLMathDataOnMod) *MathMLMATHElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLMATHElement) DATASTAR_ONRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLMATHElement) DATASTAR_FOCUSSet(b bool) *MathMLMATHElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMATHElement) DATASTAR_FOCUS() *MathMLMATHElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLMATHElement) DATASTAR_HEADER(key string, expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLMATHElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLMATHElement) DATASTAR_HEADERRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *MathMLMATHElement) DATASTAR_FETCH_URL(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *MathMLMATHElement) DATASTAR_FETCH_URLRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLMATHElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLMATHElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLMATHElement) DATASTAR_SHOWSet(b bool) *MathMLMATHElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMATHElement) DATASTAR_SHOW() *MathMLMATHElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLMATHElement) DATASTAR_INTERSECTS(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *MathMLMATHElement) DATASTAR_INTERSECTSRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *MathMLMATHElement) DATASTAR_TELEPORTSet(b bool) *MathMLMATHElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMATHElement) DATASTAR_TELEPORT() *MathMLMATHElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLMATHElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMATHElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMATHElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMATHElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLMATHElement) DATASTAR_VIEW_TRANSITION(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLMATHElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
