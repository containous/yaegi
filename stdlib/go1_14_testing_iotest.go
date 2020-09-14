// Code generated by 'github.com/containous/yaegi/extract testing/iotest'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"reflect"
	"testing/iotest"
)

func init() {
	Symbols["testing/iotest"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"DataErrReader":  reflect.ValueOf(iotest.DataErrReader),
		"ErrTimeout":     reflect.ValueOf(&iotest.ErrTimeout).Elem(),
		"HalfReader":     reflect.ValueOf(iotest.HalfReader),
		"NewReadLogger":  reflect.ValueOf(iotest.NewReadLogger),
		"NewWriteLogger": reflect.ValueOf(iotest.NewWriteLogger),
		"OneByteReader":  reflect.ValueOf(iotest.OneByteReader),
		"TimeoutReader":  reflect.ValueOf(iotest.TimeoutReader),
		"TruncateWriter": reflect.ValueOf(iotest.TruncateWriter),
	}
}
