package main

import (
	"os"
	"strings"
	"testing"
)

func main() {
	t := testing.T{}
	var tb testing.TB
	tb = &t
	tmpdir := os.Getenv("TMPDIR")
	if tmpdir == "" {
		println("FAIL")
		return
	}
	if !strings.HasPrefix(tb.TempDir(), tmpdir) {
		println("FAIL")
		return
	}
	println("PASS")
}

// Output:
// PASS