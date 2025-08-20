package mikrotik

import (
	"fmt"
	"testing"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestWirelessSecurityProfile_basic(t *testing.T) {
	// This test is skipped, until we find a way to include required packages.
	//
	// Since RouterOS 7.13, 'wireless' package is separate from the main system package
	// and there is no easy way to install it in Docker during tests.
	// see https://help.mikrotik.com/docs/spaces/ROS/pages/40992872/Packages#Packages-RouterOSpackages
	client.SkipIfRouterOSV7OrLater(t, sysResources)

	resourceName := "mikrotik_wireless_security_profile.testacc"
	name := acctest.RandomWithPrefix("security-profile")
	resource.Test(t,
		resource.TestCase{
			ProtoV5ProviderFactories: testAccProtoV5ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
					resource "mikrotik_wireless_security_profile" "testacc" {
						name = %q
						mode = "dynamic-keys"
						authentication_types = ["wpa2-psk"]
						wpa2_pre_shared_key = "supersecretkey123"
					}`, name),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet(resourceName, "id"),
						resource.TestCheckResourceAttr(resourceName, "name", name),
						resource.TestCheckResourceAttr(resourceName, "mode", "dynamic-keys"),
						resource.TestCheckResourceAttr(resourceName, "authentication_types.#", "1"),
						resource.TestCheckTypeSetElemAttr(resourceName, "authentication_types.*", "wpa2-psk"),
						resource.TestCheckResourceAttr(resourceName, "wpa2_pre_shared_key", "supersecretkey123"),
					),
				},
				{
					Config: fmt.Sprintf(`
					resource "mikrotik_wireless_security_profile" "testacc" {
						name = %q
						mode = "static-keys-optional"
						authentication_types = ["wpa-psk", "wpa2-psk"]
						wpa2_pre_shared_key = "newsupersecretkey456"
					}`, name),

					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet(resourceName, "id"),
						resource.TestCheckResourceAttr(resourceName, "name", name),
						resource.TestCheckResourceAttr(resourceName, "mode", "static-keys-optional"),
						resource.TestCheckResourceAttr(resourceName, "authentication_types.#", "2"),
						resource.TestCheckTypeSetElemAttr(resourceName, "authentication_types.*", "wpa-psk"),
						resource.TestCheckTypeSetElemAttr(resourceName, "authentication_types.*", "wpa2-psk"),
						resource.TestCheckResourceAttr(resourceName, "wpa2_pre_shared_key", "newsupersecretkey456"),
					),
				},
				{
					ImportState:       true,
					ImportStateVerify: true,
					ResourceName:      resourceName,
				},
			},
		},
	)
}
