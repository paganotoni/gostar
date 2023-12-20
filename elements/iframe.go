package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type IframeElement struct {
    *gostar.Element
}

func IFRAME(children ...fmt.Stringer) *IframeElement {
    return &IframeElement{
        Element: &gostar.Element{
            Tag: "iframe",
            IsSelfClosing: true,
            Children: children,
        },
    }
}

func (e *IframeElement) AddChildren(children ...fmt.Stringer) *IframeElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *IframeElement) TEXT(text string) *IframeElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *IframeElement) RAW(text string) *IframeElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *IframeElement) CustomData(key, value string) *IframeElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *IframeElement) CustomDataRemove(key string) *IframeElement {
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
func (e *IframeElement) ACCESSKEY(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *IframeElement) RemoveACCESSKEY(v string) *IframeElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ALLOW sets the "allow" attribute.
// Permissions policy to be applied to the iframe's contents
// Values values are constrained to:
//  * serialized_permissions_policy
func (e *IframeElement) ALLOW(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["allow"] = v
    return e
}

func (e *IframeElement) RemoveALLOW(v string) *IframeElement {
    delete(e.StringAttributes, "allow")
    return e
}
// ALLOWFULLSCREEN sets the "allowfullscreen" attribute.
// Whether to allow the iframe's contents to use requestFullscreen()
// Values values are constrained to:
//  * boolean_attribute
func (e *IframeElement) ALLOWFULLSCREEN() *IframeElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["allowfullscreen"] = struct{}{}
    return e
}

func (e *IframeElement) RemoveALLOWFULLSCREEN() *IframeElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "allowfullscreen")
    return e
}

func (e *IframeElement) SetALLOWFULLSCREEN(b bool) *IframeElement {
    if b {
        return e.ALLOWFULLSCREEN()
    }
    return e.RemoveALLOWFULLSCREEN()
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
func (e *IframeElement) AUTOCAPITALIZE(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *IframeElement) RemoveAUTOCAPITALIZE(v string) *IframeElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *IframeElement) AUTOFOCUS() *IframeElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *IframeElement) RemoveAUTOFOCUS() *IframeElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *IframeElement) SetAUTOFOCUS(b bool) *IframeElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *IframeElement) CLASS(v string) *IframeElement {
    if e.DelimitedStringAttributes == nil {
        e.DelimitedStringAttributes = map[string]*gostar.DelimitedString{}
    }
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        kv = gostar.NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = kv
    }
    kv.Add(v)
    return e
}

func (e *IframeElement) SetCLASS(v string) *IframeElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *IframeElement) RemoveCLASS(v string) *IframeElement {
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
func (e *IframeElement) CONTENTEDITABLE(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *IframeElement) RemoveCONTENTEDITABLE(v string) *IframeElement {
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
func (e *IframeElement) DIR(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *IframeElement) RemoveDIR(v string) *IframeElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *IframeElement) DRAGGABLE(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *IframeElement) RemoveDRAGGABLE(v string) *IframeElement {
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
func (e *IframeElement) ENTERKEYHINT(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *IframeElement) RemoveENTERKEYHINT(v string) *IframeElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *IframeElement) HEIGHT(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["height"] = v
    return e
}

func (e *IframeElement) RemoveHEIGHT(v string) *IframeElement {
    delete(e.StringAttributes, "height")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *IframeElement) HIDDEN(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *IframeElement) RemoveHIDDEN(v string) *IframeElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *IframeElement) ID(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *IframeElement) RemoveID(v string) *IframeElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *IframeElement) INERT() *IframeElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *IframeElement) RemoveINERT() *IframeElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *IframeElement) SetINERT(b bool) *IframeElement {
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
func (e *IframeElement) INPUTMODE(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *IframeElement) RemoveINPUTMODE(v string) *IframeElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *IframeElement) IS(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *IframeElement) RemoveIS(v string) *IframeElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *IframeElement) ITEMID(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *IframeElement) RemoveITEMID(v string) *IframeElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *IframeElement) ITEMPROP(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *IframeElement) RemoveITEMPROP(v string) *IframeElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *IframeElement) ITEMREF(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *IframeElement) RemoveITEMREF(v string) *IframeElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *IframeElement) ITEMSCOPE() *IframeElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *IframeElement) RemoveITEMSCOPE() *IframeElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *IframeElement) SetITEMSCOPE(b bool) *IframeElement {
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
func (e *IframeElement) ITEMTYPE(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *IframeElement) RemoveITEMTYPE(v string) *IframeElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *IframeElement) LANG(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *IframeElement) RemoveLANG(v string) *IframeElement {
    delete(e.StringAttributes, "lang")
    return e
}
// LOADING sets the "loading" attribute.
// Used when determining loading deferral
// Values values are constrained to:
//  * lazy
//  * lazy
//  * eager
//  * eager
func (e *IframeElement) LOADING(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["loading"] = v
    return e
}

func (e *IframeElement) RemoveLOADING(v string) *IframeElement {
    delete(e.StringAttributes, "loading")
    return e
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *IframeElement) NAME(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *IframeElement) RemoveNAME(v string) *IframeElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *IframeElement) NONCE(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *IframeElement) RemoveNONCE(v string) *IframeElement {
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
func (e *IframeElement) POPOVER(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *IframeElement) RemovePOPOVER(v string) *IframeElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//  * referrer_policy
func (e *IframeElement) REFERRERPOLICY(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["referrerpolicy"] = v
    return e
}

func (e *IframeElement) RemoveREFERRERPOLICY(v string) *IframeElement {
    delete(e.StringAttributes, "referrerpolicy")
    return e
}
// SANDBOX sets the "sandbox" attribute.
// Security rules for nested content
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * ascii_case_insensitive
//  * allow_downloads
//  * allow_downloads
//  * allow_forms
//  * allow_forms
//  * allow_modals
//  * allow_modals
//  * allow_orientation_lock
//  * allow_orientation_lock
//  * allow_pointer_lock
//  * allow_pointer_lock
//  * allow_popups
//  * allow_popups
//  * allow_popups_to_escape_sandbox
//  * allow_popups_to_escape_sandbox
//  * allow_presentation
//  * allow_presentation
//  * allow_same_origin
//  * allow_same_origin
//  * allow_scripts
//  * allow_scripts
//  * allow_top_navigation
//  * allow_top_navigation
//  * allow_top_navigation_by_user_activation
//  * allow_top_navigation_by_user_activation
//  * allow_top_navigation_to_custom_protocols
//  * allow_top_navigation_to_custom_protocols
func (e *IframeElement) SANDBOX(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["sandbox"] = v
    return e
}

func (e *IframeElement) RemoveSANDBOX(v string) *IframeElement {
    delete(e.StringAttributes, "sandbox")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *IframeElement) SLOT(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *IframeElement) RemoveSLOT(v string) *IframeElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *IframeElement) SPELLCHECK(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *IframeElement) RemoveSPELLCHECK(v string) *IframeElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *IframeElement) SRC(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *IframeElement) RemoveSRC(v string) *IframeElement {
    delete(e.StringAttributes, "src")
    return e
}
// SRCDOC sets the "srcdoc" attribute.
// A document to render in the iframe
// Values values are constrained to:
//  * an_iframe_srcdoc_document
//  * iframe
//  * srcdoc
func (e *IframeElement) SRCDOC(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["srcdoc"] = v
    return e
}

func (e *IframeElement) RemoveSRCDOC(v string) *IframeElement {
    delete(e.StringAttributes, "srcdoc")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *IframeElement) STYLE(k,v string) *IframeElement {
    if e.DelimitedKVAttributes == nil {
        e.DelimitedKVAttributes = map[string]*gostar.DelimitedKVString{}
    }
    kv, ok := e.DelimitedKVAttributes["style"]
    if !ok {
        kv = gostar.NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = kv
    }
    kv.Add(k,v)
    return e
}

func (e *IframeElement) RemoveSTYLE(k string) *IframeElement {
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
func (e *IframeElement) TABINDEX(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *IframeElement) RemoveTABINDEX(v string) *IframeElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *IframeElement) TITLE(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *IframeElement) RemoveTITLE(v string) *IframeElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *IframeElement) TRANSLATE(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *IframeElement) RemoveTRANSLATE(v string) *IframeElement {
    delete(e.StringAttributes, "translate")
    return e
}
// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *IframeElement) WIDTH(v string) *IframeElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["width"] = v
    return e
}

func (e *IframeElement) RemoveWIDTH(v string) *IframeElement {
    delete(e.StringAttributes, "width")
    return e
}
