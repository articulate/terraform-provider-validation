# Terraform Provider Validation

Terraform provides a way to add validation to variables, but you can't use other
variables in this validation. This allows you to do complex validation on variables
or resource outputs.

## Using the provider

This provider has a single _validate_ data source.

```hcl
data "validate" "foo" {
  condition     = var.foo != "" || var.bar != ""
  error_message = "You must set foo or bar."
}
```

## Requirements

* [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
* [Go](https://golang.org/doc/install) >= 1.16

## Building The Provider

To build for the current system, run `make build`. If you want to build for all
systems, run `make all`.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org)
installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `make build`.

Run the full suite of Acceptance tests, run `make testacc`.
