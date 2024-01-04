// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package Shoelace Input is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

type SLINPUTElement struct {
	*Element
}

// Create a new SLINPUTElement element.
// This will create a new element with the tag
// "sl-input" during rendering.
func SL_INPUT(children ...ElementRenderer) *SLINPUTElement {
	e := NewElement("sl-input", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SLINPUTElement{Element: e}
}

func (e *SLINPUTElement) Children(children ...ElementRenderer) *SLINPUTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SLINPUTElement) IfChildren(condition bool, children ...ElementRenderer) *SLINPUTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SLINPUTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SLINPUTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SLINPUTElement) Attr(name string, value string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SLINPUTElement) Attrs(attrs ...string) *SLINPUTElement {
	if len(attrs)%2 != 0 {
		panic("attrs must be a multiple of 2")
	}
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for i := 0; i < len(attrs); i += 2 {
		k := attrs[i]
		v := attrs[i+1]
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SLINPUTElement) AttrsMap(attrs map[string]string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SLINPUTElement) Text(text string) *SLINPUTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SLINPUTElement) TextF(format string, args ...any) *SLINPUTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SLINPUTElement) IfText(condition bool, text string) *SLINPUTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SLINPUTElement) IfTextF(condition bool, format string, args ...any) *SLINPUTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SLINPUTElement) Escaped(text string) *SLINPUTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SLINPUTElement) IfEscaped(condition bool, text string) *SLINPUTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SLINPUTElement) EscapedF(format string, args ...any) *SLINPUTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SLINPUTElement) IfEscapedF(condition bool, format string, args ...any) *SLINPUTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SLINPUTElement) CustomData(key, value string) *SLINPUTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SLINPUTElement) IfCustomData(condition bool, key, value string) *SLINPUTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SLINPUTElement) CustomDataF(key, format string, args ...any) *SLINPUTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SLINPUTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SLINPUTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SLINPUTElement) CustomDataRemove(key string) *SLINPUTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

func (e *SLINPUTElement) SIZE(c SLInputSizeChoice) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("size", string(c))
	return e
}

type SLInputSizeChoice string

const (
	// Small
	SLInputSize_small SLInputSizeChoice = "small"
	// Medium
	SLInputSize_medium SLInputSizeChoice = "medium"
	// Large
	SLInputSize_large SLInputSizeChoice = "large"
)

// Remove the attribute SIZE from the element.
func (e *SLINPUTElement) SIZERemove(c SLInputSizeChoice) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("size")
	return e
}

func (e *SLINPUTElement) TYPE(c SLInputTypeChoice) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SLInputTypeChoice string

const (
	// Date
	SLInputType_date SLInputTypeChoice = "date"
	// Datetime Local
	SLInputType_datetime_local SLInputTypeChoice = "datetime-local"
	// Email
	SLInputType_email SLInputTypeChoice = "email"
	// Number
	SLInputType_number SLInputTypeChoice = "number"
	// Password
	SLInputType_password SLInputTypeChoice = "password"
	// Search
	SLInputType_search SLInputTypeChoice = "search"
	// Tel
	SLInputType_tel SLInputTypeChoice = "tel"
	// Text
	SLInputType_text SLInputTypeChoice = "text"
	// Time
	SLInputType_time SLInputTypeChoice = "time"
	// URL
	SLInputType_url SLInputTypeChoice = "url"
)

// Remove the attribute TYPE from the element.
func (e *SLINPUTElement) TYPERemove(c SLInputTypeChoice) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

func (e *SLINPUTElement) NAME(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("name", s)
	return e
}

func (e *SLINPUTElement) IfNAME(condition bool, s string) *SLINPUTElement {
	if condition {
		e.NAME(s)
	}
	return e
}

// Remove the attribute NAME from the element.
func (e *SLINPUTElement) NAMERemove(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("name")
	return e
}

func (e *SLINPUTElement) VALUE(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("value", s)
	return e
}

func (e *SLINPUTElement) IfVALUE(condition bool, s string) *SLINPUTElement {
	if condition {
		e.VALUE(s)
	}
	return e
}

// Remove the attribute VALUE from the element.
func (e *SLINPUTElement) VALUERemove(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("value")
	return e
}

func (e *SLINPUTElement) DEFAULT_VALUE(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("defaultValue", s)
	return e
}

func (e *SLINPUTElement) IfDEFAULT_VALUE(condition bool, s string) *SLINPUTElement {
	if condition {
		e.DEFAULT_VALUE(s)
	}
	return e
}

// Remove the attribute DEFAULT_VALUE from the element.
func (e *SLINPUTElement) DEFAULT_VALUERemove(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("defaultValue")
	return e
}

func (e *SLINPUTElement) FILLED() *SLINPUTElement {
	e.FILLEDSet(true)
	return e
}

func (e *SLINPUTElement) IfFILLED(condition bool) *SLINPUTElement {
	if condition {
		e.FILLEDSet(true)
	}
	return e
}

// Set the attribute FILLED to the value b explicitly.
func (e *SLINPUTElement) FILLEDSet(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("filled", b)
	return e
}

func (e *SLINPUTElement) IfSetFILLED(condition bool, b bool) *SLINPUTElement {
	if condition {
		e.FILLEDSet(b)
	}
	return e
}

// Remove the attribute FILLED from the element.
func (e *SLINPUTElement) FILLEDRemove(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("filled")
	return e
}

func (e *SLINPUTElement) PILL() *SLINPUTElement {
	e.PILLSet(true)
	return e
}

func (e *SLINPUTElement) IfPILL(condition bool) *SLINPUTElement {
	if condition {
		e.PILLSet(true)
	}
	return e
}

// Set the attribute PILL to the value b explicitly.
func (e *SLINPUTElement) PILLSet(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("pill", b)
	return e
}

func (e *SLINPUTElement) IfSetPILL(condition bool, b bool) *SLINPUTElement {
	if condition {
		e.PILLSet(b)
	}
	return e
}

// Remove the attribute PILL from the element.
func (e *SLINPUTElement) PILLRemove(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("pill")
	return e
}

func (e *SLINPUTElement) LABEL(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("label", s)
	return e
}

func (e *SLINPUTElement) IfLABEL(condition bool, s string) *SLINPUTElement {
	if condition {
		e.LABEL(s)
	}
	return e
}

// Remove the attribute LABEL from the element.
func (e *SLINPUTElement) LABELRemove(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("label")
	return e
}

func (e *SLINPUTElement) HELP_TEXT(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("help-text", s)
	return e
}

func (e *SLINPUTElement) IfHELP_TEXT(condition bool, s string) *SLINPUTElement {
	if condition {
		e.HELP_TEXT(s)
	}
	return e
}

// Remove the attribute HELP_TEXT from the element.
func (e *SLINPUTElement) HELP_TEXTRemove(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("help-text")
	return e
}

func (e *SLINPUTElement) READONLY() *SLINPUTElement {
	e.READONLYSet(true)
	return e
}

func (e *SLINPUTElement) IfREADONLY(condition bool) *SLINPUTElement {
	if condition {
		e.READONLYSet(true)
	}
	return e
}

// Set the attribute READONLY to the value b explicitly.
func (e *SLINPUTElement) READONLYSet(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("readonly", b)
	return e
}

func (e *SLINPUTElement) IfSetREADONLY(condition bool, b bool) *SLINPUTElement {
	if condition {
		e.READONLYSet(b)
	}
	return e
}

// Remove the attribute READONLY from the element.
func (e *SLINPUTElement) READONLYRemove(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("readonly")
	return e
}

func (e *SLINPUTElement) CLEARABLE() *SLINPUTElement {
	e.CLEARABLESet(true)
	return e
}

func (e *SLINPUTElement) IfCLEARABLE(condition bool) *SLINPUTElement {
	if condition {
		e.CLEARABLESet(true)
	}
	return e
}

// Set the attribute CLEARABLE to the value b explicitly.
func (e *SLINPUTElement) CLEARABLESet(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("clearable", b)
	return e
}

func (e *SLINPUTElement) IfSetCLEARABLE(condition bool, b bool) *SLINPUTElement {
	if condition {
		e.CLEARABLESet(b)
	}
	return e
}

// Remove the attribute CLEARABLE from the element.
func (e *SLINPUTElement) CLEARABLERemove(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("clearable")
	return e
}

func (e *SLINPUTElement) DISABLED() *SLINPUTElement {
	e.DISABLEDSet(true)
	return e
}

func (e *SLINPUTElement) IfDISABLED(condition bool) *SLINPUTElement {
	if condition {
		e.DISABLEDSet(true)
	}
	return e
}

// Set the attribute DISABLED to the value b explicitly.
func (e *SLINPUTElement) DISABLEDSet(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("disabled", b)
	return e
}

func (e *SLINPUTElement) IfSetDISABLED(condition bool, b bool) *SLINPUTElement {
	if condition {
		e.DISABLEDSet(b)
	}
	return e
}

// Remove the attribute DISABLED from the element.
func (e *SLINPUTElement) DISABLEDRemove(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("disabled")
	return e
}

func (e *SLINPUTElement) PLACEHOLDER(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("placeholder", s)
	return e
}

func (e *SLINPUTElement) IfPLACEHOLDER(condition bool, s string) *SLINPUTElement {
	if condition {
		e.PLACEHOLDER(s)
	}
	return e
}

// Remove the attribute PLACEHOLDER from the element.
func (e *SLINPUTElement) PLACEHOLDERRemove(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("placeholder")
	return e
}

func (e *SLINPUTElement) PASSWORD_TOGGLE() *SLINPUTElement {
	e.PASSWORD_TOGGLESet(true)
	return e
}

func (e *SLINPUTElement) IfPASSWORD_TOGGLE(condition bool) *SLINPUTElement {
	if condition {
		e.PASSWORD_TOGGLESet(true)
	}
	return e
}

// Set the attribute PASSWORD_TOGGLE to the value b explicitly.
func (e *SLINPUTElement) PASSWORD_TOGGLESet(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("password-toggle", b)
	return e
}

func (e *SLINPUTElement) IfSetPASSWORD_TOGGLE(condition bool, b bool) *SLINPUTElement {
	if condition {
		e.PASSWORD_TOGGLESet(b)
	}
	return e
}

// Remove the attribute PASSWORD_TOGGLE from the element.
func (e *SLINPUTElement) PASSWORD_TOGGLERemove(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("password-toggle")
	return e
}

func (e *SLINPUTElement) PASSWORD_VISIBLE() *SLINPUTElement {
	e.PASSWORD_VISIBLESet(true)
	return e
}

func (e *SLINPUTElement) IfPASSWORD_VISIBLE(condition bool) *SLINPUTElement {
	if condition {
		e.PASSWORD_VISIBLESet(true)
	}
	return e
}

// Set the attribute PASSWORD_VISIBLE to the value b explicitly.
func (e *SLINPUTElement) PASSWORD_VISIBLESet(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("password-visible", b)
	return e
}

func (e *SLINPUTElement) IfSetPASSWORD_VISIBLE(condition bool, b bool) *SLINPUTElement {
	if condition {
		e.PASSWORD_VISIBLESet(b)
	}
	return e
}

// Remove the attribute PASSWORD_VISIBLE from the element.
func (e *SLINPUTElement) PASSWORD_VISIBLERemove(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("password-visible")
	return e
}

func (e *SLINPUTElement) NO_SPIN_BUTTONS() *SLINPUTElement {
	e.NO_SPIN_BUTTONSSet(true)
	return e
}

func (e *SLINPUTElement) IfNO_SPIN_BUTTONS(condition bool) *SLINPUTElement {
	if condition {
		e.NO_SPIN_BUTTONSSet(true)
	}
	return e
}

// Set the attribute NO_SPIN_BUTTONS to the value b explicitly.
func (e *SLINPUTElement) NO_SPIN_BUTTONSSet(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("no-spin-buttons", b)
	return e
}

func (e *SLINPUTElement) IfSetNO_SPIN_BUTTONS(condition bool, b bool) *SLINPUTElement {
	if condition {
		e.NO_SPIN_BUTTONSSet(b)
	}
	return e
}

// Remove the attribute NO_SPIN_BUTTONS from the element.
func (e *SLINPUTElement) NO_SPIN_BUTTONSRemove(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("no-spin-buttons")
	return e
}

func (e *SLINPUTElement) REQUIRED() *SLINPUTElement {
	e.REQUIREDSet(true)
	return e
}

func (e *SLINPUTElement) IfREQUIRED(condition bool) *SLINPUTElement {
	if condition {
		e.REQUIREDSet(true)
	}
	return e
}

// Set the attribute REQUIRED to the value b explicitly.
func (e *SLINPUTElement) REQUIREDSet(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("required", b)
	return e
}

func (e *SLINPUTElement) IfSetREQUIRED(condition bool, b bool) *SLINPUTElement {
	if condition {
		e.REQUIREDSet(b)
	}
	return e
}

// Remove the attribute REQUIRED from the element.
func (e *SLINPUTElement) REQUIREDRemove(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("required")
	return e
}

func (e *SLINPUTElement) MINLENGTH(i int) *SLINPUTElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("minlength", i)
	return e
}

func (e *SLINPUTElement) IfMINLENGTH(condition bool, i int) *SLINPUTElement {
	if condition {
		e.MINLENGTH(i)
	}
	return e
}

// Remove the attribute MINLENGTH from the element.
func (e *SLINPUTElement) MINLENGTHRemove(i int) *SLINPUTElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("minlength")
	return e
}

func (e *SLINPUTElement) MAXLENGTH(i int) *SLINPUTElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("maxlength", i)
	return e
}

func (e *SLINPUTElement) IfMAXLENGTH(condition bool, i int) *SLINPUTElement {
	if condition {
		e.MAXLENGTH(i)
	}
	return e
}

// Remove the attribute MAXLENGTH from the element.
func (e *SLINPUTElement) MAXLENGTHRemove(i int) *SLINPUTElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("maxlength")
	return e
}

func (e *SLINPUTElement) MIN(f float64) *SLINPUTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("min", f)
	return e
}

func (e *SLINPUTElement) IfMIN(condition bool, f float64) *SLINPUTElement {
	if condition {
		e.MIN(f)
	}
	return e
}

func (e *SLINPUTElement) MAX(f float64) *SLINPUTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("max", f)
	return e
}

func (e *SLINPUTElement) IfMAX(condition bool, f float64) *SLINPUTElement {
	if condition {
		e.MAX(f)
	}
	return e
}

func (e *SLINPUTElement) STEP(f float64) *SLINPUTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("step", f)
	return e
}

func (e *SLINPUTElement) IfSTEP(condition bool, f float64) *SLINPUTElement {
	if condition {
		e.STEP(f)
	}
	return e
}

func (e *SLINPUTElement) AUTOCAPITALIZE(c SLInputAutocapitalizeChoice) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("autocapitalize", string(c))
	return e
}

type SLInputAutocapitalizeChoice string

const (
	// Off
	SLInputAutocapitalize_off SLInputAutocapitalizeChoice = "off"
	// None
	SLInputAutocapitalize_none SLInputAutocapitalizeChoice = "none"
	// On
	SLInputAutocapitalize_on SLInputAutocapitalizeChoice = "on"
	// Sentences
	SLInputAutocapitalize_sentences SLInputAutocapitalizeChoice = "sentences"
	// Words
	SLInputAutocapitalize_words SLInputAutocapitalizeChoice = "words"
	// Characters
	SLInputAutocapitalize_characters SLInputAutocapitalizeChoice = "characters"
)

// Remove the attribute AUTOCAPITALIZE from the element.
func (e *SLINPUTElement) AUTOCAPITALIZERemove(c SLInputAutocapitalizeChoice) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("autocapitalize")
	return e
}

func (e *SLINPUTElement) AUTOCORRECT(c SLInputAutocorrectChoice) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("autocorrect", string(c))
	return e
}

type SLInputAutocorrectChoice string

const (
	// Off
	SLInputAutocorrect_off SLInputAutocorrectChoice = "off"
	// None
	SLInputAutocorrect_none SLInputAutocorrectChoice = "none"
	// On
	SLInputAutocorrect_on SLInputAutocorrectChoice = "on"
)

// Remove the attribute AUTOCORRECT from the element.
func (e *SLINPUTElement) AUTOCORRECTRemove(c SLInputAutocorrectChoice) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("autocorrect")
	return e
}

func (e *SLINPUTElement) AUTOCOMPLETE(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("autocomplete", s)
	return e
}

func (e *SLINPUTElement) IfAUTOCOMPLETE(condition bool, s string) *SLINPUTElement {
	if condition {
		e.AUTOCOMPLETE(s)
	}
	return e
}

// Remove the attribute AUTOCOMPLETE from the element.
func (e *SLINPUTElement) AUTOCOMPLETERemove(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("autocomplete")
	return e
}

func (e *SLINPUTElement) ENTER_KEY_HINT(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("enter-key-hint", s)
	return e
}

func (e *SLINPUTElement) IfENTER_KEY_HINT(condition bool, s string) *SLINPUTElement {
	if condition {
		e.ENTER_KEY_HINT(s)
	}
	return e
}

// Remove the attribute ENTER_KEY_HINT from the element.
func (e *SLINPUTElement) ENTER_KEY_HINTRemove(s string) *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("enter-key-hint")
	return e
}

func (e *SLINPUTElement) SPELLCHECK() *SLINPUTElement {
	e.SPELLCHECKSet(true)
	return e
}

func (e *SLINPUTElement) IfSPELLCHECK(condition bool) *SLINPUTElement {
	if condition {
		e.SPELLCHECKSet(true)
	}
	return e
}

// Set the attribute SPELLCHECK to the value b explicitly.
func (e *SLINPUTElement) SPELLCHECKSet(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("spellcheck", b)
	return e
}

func (e *SLINPUTElement) IfSetSPELLCHECK(condition bool, b bool) *SLINPUTElement {
	if condition {
		e.SPELLCHECKSet(b)
	}
	return e
}

// Remove the attribute SPELLCHECK from the element.
func (e *SLINPUTElement) SPELLCHECKRemove(b bool) *SLINPUTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("spellcheck")
	return e
}

// Merges the store with the given object

func (e *SLINPUTElement) DATASTAR_MERGE_STORE(v any) *SLINPUTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("data-merge-store", string(b))
	return e
}

// Sets the reference of the element

func (e *SLINPUTElement) DATASTAR_REF(expression string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLINPUTElement) IfDATASTAR_REF(condition bool, expression string) *SLINPUTElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SLINPUTElement) DATASTAR_REFRemove() *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SLINPUTElement) DATASTAR_BIND(key string, expression string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLINPUTElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SLINPUTElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SLINPUTElement) DATASTAR_BINDRemove() *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SLINPUTElement) DATASTAR_MODEL(expression string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLINPUTElement) IfDATASTAR_MODEL(condition bool, expression string) *SLINPUTElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SLINPUTElement) DATASTAR_MODELRemove() *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SLINPUTElement) DATASTAR_TEXT(expression string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLINPUTElement) IfDATASTAR_TEXT(condition bool, expression string) *SLINPUTElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SLINPUTElement) DATASTAR_TEXTRemove() *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SLInputDataOnMod customDataKeyModifier

// Debounces the event handler
func SLInputDataOnModDebounce(
	d time.Duration,
) SLInputDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SLInputDataOnModThrottle(
	d time.Duration,
) SLInputDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SLINPUTElement) DATASTAR_ON(key string, expression string, modifiers ...SLInputDataOnMod) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SLInputDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLINPUTElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SLInputDataOnMod) *SLINPUTElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SLINPUTElement) DATASTAR_ONRemove() *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SLINPUTElement) DATASTAR_FOCUSSet(b bool) *SLINPUTElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLINPUTElement) DATASTAR_FOCUS() *SLINPUTElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SLINPUTElement) DATASTAR_HEADER(key string, expression string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLINPUTElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SLINPUTElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SLINPUTElement) DATASTAR_HEADERRemove() *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SLINPUTElement) DATASTAR_FETCH_URL(expression string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLINPUTElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SLINPUTElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SLINPUTElement) DATASTAR_FETCH_URLRemove() *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SLINPUTElement) DATASTAR_FETCH_INDICATOR(expression string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLINPUTElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SLINPUTElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SLINPUTElement) DATASTAR_FETCH_INDICATORRemove() *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SLINPUTElement) DATASTAR_SHOWSet(b bool) *SLINPUTElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLINPUTElement) DATASTAR_SHOW() *SLINPUTElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SLINPUTElement) DATASTAR_INTERSECTS(expression string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLINPUTElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SLINPUTElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SLINPUTElement) DATASTAR_INTERSECTSRemove() *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SLINPUTElement) DATASTAR_TELEPORTSet(b bool) *SLINPUTElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLINPUTElement) DATASTAR_TELEPORT() *SLINPUTElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SLINPUTElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SLINPUTElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLINPUTElement) DATASTAR_SCROLL_INTO_VIEW() *SLINPUTElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SLINPUTElement) DATASTAR_VIEW_TRANSITION(expression string) *SLINPUTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLINPUTElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SLINPUTElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SLINPUTElement) DATASTAR_VIEW_TRANSITIONRemove() *SLINPUTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
