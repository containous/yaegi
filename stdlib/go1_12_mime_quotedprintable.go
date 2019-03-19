// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports mime/quotedprintable'. DO NOT EDIT.

import (
	"mime/quotedprintable"
	"reflect"
)

func init() {
	Value["mime/quotedprintable"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewReader": reflect.ValueOf(quotedprintable.NewReader),
		"NewWriter": reflect.ValueOf(quotedprintable.NewWriter),

		// type definitions
		"Reader": reflect.ValueOf((*quotedprintable.Reader)(nil)),
		"Writer": reflect.ValueOf((*quotedprintable.Writer)(nil)),
	}
}
