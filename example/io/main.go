package main

import (
	"io"
	"os"

	"github.com/robert-macwha/safe/pkg/safe"
)

func SafeIoRead(path string) (res safe.Result[string]) {
	safe.Handle(&res)

	file := safe.Res(os.Open(path)).Unwrap()
	defer file.Close()

	contents := safe.Res(io.ReadAll(file)).Unwrap()
	return safe.Ok(string(contents))
}

func IoRead(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
