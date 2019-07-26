// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports net/http/cookiejar'. DO NOT EDIT.

import (
	"net/http/cookiejar"
	"reflect"
)

func init() {
	Symbols["net/http/cookiejar"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New": reflect.ValueOf(cookiejar.New),

		// type definitions
		"Jar":              reflect.ValueOf((*cookiejar.Jar)(nil)),
		"Options":          reflect.ValueOf((*cookiejar.Options)(nil)),
		"PublicSuffixList": reflect.ValueOf((*cookiejar.PublicSuffixList)(nil)),

		// interface wrapper definitions
		"_PublicSuffixList": reflect.ValueOf((*_net_http_cookiejar_PublicSuffixList)(nil)),
	}
}

// _net_http_cookiejar_PublicSuffixList is an interface wrapper for PublicSuffixList type
type _net_http_cookiejar_PublicSuffixList struct {
	WPublicSuffix func(domain string) string
	WString       func() string
}

func (W _net_http_cookiejar_PublicSuffixList) PublicSuffix(domain string) string {
	return W.WPublicSuffix(domain)
}
func (W _net_http_cookiejar_PublicSuffixList) String() string { return W.WString() }
