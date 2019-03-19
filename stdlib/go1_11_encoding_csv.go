// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports encoding/csv'. DO NOT EDIT.

import (
	"encoding/csv"
	"reflect"
)

func init() {
	Value["encoding/csv"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrBareQuote":     reflect.ValueOf(&csv.ErrBareQuote).Elem(),
		"ErrFieldCount":    reflect.ValueOf(&csv.ErrFieldCount).Elem(),
		"ErrQuote":         reflect.ValueOf(&csv.ErrQuote).Elem(),
		"ErrTrailingComma": reflect.ValueOf(&csv.ErrTrailingComma).Elem(),
		"NewReader":        reflect.ValueOf(csv.NewReader),
		"NewWriter":        reflect.ValueOf(csv.NewWriter),

		// type definitions
		"ParseError": reflect.ValueOf((*csv.ParseError)(nil)),
		"Reader":     reflect.ValueOf((*csv.Reader)(nil)),
		"Writer":     reflect.ValueOf((*csv.Writer)(nil)),
	}
}
