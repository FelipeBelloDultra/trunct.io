package validator

import (
	"context"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validator interface {
	Valid(context.Context) Evaluator
}

type Evaluator map[string][]string

var EmailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`,
)

func (e *Evaluator) AddFieldError(key, message string) {
	if *e == nil {
		*e = make(map[string][]string)
	}

	if value, exists := (*e)[key]; !exists {
		(*e)[key] = append(value, message)
	}
}

func (e *Evaluator) CheckField(ok bool, key, message string) {
	if !ok {
		e.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, max int) bool {
	return utf8.RuneCountInString(value) <= max
}

func MinChars(value string, min int) bool {
	return utf8.RuneCountInString(value) >= min
}

func Matches(value string, regex *regexp.Regexp) bool {
	return regex.MatchString(value)
}

func IsURL(value string) bool {
	_, err := url.ParseRequestURI(value)

	return err == nil
}
