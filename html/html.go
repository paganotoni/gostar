package  html

import (
    "fmt"
)

type HtmlHTMLElement struct {
    *Element
}

func HTML(children ...ElementBuilder) *HtmlHTMLElement {
    return &HtmlHTMLElement{
        Element: &Element{
            Tag: "html",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *HtmlHTMLElement) Children(children ...ElementBuilder) *HtmlHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *HtmlHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *HtmlHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *HtmlHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *HtmlHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *HtmlHTMLElement) Text(text string) *HtmlHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *HtmlHTMLElement) TextF(format string, args ...any) *HtmlHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *HtmlHTMLElement) Escaped(text string) *HtmlHTMLElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *HtmlHTMLElement) EscapedF(format string, args ...any) *HtmlHTMLElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *HtmlHTMLElement) CustomData(key, value string) *HtmlHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *HtmlHTMLElement) CustomDataRemove(key string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) ACCESSKEY(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *HtmlHTMLElement) IfACCESSKEY(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.ACCESSKEY(v)
}

func (e *HtmlHTMLElement) RemoveACCESSKEY(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) AUTOCAPITALIZE(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *HtmlHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOCAPITALIZE(v)
}

func (e *HtmlHTMLElement) RemoveAUTOCAPITALIZE(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *HtmlHTMLElement) AUTOFOCUS() *HtmlHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *HtmlHTMLElement) IfAUTOFOCUS(cond bool) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOFOCUS()
}

func (e *HtmlHTMLElement) RemoveAUTOFOCUS() *HtmlHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *HtmlHTMLElement) SetAUTOFOCUS(b bool) *HtmlHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *HtmlHTMLElement) CLASS(v string) *HtmlHTMLElement {
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

func (e *HtmlHTMLElement) IfCLASS(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.CLASS(v)
}

func (e *HtmlHTMLElement) SetCLASS(v string) *HtmlHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *HtmlHTMLElement) RemoveCLASS(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) CONTENTEDITABLE(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *HtmlHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.CONTENTEDITABLE(v)
}

func (e *HtmlHTMLElement) RemoveCONTENTEDITABLE(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) DIR(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *HtmlHTMLElement) IfDIR(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.DIR(v)
}

func (e *HtmlHTMLElement) RemoveDIR(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *HtmlHTMLElement) DRAGGABLE(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *HtmlHTMLElement) IfDRAGGABLE(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.DRAGGABLE(v)
}

func (e *HtmlHTMLElement) RemoveDRAGGABLE(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) ENTERKEYHINT(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *HtmlHTMLElement) IfENTERKEYHINT(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.ENTERKEYHINT(v)
}

func (e *HtmlHTMLElement) RemoveENTERKEYHINT(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) HIDDEN(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *HtmlHTMLElement) IfHIDDEN(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.HIDDEN(v)
}

func (e *HtmlHTMLElement) RemoveHIDDEN(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *HtmlHTMLElement) ID(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *HtmlHTMLElement) IfID(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.ID(v)
}

func (e *HtmlHTMLElement) RemoveID(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *HtmlHTMLElement) INERT() *HtmlHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *HtmlHTMLElement) IfINERT(cond bool) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.INERT()
}

func (e *HtmlHTMLElement) RemoveINERT() *HtmlHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *HtmlHTMLElement) SetINERT(b bool) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) INPUTMODE(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *HtmlHTMLElement) IfINPUTMODE(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.INPUTMODE(v)
}

func (e *HtmlHTMLElement) RemoveINPUTMODE(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *HtmlHTMLElement) IS(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *HtmlHTMLElement) IfIS(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.IS(v)
}

func (e *HtmlHTMLElement) RemoveIS(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *HtmlHTMLElement) ITEMID(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *HtmlHTMLElement) IfITEMID(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMID(v)
}

func (e *HtmlHTMLElement) RemoveITEMID(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *HtmlHTMLElement) ITEMPROP(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *HtmlHTMLElement) IfITEMPROP(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMPROP(v)
}

func (e *HtmlHTMLElement) RemoveITEMPROP(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *HtmlHTMLElement) ITEMREF(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *HtmlHTMLElement) IfITEMREF(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMREF(v)
}

func (e *HtmlHTMLElement) RemoveITEMREF(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *HtmlHTMLElement) ITEMSCOPE() *HtmlHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *HtmlHTMLElement) IfITEMSCOPE(cond bool) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMSCOPE()
}

func (e *HtmlHTMLElement) RemoveITEMSCOPE() *HtmlHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *HtmlHTMLElement) SetITEMSCOPE(b bool) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) ITEMTYPE(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *HtmlHTMLElement) IfITEMTYPE(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMTYPE(v)
}

func (e *HtmlHTMLElement) RemoveITEMTYPE(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *HtmlHTMLElement) LANG(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *HtmlHTMLElement) IfLANG(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.LANG(v)
}

func (e *HtmlHTMLElement) RemoveLANG(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *HtmlHTMLElement) NONCE(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *HtmlHTMLElement) IfNONCE(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.NONCE(v)
}

func (e *HtmlHTMLElement) RemoveNONCE(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) POPOVER(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *HtmlHTMLElement) IfPOPOVER(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.POPOVER(v)
}

func (e *HtmlHTMLElement) RemovePOPOVER(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *HtmlHTMLElement) SLOT(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *HtmlHTMLElement) IfSLOT(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.SLOT(v)
}

func (e *HtmlHTMLElement) RemoveSLOT(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *HtmlHTMLElement) SPELLCHECK(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *HtmlHTMLElement) IfSPELLCHECK(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.SPELLCHECK(v)
}

func (e *HtmlHTMLElement) RemoveSPELLCHECK(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *HtmlHTMLElement) STYLE(k,v string) *HtmlHTMLElement {
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

func (e *HtmlHTMLElement) IfSTYLE(cond bool, k string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.STYLE(k, "")
}

func (e *HtmlHTMLElement) RemoveSTYLE(k string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) TABINDEX(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *HtmlHTMLElement) IfTABINDEX(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.TABINDEX(v)
}

func (e *HtmlHTMLElement) RemoveTABINDEX(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *HtmlHTMLElement) TITLE(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *HtmlHTMLElement) IfTITLE(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.TITLE(v)
}

func (e *HtmlHTMLElement) RemoveTITLE(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *HtmlHTMLElement) TRANSLATE(v string) *HtmlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *HtmlHTMLElement) IfTRANSLATE(cond bool, v string) *HtmlHTMLElement {
    if !cond {
        return e
    }
    return e.TRANSLATE(v)
}

func (e *HtmlHTMLElement) RemoveTRANSLATE(v string) *HtmlHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
