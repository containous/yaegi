// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports fmt'. DO NOT EDIT.

import (
	"fmt"
	"reflect"
)

func init() {
	Value["fmt"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Errorf":   reflect.ValueOf(fmt.Errorf),
		"Fprint":   reflect.ValueOf(fmt.Fprint),
		"Fprintf":  reflect.ValueOf(fmt.Fprintf),
		"Fprintln": reflect.ValueOf(fmt.Fprintln),
		"Fscan":    reflect.ValueOf(fmt.Fscan),
		"Fscanf":   reflect.ValueOf(fmt.Fscanf),
		"Fscanln":  reflect.ValueOf(fmt.Fscanln),
		"Print":    reflect.ValueOf(fmt.Print),
		"Printf":   reflect.ValueOf(fmt.Printf),
		"Println":  reflect.ValueOf(fmt.Println),
		"Scan":     reflect.ValueOf(fmt.Scan),
		"Scanf":    reflect.ValueOf(fmt.Scanf),
		"Scanln":   reflect.ValueOf(fmt.Scanln),
		"Sprint":   reflect.ValueOf(fmt.Sprint),
		"Sprintf":  reflect.ValueOf(fmt.Sprintf),
		"Sprintln": reflect.ValueOf(fmt.Sprintln),
		"Sscan":    reflect.ValueOf(fmt.Sscan),
		"Sscanf":   reflect.ValueOf(fmt.Sscanf),
		"Sscanln":  reflect.ValueOf(fmt.Sscanln),

		// type definitions
		"Formatter":  reflect.ValueOf((*fmt.Formatter)(nil)),
		"GoStringer": reflect.ValueOf((*fmt.GoStringer)(nil)),
		"ScanState":  reflect.ValueOf((*fmt.ScanState)(nil)),
		"Scanner":    reflect.ValueOf((*fmt.Scanner)(nil)),
		"State":      reflect.ValueOf((*fmt.State)(nil)),
		"Stringer":   reflect.ValueOf((*fmt.Stringer)(nil)),
	}
}
