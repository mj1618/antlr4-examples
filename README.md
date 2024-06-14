# ANTLR4 examples

Build the parsers using antlr4

```bash
antlr4 -Dlanguage=Go  -visitor -o ./go/calc ./Calc.g4
antlr4 -Dlanguage=TypeScript  -visitor -o ./ts/src/calc ./Calc.g4
```

## Go

```bash
go mod tidy
go run .
```

## TypeScript

```bash
go mod tidy
go run .
```
