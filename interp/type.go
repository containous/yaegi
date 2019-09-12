package interp

import (
	"reflect"
	"strconv"
)

// tcat defines interpreter type categories
type tcat uint

// Types for go language
const (
	nilT tcat = iota
	aliasT
	arrayT
	binT
	binPkgT
	boolT
	builtinT
	byteT
	chanT
	complex64T
	complex128T
	errorT
	float32T
	float64T
	funcT
	interfaceT
	intT
	int8T
	int16T
	int32T
	int64T
	mapT
	ptrT
	runeT
	srcPkgT
	stringT
	structT
	uintT
	uint8T
	uint16T
	uint32T
	uint64T
	uintptrT
	valueT
	variadicT
	maxT
)

var cats = [...]string{
	nilT:        "nilT",
	aliasT:      "aliasT",
	arrayT:      "arrayT",
	binT:        "binT",
	binPkgT:     "binPkgT",
	byteT:       "byteT",
	boolT:       "boolT",
	builtinT:    "builtinT",
	chanT:       "chanT",
	complex64T:  "complex64T",
	complex128T: "complex128T",
	errorT:      "errorT",
	float32T:    "float32",
	float64T:    "float64T",
	funcT:       "funcT",
	interfaceT:  "interfaceT",
	intT:        "intT",
	int8T:       "int8T",
	int16T:      "int16T",
	int32T:      "int32T",
	int64T:      "int64T",
	mapT:        "mapT",
	ptrT:        "ptrT",
	runeT:       "runeT",
	srcPkgT:     "srcPkgT",
	stringT:     "stringT",
	structT:     "structT",
	uintT:       "uintT",
	uint8T:      "uint8T",
	uint16T:     "uint16T",
	uint32T:     "uint32T",
	uint64T:     "uint64T",
	uintptrT:    "uintptrT",
	valueT:      "valueT",
	variadicT:   "variadicT",
}

func (c tcat) String() string {
	if c < tcat(len(cats)) {
		return cats[c]
	}
	return "Cat(" + strconv.Itoa(int(c)) + ")"
}

// structField type defines a field in a struct
type structField struct {
	name  string
	tag   string
	embed bool
	typ   *itype
}

// itype defines the internal representation of types in the interpreter
type itype struct {
	cat        tcat          // Type category
	field      []structField // Array of struct fields if structT or interfaceT
	key        *itype        // Type of key element if MapT or nil
	val        *itype        // Type of value element if chanT, mapT, ptrT, aliasT, arrayT or variadicT
	arg        []*itype      // Argument types if funcT or nil
	ret        []*itype      // Return types if funcT or nil
	method     []*node       // Associated methods or nil
	name       string        // name of type within its package for a defined type
	path       string        // for a defined type, the package import path
	size       int           // Size of array if ArrayT
	rtype      reflect.Type  // Reflection type if ValueT, or nil
	incomplete bool          // true if type must be parsed again (out of order declarations)
	untyped    bool          // true for a literal value (string or number)
	sizedef    bool          // true if array size is computed from type definition
	node       *node         // root AST node of type definition
	scope      *scope        // type declaration scope (in case of re-parse incomplete type)
}

// nodeType returns a type definition for the corresponding AST subtree
func nodeType(interp *Interpreter, sc *scope, n *node) (*itype, error) {
	var err cfgError

	if n.typ != nil && !n.typ.incomplete {
		return n.typ, err
	}

	var t = &itype{node: n, scope: sc}

	if n.anc.kind == typeSpec {
		name := n.anc.child[0].ident
		if sym := sc.sym[name]; sym != nil {
			// recover previously declared methods
			t.method = sym.typ.method
			t.path = sym.typ.path
			t.name = name
		}
	}

	switch n.kind {
	case addressExpr, starExpr:
		t.cat = ptrT
		if t.val, err = nodeType(interp, sc, n.child[0]); err != nil {
			return nil, err
		}
		t.incomplete = t.val.incomplete

	case arrayType:
		t.cat = arrayT
		if len(n.child) > 1 {
			switch {
			case n.child[0].rval.IsValid():
				// constant size
				t.size = int(n.child[0].rval.Int())
			case n.child[0].kind == ellipsisExpr:
				// [...]T expression
				t.sizedef = true
			default:
				if sym, _, ok := sc.lookup(n.child[0].ident); ok {
					// Resolve symbol to get size value
					if sym.typ != nil && sym.typ.cat == intT {
						if v, ok := sym.rval.Interface().(int); ok {
							t.size = v
						} else {
							t.incomplete = true
						}
					} else {
						t.incomplete = true
					}
				} else {
					// Evaluate constant array size expression
					if _, err = interp.cfg(n.child[0]); err != nil {
						return nil, err
					}
					t.incomplete = true
				}
			}
			if t.val, err = nodeType(interp, sc, n.child[1]); err != nil {
				return nil, err
			}
			t.incomplete = t.incomplete || t.val.incomplete
		} else {
			if t.val, err = nodeType(interp, sc, n.child[0]); err != nil {
				return nil, err
			}
			t.incomplete = t.val.incomplete
		}

	case basicLit:
		switch v := n.rval.Interface().(type) {
		case bool:
			t.cat = boolT
			t.name = "bool"
		case byte:
			t.cat = byteT
			t.name = "byte"
			t.untyped = true
		case complex64:
			t.cat = complex64T
			t.name = "complex64"
		case complex128:
			t.cat = complex128T
			t.name = "complex128"
			t.untyped = true
		case float32:
			t.cat = float32T
			t.name = "float32"
			t.untyped = true
		case float64:
			t.cat = float64T
			t.name = "float64"
			t.untyped = true
		case int:
			if isShiftOperand(n) && v >= 0 {
				t.cat = uintT
				t.name = "uint"
				n.rval = reflect.ValueOf(uint(v))
			} else {
				t.cat = intT
				t.name = "int"
			}
			t.untyped = true
		case rune:
			t.cat = runeT
			t.name = "rune"
			t.untyped = true
		case string:
			t.cat = stringT
			t.name = "string"
			t.untyped = true
		default:
			err = n.cfgErrorf("missing support for type %T: %v", v, n.rval)
		}

	case unaryExpr:
		t, err = nodeType(interp, sc, n.child[0])

	case binaryExpr:
		if a := n.anc; a.kind == defineStmt && len(a.child) > a.nleft+a.nright {
			t, err = nodeType(interp, sc, a.child[a.nleft])
		} else {
			if t, err = nodeType(interp, sc, n.child[0]); err != nil {
				return nil, err
			}
			if t.untyped {
				var t1 *itype
				t1, err = nodeType(interp, sc, n.child[1])
				if !(t1.untyped && isInt(t1.TypeOf()) && isFloat(t.TypeOf())) {
					t = t1
				}
			}
		}

	case callExpr:
		if t, err = nodeType(interp, sc, n.child[0]); err != nil {
			return nil, err
		}
		switch t.cat {
		case valueT:
			if t.rtype.NumOut() == 1 {
				t = &itype{cat: valueT, rtype: t.rtype.Out(0)}
			}
		default:
			if len(t.ret) == 1 {
				t = t.ret[0]
			}
		}

	case compositeLitExpr:
		t, err = nodeType(interp, sc, n.child[0])

	case chanType:
		t.cat = chanT
		if t.val, err = nodeType(interp, sc, n.child[0]); err != nil {
			return nil, err
		}
		t.incomplete = t.val.incomplete

	case ellipsisExpr:
		t.cat = variadicT
		if t.val, err = nodeType(interp, sc, n.child[0]); err != nil {
			return nil, err
		}
		t.incomplete = t.val.incomplete

	case funcLit:
		t, err = nodeType(interp, sc, n.child[2])

	case funcType:
		t.cat = funcT
		// Handle input parameters
		for _, arg := range n.child[0].child {
			cl := len(arg.child) - 1
			typ, err := nodeType(interp, sc, arg.child[cl])
			if err != nil {
				return nil, err
			}
			t.arg = append(t.arg, typ)
			for i := 1; i < cl; i++ {
				// Several arguments may be factorized on the same field type
				t.arg = append(t.arg, typ)
			}
			t.incomplete = t.incomplete || typ.incomplete
		}
		if len(n.child) == 2 {
			// Handle returned values
			for _, ret := range n.child[1].child {
				cl := len(ret.child) - 1
				typ, err := nodeType(interp, sc, ret.child[cl])
				if err != nil {
					return nil, err
				}
				t.ret = append(t.ret, typ)
				for i := 1; i < cl; i++ {
					// Several arguments may be factorized on the same field type
					t.ret = append(t.ret, typ)
				}
				t.incomplete = t.incomplete || typ.incomplete
			}
		}

	case identExpr:
		if sym, _, found := sc.lookup(n.ident); found {
			t = sym.typ
			if t.incomplete && t.node != n {
				m := t.method
				if t, err = nodeType(interp, sc, t.node); err != nil {
					return nil, err
				}
				t.method = m
				sym.typ = t
			}
		} else {
			t.incomplete = true
			sc.sym[n.ident] = &symbol{kind: typeSym, typ: t}
		}

	case interfaceType:
		t.cat = interfaceT
		for _, field := range n.child[0].child {
			if len(field.child) == 1 {
				typ, err := nodeType(interp, sc, field.child[0])
				if err != nil {
					return nil, err
				}
				t.field = append(t.field, structField{name: fieldName(field.child[0]), embed: true, typ: typ})
				t.incomplete = t.incomplete || typ.incomplete
			} else {
				typ, err := nodeType(interp, sc, field.child[1])
				if err != nil {
					return nil, err
				}
				t.field = append(t.field, structField{name: field.child[0].ident, typ: typ})
				t.incomplete = t.incomplete || typ.incomplete
			}
		}

	case mapType:
		t.cat = mapT
		if t.key, err = nodeType(interp, sc, n.child[0]); err != nil {
			return nil, err
		}
		if t.val, err = nodeType(interp, sc, n.child[1]); err != nil {
			return nil, err
		}
		t.incomplete = t.key.incomplete || t.val.incomplete

	case parenExpr:
		t, err = nodeType(interp, sc, n.child[0])

	case selectorExpr:
		// Resolve the left part of selector, then lookup the right part on it
		var lt *itype
		if lt, err = nodeType(interp, sc, n.child[0]); err != nil {
			return nil, err
		}
		if lt.incomplete {
			t.incomplete = true
			break
		}
		name := n.child[1].ident
		switch lt.cat {
		case binPkgT:
			pkg := interp.binPkg[lt.path]
			if v, ok := pkg[name]; ok {
				t.cat = valueT
				t.rtype = v.Type()
				if isBinType(v) { // a bin type is encoded as a pointer on nil value
					t.rtype = t.rtype.Elem()
				}
			} else {
				err = n.cfgErrorf("undefined selector %s.%s", lt.path, name)
				panic(err)
			}
		case srcPkgT:
			pkg := interp.srcPkg[lt.path]
			if s, ok := pkg[name]; ok {
				t = s.typ
			} else {
				err = n.cfgErrorf("undefined selector %s.%s", lt.path, name)
			}
		default:
			if m, _ := lt.lookupMethod(name); m != nil {
				t, err = nodeType(interp, sc, m.child[2])
			} else if bm, _, ok := lt.lookupBinMethod(name); ok {
				t = &itype{cat: valueT, rtype: bm.Type}
			} else if ti := lt.lookupField(name); len(ti) > 0 {
				t = lt.fieldSeq(ti)
			} else if bs, _, ok := lt.lookupBinField(name); ok {
				t = &itype{cat: valueT, rtype: bs.Type}
			} else {
				err = lt.node.cfgErrorf("undefined selector %s", name)
			}
		}

	case structType:
		t.cat = structT
		var incomplete, found bool
		var sym *symbol
		if sname := structName(n); sname != "" {
			if sym, _, found = sc.lookup(sname); found && sym.kind == typeSym {
				sym.typ = t
			}
		}
		for _, c := range n.child[0].child {
			switch {
			case len(c.child) == 1:
				typ, err := nodeType(interp, sc, c.child[0])
				if err != nil {
					return nil, err
				}
				t.field = append(t.field, structField{name: fieldName(c.child[0]), embed: true, typ: typ})
				incomplete = incomplete || typ.incomplete
			case len(c.child) == 2 && c.child[1].kind == basicLit:
				tag := c.child[1].rval.String()
				typ, err := nodeType(interp, sc, c.child[0])
				if err != nil {
					return nil, err
				}
				t.field = append(t.field, structField{name: fieldName(c.child[0]), embed: true, typ: typ, tag: tag})
				incomplete = incomplete || typ.incomplete
			default:
				var tag string
				l := len(c.child)
				if c.lastChild().kind == basicLit {
					tag = c.lastChild().rval.String()
					l--
				}
				typ, err := nodeType(interp, sc, c.child[l-1])
				if err != nil {
					return nil, err
				}
				incomplete = incomplete || typ.incomplete
				for _, d := range c.child[:l-1] {
					t.field = append(t.field, structField{name: d.ident, typ: typ, tag: tag})
				}
			}
		}
		t.incomplete = incomplete

	default:
		err = n.cfgErrorf("type definition not implemented: %s", n.kind)
	}

	if err == nil && t.cat == nilT && !t.incomplete {
		err = n.cfgErrorf("use of untyped nil %s", t.name)
	}

	return t, err
}

// struct name returns the name of a struct type
func structName(n *node) string {
	if n.anc.kind == typeSpec {
		return n.anc.child[0].ident
	}
	return ""
}

// fieldName returns an implicit struct field name according to node kind
func fieldName(n *node) string {
	switch n.kind {
	case selectorExpr:
		return fieldName(n.child[1])
	case starExpr:
		return fieldName(n.child[0])
	case identExpr:
		return n.ident
	default:
		return ""
	}
}

var zeroValues [maxT]reflect.Value

func init() {
	zeroValues[boolT] = reflect.ValueOf(false)
	zeroValues[byteT] = reflect.ValueOf(byte(0))
	zeroValues[complex64T] = reflect.ValueOf(complex64(0))
	zeroValues[complex128T] = reflect.ValueOf(complex128(0))
	zeroValues[errorT] = reflect.ValueOf(new(error)).Elem()
	zeroValues[float32T] = reflect.ValueOf(float32(0))
	zeroValues[float64T] = reflect.ValueOf(float64(0))
	zeroValues[intT] = reflect.ValueOf(int(0))
	zeroValues[int8T] = reflect.ValueOf(int8(0))
	zeroValues[int16T] = reflect.ValueOf(int16(0))
	zeroValues[int32T] = reflect.ValueOf(int32(0))
	zeroValues[int64T] = reflect.ValueOf(int64(0))
	zeroValues[runeT] = reflect.ValueOf(rune(0))
	zeroValues[stringT] = reflect.ValueOf("")
	zeroValues[uintT] = reflect.ValueOf(uint(0))
	zeroValues[uint8T] = reflect.ValueOf(uint8(0))
	zeroValues[uint16T] = reflect.ValueOf(uint16(0))
	zeroValues[uint32T] = reflect.ValueOf(uint32(0))
	zeroValues[uint64T] = reflect.ValueOf(uint64(0))
	zeroValues[uintptrT] = reflect.ValueOf(uintptr(0))
}

// if type is incomplete, re-parse it.
func (t *itype) finalize() (*itype, error) {
	var err cfgError
	if t.incomplete {
		sym, _, found := t.scope.lookup(t.name)
		if found && !sym.typ.incomplete {
			sym.typ.method = append(sym.typ.method, t.method...)
			return sym.typ, nil
		}
		m := t.method
		if t, err = nodeType(t.node.interp, t.scope, t.node); err != nil {
			return nil, err
		}
		if t.incomplete {
			return nil, t.node.cfgErrorf("incomplete type %s", t.name)
		}
		t.method = m
		t.node.typ = t
		if sym != nil {
			sym.typ = t
		}
	}
	return t, err
}

// id returns a unique type identificator string
func (t *itype) id() string {
	// TODO: if res is nil, build identity from String()

	res := ""
	switch t.cat {
	case valueT:
		res = t.rtype.PkgPath() + "." + t.rtype.Name()
	case ptrT:
		res = "*" + t.val.id()
	default:
		res = t.path + "." + t.name
	}
	return res
}

// zero instantiates and return a zero value object for the given type during execution
func (t *itype) zero() (v reflect.Value, err error) {
	if t, err = t.finalize(); err != nil {
		return v, err
	}
	switch t.cat {
	case aliasT:
		v, err = t.val.zero()

	case arrayT, ptrT, structT:
		v = reflect.New(t.TypeOf()).Elem()

	case valueT:
		v = reflect.New(t.rtype).Elem()

	default:
		v = zeroValues[t.cat]
	}
	return v, err
}

// fieldIndex returns the field index from name in a struct, or -1 if not found
func (t *itype) fieldIndex(name string) int {
	if t.cat == ptrT {
		return t.val.fieldIndex(name)
	}
	for i, field := range t.field {
		if name == field.name {
			return i
		}
	}
	return -1
}

// fieldSeq returns the field type from the list of field indexes
func (t *itype) fieldSeq(seq []int) *itype {
	ft := t
	for _, i := range seq {
		if ft.cat == ptrT {
			ft = ft.val
		}
		ft = ft.field[i].typ
	}
	return ft
}

// lookupField returns a list of indices, i.e. a path to access a field in a struct object
func (t *itype) lookupField(name string) []int {
	if fi := t.fieldIndex(name); fi >= 0 {
		return []int{fi}
	}

	for i, f := range t.field {
		switch f.typ.cat {
		case ptrT, structT:
			if index2 := f.typ.lookupField(name); len(index2) > 0 {
				return append([]int{i}, index2...)
			}
		}
	}

	return nil
}

// lookupBinField returns a structfield and a path to access an embedded binary field in a struct object
func (t *itype) lookupBinField(name string) (s reflect.StructField, index []int, ok bool) {
	if t.cat == ptrT {
		return t.val.lookupBinField(name)
	}
	if !isStruct(t) {
		return
	}
	s, ok = t.TypeOf().FieldByName(name)
	if !ok {
		for i, f := range t.field {
			if f.embed {
				if s2, index2, ok2 := f.typ.lookupBinField(name); ok2 {
					index = append([]int{i}, index2...)
					return s2, index, ok2
				}
			}
		}
	}
	return s, index, ok
}

// getMethod returns a pointer to the method definition
func (t *itype) getMethod(name string) *node {
	for _, m := range t.method {
		if name == m.ident {
			return m
		}
	}
	return nil
}

// lookupMethod returns a pointer to method definition associated to type t
// and the list of indices to access the right struct field, in case of an embedded method
func (t *itype) lookupMethod(name string) (*node, []int) {
	if t.cat == ptrT {
		return t.val.lookupMethod(name)
	}
	var index []int
	m := t.getMethod(name)
	if m == nil {
		for i, f := range t.field {
			if f.embed {
				if n, index2 := f.typ.lookupMethod(name); n != nil {
					index = append([]int{i}, index2...)
					return n, index
				}
			}
		}
	}
	return m, index
}

// lookupBinMethod returns a method and a path to access a field in a struct object (the receiver)
func (t *itype) lookupBinMethod(name string) (reflect.Method, []int, bool) {
	if t.cat == ptrT {
		return t.val.lookupBinMethod(name)
	}
	var index []int
	m, ok := t.TypeOf().MethodByName(name)
	if !ok {
		for i, f := range t.field {
			if f.embed {
				if m2, index2, ok2 := f.typ.lookupBinMethod(name); ok2 {
					index = append([]int{i}, index2...)
					return m2, index, ok2
				}
			}
		}
	}
	return m, index, ok
}

func exportName(s string) string {
	if canExport(s) {
		return s
	}
	return "X" + s
}

// TypeOf returns the reflection type of dynamic interpreter type t.
func (t *itype) Type(name string) reflect.Type {
	if t.rtype != nil {
		return t.rtype
	}

	if t.incomplete || t.cat == nilT {
		var err error
		if t, err = t.finalize(); err != nil {
			panic(err)
		}
	}
	if name == "" {
		name = t.name
	}

	switch t.cat {
	case arrayT, variadicT:
		if t.size > 0 {
			t.rtype = reflect.ArrayOf(t.size, t.val.Type(name))
		} else {
			t.rtype = reflect.SliceOf(t.val.Type(name))
		}
	case chanT:
		t.rtype = reflect.ChanOf(reflect.BothDir, t.val.Type(name))
	case errorT:
		t.rtype = reflect.TypeOf(new(error)).Elem()
	case funcT:
		in := make([]reflect.Type, len(t.arg))
		out := make([]reflect.Type, len(t.ret))
		for i, v := range t.arg {
			in[i] = v.Type(name)
		}
		for i, v := range t.ret {
			out[i] = v.Type(name)
		}
		t.rtype = reflect.FuncOf(in, out, false)
	case interfaceT:
		t.rtype = reflect.TypeOf(new(interface{})).Elem()
	case mapT:
		t.rtype = reflect.MapOf(t.key.TypeOf(), t.val.TypeOf())
	case ptrT:
		if t.val != nil && name != "" && t.val.name == name {
			t.val.rtype = reflect.TypeOf(new(interface{})).Elem()
		}
		t.rtype = reflect.PtrTo(t.val.Type(name))
	case structT:
		var fields []reflect.StructField
		for _, f := range t.field {
			field := reflect.StructField{Name: exportName(f.name), Type: f.typ.Type(name), Tag: reflect.StructTag(f.tag)}
			fields = append(fields, field)
		}
		t.rtype = reflect.StructOf(fields)
	default:
		if z, _ := t.zero(); z.IsValid() {
			t.rtype = z.Type()
		}
	}
	return t.rtype
}

func (t *itype) TypeOf() reflect.Type {
	return t.Type(t.name)
}

func (t *itype) frameType() (r reflect.Type) {
	var err error
	if t, err = t.finalize(); err != nil {
		panic(err)
	}
	switch t.cat {
	case arrayT, variadicT:
		if t.size > 0 {
			r = reflect.ArrayOf(t.size, t.val.frameType())
		} else {
			r = reflect.SliceOf(t.val.frameType())
		}
	case funcT:
		r = reflect.TypeOf((*node)(nil))
	case interfaceT:
		r = reflect.TypeOf((*valueInterface)(nil)).Elem()
	default:
		r = t.TypeOf()
	}
	return r
}

func (t *itype) implements(it *itype) bool {
	if t.cat == valueT {
		return t.TypeOf().Implements(it.TypeOf())
	}
	// TODO: implement method check for interpreted types
	return true
}

func defRecvType(n *node) *itype {
	if n.kind != funcDecl || len(n.child[0].child) == 0 {
		return nil
	}
	if r := n.child[0].child[0].lastChild(); r != nil {
		return r.typ
	}
	return nil
}

func isShiftOperand(n *node) bool {
	switch n.anc.action {
	case aShl, aShr, aShlAssign, aShrAssign:
		return n.anc.lastChild() == n
	}
	return false
}

func isInterface(t *itype) bool { return t.cat == interfaceT || t.TypeOf().Kind() == reflect.Interface }

func isStruct(t *itype) bool { return t.TypeOf().Kind() == reflect.Struct }

func isInt(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	}
	return false
}

func isUint(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	}
	return false
}

func isComplex(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Complex64, reflect.Complex128:
		return true
	}
	return false
}

func isFloat(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func isByteArray(t reflect.Type) bool {
	k := t.Kind()
	return (k == reflect.Array || k == reflect.Slice) && t.Elem().Kind() == reflect.Uint8
}

func isFloat32(t reflect.Type) bool { return t.Kind() == reflect.Float32 }
func isFloat64(t reflect.Type) bool { return t.Kind() == reflect.Float64 }
func isNumber(t reflect.Type) bool  { return isInt(t) || isFloat(t) || isComplex(t) }
func isString(t reflect.Type) bool  { return t.Kind() == reflect.String }
