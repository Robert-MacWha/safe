# Safe

Rust-like safe Option & Result generics in Go.

Helpful for more consise error-handling and letting you return instances instead of pointers from functions that can error.

## Features

-   `Result` and `Option` types.
-   Unwrap Handler to streamline error management.
-   SafeLint linter to ensure correct usage of this package.

## Install

```
go get github.com/robert-macwha/safe
```

### Linter Install

This linter can be installed as a golangci-lint plugin. Make sure you have golangci-lint installed.

1. Create a `.custom-gcl.yml` file with the following contents.

```yml
version: v1.57.0
plugins:
    - module: "github.com/robert-macwha/safe"
      import: "github.com/robert-macwha/safe/cmd/gcl"
      version: v0.0.4
```

2. Update your `.golangci.yml` file to include the safe plugin.

```yml
linters-settings:
    custom:
        safelint:
            type: "module"
            description: "Linter for `robert-macwha/safe` package"
```

3. Build the custom instance of golangci-lint with the following command, creating a `custom-gcl` file.

```bash
golangci-lint custom
```

4. If using vscode, update the workspace's settings.json to point to the new golangci-lint binary.

```json
{
    "go.lintTool": "golangci-lint",
    "go.alternateTools": {
        "golangci-lint": "${workspaceRoot}/custom-gcl"
    }
}
```

## Types

### `Result`

Interface encapsulates success or failure state.

```go
// Result is a generic rust-like Result type.
type Result[T any] interface {
	// IsOk returns true if the Result is Ok.
	IsOk() bool
	// IsErr returns true if the Result is Err.
	IsErr() bool
	// Unwrap returns the data if the Result is Ok, panics if the Result is Err.
	Unwrap() T
	// UnwrapErr returns the error if the Result is Err, panics if the Result
	// is Ok.
	UnwrapErr() error
	// UnwrapFmt returns a formatted error if the Result is Err, panics if the
	// Result is Ok.
	UnwrapFmt(s string) error
}
```

### `Option`

Interface represents a value that may or may not exist.

```go
// Option is a generic rust-like Option type.
type Option[T any] interface {
	// IsSome returns true if the Option is Some.
	IsSome() bool
	// IsNone returns true if the Option is None.
	IsNone() bool
	// Unwrap returns the data if the Option is Some, panics if the Option is None.
	Unwrap() T
	// UnwrapOr returns the data if the Option is Some, otherwise returns the
	// default value.
	UnwrapOr(def T) T
}
```

## Handler

The safe handler lets you safely unwrap `Options` or `Results`, catch any panics in another `Result` object, and return that `Result` from a function. This replaces the need for the `if err != nil {}` pattern and lets your code focus on the happy execution path.

See [example/io](./example/io/main.go) for an example.

## SafeLint

SafeLint is a linter designed for this package to ensure correct usage.

### UnwrapFmtAnalyzer

UnwrapFmtAnalyzer is a linter that checks whether the string passed to Result's `UnwrapFmt` function has a single "`%w`"verb.

### SafeHandlerAnalyzer

SafeHandlerAnalyzer is a linter that ensures `safe.Handle` is called before unwrapping any results or options.
