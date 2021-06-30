terraform {
  required_providers {
    validation = {
      source = "articulate/validation"
    }
  }
}

variable "foo" {
  default = ""
}

variable "bar" {
  default = ""
}

data "validation" "test" {
  condition     = var.foo != "" || var.bar != ""
  error_message = "You must set foo or bar."
}
