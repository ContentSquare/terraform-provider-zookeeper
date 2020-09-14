# zookeeper_node Resource

Create/Update zookeeper nodes

## Example Usage

```hcl
resource "zookeeper_node" "bladibla" {
  path = "/bladibla"
  data = "testme flute"
}
```

## Argument Reference

* `path` - Required Path name to read data from

* `data` - Optional Data to write to the zookeeper node

## Attribute Reference

* `path` - The znode path

* `zxid` - The Zookpeer node id

* `data` - Data written on the node
