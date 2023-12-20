package html

import (
	"fmt"
)

type SectionHTMLElement struct {
	*Element
}

func SECTION(children ...ElementBuilder) *SectionHTMLElement {
	return &SectionHTMLElement{
		Element: &Element{
			Tag:           elementTagSECTION,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *SectionHTMLElement) Children(children ...ElementBuilder) *SectionHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SectionHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SectionHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SectionHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SectionHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SectionHTMLElement) Text(text string) *SectionHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SectionHTMLElement) TextF(format string, args ...any) *SectionHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SectionHTMLElement) Escaped(text string) *SectionHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SectionHTMLElement) EscapedF(format string, args ...any) *SectionHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SectionHTMLElement) CustomData(key, value string) *SectionHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SectionHTMLElement) CustomDataRemove(key string) *SectionHTMLElement {
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
func (e *SectionHTMLElement) ACCESSKEY(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SectionHTMLElement) IfACCESSKEY(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SectionHTMLElement) RemoveACCESSKEY(v string) *SectionHTMLElement {
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
func (e *SectionHTMLElement) AUTOCAPITALIZE(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SectionHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SectionHTMLElement) RemoveAUTOCAPITALIZE(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SectionHTMLElement) AUTOFOCUS() *SectionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SectionHTMLElement) IfAUTOFOCUS(cond bool) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SectionHTMLElement) RemoveAUTOFOCUS() *SectionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *SectionHTMLElement) SetAUTOFOCUS(b bool) *SectionHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *SectionHTMLElement) CLASS(v string) *SectionHTMLElement {
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

func (e *SectionHTMLElement) IfCLASS(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SectionHTMLElement) SetCLASS(v string) *SectionHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SectionHTMLElement) RemoveCLASS(v string) *SectionHTMLElement {
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
func (e *SectionHTMLElement) CONTENTEDITABLE(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SectionHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SectionHTMLElement) RemoveCONTENTEDITABLE(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *SectionHTMLElement) DIR(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SectionHTMLElement) IfDIR(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SectionHTMLElement) RemoveDIR(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *SectionHTMLElement) DRAGGABLE(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SectionHTMLElement) IfDRAGGABLE(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SectionHTMLElement) RemoveDRAGGABLE(v string) *SectionHTMLElement {
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
func (e *SectionHTMLElement) ENTERKEYHINT(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SectionHTMLElement) IfENTERKEYHINT(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SectionHTMLElement) RemoveENTERKEYHINT(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *SectionHTMLElement) HIDDEN(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SectionHTMLElement) IfHIDDEN(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SectionHTMLElement) RemoveHIDDEN(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SectionHTMLElement) ID(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SectionHTMLElement) IfID(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SectionHTMLElement) RemoveID(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SectionHTMLElement) INERT() *SectionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SectionHTMLElement) IfINERT(cond bool) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SectionHTMLElement) RemoveINERT() *SectionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *SectionHTMLElement) SetINERT(b bool) *SectionHTMLElement {
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
func (e *SectionHTMLElement) INPUTMODE(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SectionHTMLElement) IfINPUTMODE(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SectionHTMLElement) RemoveINPUTMODE(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *SectionHTMLElement) IS(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SectionHTMLElement) IfIS(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SectionHTMLElement) RemoveIS(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SectionHTMLElement) ITEMID(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SectionHTMLElement) IfITEMID(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SectionHTMLElement) RemoveITEMID(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *SectionHTMLElement) ITEMPROP(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SectionHTMLElement) IfITEMPROP(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SectionHTMLElement) RemoveITEMPROP(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SectionHTMLElement) ITEMREF(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SectionHTMLElement) IfITEMREF(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SectionHTMLElement) RemoveITEMREF(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SectionHTMLElement) ITEMSCOPE() *SectionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SectionHTMLElement) IfITEMSCOPE(cond bool) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SectionHTMLElement) RemoveITEMSCOPE() *SectionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *SectionHTMLElement) SetITEMSCOPE(b bool) *SectionHTMLElement {
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
func (e *SectionHTMLElement) ITEMTYPE(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SectionHTMLElement) IfITEMTYPE(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SectionHTMLElement) RemoveITEMTYPE(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SectionHTMLElement) LANG(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SectionHTMLElement) IfLANG(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SectionHTMLElement) RemoveLANG(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SectionHTMLElement) NONCE(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SectionHTMLElement) IfNONCE(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SectionHTMLElement) RemoveNONCE(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *SectionHTMLElement) POPOVER(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SectionHTMLElement) IfPOPOVER(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SectionHTMLElement) RemovePOPOVER(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SectionHTMLElement) SLOT(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SectionHTMLElement) IfSLOT(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SectionHTMLElement) RemoveSLOT(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *SectionHTMLElement) SPELLCHECK(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SectionHTMLElement) IfSPELLCHECK(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SectionHTMLElement) RemoveSPELLCHECK(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SectionHTMLElement) STYLE(k, v string) *SectionHTMLElement {
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

func (e *SectionHTMLElement) IfSTYLE(cond bool, k string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SectionHTMLElement) RemoveSTYLE(k string) *SectionHTMLElement {
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
func (e *SectionHTMLElement) TABINDEX(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SectionHTMLElement) IfTABINDEX(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SectionHTMLElement) RemoveTABINDEX(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SectionHTMLElement) TITLE(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SectionHTMLElement) IfTITLE(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SectionHTMLElement) RemoveTITLE(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *SectionHTMLElement) TRANSLATE(v string) *SectionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SectionHTMLElement) IfTRANSLATE(cond bool, v string) *SectionHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SectionHTMLElement) RemoveTRANSLATE(v string) *SectionHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
