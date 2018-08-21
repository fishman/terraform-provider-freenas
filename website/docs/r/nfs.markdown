---
layout: "github"
page_title: "FreeNAS: freenas_nfs_share"
sidebar_current: "docs-freenas-resource-nfs"
description: |-
  Provides a NFS share resource.
---

# freenas_nfs_share

Provides a NFS share resource.

This resource allows you to add/remove NFS shares to your FreeNAS. When applied,
a new NFS share will be created. When destroyed, that share will be removed.

## Example Usage

```hcl
# Add a NFS share to your FreeNAS
resource "freenas_nfs_share" "some_share" {
  comment     = "some-comment"
  paths       = ["/some/mount/path"]
}
```

## Argument Reference

The following arguments are supported:

* `comment` - (Optional) A comment for the share.
* `paths` - (Required) The paths that the NFS share points to.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the created share.

## Import

NFS shares can be imported using the FreeNAS NFS share Id e.g.

```
$ terraform import freenas_nfs_share.core 10
```
