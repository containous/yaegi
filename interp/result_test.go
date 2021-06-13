package interp_test

import (
	"reflect"
	"testing"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type resultTestCase struct {
	desc, src string
	res       interface{}
}

func TestEvalFileResult(t *testing.T) {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	runResultTests(t, i, []resultTestCase{
		{desc: "bare import", src: `import "time"`, res: &interp.PackageImportResult{Path: "time"}},
		{desc: "named import", src: `import x "time"`, res: &interp.PackageImportResult{Name: "x", Path: "time"}},
		{desc: "multiple imports", src: "import (\ny \"time\"\nz \"fmt\"\n)", res: []interp.FileStatementResult{&interp.PackageImportResult{Name: "y", Path: "time"}, &interp.PackageImportResult{Name: "z", Path: "fmt"}}},

		{desc: "func", src: `func foo() {}`, res: &interp.FunctionDeclarationResult{Name: "foo"}},

		{desc: "struct type", src: `type Foo struct {}`, res: &interp.TypeDeclarationResult{Name: "Foo"}},
	})
}

func runResultTests(t *testing.T, i *interp.Interpreter, tests []resultTestCase) {
	t.Helper()

	for _, test := range tests {
		expected := test.res
		if stmtResult, ok := expected.(interp.FileStatementResult); ok {
			expected = []interp.FileStatementResult{stmtResult}
		}
		if stmtResults, ok := expected.([]interp.FileStatementResult); ok {
			expected = &interp.FileResult{Statements: stmtResults}
		}

		t.Run(test.desc, func(t *testing.T) {
			res, err := i.Eval(test.src)
			if err != nil {
				t.Fatal(err)
			}
			if !res.IsValid() {
				t.Fatal("Result is not valid")
			}
			if !reflect.DeepEqual(expected, res.Interface()) {
				t.Fatalf("Got %v, expected %v", res, expected)
			}
		})
	}
}
