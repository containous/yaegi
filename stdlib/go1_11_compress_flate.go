// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports compress/flate'. DO NOT EDIT.

import (
	"compress/flate"
	"io"
	"reflect"
)

func init() {
	Value["compress/flate"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BestCompression":    reflect.ValueOf(flate.BestCompression),
		"BestSpeed":          reflect.ValueOf(flate.BestSpeed),
		"DefaultCompression": reflect.ValueOf(flate.DefaultCompression),
		"HuffmanOnly":        reflect.ValueOf(flate.HuffmanOnly),
		"NewReader":          reflect.ValueOf(flate.NewReader),
		"NewReaderDict":      reflect.ValueOf(flate.NewReaderDict),
		"NewWriter":          reflect.ValueOf(flate.NewWriter),
		"NewWriterDict":      reflect.ValueOf(flate.NewWriterDict),
		"NoCompression":      reflect.ValueOf(flate.NoCompression),

		// type definitions
		"CorruptInputError": reflect.ValueOf((*flate.CorruptInputError)(nil)),
		"InternalError":     reflect.ValueOf((*flate.InternalError)(nil)),
		"ReadError":         reflect.ValueOf((*flate.ReadError)(nil)),
		"Reader":            reflect.ValueOf((*flate.Reader)(nil)),
		"Resetter":          reflect.ValueOf((*flate.Resetter)(nil)),
		"WriteError":        reflect.ValueOf((*flate.WriteError)(nil)),
		"Writer":            reflect.ValueOf((*flate.Writer)(nil)),

		// interface wrapper definitions
		"_Reader":   reflect.ValueOf((*_compress_flate_Reader)(nil)),
		"_Resetter": reflect.ValueOf((*_compress_flate_Resetter)(nil)),
	}
}

// _compress_flate_Reader is an interface wrapper for Reader type
type _compress_flate_Reader struct {
	WRead     func(p []byte) (n int, err error)
	WReadByte func() (byte, error)
}

func (W _compress_flate_Reader) Read(p []byte) (n int, err error) { return W.WRead(p) }
func (W _compress_flate_Reader) ReadByte() (byte, error)          { return W.WReadByte() }

// _compress_flate_Resetter is an interface wrapper for Resetter type
type _compress_flate_Resetter struct {
	WReset func(r io.Reader, dict []byte) error
}

func (W _compress_flate_Resetter) Reset(r io.Reader, dict []byte) error { return W.WReset(r, dict) }
