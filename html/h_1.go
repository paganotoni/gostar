package  html

import (
    "fmt"
)

type H1HTMLElement struct {
    *Element
}

func H1(children ...ElementBuilder) *H1HTMLElement {
    return &H1HTMLElement{
        Element: &Element{
            Tag: "h1",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *H1HTMLElement) Children(children ...ElementBuilder) *H1HTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *H1HTMLElement) IfChildren(condition bool, children ...ElementBuilder) *H1HTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *H1HTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *H1HTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *H1HTMLElement) Text(text string) *H1HTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *H1HTMLElement) TextF(format string, args ...any) *H1HTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *H1HTMLElement) Escaped(text string) *H1HTMLElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *H1HTMLElement) EscapedF(format string, args ...any) *H1HTMLElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *H1HTMLElement) CustomData(key, value string) *H1HTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *H1HTMLElement) CustomDataRemove(key string) *H1HTMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	delete(e.CustomDataAttributes, key)
	return e
}


// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//  * ordered_set_of_unique_space_separated_tokens
//  * identical_to
func (e *H1HTMLElement) ACCESSKEY(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *H1HTMLElement) IfACCESSKEY(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.ACCESSKEY(v)
}

func (e *H1HTMLElement) RemoveACCESSKEY(v string) *H1HTMLElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// AUTOCAPITALIZE sets the "autocapitalize" attribute.
// Recommended autocapitalization behavior (for supported input methods)
// Values values are constrained to:
//  * on
//  * on
//  * off
//  * off
//  * none
//  * none
//  * sentences
//  * sentences
//  * words
//  * words
//  * characters
//  * characters
func (e *H1HTMLElement) AUTOCAPITALIZE(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *H1HTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.AUTOCAPITALIZE(v)
}

func (e *H1HTMLElement) RemoveAUTOCAPITALIZE(v string) *H1HTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *H1HTMLElement) AUTOFOCUS() *H1HTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *H1HTMLElement) IfAUTOFOCUS(cond bool) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.AUTOFOCUS()
}

func (e *H1HTMLElement) RemoveAUTOFOCUS() *H1HTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *H1HTMLElement) SetAUTOFOCUS(b bool) *H1HTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *H1HTMLElement) CLASS(v string) *H1HTMLElement {
    if e.DelimitedStringAttributes == nil {
        e.DelimitedStringAttributes = map[string]*DelimitedString{}
    }
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        kv = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = kv
    }
    kv.Add(v)
    return e
}

func (e *H1HTMLElement) IfCLASS(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.CLASS(v)
}

func (e *H1HTMLElement) SetCLASS(v string) *H1HTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *H1HTMLElement) RemoveCLASS(v string) *H1HTMLElement {
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        return e
    }
    kv.Remove(v)
    return e
}
// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//  * true
//  * plaintext_only
//  * false
func (e *H1HTMLElement) CONTENTEDITABLE(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *H1HTMLElement) IfCONTENTEDITABLE(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.CONTENTEDITABLE(v)
}

func (e *H1HTMLElement) RemoveCONTENTEDITABLE(v string) *H1HTMLElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//  * ltr
//  * ltr
//  * rtl
//  * rtl
func (e *H1HTMLElement) DIR(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *H1HTMLElement) IfDIR(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.DIR(v)
}

func (e *H1HTMLElement) RemoveDIR(v string) *H1HTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *H1HTMLElement) DRAGGABLE(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *H1HTMLElement) IfDRAGGABLE(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.DRAGGABLE(v)
}

func (e *H1HTMLElement) RemoveDRAGGABLE(v string) *H1HTMLElement {
    delete(e.StringAttributes, "draggable")
    return e
}
// ENTERKEYHINT sets the "enterkeyhint" attribute.
// Hint for selecting an enter key action
// Values values are constrained to:
//  * enter
//  * enter
//  * done
//  * done
//  * go
//  * go
//  * next
//  * next
//  * previous
//  * previous
//  * search
//  * search
//  * send
//  * send
func (e *H1HTMLElement) ENTERKEYHINT(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *H1HTMLElement) IfENTERKEYHINT(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.ENTERKEYHINT(v)
}

func (e *H1HTMLElement) RemoveENTERKEYHINT(v string) *H1HTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *H1HTMLElement) HIDDEN(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *H1HTMLElement) IfHIDDEN(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.HIDDEN(v)
}

func (e *H1HTMLElement) RemoveHIDDEN(v string) *H1HTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *H1HTMLElement) ID(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *H1HTMLElement) IfID(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.ID(v)
}

func (e *H1HTMLElement) RemoveID(v string) *H1HTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *H1HTMLElement) INERT() *H1HTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *H1HTMLElement) IfINERT(cond bool) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.INERT()
}

func (e *H1HTMLElement) RemoveINERT() *H1HTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *H1HTMLElement) SetINERT(b bool) *H1HTMLElement {
    if b {
        return e.INERT()
    }
    return e.RemoveINERT()
}
// INPUTMODE sets the "inputmode" attribute.
// Hint for selecting an input modality
// Values values are constrained to:
//  * none
//  * none
//  * text
//  * text
//  * tel
//  * tel
//  * email
//  * email
//  * url
//  * url
//  * numeric
//  * numeric
//  * decimal
//  * decimal
//  * search
//  * search
func (e *H1HTMLElement) INPUTMODE(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *H1HTMLElement) IfINPUTMODE(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.INPUTMODE(v)
}

func (e *H1HTMLElement) RemoveINPUTMODE(v string) *H1HTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *H1HTMLElement) IS(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *H1HTMLElement) IfIS(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.IS(v)
}

func (e *H1HTMLElement) RemoveIS(v string) *H1HTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *H1HTMLElement) ITEMID(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *H1HTMLElement) IfITEMID(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.ITEMID(v)
}

func (e *H1HTMLElement) RemoveITEMID(v string) *H1HTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *H1HTMLElement) ITEMPROP(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *H1HTMLElement) IfITEMPROP(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.ITEMPROP(v)
}

func (e *H1HTMLElement) RemoveITEMPROP(v string) *H1HTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *H1HTMLElement) ITEMREF(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *H1HTMLElement) IfITEMREF(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.ITEMREF(v)
}

func (e *H1HTMLElement) RemoveITEMREF(v string) *H1HTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *H1HTMLElement) ITEMSCOPE() *H1HTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *H1HTMLElement) IfITEMSCOPE(cond bool) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.ITEMSCOPE()
}

func (e *H1HTMLElement) RemoveITEMSCOPE() *H1HTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *H1HTMLElement) SetITEMSCOPE(b bool) *H1HTMLElement {
    if b {
        return e.ITEMSCOPE()
    }
    return e.RemoveITEMSCOPE()
}
// ITEMTYPE sets the "itemtype" attribute.
// Item types of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
func (e *H1HTMLElement) ITEMTYPE(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *H1HTMLElement) IfITEMTYPE(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.ITEMTYPE(v)
}

func (e *H1HTMLElement) RemoveITEMTYPE(v string) *H1HTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *H1HTMLElement) LANG(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *H1HTMLElement) IfLANG(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.LANG(v)
}

func (e *H1HTMLElement) RemoveLANG(v string) *H1HTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *H1HTMLElement) NONCE(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *H1HTMLElement) IfNONCE(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.NONCE(v)
}

func (e *H1HTMLElement) RemoveNONCE(v string) *H1HTMLElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//  * auto
//  * auto
//  * manual
//  * manual
func (e *H1HTMLElement) POPOVER(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *H1HTMLElement) IfPOPOVER(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.POPOVER(v)
}

func (e *H1HTMLElement) RemovePOPOVER(v string) *H1HTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *H1HTMLElement) SLOT(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *H1HTMLElement) IfSLOT(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.SLOT(v)
}

func (e *H1HTMLElement) RemoveSLOT(v string) *H1HTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *H1HTMLElement) SPELLCHECK(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *H1HTMLElement) IfSPELLCHECK(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.SPELLCHECK(v)
}

func (e *H1HTMLElement) RemoveSPELLCHECK(v string) *H1HTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *H1HTMLElement) STYLE(k,v string) *H1HTMLElement {
    if e.DelimitedKVAttributes == nil {
        e.DelimitedKVAttributes = map[string]*DelimitedKVString{}
    }
    kv, ok := e.DelimitedKVAttributes["style"]
    if !ok {
        kv = NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = kv
    }
    kv.Add(k,v)
    return e
}

func (e *H1HTMLElement) IfSTYLE(cond bool, k string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.STYLE(k, "")
}

func (e *H1HTMLElement) RemoveSTYLE(k string) *H1HTMLElement {
    kv, ok := e.DelimitedKVAttributes["style"]
    if !ok {
        return e
    }
    kv.Remove(k)
    return e
}
// TABINDEX sets the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and
//      the relative order of the element for the purposes of sequential focus navigation
// Values values are constrained to:
//  * valid_integer
func (e *H1HTMLElement) TABINDEX(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *H1HTMLElement) IfTABINDEX(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.TABINDEX(v)
}

func (e *H1HTMLElement) RemoveTABINDEX(v string) *H1HTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *H1HTMLElement) TITLE(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *H1HTMLElement) IfTITLE(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.TITLE(v)
}

func (e *H1HTMLElement) RemoveTITLE(v string) *H1HTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *H1HTMLElement) TRANSLATE(v string) *H1HTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *H1HTMLElement) IfTRANSLATE(cond bool, v string) *H1HTMLElement {
    if !cond {
        return e
    }
    return e.TRANSLATE(v)
}

func (e *H1HTMLElement) RemoveTRANSLATE(v string) *H1HTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
