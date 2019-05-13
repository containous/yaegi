// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports io/ioutil'. DO NOT EDIT.

import (
	"io/ioutil"
	"reflect"
)

func init() {
	Value["io/ioutil"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Discard":   reflect.ValueOf(&ioutil.Discard).Elem(),
		"NopCloser": reflect.ValueOf(ioutil.NopCloser),
		"ReadAll":   reflect.ValueOf(ioutil.ReadAll),
		"ReadDir":   reflect.ValueOf(ioutil.ReadDir),
		"ReadFile":  reflect.ValueOf(ioutil.ReadFile),
		"TempDir":   reflect.ValueOf(ioutil.TempDir),
		"TempFile":  reflect.ValueOf(ioutil.TempFile),
		"WriteFile": reflect.ValueOf(ioutil.WriteFile),

		// type definitions

	}
	Wrapper["io/ioutil"] = map[string]reflect.Type{}
}
