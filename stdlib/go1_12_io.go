// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports io'. DO NOT EDIT.

import (
	"io"
	"reflect"
)

func init() {
	Value["io"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Copy":             reflect.ValueOf(io.Copy),
		"CopyBuffer":       reflect.ValueOf(io.CopyBuffer),
		"CopyN":            reflect.ValueOf(io.CopyN),
		"EOF":              reflect.ValueOf(&io.EOF).Elem(),
		"ErrClosedPipe":    reflect.ValueOf(&io.ErrClosedPipe).Elem(),
		"ErrNoProgress":    reflect.ValueOf(&io.ErrNoProgress).Elem(),
		"ErrShortBuffer":   reflect.ValueOf(&io.ErrShortBuffer).Elem(),
		"ErrShortWrite":    reflect.ValueOf(&io.ErrShortWrite).Elem(),
		"ErrUnexpectedEOF": reflect.ValueOf(&io.ErrUnexpectedEOF).Elem(),
		"LimitReader":      reflect.ValueOf(io.LimitReader),
		"MultiReader":      reflect.ValueOf(io.MultiReader),
		"MultiWriter":      reflect.ValueOf(io.MultiWriter),
		"NewSectionReader": reflect.ValueOf(io.NewSectionReader),
		"Pipe":             reflect.ValueOf(io.Pipe),
		"ReadAtLeast":      reflect.ValueOf(io.ReadAtLeast),
		"ReadFull":         reflect.ValueOf(io.ReadFull),
		"SeekCurrent":      reflect.ValueOf(io.SeekCurrent),
		"SeekEnd":          reflect.ValueOf(io.SeekEnd),
		"SeekStart":        reflect.ValueOf(io.SeekStart),
		"TeeReader":        reflect.ValueOf(io.TeeReader),
		"WriteString":      reflect.ValueOf(io.WriteString),

		// type definitions
		"ByteReader":      reflect.ValueOf((*io.ByteReader)(nil)),
		"ByteScanner":     reflect.ValueOf((*io.ByteScanner)(nil)),
		"ByteWriter":      reflect.ValueOf((*io.ByteWriter)(nil)),
		"Closer":          reflect.ValueOf((*io.Closer)(nil)),
		"LimitedReader":   reflect.ValueOf((*io.LimitedReader)(nil)),
		"PipeReader":      reflect.ValueOf((*io.PipeReader)(nil)),
		"PipeWriter":      reflect.ValueOf((*io.PipeWriter)(nil)),
		"ReadCloser":      reflect.ValueOf((*io.ReadCloser)(nil)),
		"ReadSeeker":      reflect.ValueOf((*io.ReadSeeker)(nil)),
		"ReadWriteCloser": reflect.ValueOf((*io.ReadWriteCloser)(nil)),
		"ReadWriteSeeker": reflect.ValueOf((*io.ReadWriteSeeker)(nil)),
		"ReadWriter":      reflect.ValueOf((*io.ReadWriter)(nil)),
		"Reader":          reflect.ValueOf((*io.Reader)(nil)),
		"ReaderAt":        reflect.ValueOf((*io.ReaderAt)(nil)),
		"ReaderFrom":      reflect.ValueOf((*io.ReaderFrom)(nil)),
		"RuneReader":      reflect.ValueOf((*io.RuneReader)(nil)),
		"RuneScanner":     reflect.ValueOf((*io.RuneScanner)(nil)),
		"SectionReader":   reflect.ValueOf((*io.SectionReader)(nil)),
		"Seeker":          reflect.ValueOf((*io.Seeker)(nil)),
		"StringWriter":    reflect.ValueOf((*io.StringWriter)(nil)),
		"WriteCloser":     reflect.ValueOf((*io.WriteCloser)(nil)),
		"WriteSeeker":     reflect.ValueOf((*io.WriteSeeker)(nil)),
		"Writer":          reflect.ValueOf((*io.Writer)(nil)),
		"WriterAt":        reflect.ValueOf((*io.WriterAt)(nil)),
		"WriterTo":        reflect.ValueOf((*io.WriterTo)(nil)),
	}
	Wrapper["io"] = map[string]reflect.Type{
		"ByteReader":      reflect.TypeOf((*_io_ByteReader)(nil)),
		"ByteScanner":     reflect.TypeOf((*_io_ByteScanner)(nil)),
		"ByteWriter":      reflect.TypeOf((*_io_ByteWriter)(nil)),
		"Closer":          reflect.TypeOf((*_io_Closer)(nil)),
		"ReadCloser":      reflect.TypeOf((*_io_ReadCloser)(nil)),
		"ReadSeeker":      reflect.TypeOf((*_io_ReadSeeker)(nil)),
		"ReadWriteCloser": reflect.TypeOf((*_io_ReadWriteCloser)(nil)),
		"ReadWriteSeeker": reflect.TypeOf((*_io_ReadWriteSeeker)(nil)),
		"ReadWriter":      reflect.TypeOf((*_io_ReadWriter)(nil)),
		"Reader":          reflect.TypeOf((*_io_Reader)(nil)),
		"ReaderAt":        reflect.TypeOf((*_io_ReaderAt)(nil)),
		"ReaderFrom":      reflect.TypeOf((*_io_ReaderFrom)(nil)),
		"RuneReader":      reflect.TypeOf((*_io_RuneReader)(nil)),
		"RuneScanner":     reflect.TypeOf((*_io_RuneScanner)(nil)),
		"Seeker":          reflect.TypeOf((*_io_Seeker)(nil)),
		"StringWriter":    reflect.TypeOf((*_io_StringWriter)(nil)),
		"WriteCloser":     reflect.TypeOf((*_io_WriteCloser)(nil)),
		"WriteSeeker":     reflect.TypeOf((*_io_WriteSeeker)(nil)),
		"Writer":          reflect.TypeOf((*_io_Writer)(nil)),
		"WriterAt":        reflect.TypeOf((*_io_WriterAt)(nil)),
		"WriterTo":        reflect.TypeOf((*_io_WriterTo)(nil)),
	}
}

// ByteReader is an interface wrapper for ByteReader type
type _io_ByteReader struct {
	WReadByte func() (byte, error)
}

func (W _io_ByteReader) ReadByte() (byte, error) { return W.WReadByte() }

// ByteScanner is an interface wrapper for ByteScanner type
type _io_ByteScanner struct {
	WReadByte   func() (byte, error)
	WUnreadByte func() error
}

func (W _io_ByteScanner) ReadByte() (byte, error) { return W.WReadByte() }
func (W _io_ByteScanner) UnreadByte() error       { return W.WUnreadByte() }

// ByteWriter is an interface wrapper for ByteWriter type
type _io_ByteWriter struct {
	WWriteByte func(c byte) error
}

func (W _io_ByteWriter) WriteByte(c byte) error { return W.WWriteByte(c) }

// Closer is an interface wrapper for Closer type
type _io_Closer struct {
	WClose func() error
}

func (W _io_Closer) Close() error { return W.WClose() }

// ReadCloser is an interface wrapper for ReadCloser type
type _io_ReadCloser struct {
	WClose func() error
	WRead  func(p []byte) (n int, err error)
}

func (W _io_ReadCloser) Close() error                     { return W.WClose() }
func (W _io_ReadCloser) Read(p []byte) (n int, err error) { return W.WRead(p) }

// ReadSeeker is an interface wrapper for ReadSeeker type
type _io_ReadSeeker struct {
	WRead func(p []byte) (n int, err error)
	WSeek func(offset int64, whence int) (int64, error)
}

func (W _io_ReadSeeker) Read(p []byte) (n int, err error) { return W.WRead(p) }
func (W _io_ReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return W.WSeek(offset, whence)
}

// ReadWriteCloser is an interface wrapper for ReadWriteCloser type
type _io_ReadWriteCloser struct {
	WClose func() error
	WRead  func(p []byte) (n int, err error)
	WWrite func(p []byte) (n int, err error)
}

func (W _io_ReadWriteCloser) Close() error                      { return W.WClose() }
func (W _io_ReadWriteCloser) Read(p []byte) (n int, err error)  { return W.WRead(p) }
func (W _io_ReadWriteCloser) Write(p []byte) (n int, err error) { return W.WWrite(p) }

// ReadWriteSeeker is an interface wrapper for ReadWriteSeeker type
type _io_ReadWriteSeeker struct {
	WRead  func(p []byte) (n int, err error)
	WSeek  func(offset int64, whence int) (int64, error)
	WWrite func(p []byte) (n int, err error)
}

func (W _io_ReadWriteSeeker) Read(p []byte) (n int, err error) { return W.WRead(p) }
func (W _io_ReadWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	return W.WSeek(offset, whence)
}
func (W _io_ReadWriteSeeker) Write(p []byte) (n int, err error) { return W.WWrite(p) }

// ReadWriter is an interface wrapper for ReadWriter type
type _io_ReadWriter struct {
	WRead  func(p []byte) (n int, err error)
	WWrite func(p []byte) (n int, err error)
}

func (W _io_ReadWriter) Read(p []byte) (n int, err error)  { return W.WRead(p) }
func (W _io_ReadWriter) Write(p []byte) (n int, err error) { return W.WWrite(p) }

// Reader is an interface wrapper for Reader type
type _io_Reader struct {
	WRead func(p []byte) (n int, err error)
}

func (W _io_Reader) Read(p []byte) (n int, err error) { return W.WRead(p) }

// ReaderAt is an interface wrapper for ReaderAt type
type _io_ReaderAt struct {
	WReadAt func(p []byte, off int64) (n int, err error)
}

func (W _io_ReaderAt) ReadAt(p []byte, off int64) (n int, err error) { return W.WReadAt(p, off) }

// ReaderFrom is an interface wrapper for ReaderFrom type
type _io_ReaderFrom struct {
	WReadFrom func(r io.Reader) (n int64, err error)
}

func (W _io_ReaderFrom) ReadFrom(r io.Reader) (n int64, err error) { return W.WReadFrom(r) }

// RuneReader is an interface wrapper for RuneReader type
type _io_RuneReader struct {
	WReadRune func() (r rune, size int, err error)
}

func (W _io_RuneReader) ReadRune() (r rune, size int, err error) { return W.WReadRune() }

// RuneScanner is an interface wrapper for RuneScanner type
type _io_RuneScanner struct {
	WReadRune   func() (r rune, size int, err error)
	WUnreadRune func() error
}

func (W _io_RuneScanner) ReadRune() (r rune, size int, err error) { return W.WReadRune() }
func (W _io_RuneScanner) UnreadRune() error                       { return W.WUnreadRune() }

// Seeker is an interface wrapper for Seeker type
type _io_Seeker struct {
	WSeek func(offset int64, whence int) (int64, error)
}

func (W _io_Seeker) Seek(offset int64, whence int) (int64, error) { return W.WSeek(offset, whence) }

// StringWriter is an interface wrapper for StringWriter type
type _io_StringWriter struct {
	WWriteString func(s string) (n int, err error)
}

func (W _io_StringWriter) WriteString(s string) (n int, err error) { return W.WWriteString(s) }

// WriteCloser is an interface wrapper for WriteCloser type
type _io_WriteCloser struct {
	WClose func() error
	WWrite func(p []byte) (n int, err error)
}

func (W _io_WriteCloser) Close() error                      { return W.WClose() }
func (W _io_WriteCloser) Write(p []byte) (n int, err error) { return W.WWrite(p) }

// WriteSeeker is an interface wrapper for WriteSeeker type
type _io_WriteSeeker struct {
	WSeek  func(offset int64, whence int) (int64, error)
	WWrite func(p []byte) (n int, err error)
}

func (W _io_WriteSeeker) Seek(offset int64, whence int) (int64, error) {
	return W.WSeek(offset, whence)
}
func (W _io_WriteSeeker) Write(p []byte) (n int, err error) { return W.WWrite(p) }

// Writer is an interface wrapper for Writer type
type _io_Writer struct {
	WWrite func(p []byte) (n int, err error)
}

func (W _io_Writer) Write(p []byte) (n int, err error) { return W.WWrite(p) }

// WriterAt is an interface wrapper for WriterAt type
type _io_WriterAt struct {
	WWriteAt func(p []byte, off int64) (n int, err error)
}

func (W _io_WriterAt) WriteAt(p []byte, off int64) (n int, err error) { return W.WWriteAt(p, off) }

// WriterTo is an interface wrapper for WriterTo type
type _io_WriterTo struct {
	WWriteTo func(w io.Writer) (n int64, err error)
}

func (W _io_WriterTo) WriteTo(w io.Writer) (n int64, err error) { return W.WWriteTo(w) }
