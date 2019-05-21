// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports go/constant'. DO NOT EDIT.

import (
	"go/constant"
	"reflect"
)

func init() {
	Value["go/constant"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BinaryOp":        reflect.ValueOf(constant.BinaryOp),
		"BitLen":          reflect.ValueOf(constant.BitLen),
		"Bool":            reflect.ValueOf(constant.Bool),
		"BoolVal":         reflect.ValueOf(constant.BoolVal),
		"Bytes":           reflect.ValueOf(constant.Bytes),
		"Compare":         reflect.ValueOf(constant.Compare),
		"Complex":         reflect.ValueOf(constant.Complex),
		"Denom":           reflect.ValueOf(constant.Denom),
		"Float":           reflect.ValueOf(constant.Float),
		"Float32Val":      reflect.ValueOf(constant.Float32Val),
		"Float64Val":      reflect.ValueOf(constant.Float64Val),
		"Imag":            reflect.ValueOf(constant.Imag),
		"Int":             reflect.ValueOf(constant.Int),
		"Int64Val":        reflect.ValueOf(constant.Int64Val),
		"MakeBool":        reflect.ValueOf(constant.MakeBool),
		"MakeFloat64":     reflect.ValueOf(constant.MakeFloat64),
		"MakeFromBytes":   reflect.ValueOf(constant.MakeFromBytes),
		"MakeFromLiteral": reflect.ValueOf(constant.MakeFromLiteral),
		"MakeImag":        reflect.ValueOf(constant.MakeImag),
		"MakeInt64":       reflect.ValueOf(constant.MakeInt64),
		"MakeString":      reflect.ValueOf(constant.MakeString),
		"MakeUint64":      reflect.ValueOf(constant.MakeUint64),
		"MakeUnknown":     reflect.ValueOf(constant.MakeUnknown),
		"Num":             reflect.ValueOf(constant.Num),
		"Real":            reflect.ValueOf(constant.Real),
		"Shift":           reflect.ValueOf(constant.Shift),
		"Sign":            reflect.ValueOf(constant.Sign),
		"String":          reflect.ValueOf(constant.String),
		"StringVal":       reflect.ValueOf(constant.StringVal),
		"ToComplex":       reflect.ValueOf(constant.ToComplex),
		"ToFloat":         reflect.ValueOf(constant.ToFloat),
		"ToInt":           reflect.ValueOf(constant.ToInt),
		"Uint64Val":       reflect.ValueOf(constant.Uint64Val),
		"UnaryOp":         reflect.ValueOf(constant.UnaryOp),
		"Unknown":         reflect.ValueOf(constant.Unknown),

		// type definitions
		"Kind":  reflect.ValueOf((*constant.Kind)(nil)),
		"Value": reflect.ValueOf((*constant.Value)(nil)),

		// interface wrapper definitions
		"_Value": reflect.ValueOf((*_go_constant_Value)(nil)),
	}
}

// _go_constant_Value is an interface wrapper for Value type
type _go_constant_Value struct {
	WExactString func() string
	WKind        func() constant.Kind
	WString      func() string
}

func (W _go_constant_Value) ExactString() string { return W.WExactString() }
func (W _go_constant_Value) Kind() constant.Kind { return W.WKind() }
func (W _go_constant_Value) String() string      { return W.WString() }
