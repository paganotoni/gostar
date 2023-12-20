package html

import (
	"fmt"
)

type BlockquoteHTMLElement struct {
	*Element
}

func BLOCKQUOTE(children ...ElementBuilder) *BlockquoteHTMLElement {
	return &BlockquoteHTMLElement{
		Element: &Element{
			Tag:           elementTagBLOCKQUOTE,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *BlockquoteHTMLElement) Children(children ...ElementBuilder) *BlockquoteHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *BlockquoteHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *BlockquoteHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *BlockquoteHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *BlockquoteHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *BlockquoteHTMLElement) Text(text string) *BlockquoteHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *BlockquoteHTMLElement) TextF(format string, args ...any) *BlockquoteHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *BlockquoteHTMLElement) Escaped(text string) *BlockquoteHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *BlockquoteHTMLElement) EscapedF(format string, args ...any) *BlockquoteHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
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
//   - identical_to
//   - ordered_set_of_unique_space_separated_tokens
func (e *BlockquoteHTMLElement) ACCESSKEY(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfACCESSKEY(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *BlockquoteHTMLElement) RemoveACCESSKEY(v string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) AUTOCAPITALIZE(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *BlockquoteHTMLElement) RemoveAUTOCAPITALIZE(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *BlockquoteHTMLElement) AUTOFOCUS() *BlockquoteHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *BlockquoteHTMLElement) IfAUTOFOCUS(cond bool) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *BlockquoteHTMLElement) RemoveAUTOFOCUS() *BlockquoteHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
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
//   - valid_url_potentially_surrounded_by_spaces
func (e *BlockquoteHTMLElement) CITE(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCITEKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfCITE(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.CITE(v)
}

func (e *BlockquoteHTMLElement) RemoveCITE(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeCITEKey)
	return e
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *BlockquoteHTMLElement) CLASS(v string) *BlockquoteHTMLElement {
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

func (e *BlockquoteHTMLElement) IfCLASS(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *BlockquoteHTMLElement) SetCLASS(v string) *BlockquoteHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *BlockquoteHTMLElement) RemoveCLASS(v string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) CONTENTEDITABLE(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *BlockquoteHTMLElement) RemoveCONTENTEDITABLE(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *BlockquoteHTMLElement) DIR(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfDIR(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *BlockquoteHTMLElement) RemoveDIR(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *BlockquoteHTMLElement) DRAGGABLE(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfDRAGGABLE(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *BlockquoteHTMLElement) RemoveDRAGGABLE(v string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) ENTERKEYHINT(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfENTERKEYHINT(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *BlockquoteHTMLElement) RemoveENTERKEYHINT(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *BlockquoteHTMLElement) HIDDEN(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfHIDDEN(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *BlockquoteHTMLElement) RemoveHIDDEN(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *BlockquoteHTMLElement) ID(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfID(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *BlockquoteHTMLElement) RemoveID(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *BlockquoteHTMLElement) INERT() *BlockquoteHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *BlockquoteHTMLElement) IfINERT(cond bool) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *BlockquoteHTMLElement) RemoveINERT() *BlockquoteHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
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
//   - decimal
//   - email
//   - none
//   - numeric
//   - search
//   - tel
//   - text
//   - url
func (e *BlockquoteHTMLElement) INPUTMODE(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfINPUTMODE(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *BlockquoteHTMLElement) RemoveINPUTMODE(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *BlockquoteHTMLElement) IS(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfIS(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *BlockquoteHTMLElement) RemoveIS(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *BlockquoteHTMLElement) ITEMID(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfITEMID(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *BlockquoteHTMLElement) RemoveITEMID(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *BlockquoteHTMLElement) ITEMPROP(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfITEMPROP(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *BlockquoteHTMLElement) RemoveITEMPROP(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *BlockquoteHTMLElement) ITEMREF(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfITEMREF(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *BlockquoteHTMLElement) RemoveITEMREF(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *BlockquoteHTMLElement) ITEMSCOPE() *BlockquoteHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *BlockquoteHTMLElement) IfITEMSCOPE(cond bool) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *BlockquoteHTMLElement) RemoveITEMSCOPE() *BlockquoteHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
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
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *BlockquoteHTMLElement) ITEMTYPE(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfITEMTYPE(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *BlockquoteHTMLElement) RemoveITEMTYPE(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *BlockquoteHTMLElement) LANG(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfLANG(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *BlockquoteHTMLElement) RemoveLANG(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *BlockquoteHTMLElement) NONCE(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfNONCE(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *BlockquoteHTMLElement) RemoveNONCE(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *BlockquoteHTMLElement) POPOVER(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfPOPOVER(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *BlockquoteHTMLElement) RemovePOPOVER(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *BlockquoteHTMLElement) SLOT(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfSLOT(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *BlockquoteHTMLElement) RemoveSLOT(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *BlockquoteHTMLElement) SPELLCHECK(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfSPELLCHECK(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *BlockquoteHTMLElement) RemoveSPELLCHECK(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *BlockquoteHTMLElement) STYLE(k, v string) *BlockquoteHTMLElement {
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

func (e *BlockquoteHTMLElement) IfSTYLE(cond bool, k string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *BlockquoteHTMLElement) RemoveSTYLE(k string) *BlockquoteHTMLElement {
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
func (e *BlockquoteHTMLElement) TABINDEX(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfTABINDEX(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *BlockquoteHTMLElement) RemoveTABINDEX(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *BlockquoteHTMLElement) TITLE(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfTITLE(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *BlockquoteHTMLElement) RemoveTITLE(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *BlockquoteHTMLElement) TRANSLATE(v string) *BlockquoteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *BlockquoteHTMLElement) IfTRANSLATE(cond bool, v string) *BlockquoteHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *BlockquoteHTMLElement) RemoveTRANSLATE(v string) *BlockquoteHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
