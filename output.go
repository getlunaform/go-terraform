package goterraform

import "io"

type TerraformOutput struct {
	Stderr io.ReadCloser
	Stdout io.ReadCloser
}
