// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports context'. DO NOT EDIT.

import (
	"context"
	"reflect"
)

func init() {
	Value["context"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Background":       reflect.ValueOf(context.Background),
		"Canceled":         reflect.ValueOf(&context.Canceled).Elem(),
		"DeadlineExceeded": reflect.ValueOf(&context.DeadlineExceeded).Elem(),
		"TODO":             reflect.ValueOf(context.TODO),
		"WithCancel":       reflect.ValueOf(context.WithCancel),
		"WithDeadline":     reflect.ValueOf(context.WithDeadline),
		"WithTimeout":      reflect.ValueOf(context.WithTimeout),
		"WithValue":        reflect.ValueOf(context.WithValue),

		// type definitions
		"CancelFunc": reflect.ValueOf((*context.CancelFunc)(nil)),
		"Context":    reflect.ValueOf((*context.Context)(nil)),
	}
}
