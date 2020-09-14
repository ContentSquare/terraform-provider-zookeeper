provider "zookeeper" {
  host = "localhost"
  port = 2181
}


# Create a simple znode on path /bladibla with `testme flute` as data
resource "zookeeper_node" "bladibla" {
  path = "/bladibla"
  data = "testme flute"
}

# Create a child znode on path /bladibla/testme2 with json data as data
resource "zookeeper_node" "bladibla_child" {
  path = "${zookeeper_node.bladibla.path}/testme2"
  data = jsonencode({ bla = "dibla", intval = 17 })
}


# Read child znode on path /bladibla/testme2
data "zookeeper_node" "read_bladibla_chile" {
  path = zookeeper_node.bladibla_child.path
}

# output test data
output "test_output" {
  value = jsondecode(data.zookeeper_node.read_bladibla_chile.data.0.data)
}
