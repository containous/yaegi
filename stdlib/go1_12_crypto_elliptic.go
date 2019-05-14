// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports crypto/elliptic'. DO NOT EDIT.

import (
	"crypto/elliptic"
	"math/big"
	"reflect"
)

func init() {
	Value["crypto/elliptic"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"GenerateKey": reflect.ValueOf(elliptic.GenerateKey),
		"Marshal":     reflect.ValueOf(elliptic.Marshal),
		"P224":        reflect.ValueOf(elliptic.P224),
		"P256":        reflect.ValueOf(elliptic.P256),
		"P384":        reflect.ValueOf(elliptic.P384),
		"P521":        reflect.ValueOf(elliptic.P521),
		"Unmarshal":   reflect.ValueOf(elliptic.Unmarshal),

		// type definitions
		"Curve":       reflect.ValueOf((*elliptic.Curve)(nil)),
		"CurveParams": reflect.ValueOf((*elliptic.CurveParams)(nil)),

		// interface wrapper definitions
		"_Curve": reflect.ValueOf((*_crypto_elliptic_Curve)(nil)),
	}
}

// _crypto_elliptic_Curve is an interface wrapper for Curve type
type _crypto_elliptic_Curve struct {
	WAdd            func(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) (x *big.Int, y *big.Int)
	WDouble         func(x1 *big.Int, y1 *big.Int) (x *big.Int, y *big.Int)
	WIsOnCurve      func(x *big.Int, y *big.Int) bool
	WParams         func() *elliptic.CurveParams
	WScalarBaseMult func(k []byte) (x *big.Int, y *big.Int)
	WScalarMult     func(x1 *big.Int, y1 *big.Int, k []byte) (x *big.Int, y *big.Int)
}

func (W _crypto_elliptic_Curve) Add(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) (x *big.Int, y *big.Int) {
	return W.WAdd(x1, y1, x2, y2)
}
func (W _crypto_elliptic_Curve) Double(x1 *big.Int, y1 *big.Int) (x *big.Int, y *big.Int) {
	return W.WDouble(x1, y1)
}
func (W _crypto_elliptic_Curve) IsOnCurve(x *big.Int, y *big.Int) bool { return W.WIsOnCurve(x, y) }
func (W _crypto_elliptic_Curve) Params() *elliptic.CurveParams         { return W.WParams() }
func (W _crypto_elliptic_Curve) ScalarBaseMult(k []byte) (x *big.Int, y *big.Int) {
	return W.WScalarBaseMult(k)
}
func (W _crypto_elliptic_Curve) ScalarMult(x1 *big.Int, y1 *big.Int, k []byte) (x *big.Int, y *big.Int) {
	return W.WScalarMult(x1, y1, k)
}
