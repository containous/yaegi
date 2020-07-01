// Code generated by 'goexports compress/zlib'. DO NOT EDIT.

// +build go1.13,!go1.14

package stdlib

import (
	"compress/zlib"
	"fmt"
	"go/constant"
	"go/token"
	"io"
	"reflect"
)

func init() {
	Symbols["compress/zlib"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BestCompression":    reflect.ValueOf(constant.MakeFromLiteral("9", token.INT, 0)),
		"BestSpeed":          reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"DefaultCompression": reflect.ValueOf(constant.MakeFromLiteral("-1", token.INT, 0)),
		"ErrChecksum":        reflect.ValueOf(&zlib.ErrChecksum).Elem(),
		"ErrDictionary":      reflect.ValueOf(&zlib.ErrDictionary).Elem(),
		"ErrHeader":          reflect.ValueOf(&zlib.ErrHeader).Elem(),
		"HuffmanOnly":        reflect.ValueOf(constant.MakeFromLiteral("-2", token.INT, 0)),
		"NewReader":          reflect.ValueOf(zlib.NewReader),
		"NewReaderDict":      reflect.ValueOf(zlib.NewReaderDict),
		"NewWriter":          reflect.ValueOf(zlib.NewWriter),
		"NewWriterLevel":     reflect.ValueOf(zlib.NewWriterLevel),
		"NewWriterLevelDict": reflect.ValueOf(zlib.NewWriterLevelDict),
		"NoCompression":      reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),

		// type definitions
		"Resetter": reflect.ValueOf((*zlib.Resetter)(nil)),
		"Writer":   reflect.ValueOf((*zlib.Writer)(nil)),

		// interface wrapper definitions
		"_Resetter": reflect.ValueOf((*_compress_zlib_Resetter)(nil)),
	}
}

// _compress_zlib_Resetter is an interface wrapper for Resetter type
type _compress_zlib_Resetter struct {
	Val    interface{}
	WReset func(r io.Reader, dict []byte) error
}

func (W _compress_zlib_Resetter) String() string { return fmt.Sprint(W.Val) }

func (W _compress_zlib_Resetter) Reset(r io.Reader, dict []byte) error { return W.WReset(r, dict) }
