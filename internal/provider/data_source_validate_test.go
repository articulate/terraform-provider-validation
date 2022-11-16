package provider_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceValidate(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
data "validation" "foo" {
  condition     = true
  error_message = "Foo is invalid."
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.validation.foo", "id", regexp.MustCompile("^Foo is invalid")),
				),
			},
			{
				Config: `
data "validation" "bar" {
  condition     = false
  error_message = "Bar is invalid."
}`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("Bar is invalid"),
			},
		},
	})
}
