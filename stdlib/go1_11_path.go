// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports path'. DO NOT EDIT.

import (
	"path"
	"reflect"
)

func init() {
	Value["path"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Base":          reflect.ValueOf(path.Base),
		"Clean":         reflect.ValueOf(path.Clean),
		"Dir":           reflect.ValueOf(path.Dir),
		"ErrBadPattern": reflect.ValueOf(&path.ErrBadPattern).Elem(),
		"Ext":           reflect.ValueOf(path.Ext),
		"IsAbs":         reflect.ValueOf(path.IsAbs),
		"Join":          reflect.ValueOf(path.Join),
		"Match":         reflect.ValueOf(path.Match),
		"Split":         reflect.ValueOf(path.Split),

		// type definitions

	}
	Wrapper["path"] = map[string]reflect.Type{}
}
