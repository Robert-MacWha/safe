package safe

// MapOption maps the value of an Option to a new value, leaving the error untouched.
func MapOption[T any, U any](o Option[T], f func(T) U) Option[U] {
	if o.IsNone() {
		return None[U]()
	}
	return Some(f(o.Unwrap()))
}

// MapOptionOr maps the value of an Option to a new value or returns a default value if the Option is None.
func MapOptionOr[T any, U any](o Option[T], def U, f func(T) U) U {
	if o.IsNone() {
		return def
	}
	return f(o.Unwrap())
}

// MapResult maps the value of a Result to a new value, leaving the error untouched.
func MapResult[T any, U any](r Result[T], f func(T) U) Result[U] {
	if r.IsErr() {
		return Err[U](r.UnwrapErr())
	}
	return Ok(f(r.Unwrap()))
}

// MapResultOr maps the value of a Result to a new value or returns a default value if the Result is an error.
func MapResultOr[T any, U any](r Result[T], def U, f func(T) U) U {
	if r.IsErr() {
		return def
	}
	return f(r.Unwrap())
}
