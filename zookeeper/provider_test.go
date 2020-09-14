package zookeeper

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


var testAccProviders map[string]*schema.Provider
var testAccProviderFactories func(providers *[]*schema.Provider) map[string]func() (*schema.Provider, error)
var testAccProvider *schema.Provider
var testAccProviderFunc func() *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"zookeeper": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("ZOOKEEPER_HOST"); err == "" {
		t.Fatal("ZOOKEEPER_HOST must be set for accqeptance tests")
	}
	if err := os.Getenv("ZOOKEEPER_PORT"); err == "" {
		t.Fatal("ZOOKEEPER_PORT must be set for acceptance tests")
	}
}