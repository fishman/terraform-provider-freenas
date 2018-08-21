---
layout: "github"
page_title: "FreeNAS: freenas_storage_dataset"
sidebar_current: "docs-freenas-resource-dataset"
description: |-
  Provides a storage dataset resource.
---

# freenas_storage_dataset

Provides a storage dataset resource.

This resource allows you to add/remove datasets on your FreeNAS. When applied,
a new dataset will be created. When destroyed, that dataset will be removed.

## Example Usage

```hcl
# Add a dataset to your FreeNAS
resource "freenas_storage_dataset" "some_dataset" {
  parent      = "some/parent"
  name        = "some-name"
  comment     = "some-comment"
}
```

## Argument Reference

The following arguments are supported:

* `parent` - (Required) The parent path of the new dataset.
* `name` - (Required) The name of the new dataset.
* `comment` - (Optional) A comment for the dataset.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the created .

## Import

Datasets can be imported using the dataset name.

```
$ terraform import freenas_storage_dataset.core ssd/some/dataset
```
