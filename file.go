package main

import (
	"io"
	"os"
)

func writePackage(fn string, r io.Reader) error {
	a, err := os.Create(fn)
	if err != nil {
		return err
	}
	if _, err := io.Copy(a, r); err != nil {
		os.Remove(fn)
		return err
	}
	return nil
}
