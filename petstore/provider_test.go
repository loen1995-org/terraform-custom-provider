package petstore

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

//Initializes global variables
func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"petstore": testAccProvider,
	}
}

//Tests that the provider schema is valid
func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

//Tests that the provider can be initialized
func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

//Tests that the PETSTORE_ADDRESS environment variable is set
func testAccPreCheck(t *testing.T) {
	if os.Getenv("PETSTORE_ADDRESS") == "" {
		t.Fatal("PETSTORE_ADDRESS must be set for acceptance tests")
	}

	if diags := Provider().Configure(context.Background(), &terraform.ResourceConfig{}); diags.HasError() {
		for _, d := range diags {
			if d.Severity == diag.Error {
				t.Fatalf("err: &s", d.Summary)
			}
		}
	}
}
