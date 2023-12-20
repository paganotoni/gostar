package  html

import (
    "fmt"
)

type BlockquoteHTMLElement struct {
    *Element
}

func BLOCKQUOTE(children ...fmt.Stringer) *BlockquoteHTMLElement {
    return &BlockquoteHTMLElement{
        Element: &Element{
            Tag: "blockquote",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *BlockquoteHTMLElement) Children(children ...fmt.Stringer) *BlockquoteHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *BlockquoteHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *BlockquoteHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *BlockquoteHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *BlockquoteHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *BlockquoteHTMLElement) Text(text string) *BlockquoteHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *BlockquoteHTMLElement) TextF(format string, args ...any) *BlockquoteHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *BlockquoteHTMLElement) Raw(text string) *BlockquoteHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *BlockquoteHTMLElement) RawF(format string, args ...any) *BlockquoteHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *BlockquoteHTMLElement) CustomData(key, value string) *BlockquoteHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *BlockquoteHTMLElement) CustomDataRemove(key string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) ACCESSKEY(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveACCESSKEY(v string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) AUTOCAPITALIZE(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveAUTOCAPITALIZE(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *BlockquoteHTMLElement) AUTOFOCUS() *BlockquoteHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *BlockquoteHTMLElement) RemoveAUTOFOCUS() *BlockquoteHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *BlockquoteHTMLElement) SetAUTOFOCUS(b bool) *BlockquoteHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CITE sets the "cite" attribute.
// Link to the source of the quotation or more information about the edit
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *BlockquoteHTMLElement) CITE(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["cite"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveCITE(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "cite")
    return e
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *BlockquoteHTMLElement) CLASS(v string) *BlockquoteHTMLElement {
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

func (e *BlockquoteHTMLElement) SetCLASS(v string) *BlockquoteHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *BlockquoteHTMLElement) RemoveCLASS(v string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) CONTENTEDITABLE(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveCONTENTEDITABLE(v string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) DIR(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveDIR(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *BlockquoteHTMLElement) DRAGGABLE(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveDRAGGABLE(v string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) ENTERKEYHINT(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveENTERKEYHINT(v string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) HIDDEN(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveHIDDEN(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *BlockquoteHTMLElement) ID(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveID(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *BlockquoteHTMLElement) INERT() *BlockquoteHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *BlockquoteHTMLElement) RemoveINERT() *BlockquoteHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *BlockquoteHTMLElement) SetINERT(b bool) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) INPUTMODE(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveINPUTMODE(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *BlockquoteHTMLElement) IS(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveIS(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *BlockquoteHTMLElement) ITEMID(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveITEMID(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *BlockquoteHTMLElement) ITEMPROP(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveITEMPROP(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *BlockquoteHTMLElement) ITEMREF(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveITEMREF(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *BlockquoteHTMLElement) ITEMSCOPE() *BlockquoteHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *BlockquoteHTMLElement) RemoveITEMSCOPE() *BlockquoteHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *BlockquoteHTMLElement) SetITEMSCOPE(b bool) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) ITEMTYPE(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveITEMTYPE(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *BlockquoteHTMLElement) LANG(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveLANG(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *BlockquoteHTMLElement) NONCE(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveNONCE(v string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) POPOVER(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemovePOPOVER(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *BlockquoteHTMLElement) SLOT(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveSLOT(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *BlockquoteHTMLElement) SPELLCHECK(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveSPELLCHECK(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *BlockquoteHTMLElement) STYLE(k,v string) *BlockquoteHTMLElement {
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

func (e *BlockquoteHTMLElement) RemoveSTYLE(k string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) TABINDEX(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveTABINDEX(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *BlockquoteHTMLElement) TITLE(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveTITLE(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *BlockquoteHTMLElement) TRANSLATE(v string) *BlockquoteHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *BlockquoteHTMLElement) RemoveTRANSLATE(v string) *BlockquoteHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
