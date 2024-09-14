# Safe

Rust-like safe Option & Result generics in Go.

Helpful for more consise error-handling and letting you return instances instead of pointers from functions that can error.

## Features

-   `Result` and `Option` types.
-   Unwrap Handler to streamline error management.
-   SafeLint linter to ensure correct usage of this package.

## Types

### `Result`

Interface encapsulates success or failure state.

```go
// Ok returns an ok Result.
func Ok[T any](data T) Result[T]

// Err returns an errored Result.
func Err[T any](err error) Result[T]

// AsResult returns a result from data and an error.  Helpful for converting results
// from normal functions to safe results.
func AsResult[T any](data T, err error) Result[T]

type Result[T any] struct {
	// IsOk returns true if the Result is ok.
	IsOk() bool

	// IsErr returns true if the Result is an error.
	IsErr() bool

	// Expect returns the value of an ok Result or panics with a custom error if the Result is an error.
	Expect(msg string) T

	// Unwrap returns the value of an ok Result or panics if the Result is an error.
	Unwrap() T

	// UnwrapOr returns the value of an ok Result or a default value if the Result is an error.
	UnwrapOr(def T) T

	// UnwrapErr returns the error of an errored Result or panics if the Result is ok.
	UnwrapErr() error

	String() string
}
```

### `Option`

Interface represents a value that may or may not exist. Supports JSON marshalling and unmarshalling.

```go
// Some returns a Some Option.
func Some[T any](data T) Option[T]

// None returns a None Option.
func None[T any]() Option[T]

type Option[T any] struct {
	// IsSome returns true if the Option is a Some.
	IsSome() bool

	// IsNone returns true if the Option is a None.
	IsNone() bool

	// Expect returns the value of a Some Option or panics with a custom error if the Option is a None.
	Expect(err error) T

	// Unwrap returns the value of a Some Option or panics if the Option is a None.
	Unwrap() T

	// UnwrapOr returns the value of a Some Option or a default value if the Option is a None.
	UnwrapOr(def T) T

	// Ok returns a Result containing the value of a Some Option or an error if the Option is a None.
	Ok(err error) Result[T]

	String() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

```

## Handler

The safe handler lets you safely unwrap `Options` or `Results`, catch any panics in another `Result` object, and return that `Result` from a function. This replaces the need for the `if err != nil {}` pattern and lets your code focus on the happy execution path.

See [example/io](./example/io/main.go) for an example.

## Map

The map functions can be used to transform instances of Result or Option without unwrapping them.

```go
// MapOption maps the value of an Option to a new value, leaving the error untouched.
func MapOption[T any, U any](o Option[T], f func(T) U) Option[U]

// MapOptionOr maps the value of an Option to a new value or returns a default value if the Option is None.
func MapOptionOr[T any, U any](o Option[T], def U, f func(T) U) U

// MapResult maps the value of a Result to a new value, leaving the error untouched.
func MapResult[T any, U any](r Result[T], f func(T) U) Result[U]

// MapResultOr maps the value of a Result to a new value or returns a default value if the Result is an error.
func MapResultOr[T any, U any](r Result[T], def U, f func(T) U) U
```

## SafeLint

SafeLint is a linter designed for this package to ensure correct usage.

### SafeHandlerAnalyzer

SafeHandlerAnalyzer is a linter that ensures `safe.Handle` is called before unwrapping any results or options.

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
