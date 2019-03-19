// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports expvar'. DO NOT EDIT.

import (
	"expvar"
	"reflect"
)

func init() {
	Value["expvar"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Do":        reflect.ValueOf(expvar.Do),
		"Get":       reflect.ValueOf(expvar.Get),
		"Handler":   reflect.ValueOf(expvar.Handler),
		"NewFloat":  reflect.ValueOf(expvar.NewFloat),
		"NewInt":    reflect.ValueOf(expvar.NewInt),
		"NewMap":    reflect.ValueOf(expvar.NewMap),
		"NewString": reflect.ValueOf(expvar.NewString),
		"Publish":   reflect.ValueOf(expvar.Publish),

		// type definitions
		"Float":    reflect.ValueOf((*expvar.Float)(nil)),
		"Func":     reflect.ValueOf((*expvar.Func)(nil)),
		"Int":      reflect.ValueOf((*expvar.Int)(nil)),
		"KeyValue": reflect.ValueOf((*expvar.KeyValue)(nil)),
		"Map":      reflect.ValueOf((*expvar.Map)(nil)),
		"String":   reflect.ValueOf((*expvar.String)(nil)),
		"Var":      reflect.ValueOf((*expvar.Var)(nil)),
	}
}
