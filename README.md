# go-terraform
Library to encapsulate the terraform utility to be called as a
golang library

### Why does this exist?
Between the unreliability of the internal interfaces in the terraform library and then
need to communicate with providers, we'll wrap the terraform command in bash, rather
than importing the `github.com/hashicorp/terraform` library and calling methods
directly. See https://github.com/hashicorp/terraform/issues/12582 for more info.
