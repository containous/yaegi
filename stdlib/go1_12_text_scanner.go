// Code generated by 'goexports text/scanner'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlib

import (
	"reflect"
	"text/scanner"
)

func init() {
	Symbols["text/scanner"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Char":           reflect.ValueOf(scanner.Char),
		"Comment":        reflect.ValueOf(scanner.Comment),
		"EOF":            reflect.ValueOf(scanner.EOF),
		"Float":          reflect.ValueOf(scanner.Float),
		"GoTokens":       reflect.ValueOf(scanner.GoTokens),
		"GoWhitespace":   reflect.ValueOf(int64(scanner.GoWhitespace)),
		"Ident":          reflect.ValueOf(scanner.Ident),
		"Int":            reflect.ValueOf(scanner.Int),
		"RawString":      reflect.ValueOf(scanner.RawString),
		"ScanChars":      reflect.ValueOf(scanner.ScanChars),
		"ScanComments":   reflect.ValueOf(scanner.ScanComments),
		"ScanFloats":     reflect.ValueOf(scanner.ScanFloats),
		"ScanIdents":     reflect.ValueOf(scanner.ScanIdents),
		"ScanInts":       reflect.ValueOf(scanner.ScanInts),
		"ScanRawStrings": reflect.ValueOf(scanner.ScanRawStrings),
		"ScanStrings":    reflect.ValueOf(scanner.ScanStrings),
		"SkipComments":   reflect.ValueOf(scanner.SkipComments),
		"String":         reflect.ValueOf(scanner.String),
		"TokenString":    reflect.ValueOf(scanner.TokenString),

		// type definitions
		"Position": reflect.ValueOf((*scanner.Position)(nil)),
		"Scanner":  reflect.ValueOf((*scanner.Scanner)(nil)),
	}
}
