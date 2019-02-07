package stdlib

// Code generated by 'goexports strconv'. DO NOT EDIT.

import (
	"reflect"
	"strconv"
)

func init() {
	Value["strconv"] = map[string]reflect.Value{
		"AppendBool":               reflect.ValueOf(strconv.AppendBool),
		"AppendFloat":              reflect.ValueOf(strconv.AppendFloat),
		"AppendInt":                reflect.ValueOf(strconv.AppendInt),
		"AppendQuote":              reflect.ValueOf(strconv.AppendQuote),
		"AppendQuoteRune":          reflect.ValueOf(strconv.AppendQuoteRune),
		"AppendQuoteRuneToASCII":   reflect.ValueOf(strconv.AppendQuoteRuneToASCII),
		"AppendQuoteRuneToGraphic": reflect.ValueOf(strconv.AppendQuoteRuneToGraphic),
		"AppendQuoteToASCII":       reflect.ValueOf(strconv.AppendQuoteToASCII),
		"AppendQuoteToGraphic":     reflect.ValueOf(strconv.AppendQuoteToGraphic),
		"AppendUint":               reflect.ValueOf(strconv.AppendUint),
		"Atoi":                     reflect.ValueOf(strconv.Atoi),
		"CanBackquote":             reflect.ValueOf(strconv.CanBackquote),
		"ErrRange":                 reflect.ValueOf(&strconv.ErrRange).Elem(),
		"ErrSyntax":                reflect.ValueOf(&strconv.ErrSyntax).Elem(),
		"FormatBool":               reflect.ValueOf(strconv.FormatBool),
		"FormatFloat":              reflect.ValueOf(strconv.FormatFloat),
		"FormatInt":                reflect.ValueOf(strconv.FormatInt),
		"FormatUint":               reflect.ValueOf(strconv.FormatUint),
		"IntSize":                  reflect.ValueOf(strconv.IntSize),
		"IsGraphic":                reflect.ValueOf(strconv.IsGraphic),
		"IsPrint":                  reflect.ValueOf(strconv.IsPrint),
		"Itoa":                     reflect.ValueOf(strconv.Itoa),
		"ParseBool":                reflect.ValueOf(strconv.ParseBool),
		"ParseFloat":               reflect.ValueOf(strconv.ParseFloat),
		"ParseInt":                 reflect.ValueOf(strconv.ParseInt),
		"ParseUint":                reflect.ValueOf(strconv.ParseUint),
		"Quote":                    reflect.ValueOf(strconv.Quote),
		"QuoteRune":                reflect.ValueOf(strconv.QuoteRune),
		"QuoteRuneToASCII":         reflect.ValueOf(strconv.QuoteRuneToASCII),
		"QuoteRuneToGraphic":       reflect.ValueOf(strconv.QuoteRuneToGraphic),
		"QuoteToASCII":             reflect.ValueOf(strconv.QuoteToASCII),
		"QuoteToGraphic":           reflect.ValueOf(strconv.QuoteToGraphic),
		"Unquote":                  reflect.ValueOf(strconv.Unquote),
		"UnquoteChar":              reflect.ValueOf(strconv.UnquoteChar),
	}

	Type["strconv"] = map[string]reflect.Type{
		"NumError": reflect.TypeOf((*strconv.NumError)(nil)).Elem(),
	}
}
