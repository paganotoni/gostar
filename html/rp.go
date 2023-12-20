package html

import (
	"fmt"
)

type RpHTMLElement struct {
	*Element
}

func RP(children ...ElementBuilder) *RpHTMLElement {
	return &RpHTMLElement{
		Element: &Element{
			Tag:           elementTagRP,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *RpHTMLElement) Children(children ...ElementBuilder) *RpHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *RpHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *RpHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *RpHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *RpHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *RpHTMLElement) Text(text string) *RpHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *RpHTMLElement) TextF(format string, args ...any) *RpHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *RpHTMLElement) Escaped(text string) *RpHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *RpHTMLElement) EscapedF(format string, args ...any) *RpHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *RpHTMLElement) CustomData(key, value string) *RpHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *RpHTMLElement) CustomDataRemove(key string) *RpHTMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	delete(e.CustomDataAttributes, key)
	return e
}

// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//   - identical_to
//   - ordered_set_of_unique_space_separated_tokens
func (e *RpHTMLElement) ACCESSKEY(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *RpHTMLElement) IfACCESSKEY(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *RpHTMLElement) RemoveACCESSKEY(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeACCESSKEYKey)
	return e
}

// AUTOCAPITALIZE sets the "autocapitalize" attribute.
// Recommended autocapitalization behavior (for supported input methods)
// Values values are constrained to:
//   - characters
//   - none
//   - off
//   - on
//   - sentences
//   - words
func (e *RpHTMLElement) AUTOCAPITALIZE(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *RpHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *RpHTMLElement) RemoveAUTOCAPITALIZE(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *RpHTMLElement) AUTOFOCUS() *RpHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *RpHTMLElement) IfAUTOFOCUS(cond bool) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *RpHTMLElement) RemoveAUTOFOCUS() *RpHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *RpHTMLElement) SetAUTOFOCUS(b bool) *RpHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *RpHTMLElement) CLASS(v string) *RpHTMLElement {
	if e.DelimitedStringAttributes == nil {
		e.DelimitedStringAttributes = map[string]*DelimitedString{}
	}
	kv, ok := e.DelimitedStringAttributes[attributeCLASSKey]
	if !ok {
		kv = NewSpaceDelimitedString()
		e.DelimitedStringAttributes[attributeCLASSKey] = kv
	}
	kv.Add(v)
	return e
}

func (e *RpHTMLElement) IfCLASS(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *RpHTMLElement) SetCLASS(v string) *RpHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *RpHTMLElement) RemoveCLASS(v string) *RpHTMLElement {
	kv, ok := e.DelimitedStringAttributes[attributeCLASSKey]
	if !ok {
		return e
	}
	kv.Remove(v)
	return e
}

// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//   - false
//   - plaintext_only
//   - true
func (e *RpHTMLElement) CONTENTEDITABLE(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *RpHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *RpHTMLElement) RemoveCONTENTEDITABLE(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *RpHTMLElement) DIR(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *RpHTMLElement) IfDIR(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *RpHTMLElement) RemoveDIR(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *RpHTMLElement) DRAGGABLE(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *RpHTMLElement) IfDRAGGABLE(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *RpHTMLElement) RemoveDRAGGABLE(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeDRAGGABLEKey)
	return e
}

// ENTERKEYHINT sets the "enterkeyhint" attribute.
// Hint for selecting an enter key action
// Values values are constrained to:
//   - done
//   - enter
//   - go
//   - next
//   - previous
//   - search
//   - send
func (e *RpHTMLElement) ENTERKEYHINT(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *RpHTMLElement) IfENTERKEYHINT(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *RpHTMLElement) RemoveENTERKEYHINT(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *RpHTMLElement) HIDDEN(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *RpHTMLElement) IfHIDDEN(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *RpHTMLElement) RemoveHIDDEN(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *RpHTMLElement) ID(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *RpHTMLElement) IfID(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *RpHTMLElement) RemoveID(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *RpHTMLElement) INERT() *RpHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *RpHTMLElement) IfINERT(cond bool) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *RpHTMLElement) RemoveINERT() *RpHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *RpHTMLElement) SetINERT(b bool) *RpHTMLElement {
	if b {
		return e.INERT()
	}
	return e.RemoveINERT()
}

// INPUTMODE sets the "inputmode" attribute.
// Hint for selecting an input modality
// Values values are constrained to:
//   - decimal
//   - email
//   - none
//   - numeric
//   - search
//   - tel
//   - text
//   - url
func (e *RpHTMLElement) INPUTMODE(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *RpHTMLElement) IfINPUTMODE(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *RpHTMLElement) RemoveINPUTMODE(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *RpHTMLElement) IS(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *RpHTMLElement) IfIS(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *RpHTMLElement) RemoveIS(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *RpHTMLElement) ITEMID(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *RpHTMLElement) IfITEMID(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *RpHTMLElement) RemoveITEMID(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *RpHTMLElement) ITEMPROP(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *RpHTMLElement) IfITEMPROP(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *RpHTMLElement) RemoveITEMPROP(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *RpHTMLElement) ITEMREF(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *RpHTMLElement) IfITEMREF(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *RpHTMLElement) RemoveITEMREF(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *RpHTMLElement) ITEMSCOPE() *RpHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *RpHTMLElement) IfITEMSCOPE(cond bool) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *RpHTMLElement) RemoveITEMSCOPE() *RpHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *RpHTMLElement) SetITEMSCOPE(b bool) *RpHTMLElement {
	if b {
		return e.ITEMSCOPE()
	}
	return e.RemoveITEMSCOPE()
}

// ITEMTYPE sets the "itemtype" attribute.
// Item types of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *RpHTMLElement) ITEMTYPE(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *RpHTMLElement) IfITEMTYPE(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *RpHTMLElement) RemoveITEMTYPE(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *RpHTMLElement) LANG(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *RpHTMLElement) IfLANG(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *RpHTMLElement) RemoveLANG(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *RpHTMLElement) NONCE(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *RpHTMLElement) IfNONCE(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *RpHTMLElement) RemoveNONCE(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *RpHTMLElement) POPOVER(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *RpHTMLElement) IfPOPOVER(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *RpHTMLElement) RemovePOPOVER(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *RpHTMLElement) SLOT(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *RpHTMLElement) IfSLOT(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *RpHTMLElement) RemoveSLOT(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *RpHTMLElement) SPELLCHECK(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *RpHTMLElement) IfSPELLCHECK(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *RpHTMLElement) RemoveSPELLCHECK(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *RpHTMLElement) STYLE(k, v string) *RpHTMLElement {
	if e.DelimitedKVAttributes == nil {
		e.DelimitedKVAttributes = map[string]*DelimitedKVString{}
	}
	kv, ok := e.DelimitedKVAttributes[attributeSTYLEKey]
	if !ok {
		kv = NewEqualSemicolonDelimitedKVString()
		e.DelimitedKVAttributes[attributeSTYLEKey] = kv
	}
	kv.Add(k, v)
	return e
}

func (e *RpHTMLElement) IfSTYLE(cond bool, k string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *RpHTMLElement) RemoveSTYLE(k string) *RpHTMLElement {
	kv, ok := e.DelimitedKVAttributes[attributeSTYLEKey]
	if !ok {
		return e
	}
	kv.Remove(k)
	return e
}

// TABINDEX sets the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and
//
//	the relative order of the element for the purposes of sequential focus navigation
//
// Values values are constrained to:
//   - valid_integer
func (e *RpHTMLElement) TABINDEX(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *RpHTMLElement) IfTABINDEX(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *RpHTMLElement) RemoveTABINDEX(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *RpHTMLElement) TITLE(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *RpHTMLElement) IfTITLE(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *RpHTMLElement) RemoveTITLE(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *RpHTMLElement) TRANSLATE(v string) *RpHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *RpHTMLElement) IfTRANSLATE(cond bool, v string) *RpHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *RpHTMLElement) RemoveTRANSLATE(v string) *RpHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
