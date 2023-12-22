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

// This element is used to specify the behavior of a subexpression when it is 
// activated by the user 
// The action can be to toggle the visibility of the subexpression, or to toggle 
// the selection of the subexpression, or to execute a script. 
type MACTIONElementBuilder struct {
    *ElementBuilder
}

// Create a new MACTIONElementBuilder element.
// This will create a new element with the tag
// "maction" during rendering.
func MACTION(children ...ElementRenderer) *MACTIONElementBuilder {
    return &MACTIONElementBuilder{
        ElementBuilder: &ElementBuilder{
            Tag: []byte("maction"),
            IsSelfClosing: false,
            Descendants: children,
		},
    }
}

func (e *MACTIONElementBuilder) Children(children ...ElementRenderer) *MACTIONElementBuilder {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *MACTIONElementBuilder) IfChildren(condition bool, children ...ElementRenderer) *MACTIONElementBuilder {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *MACTIONElementBuilder) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MACTIONElementBuilder {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *MACTIONElementBuilder) Text(text string) *MACTIONElementBuilder {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *MACTIONElementBuilder) TextF(format string, args ...any) *MACTIONElementBuilder {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *MACTIONElementBuilder) Escaped(text string) *MACTIONElementBuilder {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *MACTIONElementBuilder) EscapedF(format string, args ...any) *MACTIONElementBuilder {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MACTIONElementBuilder) CustomData(key, value string) *MACTIONElementBuilder {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MACTIONElementBuilder) CustomDataRemove(key string) *MACTIONElementBuilder {
	if e.CustomDataAttributes == nil {
		return e
	}
    e.CustomDataAttributes.Del(key)
	return e
}


// This attribute specifies the type of action performed when the element is 
// activated 
// Possible values are toggle, statusline, tooltip, and script. 
func(e *MACTIONElementBuilder) ACTIONTYPE(c MactionActiontypeChoice) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("actiontype", string(c))
    return e
}

type MactionActiontypeChoice string
const(

    MactionActiontype_toggle MactionActiontypeChoice = "toggle"

    MactionActiontype_statusline MactionActiontypeChoice = "statusline"

    MactionActiontype_tooltip MactionActiontypeChoice = "tooltip"

    MactionActiontype_script MactionActiontypeChoice = "script"
)

// Remove the attribute actiontype from the element.
func(e *MACTIONElementBuilder) ACTIONTYPERemove(c MactionActiontypeChoice) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("actiontype")
    return e
}


// This attribute specifies the type of selection performed when the element is 
// activated 
// Possible values are none, highlight, and unhighlight. 
func(e *MACTIONElementBuilder) SELECTION(c MactionSelectionChoice) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("selection", string(c))
    return e
}

type MactionSelectionChoice string
const(

    MactionSelection_none MactionSelectionChoice = "none"

    MactionSelection_highlight MactionSelectionChoice = "highlight"

    MactionSelection_unhighlight MactionSelectionChoice = "unhighlight"
)

// Remove the attribute selection from the element.
func(e *MACTIONElementBuilder) SELECTIONRemove(c MactionSelectionChoice) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("selection")
    return e
}


// Assigns a class name or set of class names to an element 
// You may assign the same class name or names to any number of elements 
// If you specify multiple class names, they must be separated by whitespace 
// characters. 
func(e *MACTIONElementBuilder) CLASS(s ...string) *MACTIONElementBuilder{
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
func(e *MACTIONElementBuilder) CLASSRemove(s ...string) *MACTIONElementBuilder{
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
func(e *MACTIONElementBuilder) DIR(c MactionDirChoice) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("dir", string(c))
    return e
}

type MactionDirChoice string
const(
// left-to-right 
    MactionDir_ltr MactionDirChoice = "ltr"
// right-to-left 
    MactionDir_rtl MactionDirChoice = "rtl"
)

// Remove the attribute dir from the element.
func(e *MACTIONElementBuilder) DIRRemove(c MactionDirChoice) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("dir")
    return e
}


// This attribute specifies whether the element should be rendered using 
// displaystyle rules or not 
// Possible values are true and false. 
func(e *MACTIONElementBuilder) DISPLAYSTYLE(c MactionDisplaystyleChoice) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("displaystyle", string(c))
    return e
}

type MactionDisplaystyleChoice string
const(
// displaystyle rules 
    MactionDisplaystyle_true MactionDisplaystyleChoice = "true"
// not displaystyle rules 
    MactionDisplaystyle_false MactionDisplaystyleChoice = "false"
)

// Remove the attribute displaystyle from the element.
func(e *MACTIONElementBuilder) DISPLAYSTYLERemove(c MactionDisplaystyleChoice) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("displaystyle")
    return e
}


// This attribute assigns a name to an element 
// This name must be unique in a document. 
func(e *MACTIONElementBuilder) ID(s string) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("id", s)
    return e
}

// Remove the attribute id from the element.
func(e *MACTIONElementBuilder) IDRemove(s string) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("id")
    return e
}


// This attribute specifies the background color of the element 
// Possible values are a color name or a color specification in the format defined 
// in the CSS3 Color Module [CSS3COLOR]. 
func(e *MACTIONElementBuilder) MATHBACKGROUND(s string) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("mathbackground", s)
    return e
}

// Remove the attribute mathbackground from the element.
func(e *MACTIONElementBuilder) MATHBACKGROUNDRemove(s string) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("mathbackground")
    return e
}


// This attribute specifies the color of the element 
// Possible values are a color name or a color specification in the format defined 
// in the CSS3 Color Module [CSS3COLOR]. 
func(e *MACTIONElementBuilder) MATHCOLOR(s string) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("mathcolor", s)
    return e
}

// Remove the attribute mathcolor from the element.
func(e *MACTIONElementBuilder) MATHCOLORRemove(s string) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("mathcolor")
    return e
}


// This attribute specifies the size of the element 
// Possible values are a dimension or a dimensionless number. 
func(e *MACTIONElementBuilder) MATHSIZESTR(s string) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("mathsize", s)
    return e
}

// Remove the attribute mathsizeStr from the element.
func(e *MACTIONElementBuilder) MATHSIZESTRRemove(s string) *MACTIONElementBuilder{
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
func(e *MACTIONElementBuilder) NONCE(s string) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("nonce", s)
    return e
}

// Remove the attribute nonce from the element.
func(e *MACTIONElementBuilder) NONCERemove(s string) *MACTIONElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("nonce")
    return e
}


// This attribute specifies the script level of the element 
// Possible values are an integer between 0 and 7, inclusive. 
func(e *MACTIONElementBuilder) SCRIPTLEVEL(i int) *MACTIONElementBuilder{
    if e.IntAttributes == nil {
        e.IntAttributes = treemap.New[string,int]()
    }
    e.IntAttributes.Set("scriptlevel", i)
    return e
}

// Remove the attribute scriptlevel from the element.
func(e *MACTIONElementBuilder) SCRIPTLEVELRemove(i int) *MACTIONElementBuilder{
    if e.IntAttributes == nil {
        return e
    }
    e.IntAttributes.Del("scriptlevel")
    return e
}


// This attribute offers advisory information about the element for which it is 
// set. 
func (e *MACTIONElementBuilder) STYLE(k string, v string) *MACTIONElementBuilder {
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
func (e *MACTIONElementBuilder) STYLEMap(m map[string]string) *MACTIONElementBuilder {
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
func (e *MACTIONElementBuilder) STYLEPairs(pairs ...string) *MACTIONElementBuilder {
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
func (e *MACTIONElementBuilder) STYLERemove(keys ...string) *MACTIONElementBuilder {
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
func(e *MACTIONElementBuilder) TABINDEX(i int) *MACTIONElementBuilder{
    if e.IntAttributes == nil {
        e.IntAttributes = treemap.New[string,int]()
    }
    e.IntAttributes.Set("tabindex", i)
    return e
}

// Remove the attribute tabindex from the element.
func(e *MACTIONElementBuilder) TABINDEXRemove(i int) *MACTIONElementBuilder{
    if e.IntAttributes == nil {
        return e
    }
    e.IntAttributes.Del("tabindex")
    return e
}


