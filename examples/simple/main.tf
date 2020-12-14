provider "zookeeper" {
  host = "localhost"
  port = 2181
}


# Create a simple znode on path /bladibla with `testme flute` as data
resource "zookeeper_node" "bladibla" {
  path = "/bladibla"
  data = "testme flute"
}

# Create a simple child znode on path /bladibla with `testme flute too` as data
resource "zookeeper_node" "bladibla_child" {
  path = "${zookeeper_node.bladibla.path}/child1"
  data = "testme flute too"
}

