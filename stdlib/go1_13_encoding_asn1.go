// Code generated by 'goexports encoding/asn1'. DO NOT EDIT.

// +build go1.13,!go1.14

package stdlib

import (
	"encoding/asn1"
	"reflect"
)

func init() {
	Symbols["encoding/asn1"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ClassApplication":     reflect.ValueOf(asn1.ClassApplication),
		"ClassContextSpecific": reflect.ValueOf(asn1.ClassContextSpecific),
		"ClassPrivate":         reflect.ValueOf(asn1.ClassPrivate),
		"ClassUniversal":       reflect.ValueOf(asn1.ClassUniversal),
		"Marshal":              reflect.ValueOf(asn1.Marshal),
		"MarshalWithParams":    reflect.ValueOf(asn1.MarshalWithParams),
		"NullBytes":            reflect.ValueOf(&asn1.NullBytes).Elem(),
		"NullRawValue":         reflect.ValueOf(&asn1.NullRawValue).Elem(),
		"TagBitString":         reflect.ValueOf(asn1.TagBitString),
		"TagBoolean":           reflect.ValueOf(asn1.TagBoolean),
		"TagEnum":              reflect.ValueOf(asn1.TagEnum),
		"TagGeneralString":     reflect.ValueOf(asn1.TagGeneralString),
		"TagGeneralizedTime":   reflect.ValueOf(asn1.TagGeneralizedTime),
		"TagIA5String":         reflect.ValueOf(asn1.TagIA5String),
		"TagInteger":           reflect.ValueOf(asn1.TagInteger),
		"TagNull":              reflect.ValueOf(asn1.TagNull),
		"TagNumericString":     reflect.ValueOf(asn1.TagNumericString),
		"TagOID":               reflect.ValueOf(asn1.TagOID),
		"TagOctetString":       reflect.ValueOf(asn1.TagOctetString),
		"TagPrintableString":   reflect.ValueOf(asn1.TagPrintableString),
		"TagSequence":          reflect.ValueOf(asn1.TagSequence),
		"TagSet":               reflect.ValueOf(asn1.TagSet),
		"TagT61String":         reflect.ValueOf(asn1.TagT61String),
		"TagUTCTime":           reflect.ValueOf(asn1.TagUTCTime),
		"TagUTF8String":        reflect.ValueOf(asn1.TagUTF8String),
		"Unmarshal":            reflect.ValueOf(asn1.Unmarshal),
		"UnmarshalWithParams":  reflect.ValueOf(asn1.UnmarshalWithParams),

		// type definitions
		"BitString":        reflect.ValueOf((*asn1.BitString)(nil)),
		"Enumerated":       reflect.ValueOf((*asn1.Enumerated)(nil)),
		"Flag":             reflect.ValueOf((*asn1.Flag)(nil)),
		"ObjectIdentifier": reflect.ValueOf((*asn1.ObjectIdentifier)(nil)),
		"RawContent":       reflect.ValueOf((*asn1.RawContent)(nil)),
		"RawValue":         reflect.ValueOf((*asn1.RawValue)(nil)),
		"StructuralError":  reflect.ValueOf((*asn1.StructuralError)(nil)),
		"SyntaxError":      reflect.ValueOf((*asn1.SyntaxError)(nil)),
	}
}
