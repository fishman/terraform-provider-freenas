---
layout: "github"
page_title: "Provider: FreeNAS"
sidebar_current: "docs-freenas-index"
description: |-
  The FreeNAS provider is used to interact with FreeNAS resources.
---

# FreeNAS Provider

The FreeNAS provider is used to interact with FreeNAS resources.

The provider allows you to manage your FreeNAS shares easily.
It needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the FreeNAS Provider
provider "freenas" {
  server   = "${var.freenas_server}"
  user     = "${var.freenas_user}"
  password = "${var.freenas_password}"
}

# Add an NFS share
resource "freenas_nfs_share" "share_name" {
  # ...
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `user` - (Optional) This is the FreeNAS username. It must be provided, but
  it can also be sourced from the `FREENAS_USER` environment variable.

* `password` - (Optional) This is the FreeNAS password. It must be provided, but
  it can also be sourced from the `FREENAS_PASSWORD` environment variable.

* `server` - (Optional) This is the target URL for the FreeNAS API endpoint. It must be provided, but
  it can also be sourced from the `FREENAS_SERVER` environment variable.  The value must end without a slash for example http://freenas.local.
