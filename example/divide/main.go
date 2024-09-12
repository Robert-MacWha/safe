package divide

import (
	"fmt"

	"github.com/robert-macwha/safe/pkg/safe"
)

func SafeDivide(a, b int) safe.Result[int] {
	if b == 0 {
		return safe.Err[int](fmt.Errorf("cannot divide by zero"))
	}
	return safe.Ok(a / b)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}
