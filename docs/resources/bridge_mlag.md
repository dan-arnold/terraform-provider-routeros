# routeros_bridge_mlag (Resource)


## Example Usage
```terraform
resource "routeros_bridge_mlag" "mlag" {
  bridge    = "bridge1"
  peer_port = "stack-link"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `bridge` (String) The bridge interface where MLAG is being created.
- `peer_port` (String) An interface that will be used as a peer port. Both peer devices are using inter-chassis communication over these peer ports to establish MLAG and update the host table. Peer port should be isolated on a different untagged VLAN using a pvid setting. Peer port can be configured as a bonding interface.

### Optional


### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
terraform import routeros_bridge_mlag.mlag .
```
