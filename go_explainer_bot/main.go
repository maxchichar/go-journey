package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"sort"
	"strings"
)

type explanation struct {
	PackageName string
	Imports     []string
	Types       []string
	Globals     []string
	Functions   []functionExplanation
}

type functionExplanation struct {
	Name         string
	Receiver     string
	Params       []string
	Results      []string
	Purpose      string
	ControlFlow  []string
	Calls        []string
	ReturnsValue bool
}

func main() {
	filePath := flag.String("file", "", "Path to a Go source file")
	inlineCode := flag.String("code", "", "Inline Go code to explain")
	flag.Parse()

	source, sourceName, err := readSource(*filePath, *inlineCode, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	expl, err := explainGoCode(sourceName, source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(renderExplanation(expl))
}

func readSource(filePath, inlineCode string, stdin io.Reader) (string, string, error) {
	switch {
	case filePath != "":
		b, err := os.ReadFile(filePath)
		if err != nil {
			return "", "", fmt.Errorf("failed reading %s: %w", filePath, err)
		}
		return string(b), filePath, nil
	case inlineCode != "":
		return inlineCode, "inline.go", nil
	default:
		b, err := io.ReadAll(stdin)
		if err != nil {
			return "", "", fmt.Errorf("failed reading stdin: %w", err)
		}
		if len(bytes.TrimSpace(b)) == 0 {
			return "", "", fmt.Errorf("provide input via -file, -code, or stdin")
		}
		return string(b), "stdin.go", nil
	}
}

func explainGoCode(filename, source string) (*explanation, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, source, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("unable to parse Go code: %w", err)
	}

	expl := &explanation{PackageName: file.Name.Name}

	importSet := map[string]struct{}{}
	for _, imp := range file.Imports {
		importSet[strings.Trim(imp.Path.Value, "\"")] = struct{}{}
	}
	expl.Imports = mapKeys(importSet)

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			collectGeneralDecl(fset, d, expl)
		case *ast.FuncDecl:
			expl.Functions = append(expl.Functions, explainFunction(fset, d))
		}
	}

	sort.Strings(expl.Types)
	sort.Strings(expl.Globals)
	sort.Slice(expl.Functions, func(i, j int) bool {
		return expl.Functions[i].Name < expl.Functions[j].Name
	})

	return expl, nil
}

func collectGeneralDecl(fset *token.FileSet, d *ast.GenDecl, expl *explanation) {
	for _, spec := range d.Specs {
		switch s := spec.(type) {
		case *ast.TypeSpec:
			expl.Types = append(expl.Types, s.Name.Name+" ("+nodeString(fset, s.Type)+")")
		case *ast.ValueSpec:
			names := make([]string, 0, len(s.Names))
			for _, n := range s.Names {
				names = append(names, n.Name)
			}
			prefix := strings.ToLower(d.Tok.String())
			expl.Globals = append(expl.Globals, fmt.Sprintf("%s %s", prefix, strings.Join(names, ", ")))
		}
	}
}

func explainFunction(fset *token.FileSet, fn *ast.FuncDecl) functionExplanation {
	fx := functionExplanation{Name: fn.Name.Name}
	if fn.Recv != nil && len(fn.Recv.List) > 0 {
		fx.Receiver = nodeString(fset, fn.Recv.List[0].Type)
	}

	fx.Params = collectFields(fset, fn.Type.Params)
	fx.Results = collectFields(fset, fn.Type.Results)
	fx.ReturnsValue = len(fx.Results) > 0
	fx.Purpose = inferPurpose(fn.Name.Name)

	callSet := map[string]struct{}{}
	flowSet := map[string]struct{}{}

	if fn.Body != nil {
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.IfStmt:
				flowSet["conditional branch (if/else)"] = struct{}{}
			case *ast.ForStmt:
				flowSet["loop (for)"] = struct{}{}
			case *ast.RangeStmt:
				flowSet["collection loop (range)"] = struct{}{}
			case *ast.SwitchStmt:
				flowSet["switch branching"] = struct{}{}
			case *ast.TypeSwitchStmt:
				flowSet["type switch branching"] = struct{}{}
			case *ast.SelectStmt:
				flowSet["channel select branching"] = struct{}{}
			case *ast.GoStmt:
				flowSet["goroutine launch"] = struct{}{}
			case *ast.DeferStmt:
				flowSet["deferred execution"] = struct{}{}
			case *ast.ReturnStmt:
				flowSet["explicit return"] = struct{}{}
			case *ast.CallExpr:
				callSet[nodeString(fset, x.Fun)] = struct{}{}
			}
			return true
		})
	}

	fx.Calls = mapKeys(callSet)
	fx.ControlFlow = mapKeys(flowSet)
	return fx
}

func collectFields(fset *token.FileSet, fields *ast.FieldList) []string {
	if fields == nil {
		return nil
	}

	result := make([]string, 0, len(fields.List))
	for _, f := range fields.List {
		typ := nodeString(fset, f.Type)
		if len(f.Names) == 0 {
			result = append(result, typ)
			continue
		}
		for _, n := range f.Names {
			result = append(result, fmt.Sprintf("%s %s", n.Name, typ))
		}
	}
	return result
}

func renderExplanation(expl *explanation) string {
	var b strings.Builder

	b.WriteString("Go Code Explanation\n")
	b.WriteString("===================\n")
	b.WriteString(fmt.Sprintf("Package: %s\n\n", expl.PackageName))

	if len(expl.Imports) > 0 {
		b.WriteString("Imports:\n")
		for _, imp := range expl.Imports {
			b.WriteString("- " + imp + "\n")
		}
		b.WriteString("\n")
	}

	if len(expl.Types) > 0 {
		b.WriteString("Type declarations:\n")
		for _, t := range expl.Types {
			b.WriteString("- " + t + "\n")
		}
		b.WriteString("\n")
	}

	if len(expl.Globals) > 0 {
		b.WriteString("Global declarations:\n")
		for _, g := range expl.Globals {
			b.WriteString("- " + g + "\n")
		}
		b.WriteString("\n")
	}

	if len(expl.Functions) == 0 {
		b.WriteString("No function declarations found.\n")
		return b.String()
	}

	b.WriteString("Functions:\n")
	for _, fn := range expl.Functions {
		displayName := fn.Name
		if fn.Receiver != "" {
			displayName = fmt.Sprintf("(%s).%s", fn.Receiver, fn.Name)
		}

		b.WriteString("- " + displayName + "\n")
		if len(fn.Params) > 0 {
			b.WriteString("  params: " + strings.Join(fn.Params, ", ") + "\n")
		}
		if len(fn.Results) > 0 {
			b.WriteString("  returns: " + strings.Join(fn.Results, ", ") + "\n")
		}
		b.WriteString("  likely purpose: " + fn.Purpose + "\n")
		if len(fn.ControlFlow) > 0 {
			b.WriteString("  control flow: " + strings.Join(fn.ControlFlow, "; ") + "\n")
		}
		if len(fn.Calls) > 0 {
			b.WriteString("  calls: " + strings.Join(fn.Calls, ", ") + "\n")
		}
	}

	return b.String()
}

func inferPurpose(name string) string {
	lower := strings.ToLower(name)
	switch {
	case name == "main":
		return "program entry point"
	case strings.HasPrefix(lower, "new"):
		return "constructs and returns a new value"
	case strings.HasPrefix(lower, "get"):
		return "retrieves a value"
	case strings.HasPrefix(lower, "set"):
		return "updates state with a new value"
	case strings.HasPrefix(lower, "parse"):
		return "parses input into structured data"
	case strings.HasPrefix(lower, "validate"):
		return "checks whether data is valid"
	case strings.HasPrefix(lower, "handle"):
		return "handles a request/event"
	case strings.HasPrefix(lower, "is") || strings.HasPrefix(lower, "has"):
		return "returns a boolean condition"
	case strings.HasPrefix(lower, "to"):
		return "converts from one representation to another"
	default:
		return "implements business logic for this module"
	}
}

func nodeString(fset *token.FileSet, n ast.Node) string {
	if n == nil {
		return ""
	}
	var b bytes.Buffer
	_ = printer.Fprint(&b, fset, n)
	return b.String()
}

func mapKeys(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
