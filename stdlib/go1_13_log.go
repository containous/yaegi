// Code generated by 'goexports log'. DO NOT EDIT.

// +build go1.13,!go1.14

package stdlib

import (
	"go/constant"
	"go/token"
	"log"
	"reflect"
)

func init() {
	Symbols["log"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Fatal":         reflect.ValueOf(logFatal),
		"Fatalf":        reflect.ValueOf(logFatalf),
		"Fatalln":       reflect.ValueOf(logFatalln),
		"Flags":         reflect.ValueOf(log.Flags),
		"LUTC":          reflect.ValueOf(constant.MakeFromLiteral("32", token.INT, 0)),
		"Ldate":         reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"Llongfile":     reflect.ValueOf(constant.MakeFromLiteral("8", token.INT, 0)),
		"Lmicroseconds": reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"Lshortfile":    reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"LstdFlags":     reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"Ltime":         reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"New":           reflect.ValueOf(logNew),
		"Output":        reflect.ValueOf(log.Output),
		"Panic":         reflect.ValueOf(log.Panic),
		"Panicf":        reflect.ValueOf(log.Panicf),
		"Panicln":       reflect.ValueOf(log.Panicln),
		"Prefix":        reflect.ValueOf(log.Prefix),
		"Print":         reflect.ValueOf(log.Print),
		"Printf":        reflect.ValueOf(log.Printf),
		"Println":       reflect.ValueOf(log.Println),
		"SetFlags":      reflect.ValueOf(log.SetFlags),
		"SetOutput":     reflect.ValueOf(log.SetOutput),
		"SetPrefix":     reflect.ValueOf(log.SetPrefix),
		"Writer":        reflect.ValueOf(log.Writer),

		// type definitions
		"Logger": reflect.ValueOf((*logLogger)(nil)),
	}
}
