<a href="https://terraform.io">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform Provider for Zookeeper

The Terraform Zookeeper provider is a plugin for Terraform that allows interacting with zookeeper node.
This provider is maintained by @contentsquare

## Usage

### TF 0.13+

```hcl
terraform {
  required_providers {
    zookeeper = {
      
      source  = "contentsqure/zookeeper"
      version = "~> 1.0"
    }
  }
}
```

## Examples

- [simple](./examples/simple)
- [reader](./examples/reader)

## Testing

Acceptance tests require a zookeeper node in order to create/update/delete znodes

Start a zookeeper container, and expose tcp port 2181 onto localhost

```
~# docker run --rm -p 2181:2181 -it zookeeper
```

Run the acceptance tests:

```shell
make testacc
```

