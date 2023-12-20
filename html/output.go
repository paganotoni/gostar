package html

import (
	"fmt"
)

type OutputHTMLElement struct {
	*Element
}

func OUTPUT(children ...ElementBuilder) *OutputHTMLElement {
	return &OutputHTMLElement{
		Element: &Element{
			Tag:           elementTagOUTPUT,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *OutputHTMLElement) Children(children ...ElementBuilder) *OutputHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *OutputHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *OutputHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *OutputHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *OutputHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *OutputHTMLElement) Text(text string) *OutputHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *OutputHTMLElement) TextF(format string, args ...any) *OutputHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *OutputHTMLElement) Escaped(text string) *OutputHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *OutputHTMLElement) EscapedF(format string, args ...any) *OutputHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *OutputHTMLElement) CustomData(key, value string) *OutputHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *OutputHTMLElement) CustomDataRemove(key string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) ACCESSKEY(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *OutputHTMLElement) IfACCESSKEY(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *OutputHTMLElement) RemoveACCESSKEY(v string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) AUTOCAPITALIZE(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *OutputHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *OutputHTMLElement) RemoveAUTOCAPITALIZE(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *OutputHTMLElement) AUTOFOCUS() *OutputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *OutputHTMLElement) IfAUTOFOCUS(cond bool) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *OutputHTMLElement) RemoveAUTOFOCUS() *OutputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *OutputHTMLElement) SetAUTOFOCUS(b bool) *OutputHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *OutputHTMLElement) CLASS(v string) *OutputHTMLElement {
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

func (e *OutputHTMLElement) IfCLASS(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *OutputHTMLElement) SetCLASS(v string) *OutputHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *OutputHTMLElement) RemoveCLASS(v string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) CONTENTEDITABLE(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *OutputHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *OutputHTMLElement) RemoveCONTENTEDITABLE(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *OutputHTMLElement) DIR(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *OutputHTMLElement) IfDIR(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *OutputHTMLElement) RemoveDIR(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *OutputHTMLElement) DRAGGABLE(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *OutputHTMLElement) IfDRAGGABLE(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *OutputHTMLElement) RemoveDRAGGABLE(v string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) ENTERKEYHINT(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *OutputHTMLElement) IfENTERKEYHINT(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *OutputHTMLElement) RemoveENTERKEYHINT(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// FOR sets the "for" attribute.
// Specifies controls from which the output was calculated
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *OutputHTMLElement) FOR(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORKey] = v
	return e
}

func (e *OutputHTMLElement) IfFOR(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.FOR(v)
}

func (e *OutputHTMLElement) RemoveFOR(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeFORKey)
	return e
}

// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//   - id
func (e *OutputHTMLElement) FORM(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMKey] = v
	return e
}

func (e *OutputHTMLElement) IfFORM(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.FORM(v)
}

func (e *OutputHTMLElement) RemoveFORM(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeFORMKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *OutputHTMLElement) HIDDEN(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *OutputHTMLElement) IfHIDDEN(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *OutputHTMLElement) RemoveHIDDEN(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *OutputHTMLElement) ID(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *OutputHTMLElement) IfID(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *OutputHTMLElement) RemoveID(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *OutputHTMLElement) INERT() *OutputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *OutputHTMLElement) IfINERT(cond bool) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *OutputHTMLElement) RemoveINERT() *OutputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *OutputHTMLElement) SetINERT(b bool) *OutputHTMLElement {
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
func (e *OutputHTMLElement) INPUTMODE(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *OutputHTMLElement) IfINPUTMODE(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *OutputHTMLElement) RemoveINPUTMODE(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *OutputHTMLElement) IS(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *OutputHTMLElement) IfIS(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *OutputHTMLElement) RemoveIS(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *OutputHTMLElement) ITEMID(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *OutputHTMLElement) IfITEMID(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *OutputHTMLElement) RemoveITEMID(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *OutputHTMLElement) ITEMPROP(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *OutputHTMLElement) IfITEMPROP(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *OutputHTMLElement) RemoveITEMPROP(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *OutputHTMLElement) ITEMREF(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *OutputHTMLElement) IfITEMREF(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *OutputHTMLElement) RemoveITEMREF(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *OutputHTMLElement) ITEMSCOPE() *OutputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *OutputHTMLElement) IfITEMSCOPE(cond bool) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *OutputHTMLElement) RemoveITEMSCOPE() *OutputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *OutputHTMLElement) SetITEMSCOPE(b bool) *OutputHTMLElement {
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
func (e *OutputHTMLElement) ITEMTYPE(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *OutputHTMLElement) IfITEMTYPE(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *OutputHTMLElement) RemoveITEMTYPE(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *OutputHTMLElement) LANG(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *OutputHTMLElement) IfLANG(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *OutputHTMLElement) RemoveLANG(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//   - text
func (e *OutputHTMLElement) NAME(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNAMEKey] = v
	return e
}

func (e *OutputHTMLElement) IfNAME(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.NAME(v)
}

func (e *OutputHTMLElement) RemoveNAME(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeNAMEKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *OutputHTMLElement) NONCE(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *OutputHTMLElement) IfNONCE(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *OutputHTMLElement) RemoveNONCE(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *OutputHTMLElement) POPOVER(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *OutputHTMLElement) IfPOPOVER(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *OutputHTMLElement) RemovePOPOVER(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *OutputHTMLElement) SLOT(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *OutputHTMLElement) IfSLOT(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *OutputHTMLElement) RemoveSLOT(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *OutputHTMLElement) SPELLCHECK(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *OutputHTMLElement) IfSPELLCHECK(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *OutputHTMLElement) RemoveSPELLCHECK(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *OutputHTMLElement) STYLE(k, v string) *OutputHTMLElement {
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

func (e *OutputHTMLElement) IfSTYLE(cond bool, k string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *OutputHTMLElement) RemoveSTYLE(k string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) TABINDEX(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *OutputHTMLElement) IfTABINDEX(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *OutputHTMLElement) RemoveTABINDEX(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *OutputHTMLElement) TITLE(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *OutputHTMLElement) IfTITLE(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *OutputHTMLElement) RemoveTITLE(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *OutputHTMLElement) TRANSLATE(v string) *OutputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *OutputHTMLElement) IfTRANSLATE(cond bool, v string) *OutputHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *OutputHTMLElement) RemoveTRANSLATE(v string) *OutputHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
