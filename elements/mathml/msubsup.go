// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml is generated from configuration file.
// Description:
// This MathML Core specification intends to address these issues by being as 
// accurate as possible on the visual rendering of mathematical formulas using 
// additional rules from the TeXBook’s Appendix G [TEXBOOK] and from the Open 
// Font Format [OPEN-FONT-FORMAT], [OPEN-TYPE-MATH-ILLUMINATED] 
// It also relies on modern browser implementations and web technologies [HTML] 
// [SVG] [CSS2] [DOM], clarifying interactions with them when needed or 
// introducing new low-level primitives to improve the web platform layering. 
package mathml

import(
    "fmt"
    "github.com/igrmk/treemap/v2"
)

// This element is used to display a subscript expression with a superscript 
// expression. 
type MSUBSUPElementBuilder struct {
    *ElementBuilder
}

// Create a new MSUBSUPElementBuilder element.
// This will create a new element with the tag
// "msubsup" during rendering.
func MSUBSUP(children ...ElementRenderer) *MSUBSUPElementBuilder {
    return &MSUBSUPElementBuilder{
        ElementBuilder: &ElementBuilder{
            Tag: []byte("msubsup"),
            IsSelfClosing: false,
            Descendants: children,
		},
    }
}

func (e *MSUBSUPElementBuilder) Children(children ...ElementRenderer) *MSUBSUPElementBuilder {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *MSUBSUPElementBuilder) IfChildren(condition bool, children ...ElementRenderer) *MSUBSUPElementBuilder {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *MSUBSUPElementBuilder) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MSUBSUPElementBuilder {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *MSUBSUPElementBuilder) Text(text string) *MSUBSUPElementBuilder {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *MSUBSUPElementBuilder) TextF(format string, args ...any) *MSUBSUPElementBuilder {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *MSUBSUPElementBuilder) Escaped(text string) *MSUBSUPElementBuilder {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *MSUBSUPElementBuilder) EscapedF(format string, args ...any) *MSUBSUPElementBuilder {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MSUBSUPElementBuilder) CustomData(key, value string) *MSUBSUPElementBuilder {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MSUBSUPElementBuilder) CustomDataRemove(key string) *MSUBSUPElementBuilder {
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
func(e *MSUBSUPElementBuilder) CLASS(s ...string) *MSUBSUPElementBuilder{
    if e.DelimitedStrings == nil {
        e.DelimitedStrings = treemap.New[string,*DelimitedBuilder[string]]()
    }
    ds, ok := e.DelimitedStrings.Get("class")
    if !ok {
        ds = NewDelimitedBuilder[string](" ")
        e.DelimitedStrings.Set("class", ds)
    }
    ds.Add(s...)
    return e
}

// Remove the attribute class from the element.
func(e *MSUBSUPElementBuilder) CLASSRemove(s ...string) *MSUBSUPElementBuilder{
    if e.DelimitedStrings == nil {
        return e
    }
    ds, ok := e.DelimitedStrings.Get("class")
    if !ok {
        return e
    }
    ds.Remove(s ...)
    return e
}



// This attribute specifies the text directionality of the element, merely 
// indicating what direction the text flows when surrounded by text with inherent 
// directionality (such as Arabic or Hebrew) 
// Possible values are ltr (left-to-right) and rtl (right-to-left). 
func(e *MSUBSUPElementBuilder) DIR(c MsubsupDirChoice) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("dir", string(c))
    return e
}

type MsubsupDirChoice string
const(
// left-to-right 
    MsubsupDir_ltr MsubsupDirChoice = "ltr"
// right-to-left 
    MsubsupDir_rtl MsubsupDirChoice = "rtl"
)

// Remove the attribute dir from the element.
func(e *MSUBSUPElementBuilder) DIRRemove(c MsubsupDirChoice) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("dir")
    return e
}


// This attribute specifies whether the element should be rendered using 
// displaystyle rules or not 
// Possible values are true and false. 
func(e *MSUBSUPElementBuilder) DISPLAYSTYLE(c MsubsupDisplaystyleChoice) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("displaystyle", string(c))
    return e
}

type MsubsupDisplaystyleChoice string
const(
// displaystyle rules 
    MsubsupDisplaystyle_true MsubsupDisplaystyleChoice = "true"
// not displaystyle rules 
    MsubsupDisplaystyle_false MsubsupDisplaystyleChoice = "false"
)

// Remove the attribute displaystyle from the element.
func(e *MSUBSUPElementBuilder) DISPLAYSTYLERemove(c MsubsupDisplaystyleChoice) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("displaystyle")
    return e
}


// This attribute assigns a name to an element 
// This name must be unique in a document. 
func(e *MSUBSUPElementBuilder) ID(s string) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("id", s)
    return e
}

// Remove the attribute id from the element.
func(e *MSUBSUPElementBuilder) IDRemove(s string) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("id")
    return e
}


// This attribute specifies the background color of the element 
// Possible values are a color name or a color specification in the format defined 
// in the CSS3 Color Module [CSS3COLOR]. 
func(e *MSUBSUPElementBuilder) MATHBACKGROUND(s string) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("mathbackground", s)
    return e
}

// Remove the attribute mathbackground from the element.
func(e *MSUBSUPElementBuilder) MATHBACKGROUNDRemove(s string) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("mathbackground")
    return e
}


// This attribute specifies the color of the element 
// Possible values are a color name or a color specification in the format defined 
// in the CSS3 Color Module [CSS3COLOR]. 
func(e *MSUBSUPElementBuilder) MATHCOLOR(s string) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("mathcolor", s)
    return e
}

// Remove the attribute mathcolor from the element.
func(e *MSUBSUPElementBuilder) MATHCOLORRemove(s string) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("mathcolor")
    return e
}


// This attribute specifies the size of the element 
// Possible values are a dimension or a dimensionless number. 
func(e *MSUBSUPElementBuilder) MATHSIZESTR(s string) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("mathsize", s)
    return e
}

// Remove the attribute mathsizeStr from the element.
func(e *MSUBSUPElementBuilder) MATHSIZESTRRemove(s string) *MSUBSUPElementBuilder{
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
func(e *MSUBSUPElementBuilder) NONCE(s string) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("nonce", s)
    return e
}

// Remove the attribute nonce from the element.
func(e *MSUBSUPElementBuilder) NONCERemove(s string) *MSUBSUPElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("nonce")
    return e
}


// This attribute specifies the script level of the element 
// Possible values are an integer between 0 and 7, inclusive. 
func(e *MSUBSUPElementBuilder) SCRIPTLEVEL(i int) *MSUBSUPElementBuilder{
    if e.IntAttributes == nil {
        e.IntAttributes = treemap.New[string,int]()
    }
    e.IntAttributes.Set("scriptlevel", i)
    return e
}

// Remove the attribute scriptlevel from the element.
func(e *MSUBSUPElementBuilder) SCRIPTLEVELRemove(i int) *MSUBSUPElementBuilder{
    if e.IntAttributes == nil {
        return e
    }
    e.IntAttributes.Del("scriptlevel")
    return e
}


// This attribute offers advisory information about the element for which it is 
// set. 
func (e *MSUBSUPElementBuilder) STYLE(k string, v string) *MSUBSUPElementBuilder {
    if e.KVStrings == nil {
        e.KVStrings = treemap.New[string,*KVBuilder]()
    }
    kv, ok := e.KVStrings.Get("style")
    if !ok {
        kv = NewKVBuilder(":", ";")
        e.KVStrings.Set("style", kv)
    }
    kv.Add(k, v)
    return e
}

// Add the attributes in the map to the element.
func (e *MSUBSUPElementBuilder) STYLEMap(m map[string]string) *MSUBSUPElementBuilder {
    if e.KVStrings == nil {
        e.KVStrings = treemap.New[string,*KVBuilder]()
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
func (e *MSUBSUPElementBuilder) STYLEPairs(pairs ...string) *MSUBSUPElementBuilder {
    if len(pairs) % 2 != 0 {
        panic("Must have an even number of pairs")
    }
    if e.KVStrings == nil {
        e.KVStrings = treemap.New[string,*KVBuilder]()
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


// Remove the attribute style from the element.
func (e *MSUBSUPElementBuilder) STYLERemove(keys ...string) *MSUBSUPElementBuilder {
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
func(e *MSUBSUPElementBuilder) TABINDEX(i int) *MSUBSUPElementBuilder{
    if e.IntAttributes == nil {
        e.IntAttributes = treemap.New[string,int]()
    }
    e.IntAttributes.Set("tabindex", i)
    return e
}

// Remove the attribute tabindex from the element.
func(e *MSUBSUPElementBuilder) TABINDEXRemove(i int) *MSUBSUPElementBuilder{
    if e.IntAttributes == nil {
        return e
    }
    e.IntAttributes.Del("tabindex")
    return e
}


