package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	FieldErrors    map[string]string
	NonFieldErrors []string
}

// check whether errors exist
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {
	//initialise map if map doesn't exist
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	if _, ok := v.FieldErrors[key]; !ok {
		v.FieldErrors[key] = message
	}
}

// check if field isn't valid. add error to the map if not valid(!ok)
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// check for blank field inputs
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// check for maximum allowed characters (n)
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// return true if (value) is in (permittedValues), otherwise false
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

// Checking for minimum characters is satsified
func MinChars(value string, minChars int) bool {
	return utf8.RuneCountInString(value) >= minChars
}

// Checks whether the email is valid or not.
func Matches(email string, rx *regexp.Regexp) bool {
	return rx.MatchString(email)
}

func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}
