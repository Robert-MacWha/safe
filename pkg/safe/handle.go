package safe

// unwrapError is the error type emitted by this safe package when a item is
// unwrapped incorrectly. It can be caught by the Handle function.
type unwrapError struct {
	err error
}

// Handle catches any panic errors issued by the Result and wraps them in a
// Result
func Handle[T any](res *Result[T]) {
	if r := recover(); r != nil {
		err, ok := r.(unwrapError)
		if !ok {
			panic(err)
		}

		*res = result[T]{err: err.err}
	}
}
