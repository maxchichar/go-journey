# Go Explainer Bot

A small CLI bot that explains Go code by parsing its Abstract Syntax Tree (AST).

## What it explains

- Package name
- Imports
- Type/global declarations
- Functions (params, return values, purpose guess)
- Control-flow constructs (if, for, switch, goroutines, defer, etc.)
- Function calls

## Usage

```bash
go run ./go_explainer_bot -file path/to/file.go
```

Or pass inline code:

```bash
go run ./go_explainer_bot -code 'package main\nfunc main(){}'
```

Or pipe from stdin:

```bash
cat path/to/file.go | go run ./go_explainer_bot
```
