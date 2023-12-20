package  html

import (
    "fmt"
)

type TableHTMLElement struct {
    *Element
}

func TABLE(children ...ElementBuilder) *TableHTMLElement {
    return &TableHTMLElement{
        Element: &Element{
            Tag: "table",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *TableHTMLElement) Children(children ...ElementBuilder) *TableHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *TableHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TableHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *TableHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TableHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *TableHTMLElement) Text(text string) *TableHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *TableHTMLElement) TextF(format string, args ...any) *TableHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *TableHTMLElement) Escaped(text string) *TableHTMLElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *TableHTMLElement) EscapedF(format string, args ...any) *TableHTMLElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *TableHTMLElement) CustomData(key, value string) *TableHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TableHTMLElement) CustomDataRemove(key string) *TableHTMLElement {
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
func (e *TableHTMLElement) ACCESSKEY(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *TableHTMLElement) IfACCESSKEY(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.ACCESSKEY(v)
}

func (e *TableHTMLElement) RemoveACCESSKEY(v string) *TableHTMLElement {
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
func (e *TableHTMLElement) AUTOCAPITALIZE(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *TableHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOCAPITALIZE(v)
}

func (e *TableHTMLElement) RemoveAUTOCAPITALIZE(v string) *TableHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *TableHTMLElement) AUTOFOCUS() *TableHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *TableHTMLElement) IfAUTOFOCUS(cond bool) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOFOCUS()
}

func (e *TableHTMLElement) RemoveAUTOFOCUS() *TableHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *TableHTMLElement) SetAUTOFOCUS(b bool) *TableHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *TableHTMLElement) CLASS(v string) *TableHTMLElement {
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

func (e *TableHTMLElement) IfCLASS(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.CLASS(v)
}

func (e *TableHTMLElement) SetCLASS(v string) *TableHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *TableHTMLElement) RemoveCLASS(v string) *TableHTMLElement {
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
func (e *TableHTMLElement) CONTENTEDITABLE(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *TableHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.CONTENTEDITABLE(v)
}

func (e *TableHTMLElement) RemoveCONTENTEDITABLE(v string) *TableHTMLElement {
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
func (e *TableHTMLElement) DIR(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *TableHTMLElement) IfDIR(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.DIR(v)
}

func (e *TableHTMLElement) RemoveDIR(v string) *TableHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *TableHTMLElement) DRAGGABLE(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *TableHTMLElement) IfDRAGGABLE(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.DRAGGABLE(v)
}

func (e *TableHTMLElement) RemoveDRAGGABLE(v string) *TableHTMLElement {
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
func (e *TableHTMLElement) ENTERKEYHINT(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *TableHTMLElement) IfENTERKEYHINT(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.ENTERKEYHINT(v)
}

func (e *TableHTMLElement) RemoveENTERKEYHINT(v string) *TableHTMLElement {
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
func (e *TableHTMLElement) HIDDEN(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *TableHTMLElement) IfHIDDEN(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.HIDDEN(v)
}

func (e *TableHTMLElement) RemoveHIDDEN(v string) *TableHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *TableHTMLElement) ID(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *TableHTMLElement) IfID(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.ID(v)
}

func (e *TableHTMLElement) RemoveID(v string) *TableHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *TableHTMLElement) INERT() *TableHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *TableHTMLElement) IfINERT(cond bool) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.INERT()
}

func (e *TableHTMLElement) RemoveINERT() *TableHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *TableHTMLElement) SetINERT(b bool) *TableHTMLElement {
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
func (e *TableHTMLElement) INPUTMODE(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *TableHTMLElement) IfINPUTMODE(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.INPUTMODE(v)
}

func (e *TableHTMLElement) RemoveINPUTMODE(v string) *TableHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *TableHTMLElement) IS(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *TableHTMLElement) IfIS(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.IS(v)
}

func (e *TableHTMLElement) RemoveIS(v string) *TableHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *TableHTMLElement) ITEMID(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *TableHTMLElement) IfITEMID(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMID(v)
}

func (e *TableHTMLElement) RemoveITEMID(v string) *TableHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *TableHTMLElement) ITEMPROP(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *TableHTMLElement) IfITEMPROP(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMPROP(v)
}

func (e *TableHTMLElement) RemoveITEMPROP(v string) *TableHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *TableHTMLElement) ITEMREF(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *TableHTMLElement) IfITEMREF(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMREF(v)
}

func (e *TableHTMLElement) RemoveITEMREF(v string) *TableHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *TableHTMLElement) ITEMSCOPE() *TableHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *TableHTMLElement) IfITEMSCOPE(cond bool) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMSCOPE()
}

func (e *TableHTMLElement) RemoveITEMSCOPE() *TableHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *TableHTMLElement) SetITEMSCOPE(b bool) *TableHTMLElement {
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
func (e *TableHTMLElement) ITEMTYPE(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *TableHTMLElement) IfITEMTYPE(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMTYPE(v)
}

func (e *TableHTMLElement) RemoveITEMTYPE(v string) *TableHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TableHTMLElement) LANG(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *TableHTMLElement) IfLANG(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.LANG(v)
}

func (e *TableHTMLElement) RemoveLANG(v string) *TableHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *TableHTMLElement) NONCE(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *TableHTMLElement) IfNONCE(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.NONCE(v)
}

func (e *TableHTMLElement) RemoveNONCE(v string) *TableHTMLElement {
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
func (e *TableHTMLElement) POPOVER(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *TableHTMLElement) IfPOPOVER(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.POPOVER(v)
}

func (e *TableHTMLElement) RemovePOPOVER(v string) *TableHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *TableHTMLElement) SLOT(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *TableHTMLElement) IfSLOT(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.SLOT(v)
}

func (e *TableHTMLElement) RemoveSLOT(v string) *TableHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *TableHTMLElement) SPELLCHECK(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *TableHTMLElement) IfSPELLCHECK(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.SPELLCHECK(v)
}

func (e *TableHTMLElement) RemoveSPELLCHECK(v string) *TableHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TableHTMLElement) STYLE(k,v string) *TableHTMLElement {
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

func (e *TableHTMLElement) IfSTYLE(cond bool, k string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.STYLE(k, "")
}

func (e *TableHTMLElement) RemoveSTYLE(k string) *TableHTMLElement {
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
func (e *TableHTMLElement) TABINDEX(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *TableHTMLElement) IfTABINDEX(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.TABINDEX(v)
}

func (e *TableHTMLElement) RemoveTABINDEX(v string) *TableHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *TableHTMLElement) TITLE(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *TableHTMLElement) IfTITLE(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.TITLE(v)
}

func (e *TableHTMLElement) RemoveTITLE(v string) *TableHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *TableHTMLElement) TRANSLATE(v string) *TableHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *TableHTMLElement) IfTRANSLATE(cond bool, v string) *TableHTMLElement {
    if !cond {
        return e
    }
    return e.TRANSLATE(v)
}

func (e *TableHTMLElement) RemoveTRANSLATE(v string) *TableHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
