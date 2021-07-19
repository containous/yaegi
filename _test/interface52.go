package main

import (
	"testing"
)

func main() {
	t := testing.T{}
	var tb testing.TB
	tb = &t
	//tmpdir := os.Getenv("TMPDIR")
	//if tmpdir == "" {
	//	println("FAIL")
	//	return
	//}
	// fmt.Fprintln(os.Stdout, "tmpdir:", tmpdir, "testing tmpdir:", tb.TempDir())
	// if !strings.HasPrefix(tb.TempDir(), tmpdir) {
	if tb.TempDir() == "" {
		println("FAIL")
		return
	}
	println("PASS")
}

// Output:
// PASS
