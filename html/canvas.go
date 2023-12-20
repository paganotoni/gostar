package html

import (
	"fmt"
)

type CanvasHTMLElement struct {
	*Element
}

func CANVAS(children ...ElementBuilder) *CanvasHTMLElement {
	return &CanvasHTMLElement{
		Element: &Element{
			Tag:           elementTagCANVAS,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *CanvasHTMLElement) Children(children ...ElementBuilder) *CanvasHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *CanvasHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *CanvasHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *CanvasHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *CanvasHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *CanvasHTMLElement) Text(text string) *CanvasHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *CanvasHTMLElement) TextF(format string, args ...any) *CanvasHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *CanvasHTMLElement) Escaped(text string) *CanvasHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *CanvasHTMLElement) EscapedF(format string, args ...any) *CanvasHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *CanvasHTMLElement) CustomData(key, value string) *CanvasHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *CanvasHTMLElement) CustomDataRemove(key string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) ACCESSKEY(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *CanvasHTMLElement) IfACCESSKEY(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *CanvasHTMLElement) RemoveACCESSKEY(v string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) AUTOCAPITALIZE(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *CanvasHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *CanvasHTMLElement) RemoveAUTOCAPITALIZE(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *CanvasHTMLElement) AUTOFOCUS() *CanvasHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *CanvasHTMLElement) IfAUTOFOCUS(cond bool) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *CanvasHTMLElement) RemoveAUTOFOCUS() *CanvasHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *CanvasHTMLElement) SetAUTOFOCUS(b bool) *CanvasHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *CanvasHTMLElement) CLASS(v string) *CanvasHTMLElement {
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

func (e *CanvasHTMLElement) IfCLASS(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *CanvasHTMLElement) SetCLASS(v string) *CanvasHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *CanvasHTMLElement) RemoveCLASS(v string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) CONTENTEDITABLE(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *CanvasHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *CanvasHTMLElement) RemoveCONTENTEDITABLE(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *CanvasHTMLElement) DIR(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *CanvasHTMLElement) IfDIR(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *CanvasHTMLElement) RemoveDIR(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *CanvasHTMLElement) DRAGGABLE(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *CanvasHTMLElement) IfDRAGGABLE(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *CanvasHTMLElement) RemoveDRAGGABLE(v string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) ENTERKEYHINT(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *CanvasHTMLElement) IfENTERKEYHINT(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *CanvasHTMLElement) RemoveENTERKEYHINT(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *CanvasHTMLElement) HEIGHT(v int) *CanvasHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeHEIGHTKey] = v
	return e
}

func (e *CanvasHTMLElement) IfHEIGHT(cond bool, v int) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.HEIGHT(v)
}

func (e *CanvasHTMLElement) RemoveHEIGHT(v int) *CanvasHTMLElement {
	delete(e.IntAttributes, attributeHEIGHTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *CanvasHTMLElement) HIDDEN(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *CanvasHTMLElement) IfHIDDEN(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *CanvasHTMLElement) RemoveHIDDEN(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *CanvasHTMLElement) ID(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *CanvasHTMLElement) IfID(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *CanvasHTMLElement) RemoveID(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *CanvasHTMLElement) INERT() *CanvasHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *CanvasHTMLElement) IfINERT(cond bool) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *CanvasHTMLElement) RemoveINERT() *CanvasHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *CanvasHTMLElement) SetINERT(b bool) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) INPUTMODE(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *CanvasHTMLElement) IfINPUTMODE(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *CanvasHTMLElement) RemoveINPUTMODE(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *CanvasHTMLElement) IS(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *CanvasHTMLElement) IfIS(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *CanvasHTMLElement) RemoveIS(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *CanvasHTMLElement) ITEMID(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *CanvasHTMLElement) IfITEMID(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *CanvasHTMLElement) RemoveITEMID(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *CanvasHTMLElement) ITEMPROP(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *CanvasHTMLElement) IfITEMPROP(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *CanvasHTMLElement) RemoveITEMPROP(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *CanvasHTMLElement) ITEMREF(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *CanvasHTMLElement) IfITEMREF(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *CanvasHTMLElement) RemoveITEMREF(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *CanvasHTMLElement) ITEMSCOPE() *CanvasHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *CanvasHTMLElement) IfITEMSCOPE(cond bool) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *CanvasHTMLElement) RemoveITEMSCOPE() *CanvasHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *CanvasHTMLElement) SetITEMSCOPE(b bool) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) ITEMTYPE(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *CanvasHTMLElement) IfITEMTYPE(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *CanvasHTMLElement) RemoveITEMTYPE(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *CanvasHTMLElement) LANG(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *CanvasHTMLElement) IfLANG(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *CanvasHTMLElement) RemoveLANG(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *CanvasHTMLElement) NONCE(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *CanvasHTMLElement) IfNONCE(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *CanvasHTMLElement) RemoveNONCE(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *CanvasHTMLElement) POPOVER(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *CanvasHTMLElement) IfPOPOVER(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *CanvasHTMLElement) RemovePOPOVER(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *CanvasHTMLElement) SLOT(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *CanvasHTMLElement) IfSLOT(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *CanvasHTMLElement) RemoveSLOT(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *CanvasHTMLElement) SPELLCHECK(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *CanvasHTMLElement) IfSPELLCHECK(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *CanvasHTMLElement) RemoveSPELLCHECK(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *CanvasHTMLElement) STYLE(k, v string) *CanvasHTMLElement {
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

func (e *CanvasHTMLElement) IfSTYLE(cond bool, k string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *CanvasHTMLElement) RemoveSTYLE(k string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) TABINDEX(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *CanvasHTMLElement) IfTABINDEX(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *CanvasHTMLElement) RemoveTABINDEX(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *CanvasHTMLElement) TITLE(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *CanvasHTMLElement) IfTITLE(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *CanvasHTMLElement) RemoveTITLE(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *CanvasHTMLElement) TRANSLATE(v string) *CanvasHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *CanvasHTMLElement) IfTRANSLATE(cond bool, v string) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *CanvasHTMLElement) RemoveTRANSLATE(v string) *CanvasHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *CanvasHTMLElement) WIDTH(v int) *CanvasHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeWIDTHKey] = v
	return e
}

func (e *CanvasHTMLElement) IfWIDTH(cond bool, v int) *CanvasHTMLElement {
	if !cond {
		return e
	}
	return e.WIDTH(v)
}

func (e *CanvasHTMLElement) RemoveWIDTH(v int) *CanvasHTMLElement {
	delete(e.IntAttributes, attributeWIDTHKey)
	return e
}
