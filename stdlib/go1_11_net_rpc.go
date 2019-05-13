// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports net/rpc'. DO NOT EDIT.

import (
	"net/rpc"
	"reflect"
)

func init() {
	Value["net/rpc"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Accept":             reflect.ValueOf(rpc.Accept),
		"DefaultDebugPath":   reflect.ValueOf(rpc.DefaultDebugPath),
		"DefaultRPCPath":     reflect.ValueOf(rpc.DefaultRPCPath),
		"DefaultServer":      reflect.ValueOf(&rpc.DefaultServer).Elem(),
		"Dial":               reflect.ValueOf(rpc.Dial),
		"DialHTTP":           reflect.ValueOf(rpc.DialHTTP),
		"DialHTTPPath":       reflect.ValueOf(rpc.DialHTTPPath),
		"ErrShutdown":        reflect.ValueOf(&rpc.ErrShutdown).Elem(),
		"HandleHTTP":         reflect.ValueOf(rpc.HandleHTTP),
		"NewClient":          reflect.ValueOf(rpc.NewClient),
		"NewClientWithCodec": reflect.ValueOf(rpc.NewClientWithCodec),
		"NewServer":          reflect.ValueOf(rpc.NewServer),
		"Register":           reflect.ValueOf(rpc.Register),
		"RegisterName":       reflect.ValueOf(rpc.RegisterName),
		"ServeCodec":         reflect.ValueOf(rpc.ServeCodec),
		"ServeConn":          reflect.ValueOf(rpc.ServeConn),
		"ServeRequest":       reflect.ValueOf(rpc.ServeRequest),

		// type definitions
		"Call":        reflect.ValueOf((*rpc.Call)(nil)),
		"Client":      reflect.ValueOf((*rpc.Client)(nil)),
		"ClientCodec": reflect.ValueOf((*rpc.ClientCodec)(nil)),
		"Request":     reflect.ValueOf((*rpc.Request)(nil)),
		"Response":    reflect.ValueOf((*rpc.Response)(nil)),
		"Server":      reflect.ValueOf((*rpc.Server)(nil)),
		"ServerCodec": reflect.ValueOf((*rpc.ServerCodec)(nil)),
		"ServerError": reflect.ValueOf((*rpc.ServerError)(nil)),
	}
	Wrapper["net/rpc"] = map[string]reflect.Type{
		"ClientCodec": reflect.TypeOf((*_net_rpc_ClientCodec)(nil)),
		"ServerCodec": reflect.TypeOf((*_net_rpc_ServerCodec)(nil)),
	}
}

// _net_rpc_ClientCodec is an interface wrapper for ClientCodec type
type _net_rpc_ClientCodec struct {
	WClose              func() error
	WReadResponseBody   func(a0 interface{}) error
	WReadResponseHeader func(a0 *rpc.Response) error
	WWriteRequest       func(a0 *rpc.Request, a1 interface{}) error
}

func (W _net_rpc_ClientCodec) Close() error                          { return W.WClose() }
func (W _net_rpc_ClientCodec) ReadResponseBody(a0 interface{}) error { return W.WReadResponseBody(a0) }
func (W _net_rpc_ClientCodec) ReadResponseHeader(a0 *rpc.Response) error {
	return W.WReadResponseHeader(a0)
}
func (W _net_rpc_ClientCodec) WriteRequest(a0 *rpc.Request, a1 interface{}) error {
	return W.WWriteRequest(a0, a1)
}

// _net_rpc_ServerCodec is an interface wrapper for ServerCodec type
type _net_rpc_ServerCodec struct {
	WClose             func() error
	WReadRequestBody   func(a0 interface{}) error
	WReadRequestHeader func(a0 *rpc.Request) error
	WWriteResponse     func(a0 *rpc.Response, a1 interface{}) error
}

func (W _net_rpc_ServerCodec) Close() error                         { return W.WClose() }
func (W _net_rpc_ServerCodec) ReadRequestBody(a0 interface{}) error { return W.WReadRequestBody(a0) }
func (W _net_rpc_ServerCodec) ReadRequestHeader(a0 *rpc.Request) error {
	return W.WReadRequestHeader(a0)
}
func (W _net_rpc_ServerCodec) WriteResponse(a0 *rpc.Response, a1 interface{}) error {
	return W.WWriteResponse(a0, a1)
}
