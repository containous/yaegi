// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports go/build'. DO NOT EDIT.

import (
	"go/build"
	"reflect"
)

func init() {
	Value["go/build"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AllowBinary":   reflect.ValueOf(build.AllowBinary),
		"ArchChar":      reflect.ValueOf(build.ArchChar),
		"Default":       reflect.ValueOf(&build.Default).Elem(),
		"FindOnly":      reflect.ValueOf(build.FindOnly),
		"IgnoreVendor":  reflect.ValueOf(build.IgnoreVendor),
		"Import":        reflect.ValueOf(build.Import),
		"ImportComment": reflect.ValueOf(build.ImportComment),
		"ImportDir":     reflect.ValueOf(build.ImportDir),
		"IsLocalImport": reflect.ValueOf(build.IsLocalImport),
		"ToolDir":       reflect.ValueOf(&build.ToolDir).Elem(),

		// type definitions
		"Context":              reflect.ValueOf((*build.Context)(nil)),
		"ImportMode":           reflect.ValueOf((*build.ImportMode)(nil)),
		"MultiplePackageError": reflect.ValueOf((*build.MultiplePackageError)(nil)),
		"NoGoError":            reflect.ValueOf((*build.NoGoError)(nil)),
		"Package":              reflect.ValueOf((*build.Package)(nil)),
	}
}
