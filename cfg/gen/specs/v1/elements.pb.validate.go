// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: specs/v1/elements.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Attribute with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Attribute) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Attribute with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AttributeMultiError, or nil
// if none found.
func (m *Attribute) ValidateAll() error {
	return m.validate(true)
}

func (m *Attribute) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Name

	// no validation rules for Description

	if all {
		switch v := interface{}(m.GetType()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeValidationError{
					field:  "Type",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeValidationError{
					field:  "Type",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeValidationError{
				field:  "Type",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AttributeMultiError(errors)
	}

	return nil
}

// AttributeMultiError is an error wrapping multiple validation errors returned
// by Attribute.ValidateAll() if the designated constraints aren't met.
type AttributeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeMultiError) AllErrors() []error { return m }

// AttributeValidationError is the validation error returned by
// Attribute.Validate if the designated constraints aren't met.
type AttributeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeValidationError) ErrorName() string { return "AttributeValidationError" }

// Error satisfies the builtin error interface
func (e AttributeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttribute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeValidationError{}

// Validate checks the field values on Element with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Element) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Element with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ElementMultiError, or nil if none found.
func (m *Element) ValidateAll() error {
	return m.validate(true)
}

func (m *Element) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Tag

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for NoChildren

	for idx, item := range m.GetAttributes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ElementValidationError{
						field:  fmt.Sprintf("Attributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ElementValidationError{
						field:  fmt.Sprintf("Attributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ElementValidationError{
					field:  fmt.Sprintf("Attributes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ElementMultiError(errors)
	}

	return nil
}

// ElementMultiError is an error wrapping multiple validation errors returned
// by Element.ValidateAll() if the designated constraints aren't met.
type ElementMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ElementMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ElementMultiError) AllErrors() []error { return m }

// ElementValidationError is the validation error returned by Element.Validate
// if the designated constraints aren't met.
type ElementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ElementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ElementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ElementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ElementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ElementValidationError) ErrorName() string { return "ElementValidationError" }

// Error satisfies the builtin error interface
func (e ElementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sElement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ElementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ElementValidationError{}

// Validate checks the field values on Namespace with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Namespace) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Namespace with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NamespaceMultiError, or nil
// if none found.
func (m *Namespace) ValidateAll() error {
	return m.validate(true)
}

func (m *Namespace) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Description

	for idx, item := range m.GetGlobalAttributes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NamespaceValidationError{
						field:  fmt.Sprintf("GlobalAttributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NamespaceValidationError{
						field:  fmt.Sprintf("GlobalAttributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NamespaceValidationError{
					field:  fmt.Sprintf("GlobalAttributes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetElements() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NamespaceValidationError{
						field:  fmt.Sprintf("Elements[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NamespaceValidationError{
						field:  fmt.Sprintf("Elements[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NamespaceValidationError{
					field:  fmt.Sprintf("Elements[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return NamespaceMultiError(errors)
	}

	return nil
}

// NamespaceMultiError is an error wrapping multiple validation errors returned
// by Namespace.ValidateAll() if the designated constraints aren't met.
type NamespaceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NamespaceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NamespaceMultiError) AllErrors() []error { return m }

// NamespaceValidationError is the validation error returned by
// Namespace.Validate if the designated constraints aren't met.
type NamespaceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NamespaceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NamespaceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NamespaceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NamespaceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NamespaceValidationError) ErrorName() string { return "NamespaceValidationError" }

// Error satisfies the builtin error interface
func (e NamespaceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNamespace.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NamespaceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NamespaceValidationError{}

// Validate checks the field values on Attribute_KV with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Attribute_KV) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Attribute_KV with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Attribute_KVMultiError, or
// nil if none found.
func (m *Attribute_KV) ValidateAll() error {
	return m.validate(true)
}

func (m *Attribute_KV) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for KeyValueDelimiter

	// no validation rules for PairDelimiter

	if len(errors) > 0 {
		return Attribute_KVMultiError(errors)
	}

	return nil
}

// Attribute_KVMultiError is an error wrapping multiple validation errors
// returned by Attribute_KV.ValidateAll() if the designated constraints aren't met.
type Attribute_KVMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Attribute_KVMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Attribute_KVMultiError) AllErrors() []error { return m }

// Attribute_KVValidationError is the validation error returned by
// Attribute_KV.Validate if the designated constraints aren't met.
type Attribute_KVValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Attribute_KVValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Attribute_KVValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Attribute_KVValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Attribute_KVValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Attribute_KVValidationError) ErrorName() string { return "Attribute_KVValidationError" }

// Error satisfies the builtin error interface
func (e Attribute_KVValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttribute_KV.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Attribute_KVValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Attribute_KVValidationError{}

// Validate checks the field values on Attribute_Choice with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Attribute_Choice) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Attribute_Choice with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Attribute_ChoiceMultiError, or nil if none found.
func (m *Attribute_Choice) ValidateAll() error {
	return m.validate(true)
}

func (m *Attribute_Choice) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Description

	if len(errors) > 0 {
		return Attribute_ChoiceMultiError(errors)
	}

	return nil
}

// Attribute_ChoiceMultiError is an error wrapping multiple validation errors
// returned by Attribute_Choice.ValidateAll() if the designated constraints
// aren't met.
type Attribute_ChoiceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Attribute_ChoiceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Attribute_ChoiceMultiError) AllErrors() []error { return m }

// Attribute_ChoiceValidationError is the validation error returned by
// Attribute_Choice.Validate if the designated constraints aren't met.
type Attribute_ChoiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Attribute_ChoiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Attribute_ChoiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Attribute_ChoiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Attribute_ChoiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Attribute_ChoiceValidationError) ErrorName() string { return "Attribute_ChoiceValidationError" }

// Error satisfies the builtin error interface
func (e Attribute_ChoiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttribute_Choice.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Attribute_ChoiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Attribute_ChoiceValidationError{}

// Validate checks the field values on Attribute_Choices with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Attribute_Choices) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Attribute_Choices with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Attribute_ChoicesMultiError, or nil if none found.
func (m *Attribute_Choices) ValidateAll() error {
	return m.validate(true)
}

func (m *Attribute_Choices) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetChoices() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, Attribute_ChoicesValidationError{
						field:  fmt.Sprintf("Choices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, Attribute_ChoicesValidationError{
						field:  fmt.Sprintf("Choices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Attribute_ChoicesValidationError{
					field:  fmt.Sprintf("Choices[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return Attribute_ChoicesMultiError(errors)
	}

	return nil
}

// Attribute_ChoicesMultiError is an error wrapping multiple validation errors
// returned by Attribute_Choices.ValidateAll() if the designated constraints
// aren't met.
type Attribute_ChoicesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Attribute_ChoicesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Attribute_ChoicesMultiError) AllErrors() []error { return m }

// Attribute_ChoicesValidationError is the validation error returned by
// Attribute_Choices.Validate if the designated constraints aren't met.
type Attribute_ChoicesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Attribute_ChoicesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Attribute_ChoicesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Attribute_ChoicesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Attribute_ChoicesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Attribute_ChoicesValidationError) ErrorName() string {
	return "Attribute_ChoicesValidationError"
}

// Error satisfies the builtin error interface
func (e Attribute_ChoicesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttribute_Choices.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Attribute_ChoicesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Attribute_ChoicesValidationError{}

// Validate checks the field values on Attribute_Type with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Attribute_Type) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Attribute_Type with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Attribute_TypeMultiError,
// or nil if none found.
func (m *Attribute_Type) ValidateAll() error {
	return m.validate(true)
}

func (m *Attribute_Type) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Type.(type) {
	case *Attribute_Type_String_:
		if v == nil {
			err := Attribute_TypeValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for String_
	case *Attribute_Type_Delimited:
		if v == nil {
			err := Attribute_TypeValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Delimited
	case *Attribute_Type_Kv:
		if v == nil {
			err := Attribute_TypeValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetKv()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, Attribute_TypeValidationError{
						field:  "Kv",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, Attribute_TypeValidationError{
						field:  "Kv",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetKv()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Attribute_TypeValidationError{
					field:  "Kv",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Attribute_Type_Bool:
		if v == nil {
			err := Attribute_TypeValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Bool
	case *Attribute_Type_Integer:
		if v == nil {
			err := Attribute_TypeValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Integer
	case *Attribute_Type_Number:
		if v == nil {
			err := Attribute_TypeValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Number
	case *Attribute_Type_Choices:
		if v == nil {
			err := Attribute_TypeValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetChoices()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, Attribute_TypeValidationError{
						field:  "Choices",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, Attribute_TypeValidationError{
						field:  "Choices",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetChoices()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Attribute_TypeValidationError{
					field:  "Choices",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Attribute_Type_Rune:
		if v == nil {
			err := Attribute_TypeValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Rune
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return Attribute_TypeMultiError(errors)
	}

	return nil
}

// Attribute_TypeMultiError is an error wrapping multiple validation errors
// returned by Attribute_Type.ValidateAll() if the designated constraints
// aren't met.
type Attribute_TypeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Attribute_TypeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Attribute_TypeMultiError) AllErrors() []error { return m }

// Attribute_TypeValidationError is the validation error returned by
// Attribute_Type.Validate if the designated constraints aren't met.
type Attribute_TypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Attribute_TypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Attribute_TypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Attribute_TypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Attribute_TypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Attribute_TypeValidationError) ErrorName() string { return "Attribute_TypeValidationError" }

// Error satisfies the builtin error interface
func (e Attribute_TypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttribute_Type.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Attribute_TypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Attribute_TypeValidationError{}
