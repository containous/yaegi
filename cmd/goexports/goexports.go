//go:generate go build

// This program generates code to register binary program symbols to the interpreter.
// See stdlib.go for usage

package main

import (
	"bytes"
	"go/constant"
	"go/format"
	"go/importer"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
)

const model = `package {{.Dest}}

// Code generated by 'goexports {{.PkgName}}'. DO NOT EDIT.

import (
	"{{.PkgName}}"
	"reflect"
)

func init() {
	Value["{{.PkgName}}"] = map[string]reflect.Value{
		{{range $key, $value := .Val -}}
			"{{$key}}": reflect.ValueOf({{$value}}),
		{{end}}
	}

	Type["{{.PkgName}}"] = map[string]reflect.Type{
		{{range $key, $value := .Typ -}}
			"{{$key}}": reflect.TypeOf((*{{$value}})(nil)).Elem(),
		{{end}}
	}
}
`

func genContent(dest, pkgName string) ([]byte, error) {
	p, err := importer.For("source", nil).Import(pkgName)
	if err != nil {
		return nil, err
	}

	typ := map[string]string{}
	val := map[string]string{}
	sc := p.Scope()
	for _, name := range sc.Names() {
		o := sc.Lookup(name)
		if !o.Exported() {
			continue
		}
		pname := path.Base(pkgName) + "." + name
		switch o := o.(type) {
		case *types.Const:
			val[name] = fixConst(pname, o.Val())
		case *types.Func, *types.Var:
			val[name] = pname
		case *types.TypeName:
			typ[name] = pname
		}
	}

	base := template.New("goexports")
	parse, err := base.Parse(model)
	if err != nil {
		return nil, err
	}

	b := &bytes.Buffer{}
	data := map[string]interface{}{
		"Dest":    dest,
		"PkgName": pkgName,
		"Val":     val,
		"Typ":     typ,
	}
	err = parse.Execute(b, data)
	if err != nil {
		return nil, err
	}

	// gofmt
	source, err := format.Source(b.Bytes())
	if err != nil {
		return nil, err
	}
	return source, nil
}

// fixConst checks untyped constant value, converting it if necessary to avoid overflow
func fixConst(name string, val constant.Value) string {
	if val.Kind() == constant.Int {
		str := val.ExactString()
		i, err := strconv.ParseInt(str, 0, 64)
		if err == nil {
			switch {
			case i == int64(int32(i)):
				return name
			case i == int64(uint32(i)):
				return "uint32(" + name + ")"
			default:
				return "int64(" + name + ")"
			}
		}
		_, err = strconv.ParseUint(str, 0, 64)
		if err == nil {
			return "uint64(" + name + ")"
		}
	}
	return name
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dest := path.Base(dir)

	for _, pkg := range os.Args[1:] {
		content, err := genContent(dest, pkg)
		if err != nil {
			log.Fatal(err)
		}

		var oFile string
		if pkg == "syscall" {
			os, arch := os.Getenv("GOOS"), os.Getenv("GOARCH")
			oFile = strings.Replace(pkg, "/", "_", -1) + "_" + os + "_" + arch + ".go"
		} else {
			oFile = strings.Replace(pkg, "/", "_", -1) + ".go"
		}

		err = ioutil.WriteFile(oFile, content, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}
}
