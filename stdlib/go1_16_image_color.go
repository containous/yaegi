// Code generated by 'yaegi extract image/color'. DO NOT EDIT.

// +build go1.16

package stdlib

import (
	"image/color"
	"reflect"
)

func init() {
	Symbols["image/color/color"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Alpha16Model": reflect.ValueOf(&color.Alpha16Model).Elem(),
		"AlphaModel":   reflect.ValueOf(&color.AlphaModel).Elem(),
		"Black":        reflect.ValueOf(&color.Black).Elem(),
		"CMYKModel":    reflect.ValueOf(&color.CMYKModel).Elem(),
		"CMYKToRGB":    reflect.ValueOf(color.CMYKToRGB),
		"Gray16Model":  reflect.ValueOf(&color.Gray16Model).Elem(),
		"GrayModel":    reflect.ValueOf(&color.GrayModel).Elem(),
		"ModelFunc":    reflect.ValueOf(color.ModelFunc),
		"NRGBA64Model": reflect.ValueOf(&color.NRGBA64Model).Elem(),
		"NRGBAModel":   reflect.ValueOf(&color.NRGBAModel).Elem(),
		"NYCbCrAModel": reflect.ValueOf(&color.NYCbCrAModel).Elem(),
		"Opaque":       reflect.ValueOf(&color.Opaque).Elem(),
		"RGBA64Model":  reflect.ValueOf(&color.RGBA64Model).Elem(),
		"RGBAModel":    reflect.ValueOf(&color.RGBAModel).Elem(),
		"RGBToCMYK":    reflect.ValueOf(color.RGBToCMYK),
		"RGBToYCbCr":   reflect.ValueOf(color.RGBToYCbCr),
		"Transparent":  reflect.ValueOf(&color.Transparent).Elem(),
		"White":        reflect.ValueOf(&color.White).Elem(),
		"YCbCrModel":   reflect.ValueOf(&color.YCbCrModel).Elem(),
		"YCbCrToRGB":   reflect.ValueOf(color.YCbCrToRGB),

		// type definitions
		"Alpha":   reflect.ValueOf((*color.Alpha)(nil)),
		"Alpha16": reflect.ValueOf((*color.Alpha16)(nil)),
		"CMYK":    reflect.ValueOf((*color.CMYK)(nil)),
		"Color":   reflect.ValueOf((*color.Color)(nil)),
		"Gray":    reflect.ValueOf((*color.Gray)(nil)),
		"Gray16":  reflect.ValueOf((*color.Gray16)(nil)),
		"Model":   reflect.ValueOf((*color.Model)(nil)),
		"NRGBA":   reflect.ValueOf((*color.NRGBA)(nil)),
		"NRGBA64": reflect.ValueOf((*color.NRGBA64)(nil)),
		"NYCbCrA": reflect.ValueOf((*color.NYCbCrA)(nil)),
		"Palette": reflect.ValueOf((*color.Palette)(nil)),
		"RGBA":    reflect.ValueOf((*color.RGBA)(nil)),
		"RGBA64":  reflect.ValueOf((*color.RGBA64)(nil)),
		"YCbCr":   reflect.ValueOf((*color.YCbCr)(nil)),

		// interface wrapper definitions
		"_Color": reflect.ValueOf((*_image_color_Color)(nil)),
		"_Model": reflect.ValueOf((*_image_color_Model)(nil)),
	}
}

// _image_color_Color is an interface wrapper for Color type
type _image_color_Color struct {
	WRGBA func() (r uint32, g uint32, b uint32, a uint32)
}

func (W _image_color_Color) RGBA() (r uint32, g uint32, b uint32, a uint32) { return W.WRGBA() }

// _image_color_Model is an interface wrapper for Model type
type _image_color_Model struct {
	WConvert func(c color.Color) color.Color
}

func (W _image_color_Model) Convert(c color.Color) color.Color { return W.WConvert(c) }
