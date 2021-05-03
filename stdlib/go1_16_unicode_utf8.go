// Code generated by 'yaegi extract unicode/utf8'. DO NOT EDIT.

// +build go1.16

package stdlib

import (
	"go/constant"
	"go/token"
	"reflect"
	"unicode/utf8"
)

func init() {
	Symbols["unicode/utf8/utf8"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"DecodeLastRune":         reflect.ValueOf(utf8.DecodeLastRune),
		"DecodeLastRuneInString": reflect.ValueOf(utf8.DecodeLastRuneInString),
		"DecodeRune":             reflect.ValueOf(utf8.DecodeRune),
		"DecodeRuneInString":     reflect.ValueOf(utf8.DecodeRuneInString),
		"EncodeRune":             reflect.ValueOf(utf8.EncodeRune),
		"FullRune":               reflect.ValueOf(utf8.FullRune),
		"FullRuneInString":       reflect.ValueOf(utf8.FullRuneInString),
		"MaxRune":                reflect.ValueOf(constant.MakeFromLiteral("1114111", token.INT, 0)),
		"RuneCount":              reflect.ValueOf(utf8.RuneCount),
		"RuneCountInString":      reflect.ValueOf(utf8.RuneCountInString),
		"RuneError":              reflect.ValueOf(constant.MakeFromLiteral("65533", token.INT, 0)),
		"RuneLen":                reflect.ValueOf(utf8.RuneLen),
		"RuneSelf":               reflect.ValueOf(constant.MakeFromLiteral("128", token.INT, 0)),
		"RuneStart":              reflect.ValueOf(utf8.RuneStart),
		"UTFMax":                 reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"Valid":                  reflect.ValueOf(utf8.Valid),
		"ValidRune":              reflect.ValueOf(utf8.ValidRune),
		"ValidString":            reflect.ValueOf(utf8.ValidString),
	}
}
