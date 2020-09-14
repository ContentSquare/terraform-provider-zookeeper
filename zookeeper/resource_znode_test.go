package zookeeper

import (
	fmt "fmt"
	"github.com/contentsquare/terraform-provider-zookeeper/api/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccZookeeperNode_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		//ProviderFactories: testAccProviderFactories(&providers),
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZookeeperNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckZookeeperNodeBasicResource("/bladibla", "some nice test data"),
				Check: testResourceTopic_initialCheck,
			},
		},
	})
}

func TestAccZookeeperNode_update(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		//ProviderFactories: testAccProviderFactories(&providers),
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZookeeperNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckZookeeperNodeBasicResource("/bladibla", "some nice test data"),
				Check: testResourceTopic_initialCheck,
			},
			{
				Config: testAccCheckZookeeperNodeUpdateResource("/bladibla", "new data for tests"),
				Check: testResourceTopic_updatedcheck,
			},
		},
	})
}

func testAccCheckZookeeperNodeDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zookeeper_node" {
			continue
		}

		path := rs.Primary.ID

		err := c.DeleteZnode(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckZookeeperNodeBasicResource(path, data string) string {
	return fmt.Sprintf(`
	resource "zookeeper_node" "new" {
		path = "%s"
		data = "%s"
	}`, path, data)
}

func testAccCheckZookeeperNodeUpdateResource(path, data string) string {
	return fmt.Sprintf(`
	resource "zookeeper_node" "new" {
		path = "%s"
		data = "%s"
	}`, path, data)
}

func testResourceTopic_initialCheck(s *terraform.State) error {
	resourceState := s.Modules[0].Resources["zookeeper_node.new"]
	if resourceState == nil {
		return fmt.Errorf("resource not found in state")
	}

	instanceState := resourceState.Primary
	if instanceState == nil {
		return fmt.Errorf("resource has no primary instance")
	}

	path := instanceState.ID

	if path != instanceState.Attributes["path"] {
		return fmt.Errorf("id doesn't match path")
	}

	c := testAccProvider.Meta().(*client.Client)
	znode, err := c.ReadZNode(path)
	if err != nil {
		return err
	}

	if znode.Path != path {
		return fmt.Errorf("znode path is not equal to path. %s != %s", znode.Path, path)
	}

	if string(znode.Data) != "some nice test data" {
		return fmt.Errorf("znode data is not equal to expected. %s != %s", string(znode.Data), "some nice test data")
	}

	return nil
}

func testResourceTopic_updatedcheck(s *terraform.State) error {
	resourceState := s.Modules[0].Resources["zookeeper_node.new"]
	if resourceState == nil {
		return fmt.Errorf("resource not found in state")
	}

	instanceState := resourceState.Primary
	if instanceState == nil {
		return fmt.Errorf("resource has no primary instance")
	}

	path := instanceState.ID

	if path != instanceState.Attributes["path"] {
		return fmt.Errorf("id doesn't match path")
	}

	c := testAccProvider.Meta().(*client.Client)
	znode, err := c.ReadZNode(path)
	if err != nil {
		return err
	}

	if znode.Path != path {
		return fmt.Errorf("znode path is not equal to path. %s != %s", znode.Path, path)
	}

	if string(znode.Data) != "new data for tests" {
		return fmt.Errorf("znode data is not equal to expected. %s != %s", string(znode.Data), "new data for tests")
	}

	return nil
}
