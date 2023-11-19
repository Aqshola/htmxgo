package types

import (
	"net/http"
	"regexp"
)

type Route struct {
	Method  string
	Regex   *regexp.Regexp
	Handler http.HandlerFunc
}

type CtrKey struct{}
