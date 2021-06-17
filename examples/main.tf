terraform {
  required_providers {
    validate = {
      source = "articulate/validate"
    }
  }
}

variable "foo" {
  default = ""
}

variable "bar" {
  default = ""
}

data "validate" "test" {
  condition     = var.foo != "" || var.bar != ""
  error_message = "You must set foo or bar."
}
