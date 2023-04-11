// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: http/v1/http.proto

package v1

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

// Validate checks the field values on PageRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageRequestMultiError, or
// nil if none found.
func (m *PageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return PageRequestMultiError(errors)
	}

	return nil
}

// PageRequestMultiError is an error wrapping multiple validation errors
// returned by PageRequest.ValidateAll() if the designated constraints aren't met.
type PageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageRequestMultiError) AllErrors() []error { return m }

// PageRequestValidationError is the validation error returned by
// PageRequest.Validate if the designated constraints aren't met.
type PageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageRequestValidationError) ErrorName() string { return "PageRequestValidationError" }

// Error satisfies the builtin error interface
func (e PageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageRequestValidationError{}

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUserName()); l < 3 || l > 11 {
		err := LoginRequestValidationError{
			field:  "UserName",
			reason: "value length must be between 3 and 11 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 6 {
		err := LoginRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCaptcha()) != 6 {
		err := LoginRequestValidationError{
			field:  "Captcha",
			reason: "value length must be 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if utf8.RuneCountInString(m.GetCaptchaID()) < 6 {
		err := LoginRequestValidationError{
			field:  "CaptchaID",
			reason: "value length must be at least 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginResponseMultiError, or
// nil if none found.
func (m *LoginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ID

	// no validation rules for UID

	// no validation rules for UserName

	// no validation rules for NickName

	// no validation rules for Birth

	// no validation rules for Avatar

	// no validation rules for RoleID

	// no validation rules for RoleName

	// no validation rules for Phone

	// no validation rules for Wechat

	// no validation rules for Email

	// no validation rules for State

	// no validation rules for Motto

	// no validation rules for Token

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return LoginResponseMultiError(errors)
	}

	return nil
}

// LoginResponseMultiError is an error wrapping multiple validation errors
// returned by LoginResponse.ValidateAll() if the designated constraints
// aren't met.
type LoginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginResponseMultiError) AllErrors() []error { return m }

// LoginResponseValidationError is the validation error returned by
// LoginResponse.Validate if the designated constraints aren't met.
type LoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResponseValidationError) ErrorName() string { return "LoginResponseValidationError" }

// Error satisfies the builtin error interface
func (e LoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResponseValidationError{}

// Validate checks the field values on CaptchaResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CaptchaResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CaptchaResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CaptchaResponseMultiError, or nil if none found.
func (m *CaptchaResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CaptchaResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CaptchaID

	// no validation rules for PicPath

	// no validation rules for CaptchaLength

	if len(errors) > 0 {
		return CaptchaResponseMultiError(errors)
	}

	return nil
}

// CaptchaResponseMultiError is an error wrapping multiple validation errors
// returned by CaptchaResponse.ValidateAll() if the designated constraints
// aren't met.
type CaptchaResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CaptchaResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CaptchaResponseMultiError) AllErrors() []error { return m }

// CaptchaResponseValidationError is the validation error returned by
// CaptchaResponse.Validate if the designated constraints aren't met.
type CaptchaResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CaptchaResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CaptchaResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CaptchaResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CaptchaResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CaptchaResponseValidationError) ErrorName() string { return "CaptchaResponseValidationError" }

// Error satisfies the builtin error interface
func (e CaptchaResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCaptchaResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CaptchaResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CaptchaResponseValidationError{}

// Validate checks the field values on MenuResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MenuResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MenuResponseMultiError, or
// nil if none found.
func (m *MenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MenuResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Path

	// no validation rules for Name

	// no validation rules for Component

	// no validation rules for Redirect

	if all {
		switch v := interface{}(m.GetMeta()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MenuResponseValidationError{
					field:  "Meta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MenuResponseValidationError{
					field:  "Meta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MenuResponseValidationError{
				field:  "Meta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetChildren() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuResponseValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuResponseValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuResponseValidationError{
					field:  fmt.Sprintf("Children[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MenuResponseMultiError(errors)
	}

	return nil
}

// MenuResponseMultiError is an error wrapping multiple validation errors
// returned by MenuResponse.ValidateAll() if the designated constraints aren't met.
type MenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MenuResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MenuResponseMultiError) AllErrors() []error { return m }

// MenuResponseValidationError is the validation error returned by
// MenuResponse.Validate if the designated constraints aren't met.
type MenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MenuResponseValidationError) ErrorName() string { return "MenuResponseValidationError" }

// Error satisfies the builtin error interface
func (e MenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MenuResponseValidationError{}

// Validate checks the field values on MenuResponse_Meta with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MenuResponse_Meta) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MenuResponse_Meta with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MenuResponse_MetaMultiError, or nil if none found.
func (m *MenuResponse_Meta) ValidateAll() error {
	return m.validate(true)
}

func (m *MenuResponse_Meta) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ID

	// no validation rules for ParentID

	// no validation rules for RoleIDs

	// no validation rules for Title

	// no validation rules for Icon

	// no validation rules for Hidden

	// no validation rules for KeepAlive

	// no validation rules for Sort

	if len(errors) > 0 {
		return MenuResponse_MetaMultiError(errors)
	}

	return nil
}

// MenuResponse_MetaMultiError is an error wrapping multiple validation errors
// returned by MenuResponse_Meta.ValidateAll() if the designated constraints
// aren't met.
type MenuResponse_MetaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MenuResponse_MetaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MenuResponse_MetaMultiError) AllErrors() []error { return m }

// MenuResponse_MetaValidationError is the validation error returned by
// MenuResponse_Meta.Validate if the designated constraints aren't met.
type MenuResponse_MetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MenuResponse_MetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MenuResponse_MetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MenuResponse_MetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MenuResponse_MetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MenuResponse_MetaValidationError) ErrorName() string {
	return "MenuResponse_MetaValidationError"
}

// Error satisfies the builtin error interface
func (e MenuResponse_MetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMenuResponse_Meta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MenuResponse_MetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MenuResponse_MetaValidationError{}
