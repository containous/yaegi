// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports index/suffixarray'. DO NOT EDIT.

import (
	"index/suffixarray"
	"reflect"
)

func init() {
	Value["index/suffixarray"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New": reflect.ValueOf(suffixarray.New),

		// type definitions
		"Index": reflect.ValueOf((*suffixarray.Index)(nil)),
	}
	Wrapper["index/suffixarray"] = map[string]reflect.Type{}
}
