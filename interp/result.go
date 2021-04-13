package interp

type FileResult struct {
	Statements []FileStatementResult
}

type FileStatementResult interface {
	isFileStatementResult()
}

func (*PackageImportResult) isFileStatementResult()       {}
func (*FunctionDeclarationResult) isFileStatementResult() {}
func (*TypeDeclarationResult) isFileStatementResult()     {}

type PackageImportResult struct {
	Name, Path string
}

type FunctionDeclarationResult struct {
	Name string
}

type TypeDeclarationResult struct {
	Name string
}
