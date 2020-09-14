# zookeeper_node Data Source

Reads data from a zookeeper node

## Example Usage

```hcl
data "zookeeper_node" "bladibla" {
  path = "/bladibla"
}
```

## Argument Reference

* `path` - Required Path name to read data from

## Attribute Reference

* `data` - List of read data

```hcl
{
  "data" = [
    {
      "data" = "{\"bla\":\"dibla\",\"intval\":17}"
      "zxid" = 46
    },
  ]
  "id" = "/bladibla"
  "path" = "/bladibla"
}
```
