package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"os/exec"

	"github.com/go2hx/controlflow"
)

func main() {
	tests := []string{
		//"goto0",
		//"break1",
		"test",
	}
	for _, test := range tests {
		create("tests/unit", test+".go")
	}
}

func create(filePath string, fileName string) {
	// Parse the code
	code, err := os.ReadFile(filePath + "/" + fileName)
	if err != nil {
		panic(err)
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "code.go", code, 0)
	if err != nil {
		panic(err)
	}

	// List of statements in the function's body
	body := file.Decls[0].(*ast.FuncDecl).Body.List

	// Rewrite the function body without goto statements
	newBody := controlflow.ElimGotos(body)

	// Print the result
	config := printer.Config{Mode: printer.UseSpaces, Tabwidth: 4}
	// new file
	newPath := filePath + "mod/" + fileName
	modFile, err := os.Create(newPath)
	if err != nil {
		panic(err)
	}
	// wrap it in main func
	modFile.WriteString("package main\nfunc main() {\n")
	config.Fprint(modFile, token.NewFileSet(), newBody)
	modFile.WriteString("\n}")

	oldPath := filePath + "/" + fileName
	run(oldPath, newPath)
}

func run(oldPath, newPath string) {
	runCmd(oldPath)
	runCmd(newPath)
}

func runCmd(runPath string) {
	fmt.Println(runPath + ":")
	out, err := exec.Command("go", "run", runPath).CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		panic(out)
	}
}
