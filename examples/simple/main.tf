provider "zookeeper" {
  host = "localhost"
  port = 2181
}


# Create a simple znode on path /bladibla with `testme flute` as data
resource "zookeeper_node" "bladibla" {
  path = "/bladibla"
  data = "testme flute"
}