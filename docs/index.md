# Zookeeper Provider

Provider to read and write data to/from zookeeper

## Example Usage

```hcl
provider "zookeeper" {
  host = "localhost"
  port = 2181
}
```

## Argument Reference

* `host` - The zookeeper host. defaults to localhost

Read from env `ZOOKEEPER_HOST`

* `port` - The zookeeper port. defaults to 2181

Read from env `ZOOKEEPER_PORT`